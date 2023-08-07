package main

// func main() {
// 	cords := `[[90.414820003,23.739610355],[90.414775352,23.740601612],[90.414435081,23.74054328],[90.414386471,23.74078633],[90.414308695,23.74174881],[90.414250363,23.742205744],[90.414473969,23.742244632],[90.414473969,23.742468239],[90.414386471,23.742623791],[90.414094811,23.742672401],[90.413958702,23.742711289],[90.413409803,23.743122963],[90.414268494,23.744163513],[90.415977478,23.742969513],[90.419876099,23.739946365],[90.418356038,23.73733105],[90.417114478,23.738238609],[90.416795987,23.738469233],[90.416553554,23.738634587],[90.41570266,23.739169808],[90.415445966,23.7393047],[90.414889796,23.739565783],[90.414820003,23.739610355]]`
// 	var cordinates [][]float64
// 	json.Unmarshal([]byte(cords), &cordinates)

// 	polygon := h3.GeoPolygon{
// 		Geofence: make([]h3.GeoCoord, 0),

// 	}

// 	for i, _ := range cordinates {

// 		latlng := h3.GeoCoord{Latitude: cordinates[i][1], Longitude: cordinates[i][0]}
// 		polygon.Geofence = append(polygon.Geofence, latlng)
// 	}
// 	// fmt.Print(polygon)
// 	hex := h3.Polyfill(polygon, 13)
// 	for _, i := range hex {
// 		fmt.Println(h3.ToString(i))
// 	}

// }

//https://api.mapbox.com/geocoding/v5/mapbox.places/uttora.json?limit=5&proximity=90.41086771188316%2C23.75489867731484&language=en-GB&access_token=pk.eyJ1IjoibWFwYm94IiwiYSI6ImNpejY4NXFhYTA2bTMyeW44ZG0ybXBkMHkifQ.gUGbDOPUN1v1fTs5SeOR4A

// import (
// 	"fmt"

// 	"github.com/mmadfox/go-geojson2h3"
// 	"github.com/tidwall/geojson"
// 	"github.com/uber/h3-go/v3"
// )

// func main() {
// 	resolution := 9
// 	object, err := geojson.Parse(`{
// 		"type": "FeatureCollection",
// 		"features": [
// 		  {
// 			"type": "Feature",
// 			"properties": {},
// 			"geometry": {
// 			  "coordinates": [
// 				[
// 				  [
// 					90.4087946276178,
// 					23.75022861018195
// 				  ],
// 				  [
// 					90.40797923607715,
// 					23.748922515760526
// 				  ],
// 				  [
// 					90.40870879692852,
// 					23.74845114012024
// 				  ],
// 				  [
// 					90.41226004166458,
// 					23.745740697079
// 				  ],
// 				  [
// 					90.41251753373086,
// 					23.749717958283455
// 				  ],
// 				  [
// 					90.4087946276178,
// 					23.75022861018195
// 				  ]
// 				]
// 			  ],
// 			  "type": "Polygon"
// 			}
// 		  }
// 		]
// 	  }`, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	indexes, err := geojson2h3.ToH3(resolution, object)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, index := range indexes {
// 		fmt.Printf("h3index: %s\n", h3.ToString(index))
// 	}
// 	h3.Polyfill()
// }
