package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: d.ShortDesc(d.DateDesc),
	Long:  d.DateDesc,
	Run: func(_cmd *cobra.Command, _args []string) {
		command := u.NewGenericCommand()
		command.AddTargetedRoutine(u.NodeControllerTarget.Name, "date --rfc-3339=seconds")
		command.AddTargetedRoutine(u.EdgeProcessorTarget.Name, "date --rfc-3339=seconds")
		command.AddTargetedRoutine(u.WagmanTarget.Name, "wagman-client date")

		output := command.Execute(targets)
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
