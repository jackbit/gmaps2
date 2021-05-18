package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	KM_IN_MI           = 0.621371192
	KM_IN_NM           = 0.539957
	DEGREES_PER_RADIAN = 57.2957795
	KM_EARTH_RADIUS    = 6371.0
	MATH_PI            = 3.141592653589793
)

// ReadJSON to read file bytes from a filepath
func ReadJSON(file string) ([]byte, error) {
	var content []byte
	jsonstat, err := os.Stat(file)
	if err != nil {
		return content, err
	}
	if jsonstat.IsDir() {
		return content, errors.New(fmt.Sprintf("ERROR: %s should be a file\n", file))
	}

	f, err := os.Open(file)
	if err != nil {
		return content, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

// SaveJSON to save data bytes to a filepath
func SaveJSON(v interface{}, outputfile string) error {
	jsondata, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outputfile, jsondata, 0666)
}

func StringToLatLng(point string) []float64 {
	points := strings.Split(point, ",")
	lat, _ := strconv.ParseFloat(strings.TrimSpace(points[0]), 64)
	lng, _ := strconv.ParseFloat(strings.TrimSpace(points[1]), 64)
	return []float64{lat, lng}
}

// PointToBound to returns coordinates of the southwest and northeast corners of a box
// with the given point at its center. The radius is the shortest distance
// from the center point to any side of the box (the length of each side
// is twice the radius).
func PointToBound(point []float64, radius float64, unit string) (float64, float64, float64, float64) {
	lat := point[0]
	lng := point[1]

	if unit == "" {
		unit = "km"
	}

	latDistance := LatitudeDegreeDistance(unit)
	latRadius := radius / latDistance
	radian := lat * (MATH_PI / 180)
	lngRadius := radius / (latDistance * math.Cos(radian))

	latSW := lat - latRadius
	lngSW := lng - lngRadius

	latNE := lat + latRadius
	lngNE := lng + lngRadius

	return latSW, lngSW, latNE, lngNE
}

// LatitudeDegreeDistance to return latitude distance by unit metric
// available unit is km (kilometer), mi (miles), nm (nautical miles)
func LatitudeDegreeDistance(unit string) float64 {
	var eartradius float64
	switch unit {
	case "km":
		eartradius = KM_EARTH_RADIUS
	case "mi":
		eartradius = KM_EARTH_RADIUS * KM_IN_MI
	case "nm":
		eartradius = KM_EARTH_RADIUS * KM_IN_NM
	}
	return (2 * MATH_PI * eartradius / 360)
}

// PointToPolygon to convert point (lat,lng) to polygon NE-SW
func PointToPolygon(point []float64, radius float64, unit string) [][][]float64 {
	latSW, lngSW, latNE, lngNE := PointToBound(point, radius, unit)
	box := [][]float64{
		{lngNE, latSW},
		{lngNE, latSW},
		{lngNE, latNE},
		{lngSW, latNE},
		{lngSW, latSW},
	}
	return [][][]float64{box}
}

// FloatPrecission to return precission decimal
// example
func FloatPrecision(degree float64) float64 {
	precs := strings.Split(fmt.Sprintf("%v", degree), ".")
	if len(precs) == 1 {
		return 1.0
	}
	return 1.0 / math.Pow10(len(precs[1]))
}
