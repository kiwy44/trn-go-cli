package olvas

import (
	"errors"
	"fmt"

	"slices"

	"github.com/kiwy44/trn-go-cli/pkg"
	"github.com/spf13/cobra"
)

// Paraméterek
type options struct {
	naploFajl string
	sulyossag string
}

func NewCmd() *cobra.Command {
	o := &options{}

	cmd := &cobra.Command{
		Use:   "olvas",
		Short: "Naplóbejegyzések olvasása",
		Long:  `Bizonyos sulyosságú bejegyzések olvasása napló fájlból`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(o)
		},
	}

	// Paraméterek beállítása
	cmd.Flags().StringVarP(&o.naploFajl, "naplo-fajl", "f", "./naplo.log", "Napló fájl elérési útja")
	cmd.Flags().StringVarP(&o.sulyossag, "suly", "s", "INFO", "Berjegyzés súlyossága. Alapértelmezett: INFO. Lehetőségek: ERROR, WARN, INFO")

	return cmd
}

func run(o *options) error {
	// Súlyosságok
	sulyok := []string{"ERROR", "INFO", "WARN"}

	// Naplófájl paraméter ellenőrzése
	if o.naploFajl == "" {
		err := errors.New("nincs naplófájl megadva")
		return fmt.Errorf("%w", err)
	}

	// Súlyosság paraméter ellenőrzése
	if !slices.Contains(sulyok, o.sulyossag) {
		err := errors.New("nem megfelelő súlyosság")
		return fmt.Errorf("%w: %s", err, o.sulyossag)
	}

	// Napló olvasás függvény meghívása
	pkg.NaploOlvaso(o.naploFajl, o.sulyossag)
	return nil
}
