package cmd

import (
	"errors"
	"gmaps2/geos2"

	"github.com/spf13/cobra"
)

// s2polylineCMD represents the geojson command.
var s2polylineCMD = &cobra.Command{
	Use:   "s2polyline -f 'input/linestring.json'",
	Short: "s2polyline convert geojson linestring to s2polyline",
	RunE: func(cmd *cobra.Command, args []string) (runError error) {
		file, _ := cmd.Flags().GetString("file")
		intersect, _ := cmd.Flags().GetString("intersect")
		contain, _ := cmd.Flags().GetString("contain")

		if file != "" {
			return geos2.S2Polyline(geos2.S2PolylineArgs{
				File:      file,
				Intersect: intersect,
				Contain:   contain,
			})
		}

		return errors.New("Missing file")
	},
}

// init function initialises the command options.
func init() {
	s2polylineCMD.Flags().StringP("file", "f", "", "json source for geo convertion")
	s2polylineCMD.Flags().StringP("intersect", "i", "", "intersect s2.cell from lat,lng")
	s2polylineCMD.Flags().StringP("contain", "c", "", "contain s2.cell from lat,lng")
	mainCMD.AddCommand(s2polylineCMD)
}
