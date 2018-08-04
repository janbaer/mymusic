package cmd

import (
	"fmt"

	"github.com/janbaer/mp3db/constants"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the current version of the program",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", constants.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
