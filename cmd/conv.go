package cmd

import (
	"github.com/spf13/cobra"
)

var Conv = &cobra.Command{
	Use:   "conv",
	Short: "conv",
}

func init() {
	rootCmd.AddCommand(Conv)
	// latlon := contract.GetLatLon()

	//runCmd.Flags().StringVarP(, "lat", "", "", "config path")
}
