package cmd

import (
	"errors"
	"gmaps2/geojson"

	"github.com/spf13/cobra"
)

// geojsonCMD represents the geojson command.
var geojsonCMD = &cobra.Command{
	Use:   "geojson -t 'line' -f 'output/direction.json'",
	Short: "geojson convert google direction result to geojson format",
	RunE: func(cmd *cobra.Command, args []string) (runError error) {
		file, _ := cmd.Flags().GetString("file")
		typegeo, _ := cmd.Flags().GetString("type")
		radius, _ := cmd.Flags().GetFloat64("radius")
		unit, _ := cmd.Flags().GetString("unit")

		if file != "" {
			var err error
			switch typegeo {
			case "line":
				err = geojson.SaveLineString(file)
			case "multilines":
				err = geojson.SaveMultiLineString(file)
			case "multiboxes":
				err = geojson.SaveMultiBoxes(file)
			case "linetogon":
				err = geojson.LineToPolygon(file)
			case "multipolygon":
				err = geojson.SaveMultiPolygon(radius, unit, file)
			}

			return err
		}

		return errors.New("Missing file")
	},
}

// init function initialises the command options.
func init() {
	geojsonCMD.Flags().StringP("file", "f", "", "json source for geo convertion")
	geojsonCMD.Flags().Float64P("radius", "r", 0.0, "radius in float64")
	geojsonCMD.Flags().StringP("unit", "u", "km", "unit radius (km, mi, nm)")
	geojsonCMD.Flags().StringP("type", "t", "line", "type=line|multilines|multiboxes|linetogon|multipolygon")
	mainCMD.AddCommand(geojsonCMD)
}
