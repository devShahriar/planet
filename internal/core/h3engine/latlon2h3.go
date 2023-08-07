package h3engine

import (
	"fmt"

	"github.com/uber/h3-go/v3"
)

func ConvertToH3(lat float64, lon float64, resolution int) {
	//latlon := h3.NewLatLng(lat, lon)
	// cell := h3.FromGeo(latlon, resolution)

	// fmt.Printf("Input Lat : %v | Lon : %v\n", lat, lon)
	// fmt.Printf(contract.Purple+"Output h3 index => %+s \n", h3.IndexToString(cell))

	geocord := h3.GeoCoord{Latitude: lat, Longitude: lon}
	h3_index := h3.FromGeo(geocord, resolution)
	fmt.Print(h3.ToString(h3_index))
}
