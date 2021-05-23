package geojson

import (
	"fmt"
	"gmaps2/utils"
	"path/filepath"
)

type PolygonPoint [][][]float64

// MultiPolygonGeo represents GeoJSON for MultiPolygon
type MultiPolygonGeo struct {
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	Coordinates []PolygonPoint `json:"coordinates"`
}

// SaveMultiPolygon to convert direction polyline to multiple polygons
// and save to output/multipolygon.json
func SaveMultiPolygon(radius float64, unit, jsonfile string) error {
	linestring, err := JsonToLineString(jsonfile)
	if err != nil {
		return err
	}

	var polygons []PolygonPoint
	for _, latlng := range linestring.MapsLatLng() {
		polygon := utils.PointToPolygon([]float64{latlng.Lat, latlng.Lng}, radius, unit)
		polygons = append(polygons, polygon)
	}

	polygons = append(polygons, polygons[0])

	err = utils.SaveJSON(
		GeoJSON{
			Type: "Feature",
			Geometry: MultiPolygonGeo{
				Name:        "linepolygon",
				Type:        "MultiPolygon",
				Coordinates: polygons,
			},
		},
		filepath.Join("output", "multipolygon.json"),
	)

	if err == nil {
		fmt.Println("Result is saved on output/multipolygon.json")
	}

	// var disolves []GeoJSON
	// for _, poly := range polygons {
	// 	disolves = append(
	// 		disolves,
	// 		GeoJSON{
	// 			Type: "Feature",
	// 			Geometry: map[string]interface{}{
	// 				"coordinates": poly,
	// 			},
	// 		},
	// 	)
	// }

	err = utils.SaveJSON(
		map[string]interface{}{
			"type": "FeatureCollection",
			"features": GeoJSON{
				Type: "Feature",
				Geometry: map[string]interface{}{
					"type":        "Polygon",
					"coordinates": polygons,
				},
			},
		},
		filepath.Join("output", "disolves.json"),
	)

	return err
}
