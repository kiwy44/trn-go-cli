package naplo

import (
	"github.com/kiwy44/trn-go-cli/cmd/naplo/ir"
	"github.com/kiwy44/trn-go-cli/cmd/naplo/olvas"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var logCmd = &cobra.Command{
		Use:           "naplo",
		Short:         "Napló fájl olvasás és írás",
		Long:          `Parancsok naplóbejegyzéseket tartalmazó fájlok olvasásához és írásához`,
		SilenceErrors: false,
		SilenceUsage:  false,
	}
	// Parancsok
	logCmd.AddCommand(olvas.NewCmd())
	logCmd.AddCommand(ir.NewCmd())
	return logCmd
}
