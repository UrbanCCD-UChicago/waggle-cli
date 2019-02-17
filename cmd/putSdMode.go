package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var putSdModeCmd = &cobra.Command{
	Use:   "put-sd-mode",
	Short: d.ShortDesc(d.PutSDMode),
	Long:  d.PutSDMode,
	Run: func(_cmd *cobra.Command, _args []string) {
		userCmd := "ssh ep -x halt && " +
			"wagman-client bs 1 sd && " +
			"wagman-client stop 1 30 && " +
			"wagman-client bs 0 sd && " +
			"wagman-client start 1 && " +
			"wagman-client stop 0 60 && " +
			"shutdown -h now"

		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, userCmd)
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(putSdModeCmd)
}
