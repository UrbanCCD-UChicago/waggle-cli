package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var disableEdgeProcessor = &cobra.Command{
	Use:   "disable-edge-processor",
	Short: d.ShortDesc(d.DisbleEdgeProcessorDesc),
	Long:  d.DisbleEdgeProcessorDesc,
	Run: func(_cmd *cobra.Command, args []string) {
		edegeProcessorID := args[0]
		userCmd := fmt.Sprintf("wagman-client stop %s && wagman-client disable %s", edegeProcessorID, edegeProcessorID)

		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, userCmd)
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(disableEdgeProcessor)
}
