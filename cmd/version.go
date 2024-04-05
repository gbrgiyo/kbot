package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appVersion string = "Version"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `The version command prints the version number of kbot.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
