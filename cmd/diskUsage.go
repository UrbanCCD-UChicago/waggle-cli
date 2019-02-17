package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var diskUsageCmd = &cobra.Command{
	Use:   "disk-usage",
	Short: d.ShortDesc(d.DiskUsage),
	Long:  d.DiskUsage,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "df -h | grep mmcblk")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "df -h | grep mmcblk")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(diskUsageCmd)
}
