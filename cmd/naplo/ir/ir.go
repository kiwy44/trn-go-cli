package ir

import (
	"errors"
	"fmt"

	"slices"

	"github.com/spf13/cobra"
)

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

	cmd.Flags().StringVarP(&o.naploFajl, "naplo-fajl", "f", "./naplo.log", "Napló fájl elérési útja")
	cmd.Flags().StringVarP(&o.sulyossag, "suly", "s", "INFO", "Berjegyzés súlyossága. Alapértelmezett: INFO. Lehetőségek: ERROR, WARNING, INFO")
	cmd.Flags().StringVarP(&o.bejegyzes, "bejegyzes", "b", "", "Bejegyzés, üzenet")

	return cmd
}

func run(o *options) error {

	sulyok := []string{"ERROR", "INFO", "WARNING"}

	if o.naploFajl == "" {
		err := errors.New("nincs naplófájl megadva")
		return fmt.Errorf("%w", err)
	}

	if !slices.Contains(sulyok, o.sulyossag) {
		err := errors.New("nem megfelelő súlyosság")
		return fmt.Errorf("%w: %s", err, o.sulyossag)
	}

	if o.bejegyzes == "" {
		err := errors.New("bejegyzés, üzenet nem lehet üres")
		return fmt.Errorf("%w", err)
	}

	pkg.logIro(o.naploFajl, o.sulyossag, o.bejegyzes)
	return nil
}
