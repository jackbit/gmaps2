package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const appName = "gmaps2"
const appVersion = "0.1.0"

var mainCMD = &cobra.Command{
	Use: "gmaps2",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := mainCMD.Execute(); err != nil {
		fmt.Println(err)
	}
}
