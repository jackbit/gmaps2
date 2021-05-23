package geos2

import (
	"fmt"
	"gmaps2/geojson"
	"gmaps2/utils"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/geo/s2"
	"github.com/paulsmith/gogeos/geos"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/planar"
	"googlemaps.github.io/maps"

	geo "github.com/toyo/go-latlong"
)

// PolygonGEOS1 returns polyline to polygon by GEOOS
func PolygonGEOS1(args S2PolylineArgs) error {
	linestring, err := geojson.JsonToLineString(args.File)
	if err != nil {
		return err
	}

	linecoords := linestring.MapsLatLng()
	line, err := geos.FromWKT(lineToWKT(linecoords))
	if err != nil {
		return err
	}

	buf, err := line.Buffer(0.5)
	if err != nil {
		return err
	}

	return bufferToPolygons(buf.String())
}

// PolygonGEOS2 returns polyline to polygon by GEOOS
func PolygonGEOS2(args S2PolylineArgs) error {
	linestring, err := geojson.JsonToLineString(args.File)
	if err != nil {
		return err
	}

	linecoords := linestring.MapsLatLng()
	geometry, _ := wkt.UnmarshalString(lineToWKT(linecoords))
	fmt.Println("")
	fmt.Println(wkt.MarshalString(geometry))

	bound := geometry.Bound()
	polygon := bound.ToPolygon()
	fmt.Println("")
	fmt.Println(wkt.MarshalString(polygon))

	strategy := planar.NormalStrategy()
	convexhull, err := strategy.ConvexHull(geometry)
	if err != nil {
		return err
	}
	fmt.Println("")
	fmt.Println("CONVEXHULL")
	fmt.Println(wkt.MarshalString(convexhull))

	buffer := strategy.Buffer(geometry, 0.001, 0)
	fmt.Println("")
	fmt.Println("BUFFER radius=0.001, quadsegs=0")
	bufferstr := wkt.MarshalString(buffer)
	fmt.Println(bufferstr)
	bufferToPolygons(bufferstr)

	return nil
}

// PolygonGEOS3 returns S2Polygon from GEOSPolygon
func PolygonGEOS3(args S2PolylineArgs) error {
	linestring, err := geojson.JsonToLineString(args.File)
	if err != nil {
		return err
	}

	linecoords := linestring.MapsLatLng()
	geometry, _ := wkt.UnmarshalString(lineToWKT(linecoords))
	strategy := planar.NormalStrategy()
	buffer := strategy.Buffer(geometry, 0.001, 0)
	bufferstr := wkt.MarshalString(buffer)
	bufferToS2Polygon(bufferstr)

	return nil
}

// linesToWKT to convert linestring coordinates to WKT format
func lineToWKT(coords []maps.LatLng) string {
	wkts := []string{}
	for _, coord := range coords {
		wkts = append(wkts, fmt.Sprintf("%v %v", coord.Lng, coord.Lat))
	}

	return fmt.Sprintf("LINESTRING (%v)", strings.Join(wkts, ", "))
}

// convertWKTPolygon to convert WKTPolygon to Coordinates
func convertWKTPolygon(buffstring string) [][]float64 {
	xpressionpolygon := regexp.MustCompile("POLYGON|(\\(\\()|(\\)\\))")
	buffstring = xpressionpolygon.ReplaceAllString(buffstring, "")
	textcoords := strings.Split(buffstring, ",")

	var coordinates [][]float64
	for _, textcoord := range textcoords {
		coords := strings.Split(textcoord, " ")
		latitude, _ := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
		longitude, _ := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
		coordinates = append(coordinates, []float64{longitude, latitude})
	}
	return coordinates
}

// bufferToPolygons to convert and save poline WKT to PolygonGeoJSON
func bufferToPolygons(buffstring string) error {
	geopolygon := struct {
		Type        string      `json:"type"`
		Coordinates [][]float64 `json:"coordinates"`
	}{
		Type:        "Polygon",
		Coordinates: convertWKTPolygon(buffstring),
	}

	err := utils.SaveJSON(
		geojson.GeoJSON{
			Type:     "Feature",
			Geometry: geopolygon,
		},
		filepath.Join("output", "wktpolygon.json"),
	)
	if err == nil {
		fmt.Println("Result is saved on output/wktpolygon.json")
	}
	return nil
}

func wktToPoints(buffstring string) geo.Polygon {
	xpressionpolygon := regexp.MustCompile("POLYGON|(\\(\\()|(\\)\\))")
	buffstring = xpressionpolygon.ReplaceAllString(buffstring, "")
	textcoords := strings.Split(buffstring, ",")

	var multipoint geo.MultiPoint
	for _, textcoord := range textcoords {
		coords := strings.Split(textcoord, " ")
		latitude, _ := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
		longitude, _ := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
		point := geo.NewPoint(
			geo.NewAngle(
				latitude,
				utils.FloatPrecision(latitude),
			),
			geo.NewAngle(
				longitude,
				utils.FloatPrecision(longitude),
			),
			nil,
		)
		multipoint = append(multipoint, point)
	}
	return geo.Polygon{geo.LineString{multipoint}}
}

// bufferToS2Polygon to convert polygon WKT to S2Polygon
func bufferToS2Polygon(buffstring string) error {
	polygon := wktToPoints(buffstring)
	s2region := polygon.S2Region()
	coverer := &s2.RegionCoverer{
		MinLevel: 15,
		MaxLevel: 20,
		MaxCells: 50,
	}
	covering := coverer.Covering(s2region)
	fmt.Println("S2Region Recovering:")
	for _, cellID := range covering {
		fmt.Println(cellID.ToToken())
	}
	return nil
}
