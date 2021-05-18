package geos2

import (
	"fmt"
	"gmaps2/geojson"
	"gmaps2/utils"
	"strings"

	"github.com/golang/geo/s2"
	geo "github.com/toyo/go-latlong"
)

type S2PolylineArgs struct {
	Intersect string
	Contain   string
	File      string
}

// S2Polyline returns polyline information including googles2
// and execute operator like intersection
func S2Polyline(args S2PolylineArgs) error {
	linestring, err := geojson.JsonToLineString(args.File)
	if err != nil {
		return err
	}

	s2polyline := linestring.S2Polyline()
	firstpoint := linestring.MapsLatLng()[0]

	var tokens []string
	for _, cellid := range s2polyline.CellUnionBound() {
		celltoken := cellid.ToToken()
		tokens = append(tokens, celltoken)
	}

	joinedtokens := strings.Join(tokens, ",")
	fmt.Println(fmt.Sprintf("RectBound: %v", linestring.RectBound()))
	fmt.Println(fmt.Sprintf("CellUnionBound Tokens: %v", joinedtokens))

	url := fmt.Sprintf(
		"https://s2.sidewalklabs.com/regioncoverer/?center=%v,%v&zoom=11&cells=%v",
		firstpoint.Lat,
		firstpoint.Lng,
		joinedtokens,
	)

	fmt.Println(fmt.Sprintf("Visual URL: %v", url))

	cap := s2polyline.CapBound()
	fmt.Println(fmt.Sprintf("Center: %v", cap.Center().String()))
	fmt.Println(fmt.Sprintf("Radius: %v", cap.Radius().String()))

	if args.Intersect != "" {
		polylineIntersectPoint(linestring, utils.StringToLatLng(args.Intersect))
	}

	if args.Contain != "" {
		polylineContainCell(linestring, utils.StringToLatLng(args.Contain))
	}

	return nil
}

// polylineIntersectPoint intersects s2.Polyline with point
func polylineIntersectPoint(linestring geo.LineString, point []float64) error {
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(point[0], point[1]))
	isTrue := linestring.IntersectsCell(cell)

	fmt.Println(fmt.Sprintf("Intersected: %v | Cell ID: %v", isTrue, cell.ID().ToToken()))

	return nil
}

// polylineContainCell contains point in s2.Polyline
func polylineContainCell(linestring geo.LineString, point []float64) error {
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(point[0], point[1]))
	isTrue := linestring.ContainsCell(cell)

	fmt.Println(fmt.Sprintf("Contained: %v | Cell ID: %v", isTrue, cell.ID().ToToken()))

	return nil
}
