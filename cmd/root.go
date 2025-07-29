package cmd

import (
	"log"

	"github.com/kiwy44/trn-go-cli/cmd/naplo"
	"github.com/spf13/cobra"
)

var appVersion bool

var rootCmd = &cobra.Command{
	Use:           "trn-go-cli",
	Short:         "Cli hasznos dolgoknak",
	Long:          `Parancssori alkalmazás, hogy a hasznos dolgokat egyben kezelhessem`,
	SilenceErrors: false,
	SilenceUsage:  false,
	Run: func(cmd *cobra.Command, args []string) {
		if appVersion {
			log.Printf("Verzió szám: %s\n", "1.0.0")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// További parancsok, lehetőségek
func init() {
	rootCmd.Flags().BoolVarP(&appVersion, "verzio", "v", false, "Mutasd a verzió számot")
	rootCmd.AddCommand(naplo.NewCmd())
}
