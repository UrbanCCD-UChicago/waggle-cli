package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var listProcessesCmd = &cobra.Command{
	Use:   "list-processes",
	Short: d.ShortDesc(d.ListProcesses),
	Long:  d.ListProcesses,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "systemctl | grep \"waggle\\|rabbit\" | tr -s \" \"")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "systemctl | grep \"waggle\\|rabbit\" | tr -s \" \"")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(listProcessesCmd)
}
