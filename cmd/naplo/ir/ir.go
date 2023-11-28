package ir

import (
	"errors"
	"fmt"

	"slices"

	"github.com/cloudsteak/trn-go-cli/pkg"
	"github.com/spf13/cobra"
)

// Paraméterek
type options struct {
	naploFajl string
	sulyossag string
	bejegyzes string
}

func NewCmd() *cobra.Command {
	o := &options{}

	cmd := &cobra.Command{
		Use:   "ir",
		Short: "Naplóbejegyzés írása",
		Long:  `Bármilyen üzenet írása naplófájlba`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(o)
		},
	}

	// Paraméterek beállítása
	cmd.Flags().StringVarP(&o.naploFajl, "naplo-fajl", "f", "naplo.log", "Napló fájl elérési útja")
	cmd.Flags().StringVarP(&o.sulyossag, "suly", "s", "INFO", "Berjegyzés súlyossága. Alapértelmezett: INFO. Lehetőségek: ERROR, WARN, INFO")
	cmd.Flags().StringVarP(&o.bejegyzes, "bejegyzes", "b", "", "Bejegyzés, üzenet")

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

	// Bejegyzés ellenőrzése
	if o.bejegyzes == "" {
		err := errors.New("bejegyzés, üzenet nem lehet üres")
		return fmt.Errorf("%w", err)
	}

	// Napló írás függvény meghívása
	pkg.NaploIro(o.naploFajl, o.sulyossag, o.bejegyzes)
	return nil
}
