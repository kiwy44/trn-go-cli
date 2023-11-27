package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var appVersion bool

var rootCmd = &cobra.Command{
	Use:           "cool-cli",
	Short:         "Cli csak úgy",
	Long:          `Parancssori alkalmazás, csak úgy`,
	SilenceErrors: false,
	SilenceUsage:  false,
	Run: func(cmd *cobra.Command, args []string) {
		if appVersion {
			log.Printf("Version number: %s\n", "1.0.0")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&appVersion, "version", "v", false, "Get version number")
}
