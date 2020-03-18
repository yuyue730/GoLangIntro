package model

import "encoding/json"

// Describing a car's detail information
type Car struct {
	Name         string
	ImageURL     string
	Size         string
	Fuel         float64
	Transmission string
	Displacement float64
	MaxSpeed     float64
	Acceleration float64
	Price        float64
}

func FromJsonObj(o interface{}) (Car, error) {
	var car Car
	s, err := json.Marshal(o)
	if err != nil {
		return car, err
	}

	err = json.Unmarshal(s, &car)
	return car, err
}
