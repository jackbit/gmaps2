package direction

import (
	"encoding/json"
	"gmaps2/utils"
)

type Parameter struct {
	Origin        string
	Destination   string
	Waypoints     string
	Avoid         string
	DepartureTime string
}

// NewConfig to initialize Direction struct by config file
// and return Parameter object or error
func NewConfig(configfile string) (Parameter, error) {
	var params Parameter
	content, err := utils.ReadJSON(configfile)
	if err != nil {
		return params, err
	}

	if err = json.Unmarshal(content, &params); err != nil {
		return params, err
	}

	return params, nil
}

// New to initialize Direction struct by origin, destination and waypoints
// and return Parameter object
func NewPoints(origin, destination, avoid, waypoints, departuretime string) Parameter {
	return Parameter{
		Origin:        origin,
		Destination:   destination,
		Avoid:         avoid,
		Waypoints:     waypoints,
		DepartureTime: departuretime,
	}
}
