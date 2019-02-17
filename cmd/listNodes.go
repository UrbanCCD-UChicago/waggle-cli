package cmd

import (
	"fmt"

	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var (
	aliveOnly bool
	format    string
)

var listNodesCmd = &cobra.Command{
	Use:   "list-nodes",
	Short: d.ShortDesc(d.ListNodes),
	Long:  d.ListNodes,
	Run: func(_cmd *cobra.Command, _args []string) {
		var nodes []u.Node

		if aliveOnly {
			nodes = u.GetLiveNodes()
		} else {
			nodes = u.GetAllNodes()
		}

		for _, n := range nodes {
			if format == "csv" {
				fmt.Println(n.AsCSVRow())
			} else {
				fmt.Println(n.AsTableRow())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listNodesCmd)
	listNodesCmd.Flags().BoolVar(&aliveOnly, "alive-only", false, "filter to only live nodes")
	listNodesCmd.Flags().StringVar(&format, "format", "", "format the output")
}
