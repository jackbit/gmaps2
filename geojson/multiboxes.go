package geojson

import (
	"encoding/json"
	"fmt"
	"gmaps2/utils"
	"path/filepath"
)

// BoundingBox represents GeoJSON for Polygon
// coordinates will be [[lng,lat], [lng,lat]]
type BoundingBox struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// SavePolygonPoint to return coordinates as bounding box
// and save to output/multiboxes.json
func SaveMultiBoxes(jsonfile string) error {
	content, err := utils.ReadJSON(jsonfile)
	if err != nil {
		return err
	}
	var params struct {
		Origin      string  `json:"origin"`
		Destination string  `json:"destination"`
		Unit        string  `json:"unit"`
		Radius      float64 `json:"radius"`
	}
	if err = json.Unmarshal(content, &params); err != nil {
		return err
	}

	var geos []GeoJSON
	pointList := [][]string{{"origin", params.Origin}, {"destination", params.Destination}}
	for _, pointItem := range pointList {
		point := utils.StringToLatLng(pointItem[1])
		polygon := utils.PointToPolygon(point, params.Radius, params.Unit)
		geos = append(
			geos,
			GeoJSON{
				Type: "Feature",
				Geometry: BoundingBox{
					Name:        pointItem[0],
					Type:        "Polygon",
					Coordinates: polygon,
				},
			},
		)
	}

	err = utils.SaveJSON(
		geos,
		filepath.Join("output", "multiboxes.json"),
	)
	if err == nil {
		fmt.Println("Result is saved on output/multiboxes.json")
	}

	return err
}
