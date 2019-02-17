package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var resetWagmanCmd = &cobra.Command{
	Use:   "reset-wagman",
	Short: d.ShortDesc(d.ResetWagmanDesc),
	Long:  d.ResetWagmanDesc,
	Run: func(_cmd *cobra.Command, _args []string) {
		userCmd := "ssh ep -x halt && " +
			"sleep 45 && " +
			"/usr/lib/waggle/nodecontroller/scripts/reset-wagman && " +
			"sleep 15 && " +
			"date && " +
			"ls -l /dev/waggle_sysmon "

		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, userCmd)
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(resetWagmanCmd)
}
