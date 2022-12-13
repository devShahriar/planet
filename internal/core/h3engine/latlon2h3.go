package h3engine

import (
	"fmt"

	"github.com/devShahriar/planet/contract"
	"github.com/uber/h3-go/v4"
)

func ConvertToH3(lat float64, lon float64, resolution int) {
	latlon := h3.NewLatLng(lat, lon)
	cell := h3.LatLngToCell(latlon, resolution)

	fmt.Printf("Input Lat : %v | Lon : %v\n", lat, lon)
	fmt.Printf(contract.Purple+"Output h3 index => %+s \n", cell)
}
