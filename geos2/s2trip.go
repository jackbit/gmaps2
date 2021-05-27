package geos2

import (
	"context"
	"errors"
	"fmt"
	"gmaps2/utils"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/golang/geo/s2"
	"github.com/toyo/go-latlong"
	"googlemaps.github.io/maps"
)

type S2Trip struct {
	PickupLocation string
	DropLocation   string
	StartLocation  string
	EndLocation    string
	RediusLevel    int
	TripLevelMax   int
	TripLevelMin   int
	TripCellMax    int
}

// LineString to convert google direction to LineString
func (s S2Trip) LineString() (latlong.LineString, error) {
	var linestring latlong.LineString
	var routes []maps.Route
	var route maps.Route
	var polyline []maps.LatLng
	apikey := envy.Get("GMAP_API_KEY", "")
	if apikey == "" {
		return linestring, errors.New("GMAP_API_KEY is missing in .env")
	}

	directionrequest := &maps.DirectionsRequest{
		Origin:      s.StartLocation,
		Destination: s.EndLocation,
	}

	gmap, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		return linestring, err
	}

	routes, _, err = gmap.Directions(context.Background(), directionrequest)
	if err != nil {
		return linestring, err
	}

	route = routes[0]
	polyline, err = maps.DecodePolyline(route.OverviewPolyline.Points)
	if err != nil {
		return linestring, err
	}

	var multipoint latlong.MultiPoint
	for _, latlng := range polyline {
		latlngpoint := latlong.NewPointFromS2Point(s2.PointFromLatLng(s2.LatLngFromDegrees(latlng.Lat, latlng.Lng)))
		multipoint = append(multipoint, latlngpoint)
	}
	linestring = latlong.LineString{multipoint}
	return linestring, nil

}

// LineStringToS2Covering to convert and display LineString to S2Covering
func (s S2Trip) LineStringToS2Covering(linestring latlong.LineString) s2.CellUnion {
	region := linestring.S2Region()
	coverer := &s2.RegionCoverer{
		MinLevel: s.TripLevelMin,
		MaxLevel: s.TripLevelMax,
		MaxCells: s.TripCellMax,
	}
	covering := coverer.Covering(region)
	var tokens []string
	var parenttokens []string
	parentcheck := map[string]bool{}
	for _, cellid := range covering {
		tokens = append(tokens, cellid.ToToken())
		parenttoken := cellid.Parent(s.RediusLevel).ToToken()
		if !parentcheck[parenttoken] {
			parentcheck[parenttoken] = true
			parenttokens = append(parenttokens, parenttoken)
		}
	}
	fmt.Println("Trip Covering")
	fmt.Println(strings.Join(tokens, ","))
	fmt.Println("Trip Parents:")
	fmt.Println(strings.Join(parenttokens, ","))
	cell := s2.CellFromCellID(covering[0])
	fmt.Println(fmt.Sprintf("Cell (meter): %v", (float64(cell.SizeIJ())/100000.0)*1000))
	return covering
}

func (s S2Trip) RectToS2Covering(pointtext string, tripcovering s2.CellUnion) s2.Rect {
	coverer := &s2.RegionCoverer{
		MinLevel: s.RediusLevel,
		MaxLevel: s.RediusLevel,
		MaxCells: s.TripCellMax,
	}
	pointlatlng := utils.StringToS2LatLng(pointtext)
	rect := s2.RectFromLatLng(pointlatlng)

	covering := coverer.Covering(rect)
	var tokens []string
	for _, cellid := range covering {
		tokens = append(tokens, cellid.ToToken())

		for _, edgeparent := range cellid.EdgeNeighbors() {
			tokens = append(tokens, edgeparent.ToToken())
		}
	}

	fmt.Println(fmt.Sprintf("Cell Tokens: %v", strings.Join(tokens, ",")))
	cell := s2.CellFromCellID(covering[0])
	fmt.Println(fmt.Sprintf("Cell (meter): %v", (float64(cell.SizeIJ())/100000.0)*1000))
	fmt.Println("")

	var nearestcellid s2.CellID
	var nearestdistance float64 = 0.0
	// cellonly := s2.CellFromLatLng(utils.StringToS2LatLng(pointtext))
	for _, tripcellid := range tripcovering {
		// distance := cellonly.DistanceToCell(s2.CellFromCellID(tripcellid))
		distance := s2.CellFromCellID(tripcellid).Distance(s2.PointFromLatLng(pointlatlng))
		kmdistance := 6371.01 * float64(distance.Angle())
		if nearestdistance > kmdistance || nearestdistance == 0.0 {
			nearestdistance = kmdistance
			nearestcellid = tripcellid
		}
	}

	fmt.Println(fmt.Sprintf("Nearest Distance is %v km, cellid: %v, latlng: %v", nearestdistance, nearestcellid.ToToken(), nearestcellid.LatLng().String()))

	return rect
}
