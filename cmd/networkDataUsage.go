package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var networkDataUsageCmd = &cobra.Command{
	Use:   "network-data-usage",
	Short: d.ShortDesc(d.NetworkDataUsage),
	Long:  d.NetworkDataUsage,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "ifconfig ppp0 | grep \"TX packets\" | tr -s \" \"")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(networkDataUsageCmd)
}
