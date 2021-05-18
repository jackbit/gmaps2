package geojson

import (
	"encoding/json"
	"fmt"
	"gmaps2/direction"
	"gmaps2/utils"
	"path/filepath"
)

// MultiLineString represents GeoJSON for MultiLineString
// coordinates will be [ [[lng,lat], [lng,lat]], ..., [[lng,lat], [lng,lat]] ]
type MultiLineString struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// SaveMultiLineString to convert polyline to Multi Line GeoJSON format
// and save in output/multiline.json
func SaveMultiLineString(jsonfile string) error {
	var routePolyline direction.RoutePolyline
	content, err := utils.ReadJSON(jsonfile)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, &routePolyline); err != nil {
		return err
	}

	multilines := &MultiLineString{
		Type: "MultiLineString",
	}

	iterlimit := len(routePolyline.Polyline) - 1
	i := 0
	for i < iterlimit {
		poly1 := routePolyline.Polyline[i]
		poly2 := routePolyline.Polyline[i+1]
		coordinatePoly1 := []float64{poly1.Lng, poly1.Lat}
		coordinatePoly2 := []float64{poly2.Lng, poly2.Lat}
		coordinates := [][]float64{coordinatePoly1, coordinatePoly2}
		multilines.Coordinates = append(multilines.Coordinates, coordinates)
		i += 1
	}

	err = utils.SaveJSON(
		GeoJSON{
			Type:     "Feature",
			Geometry: multilines,
		},
		filepath.Join("output", "multiline.json"),
	)
	if err == nil {
		fmt.Println("Result is saved on output/multiline.json")
	}

	return err
}
