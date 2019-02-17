package cmd

import (
	d "github.com/UrbanCCD-UChicago/waggle/descriptions"
	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var refreshCacheCmd = &cobra.Command{
	Use:   "refresh-cache",
	Short: d.ShortDesc(d.RefreshCache),
	Long:  d.RefreshCache,
	Run: func(_cmd *cobra.Command, _args []string) {
		u.LoadDatabase()
		u.PrintVerbose("ok")
	},
}

func init() {
	rootCmd.AddCommand(refreshCacheCmd)
}
