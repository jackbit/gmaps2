package geojson

import (
	"encoding/json"
	"fmt"
	"gmaps2/utils"
	"path/filepath"

	geo "github.com/toyo/go-latlong"
)

// LineToPolygon to convert direction polyline to polygon geojson
// and save to output/linetogon.json
func LineToPolygon(jsonfile string) error {
	linestring, err := JsonToLineString(jsonfile)
	if err != nil {
		return err
	}

	polygon := linestring.Polygon()
	feature := polygon.NewGeoJSONFeature(nil)
	err = utils.SaveJSON(
		feature,
		filepath.Join("output", "linetogon.json"),
	)
	if err == nil {
		fmt.Println("Result is saved on output/linetogon.json")
	}

	return err
}

// JsonToLineString to convert GeoJSON LineString to Geometric LineString
func JsonToLineString(jsonfile string) (geo.LineString, error) {
	var linestring geo.LineString
	content, err := utils.ReadJSON(jsonfile)
	if err != nil {
		return linestring, err
	}
	var geom geo.GeoJSONGeometry
	if err = json.Unmarshal(content, &geom); err != nil {
		return linestring, err
	}

	linestring = geom.Geo().(geo.LineString)

	return linestring, nil
}
