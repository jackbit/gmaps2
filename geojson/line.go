package geojson

import (
	"encoding/json"
	"fmt"
	"gmaps2/direction"
	"gmaps2/utils"
	"path/filepath"
)

// LineString represents GeoJSON for LineString
// coordinates will be [[lng,lat], [lng,lat]]
type LineString struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

// SaveLineString to convert polyline to Multi Line GeoJSON format
// and save in output/linestring.json
func SaveLineString(jsonfile string) error {
	var routePolyline direction.RoutePolyline
	content, err := utils.ReadJSON(jsonfile)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, &routePolyline); err != nil {
		return err
	}

	linestring := &LineString{
		Type: "LineString",
	}

	for _, latlng := range routePolyline.Polyline {
		linestring.Coordinates = append(
			linestring.Coordinates,
			[]float64{latlng.Lng, latlng.Lat},
		)
	}

	err = utils.SaveJSON(
		GeoJSON{
			Type:     "Feature",
			Geometry: linestring,
		},
		filepath.Join("output", "linestring.json"),
	)
	if err == nil {
		fmt.Println("Result is saved on output/linestring.json")
	}

	return err
}
