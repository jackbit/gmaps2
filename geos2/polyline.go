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

	cellpoint := utils.StringToLatLng(args.Contain)
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(cellpoint[0], cellpoint[1]))

	// if args.Intersect != "" {
	polylineIntersectPoint(linestring, cell)
	// }
	// if args.Contain != "" {
	polylineContainCell(linestring, cell)
	// }

	s2region := linestring.S2Region()
	coverer := &s2.RegionCoverer{
		MinLevel: 15,
		MaxLevel: 20,
		MaxCells: 50,
	}
	covering := coverer.Covering(s2region)
	fmt.Println("S2Region Recovering:")
	var celltokens []string
	for _, cellID := range covering {
		celltokens = append(celltokens, cellID.ToToken())
	}
	fmt.Println(strings.Join(celltokens, ","))

	fmt.Println(fmt.Sprintf("CellID %v intersects S2Cover ? %v", cell.ID().ToToken(), covering.IntersectsCellID(cell.ID())))
	fmt.Println(fmt.Sprintf("CellID %v contains S2Cover ? %v", cell.ID().ToToken(), covering.ContainsCellID(cell.ID())))

	return nil
}

// polylineIntersectPoint intersects s2.Polyline with point
func polylineIntersectPoint(linestring geo.LineString, cell s2.Cell) error {
	isTrue := linestring.IntersectsCell(cell)

	fmt.Println(fmt.Sprintf("s2Polyline intersects: cell ID %v ? %v", cell.ID().ToToken(), isTrue))

	return nil
}

// polylineContainCell contains point in s2.Polyline
func polylineContainCell(linestring geo.LineString, cell s2.Cell) error {
	isTrue := linestring.ContainsCell(cell)

	fmt.Println(fmt.Sprintf("s2Polyline contains cell ID %v ? %v", cell.ID().ToToken(), isTrue))

	return nil
}
