package cmd

import (
	"errors"
	"gmaps2/geos2"

	"github.com/spf13/cobra"
)

// geosCMD represents the geojson command.
var geosCMD = &cobra.Command{
	Use:   "geos -f 'input/linestring.json'",
	Short: "geos convert linestring to polygon by GEOS",
	RunE: func(cmd *cobra.Command, args []string) (runError error) {
		file, _ := cmd.Flags().GetString("file")
		sample, _ := cmd.Flags().GetString("sample")

		if file != "" {
			switch sample {
			case "1":
				return geos2.PolygonGEOS1(geos2.S2PolylineArgs{
					File: file,
				})
			case "2":
				return geos2.PolygonGEOS2(geos2.S2PolylineArgs{
					File: file,
				})
			case "3":
				return geos2.PolygonGEOS3(geos2.S2PolylineArgs{
					File: file,
				})
			}
			return nil
		}

		return errors.New("Missing file")
	},
}

// init function initialises the command options.
func init() {
	geosCMD.Flags().StringP("file", "f", "", "json source for geo convertion")
	geosCMD.Flags().StringP("sample", "s", "", "example operation in integer")
	mainCMD.AddCommand(geosCMD)
}
