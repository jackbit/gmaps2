package geos2

import (
	"fmt"
	"gmaps2/geojson"
	"gmaps2/utils"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/paulsmith/gogeos/geos"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/planar"
	"googlemaps.github.io/maps"
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

// linesToWKT to convert linestring coordinates to WKT format
func lineToWKT(coords []maps.LatLng) string {
	wkts := []string{}
	for _, coord := range coords {
		wkts = append(wkts, fmt.Sprintf("%v %v", coord.Lng, coord.Lat))
	}

	return fmt.Sprintf("LINESTRING (%v)", strings.Join(wkts, ", "))
}

// bufferToPolygons to convert and save poline WKT to PolygonGeoJSON
func bufferToPolygons(buffstring string) error {
	xpressionpolygon := regexp.MustCompile("POLYGON|(\\(\\()|(\\)\\))")
	buffstring = xpressionpolygon.ReplaceAllString(buffstring, "")
	textcoords := strings.Split(buffstring, ",")

	geoline := &geojson.LineString{
		Type: "LineString",
	}
	for _, textcoord := range textcoords {
		coords := strings.Split(textcoord, " ")
		latitude, _ := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
		longitude, _ := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
		geoline.Coordinates = append(geoline.Coordinates, []float64{longitude, latitude})
	}

	err := utils.SaveJSON(
		geojson.GeoJSON{
			Type:     "Feature",
			Geometry: geoline,
		},
		filepath.Join("output", "wktpolygon.json"),
	)
	if err == nil {
		fmt.Println("Result is saved on output/wktpolygon.json")
	}
	return nil
}
