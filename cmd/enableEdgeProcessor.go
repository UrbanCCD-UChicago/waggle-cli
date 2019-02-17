package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var enableEdgeProcessorCmd = &cobra.Command{
	Use:   "enable-edge-processor",
	Short: d.ShortDesc(d.EnableEdgeProcessorDesc),
	Long:  d.EnableEdgeProcessorDesc,
	Run: func(_cmd *cobra.Command, args []string) {
		epID := args[0]
		userCmd := fmt.Sprintf("waggle-client bs %s sd && wagman-client enable %s && wagman-client start %s", epID, epID, epID)

		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, userCmd)
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(enableEdgeProcessorCmd)
}
