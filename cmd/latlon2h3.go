package cmd

import (
	"fmt"
	"strconv"

	"github.com/devShahriar/planet/contract"
	"github.com/devShahriar/planet/internal/core/h3engine"
	"github.com/spf13/cobra"
)

var latlon2h3 = &cobra.Command{
	Use:     "latlon2h3",
	Short:   "latlon2h3",
	Example: "planet conv latlon2h3 --lat <latitude> --lon <longitude> -r <resolution>",
	Run: func(cmd *cobra.Command, args []string) {

		input := contract.GetInputPayload()

		latFloat, err := strconv.ParseFloat(input.Lat, 64)
		if err != nil {
			fmt.Println("Invalid latitude value", err)
		}
		lonFloat, err := strconv.ParseFloat(input.Lon, 64)
		if err != nil {
			fmt.Println("Invalid lontitude value", err)
		}

		resolutionInt, err := strconv.Atoi(input.Resolution)
		if err != nil {
			fmt.Println("Invalid resolution value type int", err)
		}

		h3engine.ConvertToH3(latFloat, lonFloat, resolutionInt)
	},
}

func init() {
	Conv.AddCommand(latlon2h3)
	input := contract.GetInputPayload()
	latlon2h3.Flags().StringVarP(&input.Lat, "lat", "", "", "Input latitude")
	latlon2h3.Flags().StringVarP(&input.Lon, "lon", "", "", "Input lontidue")
	latlon2h3.Flags().StringVarP(&input.Resolution, "resolution", "r", "", "Input resolution")

}
