package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: d.ShortDesc(d.Hostname),
	Long:  d.LsDev,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "cat /etc/hostname")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "cat /etc/hostname")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(hostnameCmd)
}
