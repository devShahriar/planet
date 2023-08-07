package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/uber/h3-go/v3"
)

func main() {
	// Read the GeoJSON file
	filePath := "/Users/shahriar/Code/planet/poc/h3_poc/dhaka.geojson"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading GeoJSON file:", err)
		return
	}

	// Parse the GeoJSON
	var geoJSON struct {
		Features []struct {
			Geometry struct {
				Type        string
				Coordinates [][][]float64
			}
		}
	}
	err = json.Unmarshal(data, &geoJSON)
	if err != nil {
		fmt.Println("Error parsing GeoJSON:", err)
		return
	}

	// Create a slice to store the H3 indexes
	h3Indexes := make([]string, 0)

	// Iterate through each feature in the GeoJSON
	for _, feature := range geoJSON.Features {
		geometry := feature.Geometry
		if geometry.Type == "Polygon" {
			// Convert the polygon to H3 indexes
			polygon := geometry.Coordinates
			h3Indexes = append(h3Indexes, h3IndexesFromPolygon(polygon[0])...)
		} else if geometry.Type == "MultiPolygon" {
			// Convert the multipolygon to H3 indexes
			multipolygon := geometry.Coordinates
			for _, polygon := range multipolygon {
				h3Indexes = append(h3Indexes, h3IndexesFromPolygon(polygon)...)
			}
		}
	}

	// Print the H3 indexes
	fmt.Println("lat,lon")
	for _, index := range h3Indexes {
		//fmt.Println(index)
		geo := h3.ToGeo(h3.FromString(index))
		out := fmt.Sprintf(`"%v,%v",`, geo.Latitude, geo.Longitude)

		fmt.Println(out)
	}
}

// Convert a polygon to H3 indexes
func h3IndexesFromPolygon(cords [][]float64) []string {
	// Convert the polygon coordinates to GeoCoord
	polygon := h3.GeoPolygon{
		Geofence: make([]h3.GeoCoord, 0),
	}
	for i, _ := range cords {

		latlng := h3.GeoCoord{Latitude: cords[i][1], Longitude: cords[i][0]}
		polygon.Geofence = append(polygon.Geofence, latlng)
	}

	// Convert the coordinates to H3 indexes
	resolution := 10 // Replace with desired resolution
	//polygonVertices := h3.GeoPolygon{Geofence: polygon}
	indexes := h3.Polyfill(polygon, resolution)

	// Convert the H3 indexes to strings
	indexStrings := make([]string, len(indexes))
	for i, index := range indexes {
		indexStrings[i] = h3.ToString(index)
	}

	return indexStrings
}
