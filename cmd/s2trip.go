package cmd

import (
	"errors"
	"fmt"
	"gmaps2/geos2"

	"github.com/spf13/cobra"
)

// s2tripCMD represents the driver trip with customer pickup and drop locations.
var s2tripCMD = &cobra.Command{
	Use:   "s2trip -o 'lat,lng' -d 'lat,lng'",
	Short: "s2trip to demo driver trip with passenger pickup and drop location in s2geometry",
	RunE: func(cmd *cobra.Command, args []string) (runError error) {
		pickuplocation, _ := cmd.Flags().GetString("pickup")
		droplocation, _ := cmd.Flags().GetString("drop")
		startlocation, _ := cmd.Flags().GetString("start")
		endlocation, _ := cmd.Flags().GetString("end")

		radius, _ := cmd.Flags().GetInt("radius")
		tripmaxlevel, _ := cmd.Flags().GetInt("tripmaxlevel")
		tripminlevel, _ := cmd.Flags().GetInt("tripminlevel")
		tripmaxcell, _ := cmd.Flags().GetInt("tripmaxcell")

		if pickuplocation == "" || droplocation == "" {
			return errors.New("Pickup and drop locations are required")
		}

		if startlocation == "" || endlocation == "" {
			return errors.New("Start and end locations are required")
		}

		s2trip := geos2.S2Trip{
			PickupLocation: pickuplocation,
			DropLocation:   droplocation,
			StartLocation:  startlocation,
			EndLocation:    endlocation,
			RediusLevel:    radius,
			TripLevelMax:   tripmaxlevel,
			TripLevelMin:   tripminlevel,
			TripCellMax:    tripmaxcell,
		}
		linestring, err := s2trip.LineString()
		if err != nil {
			return err
		}

		covering := s2trip.LineStringToS2Covering(linestring)
		fmt.Println("")
		fmt.Println("Pickup Location")
		s2trip.RectToS2Covering(s2trip.PickupLocation, covering)
		fmt.Println("")
		fmt.Println("Drop Location")
		s2trip.RectToS2Covering(s2trip.DropLocation, covering)
		return nil
	},
}

// init function initialises the command options.
func init() {
	s2tripCMD.Flags().StringP("pickup", "p", "", "pickup point (lat,lng) in float64")
	s2tripCMD.Flags().StringP("drop", "d", "", "drop point (lat,lng) in float64")
	s2tripCMD.Flags().StringP("start", "s", "", "start trip point (lat,lng) in float64")
	s2tripCMD.Flags().StringP("end", "e", "", "end trip point (lat,lng) in float64")
	s2tripCMD.Flags().IntP("radius", "r", 13, "cell level as radius for customer (0-30), default: 13 (1.2 km)")
	s2tripCMD.Flags().IntP("tripmaxlevel", "m", 16, "max level recovering (0-30), default: 16 (153m) ")
	s2tripCMD.Flags().IntP("tripminlevel", "n", 16, "min level recovering (0-30), default: 16 (153m) ")
	s2tripCMD.Flags().IntP("tripmaxcell", "c", 50, "max cells recovering, default: 50")
	mainCMD.AddCommand(s2tripCMD)
}
