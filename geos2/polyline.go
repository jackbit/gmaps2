package geos2

import (
	"fmt"
	"gmaps2/geojson"
)

func S2Polyline(jsonfile string) error {
	linestring, err := geojson.JsonToLineString(jsonfile)
	if err != nil {
		return err
	}

	s2polyline := linestring.S2Polyline()
	fmt.Println("Cell ID Tokens: ")

	for _, cellid := range s2polyline.CellUnionBound() {
		fmt.Println(cellid.ToToken())
	}

	return nil
}
