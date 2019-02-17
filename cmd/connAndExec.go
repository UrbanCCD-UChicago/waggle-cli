package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var connAndExecCmd = &cobra.Command{
	Use:   "connect-and-execute",
	Short: d.ShortDesc(d.ConnAndExecDesc),
	Long:  d.ConnAndExecDesc,
	Run: func(_cmd *cobra.Command, args []string) {
		userCmd := args[0]

		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, userCmd)
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, userCmd)
		command.AddTargetedRoutine(u.WagmanTarget.Name, userCmd)

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(connAndExecCmd)
	// for this command node and system are required
	rootCmd.MarkFlagRequired("node")
	rootCmd.MarkFlagRequired("system")
}
