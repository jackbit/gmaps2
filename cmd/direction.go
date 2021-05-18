package cmd

import (
	"gmaps2/direction"

	"github.com/spf13/cobra"
)

// directionCMD represents the google direction command.
var directionCMD = &cobra.Command{
	Use:   "direction -o '7.123,144.567' -d '7.456,144.890' -w '7.234,144.678|7.345,144.789'",
	Short: "search direction from google map api",
	RunE: func(cmd *cobra.Command, args []string) (runError error) {
		origin, _ := cmd.Flags().GetString("origin")
		destination, _ := cmd.Flags().GetString("destination")
		waypoints, _ := cmd.Flags().GetString("waypoints")
		avoid, _ := cmd.Flags().GetString("avoid")
		configfile, _ := cmd.Flags().GetString("configfile")
		departuretime, _ := cmd.Flags().GetString("departuretime")

		var params direction.Parameter
		var err error
		if configfile != "" {
			params, err = direction.NewConfig(configfile)
			if err != nil {
				return err
			}
		} else {
			params = direction.NewPoints(origin, destination, avoid, waypoints, departuretime)
		}

		return direction.Search(params)
	},
}

// init function initialises the command options.
func init() {
	directionCMD.Flags().StringP("origin", "o", "", "origin latitude,longitude")
	directionCMD.Flags().StringP("destination", "d", "", "destination latitude,longitude")
	directionCMD.Flags().StringP("waypoints", "w", "", "waypoint addrlatitude,longitude with '|' as separator")
	directionCMD.Flags().StringP("avoid", "a", "", "avoid=tolls|highways|ferries")
	directionCMD.Flags().StringP("departuretime", "t", "", "departure time in unixtimestamp")
	directionCMD.Flags().StringP("configfile", "c", "", "config file")
	mainCMD.AddCommand(directionCMD)
}
