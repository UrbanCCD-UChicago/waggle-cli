package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var uptimeCmd = &cobra.Command{
	Use:   "uptime",
	Short: d.ShortDesc(d.UptimeDesc),
	Long:  d.UptimeDesc,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "uptime")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "uptime")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "wagman-client up")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(uptimeCmd)
}
