package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var eeresetWagmanCmd = &cobra.Command{
	Use:   "eereset-wagman",
	Short: d.ShortDesc(d.EEResetWagmanDesc),
	Long:  d.EEResetWagmanDesc,
	Run: func(_cmd *cobra.Command, args []string) {
		eeresetCmd := "ssh ep -x halt && " +
			"sleep 45 && " +
			"/usr/lib/waggle/nodecontroller/scripts/reset-wagman && " +
			"sleep 15 && " +
			"wagman-client eereset && " +
			"wagman-client reset"

		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, eeresetCmd)
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(eeresetWagmanCmd)
}
