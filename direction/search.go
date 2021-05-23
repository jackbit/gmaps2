package direction

import (
	"context"
	"errors"
	"fmt"
	"gmaps2/utils"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/envy"
	"googlemaps.github.io/maps"
)

const URL = "https://maps.googleapis.com/maps/api/directions/json"

var GMAP_API_KEY = envy.Get("GMAP_API_KEY", "")

type RoutePolyline struct {
	Route    maps.Route    `json:"route"`
	Polyline []maps.LatLng `json:"polyline"`
}

// Search to search direction from google api
// and save result in output/direction.json
func Search(params Parameter) error {
	routes, err := getRoutes(params)
	if err != nil {
		return err
	}

	var route maps.Route
	var polyline []maps.LatLng
	if len(routes) > 0 {
		route = routes[0]
		polyline, err = maps.DecodePolyline(route.OverviewPolyline.Points)
		if err != nil {
			return err
		}
	}

	err = utils.SaveJSON(
		&RoutePolyline{route, polyline},
		filepath.Join("output", "direction2.json"),
	)
	if err == nil {
		fmt.Println("Result is saved on output/direction.json")
	}

	return err
}

// getRoutes to request parameter to google direction api
// and returns available routes
func getRoutes(params Parameter) ([]maps.Route, error) {
	var routes []maps.Route
	if params.Origin == "" {
		return routes, errors.New("Origin is required")
	}
	if params.Destination == "" {
		return routes, errors.New("Destination is required")
	}
	if params.Destination == "" {
		return routes, errors.New("Destination is required")
	}

	apikey := envy.Get("GMAP_API_KEY", "")
	if apikey == "" {
		return routes, errors.New("GMAP_API_KEY is missing in .env")
	}

	directionrequest := &maps.DirectionsRequest{
		Origin:      params.Origin,
		Destination: params.Destination,
	}

	if params.Avoid != "" {
		directionrequest.Avoid = []maps.Avoid{maps.Avoid(params.Avoid)}
	}
	if params.Waypoints != "" {
		directionrequest.Waypoints = strings.Split(params.Waypoints, "|")
	}
	if params.DepartureTime != "" {
		directionrequest.DepartureTime = params.DepartureTime
	}

	gmap, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		return routes, err
	}

	routes, _, err = gmap.Directions(context.Background(), directionrequest)
	return routes, err
}
