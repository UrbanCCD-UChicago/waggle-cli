package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var lsDevCmd = &cobra.Command{
	Use:   "ls-dev",
	Short: d.ShortDesc(d.LsDev),
	Long:  d.LsDev,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "ls -al /dev/waggle_*")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "ls -al /dev/waggle_*")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(lsDevCmd)
}
