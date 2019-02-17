package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var currentUsageCmd = &cobra.Command{
	Use:   "current-usage",
	Short: d.ShortDesc(d.CurrentUsage),
	Long:  d.CurrentUsage,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "wagman-client cu")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "wagman-client cu")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(currentUsageCmd)
}
