package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var enableWagglePluginCmd = &cobra.Command{
	Use:   "enable-waggle-plugin",
	Short: d.ShortDesc(d.EnableWagglePluginDesc),
	Long:  d.EnableWagglePluginDesc,
	Run: func(_cmd *cobra.Command, args []string) {
		name := args[0]
		userCmd := fmt.Sprintf("systemctl enable waggle-plugin-%s && systemctl start waggle-plugin-%s", name, name)

		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, userCmd)
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(enableWagglePluginCmd)
}
