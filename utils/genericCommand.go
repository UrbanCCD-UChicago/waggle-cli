package utils

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

// GenericCommand wraps the functionality that is present in the majority of the commands that
// would otherwise be written out in long form -- and most likely copy/pasted from one to another.
//
// Routines is a map of target name keys and command/routine values. For each target passed into
// the command's execution, there should be a configured routine for that target. If there is no
// routine acceptable for that targeted system then set `""` (the empty string) as the value -- this
// will be the case for many commands on the wagman system target.
//
// Results is another map of target names and an array of output strings. For each routine we
// expect some kind of output. This allows us to split the output by new lines and append those
// lines to the array. It then makes it trivial to either visually grok the results of the command
// or process the results further down a pipeline.
//
// Examples:
//
// 	getDatesCommand := NewGenericCommand()
// 	getDatesCommand.AddTargetedRoutine(u.NodeControllerTarget.Name, "date")
// 	getDatesCommand.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "date")
// 	getDatesCommand.AddTargetedRoutine(u.WagmanTarget.Name, "wagman-client date")
//
// 	lsDevCommand := NewGenericCommand()
// 	lsDevCommand.AddTargetedRoutine(u.NodeControllerTarget.Name, "ls -l /dev/")
// 	lsDevCommand.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "ls -l /dev/")
// 	lsDevCommand.AddTargetedRoutine(u.WagmanTarget.Name, "") // not available on wagman
type GenericCommand struct {
	Routines map[string]string
	Results  map[string][]string
}

// NewGenericCommand creates a new, blank GenericCommand struct.
func NewGenericCommand() GenericCommand {
	return GenericCommand{
		Routines: make(map[string]string),
		Results:  make(map[string][]string),
	}
}

// AddTargetedRoutine adds a new routine for a given target for the command.
func (c GenericCommand) AddTargetedRoutine(targetName string, routine string) {
	c.Routines[targetName] = routine
}

// Execute runs the routines on the given targets and returns a JSON string of results.
func (c GenericCommand) Execute(targets map[string]string) string {
	var wg sync.WaitGroup

	for targetName, hostname := range targets {
		if localCmd := c.Routines[targetName]; localCmd != "" {
			wg.Add(1)

			go func(lCmd string, lHostname string, lResults map[string][]string, lTargetName string) {
				defer wg.Done()

				lfCmd := fmt.Sprintf("'%s'", lCmd)
				out, err := exec.Command("ssh", "-q", lHostname, "bash", "-c", lfCmd).CombinedOutput()
				if err != nil {
					lResults[lTargetName] = append(lResults[lTargetName], fmt.Sprintf("%s", err))
				} else {
					for _, l := range strings.Split(string(out), "\n") {
						if l != "" {
							lResults[lTargetName] = append(lResults[lTargetName], l)
						}
					}
				}
			}(localCmd, hostname, c.Results, targetName)
		}
	}

	wg.Wait()

	data, _ := json.MarshalIndent(c.Results, "", "  ")
	return fmt.Sprintf("%s", data)
}
