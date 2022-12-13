package contract

type Point struct {
	Lat float64
	Lon float64
}

type InputData struct {
	Lat        string
	Lon        string
	Resolution string
}

func GetLatLon() *Point {
	return &Point{}
}

var Input *InputData

func GetInputPayload() *InputData {
	if Input == nil {
		Input = &InputData{}
		return Input
	}
	return Input
}
