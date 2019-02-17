package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var diableWagglePlunginCmd = &cobra.Command{
	Use:   "disable-waggle-plugin",
	Short: d.ShortDesc(d.DisableWagglePluginDesc),
	Long:  d.DisableWagglePluginDesc,
	Run: func(_cmd *cobra.Command, args []string) {
		name := args[0]
		userCmd := fmt.Sprintf("systemctl stop waggle-plugin-%s && systemctl disable waggle-plugin-%s", name, name)

		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, userCmd)
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(diableWagglePlunginCmd)
}
