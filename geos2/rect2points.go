package geos2

import (
	"fmt"
	"gmaps2/geojson"
	"gmaps2/utils"
	"strings"

	"github.com/golang/geo/s2"
)

type RectPoints struct {
	Origin           string
	Destination      string
	Level            int
	CoveringMaxLevel int
	CoveringMinLevel int
	CoveringMaxCells int
	File             string
}

func (r RectPoints) coveringPoint(pointtext string) s2.Rect {
	coverer := &s2.RegionCoverer{
		MinLevel: r.Level,
		MaxLevel: r.Level,
		MaxCells: r.CoveringMaxCells,
	}
	rect := s2.RectFromLatLng(
		utils.StringToS2LatLng(pointtext),
	)
	covering := coverer.Covering(rect)
	var tokens []string
	for _, cellid := range covering {
		tokens = append(tokens, cellid.ToToken())
	}

	fmt.Println(fmt.Sprintf("Cell Tokens: %v", strings.Join(tokens, ",")))

	cell := s2.CellFromCellID(covering[0])
	fmt.Println(fmt.Sprintf("Cell (meter): %v", (float64(cell.SizeIJ())/100000.0)*1000))
	fmt.Println("")
	return rect
}

// Rect2Points to create rectal region and discover by 2 points
func Rect2Points(rectpoints RectPoints) error {
	fmt.Println("Origin:")
	originrect := rectpoints.coveringPoint(rectpoints.Origin)

	fmt.Println("Destination:")
	destrect := rectpoints.coveringPoint(rectpoints.Destination)

	triprects := originrect.Union(destrect)

	tripcoverer := &s2.RegionCoverer{
		MinLevel: rectpoints.CoveringMinLevel,
		MaxLevel: rectpoints.CoveringMaxLevel,
		MaxCells: rectpoints.CoveringMaxCells,
	}

	tripcovering := tripcoverer.Covering(triprects)
	var tokens []string
	for _, cellid := range tripcovering {
		tokens = append(tokens, cellid.ToToken())
	}
	fmt.Println("")
	fmt.Println("Cells from Origin to Destination:")
	fmt.Println(strings.Join(tokens, ","))
	cell := s2.CellFromCellID(tripcovering[0])
	fmt.Println(fmt.Sprintf("Cell (meter): %v", (float64(cell.SizeIJ())/100000.0)*1000))

	if rectpoints.File != "" {
		linestring, err := geojson.JsonToLineString(rectpoints.File)
		if err != nil {
			return err
		}

		directionregion := linestring.S2Region()
		directioncoverer := &s2.RegionCoverer{
			MinLevel: rectpoints.CoveringMinLevel,
			MaxLevel: rectpoints.CoveringMaxLevel,
			MaxCells: rectpoints.CoveringMaxCells,
		}
		directioncovering := directioncoverer.Covering(directionregion)
		var directiontokens []string
		for _, cellid := range directioncovering {
			directiontokens = append(directiontokens, cellid.ToToken())
		}
		fmt.Println("Direction Covering")
		fmt.Println(strings.Join(directiontokens, ","))
	}

	return nil
}
