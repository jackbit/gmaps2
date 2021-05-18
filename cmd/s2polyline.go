package cmd

import (
	"errors"
	"gmaps2/geos2"

	"github.com/spf13/cobra"
)

// geos2CMD represents the geojson command.
var geos2CMD = &cobra.Command{
	Use:   "geos2 -t 'polyline' -f 'input/linestring.json'",
	Short: "geos2 geojson to s2 and run operation",
	RunE: func(cmd *cobra.Command, args []string) (runError error) {
		file, _ := cmd.Flags().GetString("file")
		typegeo, _ := cmd.Flags().GetString("type")

		if file != "" {
			var err error
			switch typegeo {
			case "polyline":
				err = geos2.S2Polyline(file)
			}

			return err
		}

		return errors.New("Missing file")
	},
}

// init function initialises the command options.
func init() {
	geos2CMD.Flags().StringP("file", "f", "", "json source for geo convertion")
	geos2CMD.Flags().StringP("type", "t", "polyline", "type=polyline|polygon")
	mainCMD.AddCommand(geos2CMD)
}
