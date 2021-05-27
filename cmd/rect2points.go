package cmd

import (
	"errors"
	"gmaps2/geos2"

	"github.com/spf13/cobra"
)

// s2rectsCMD represents the geojson command.
var s2rectsCMD = &cobra.Command{
	Use:   "s2rects -o 'lat,lng' -d 'lat,lng'",
	Short: "s2rects create rectal region between original and destination points",
	RunE: func(cmd *cobra.Command, args []string) (runError error) {
		origin, _ := cmd.Flags().GetString("origin")
		destination, _ := cmd.Flags().GetString("destination")
		level, _ := cmd.Flags().GetInt("level")
		maxlevel, _ := cmd.Flags().GetInt("maxlevel")
		minlevel, _ := cmd.Flags().GetInt("minlevel")
		maxcell, _ := cmd.Flags().GetInt("maxcell")
		file, _ := cmd.Flags().GetString("file")

		if origin == "" || destination == "" {
			return errors.New("Origin and destination are required to create rectal region")
		}

		return geos2.Rect2Points(geos2.RectPoints{
			Origin:           origin,
			Destination:      destination,
			Level:            level,
			CoveringMaxLevel: maxlevel,
			CoveringMinLevel: minlevel,
			CoveringMaxCells: maxcell,
			File:             file,
		})
	},
}

// init function initialises the command options.
func init() {
	s2rectsCMD.Flags().StringP("origin", "o", "", "origin point (lat,lng) in float64")
	s2rectsCMD.Flags().StringP("destination", "d", "", "origin point (lat,lng) in float64")
	s2rectsCMD.Flags().IntP("level", "l", 13, "origin or destination level per cell (0-20), default: 13 (1.2 km)")
	s2rectsCMD.Flags().IntP("maxlevel", "m", 14, "max level recovering (0-30), default: 14 (613m) ")
	s2rectsCMD.Flags().IntP("minlevel", "n", 14, "min level recovering (0-30), default: 14 (613m) ")
	s2rectsCMD.Flags().IntP("maxcell", "c", 500, "max cells recovering, default: 500")
	s2rectsCMD.Flags().StringP("file", "f", "input/linestring2.json", "polyline file")
	mainCMD.AddCommand(s2rectsCMD)
}
