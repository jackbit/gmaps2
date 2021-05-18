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
			}

			return err
		}

		return errors.New("Missing file")
	},
}

// init function initialises the command options.
func init() {
	geojsonCMD.Flags().StringP("file", "f", "", "json source for geo convertion")
	geojsonCMD.Flags().StringP("type", "t", "line", "type=line|multilines|multiboxes|linetogon")
	mainCMD.AddCommand(geojsonCMD)
}
