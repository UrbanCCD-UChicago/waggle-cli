package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var failCountsCmd = &cobra.Command{
	Use:   "fail-counts",
	Short: d.ShortDesc(d.FailCounts),
	Long:  d.FailCounts,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "wagman-client fc")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "wagman-client fc")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(failCountsCmd)
}
