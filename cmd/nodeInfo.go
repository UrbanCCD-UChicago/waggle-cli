package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var nodeInfoCmd = &cobra.Command{
	Use:   "node-info",
	Short: d.ShortDesc(d.NodeInfoDesc),
	Long:  d.NodeInfoDesc,
	Run: func(cmd *cobra.Command, args []string) {
		var output string

		if env.Name == u.AdminEnv.Name {
			output = fmt.Sprintf("%s", node)

		} else {
			command := u.NewGenericCommand()
			command.AddTargetedRoutine(u.NodeControllerTarget.Name, "echo TBD")
			command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "echo TBD")
			command.AddTargetedRoutine(u.WagmanTarget.Name, "echo TBD")

			output = command.Execute(targets)
		}

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(nodeInfoCmd)
}
