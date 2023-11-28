# Go - Parancssori alkalmazás (Cobra-Cli)

Parancssori alkalmazást készítünk, amellyel nagyon sok lehetőségünk van olyan funkciókat egyben használni, amelyeket máskülönben sok-sok külön eszközzel bvennénk igénybe. Ehhez a [Cobra](https://github.com/spf13/cobra)-t használjuk. Olyan alkalmazást hozunk létre most, amely bizonyos súlyosságú hibákat képes fájlba írni.

## Előfeltételek

Az alábbi helyen megtalálod az előkészületeket a Go-ban való fejlesztéshez: https://github.com/cloudsteak/golang-basics

## Projekt létrehozás

1. Nyiss egy parancssort (CMD)
2. Navigálj abba a mappába ahol a kódod fogod tárolni a helyi gépeden.
3. Hozd létre a projekted mappáját. Pl.: `trn-go-cli`

```bash
mkdir trn-go-cli
```

4. Lépj be a mappába

```bash
cd trn-go-cli
```

5. Készítsd el a projekted alap struktúráját (`github.com/cloudsteak/trn-go-cli` helyett használd a saját kódodhoz tartozó github elérhetőséget)

```bash
go mod init github.com/cloudsteak/trn-go-cli
```

6. indítsd el innen a Visual Studio Code-ot.

```bash
code .
```

## Go alkalmazás

1. Hozzuk létre at alap mappa struktúrát

```go
mkdir -p cmd/naplo/ir
mkdir -p cmd/naplo/olvas
mkdir -p pkg
```

2. Telepítsük a Cobra-t (parancssorból a projekt mappában)

```go
go get -u github.com/spf13/cobra@latest
```

3. A `pkg` mappában hozd létre a `logging.go` fájlt. Ebben létrehozunk egy napló olvasó és egy napló író függvényt:

```go
package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func NaploOlvaso(naploFajl string, sullyossag string) {
	// Fájl megnyitása
	f, err := os.Open(naploFajl)
	if err != nil {
		log.Fatal(err)
	}

	// Fájl zárás. Mér a tartalma megvan
	defer f.Close()

	//Súlyosság keresése a fájl tartalmában
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		// Ha megvan a megfeleő súlyosság, akkor kiírjuk a képernyőre
		if strings.Contains(s, " - "+sullyossag+" - ") {
			fmt.Println(s)
		}
	}
}

func NaploIro(naploFajl string, sulyossag string, bejegyzes string) {
	// Fájl megnyitása. Ha nem létezik, akkor létrehozzuk
	f, err := os.OpenFile(naploFajl,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Teljes bejegyzés összeállítása
	naploBejegyzes := (time.Now().Format(time.DateTime)) + " - " + sulyossag + " - " + bejegyzes + "\n"
	// Bejegyzés fájlba írása
	l, err := f.WriteString(naploBejegyzes)
	if err != nil {
		log.Fatalln(err)
		defer f.Close()
		return
	}

	// Visszajelzés a felhasználónak
	fmt.Println(l, "karakter kiírva fájlba")
	// Fájl zárása
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
		defer f.Close()
		return
	}
}
```

4. Hozd, `cmd/naplo/ir/ir.go` fájlt:

```go
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

```

5. Hozd, `cmd/naplo/olvas/olvas.go` fájlt:

```go
package olvas

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

```

6. Hozd, `cmd/naplo/naplo.go` fájlt:

```go
package naplo

import (
	"github.com/cloudsteak/trn-go-cli/cmd/naplo/ir"
	"github.com/cloudsteak/trn-go-cli/cmd/naplo/olvas"
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

```

7. Hozd, `cmd/root.go` fájlt:

```go
package cmd

import (
	"log"

	"github.com/cloudsteak/trn-go-cli/cmd/naplo"
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


```

8. Hozd létre a `main.go` fájlt a projekt gyökerében. A szerkesztőben a `main.go` fájlba illeszd bel az alábbi kódot:

```go
package main

import (
	"github.com/cloudsteak/trn-go-cli/cmd"
)

func main() {
	cmd.Execute()
}
```

9. Mentsd el.

10. Tereminalban futtasd az alábbi parancsot:

```bash
go get .
```

```bash
go run . -h
```

# Alkalmazás fordítása (build)

Ha szeretnénk az alkalmazásunkat máshol is futtatni, anélkül, hogy minden fejlesztői eszközt és függőséget telepíteni kellene, akkor azt egy csomagba le is tudjuk fordítani (build). Ehhez az alábbi parancsot kell futtatni: `go build`

Eredményképpen Windows-on egy exe fájlt kapunk, amit futtathatunk a Go fejlesztői környezewten kívül is.

## Meglévő kód használata

1. Terminalban belépek a projekt mappába
2. Terminal-ban lefuttatom az alábbi parancsot:

```bash
go get .
```

## Naplóbejegyzések hozzáadása a fájlhoz

Adjunk hozzá néhány bejegyzést. Használjuk ehhez a fordított exe-t

```bash
./trn-go-cli.exe naplo ir -b "Sikeresen bejelentkezett." -s "INFO"
./trn-go-cli.exe naplo ir -b "A keresett oldal nem található. Kérjük, ellenőrizze az URL-t, és próbálja meg újra." -s "ERROR"
./trn-go-cli.exe naplo ir -b "Sikeresen frissítette a felhasználói adatait." -s "INFO"
./trn-go-cli.exe naplo ir -b "Az akkumulátor töltöttsége alacsony. Kérjük, csatlakoztassa a készüléket egy töltőhöz." -s "WARN"
./trn-go-cli.exe naplo ir -b "Az Ön munkamenete lejárt. Kérjük, jelentkezzen be újra a folytatáshoz." -s "WARN"
./trn-go-cli.exe naplo ir -b "Váratlan rendszerhiba történt. Próbálkozzon újra később, vagy lépjen kapcsolatba a támogatással." -s "ERROR"
./trn-go-cli.exe naplo ir -b "Sikeresen elfogadta a szerződési feltételeket. Most már hozzáférhet az alkalmazás teljes funkcionalitásához." -s "INFO"
./trn-go-cli.exe naplo ir -b "Nem sikerült kapcsolódni a szerverhez. Kérjük, ellenőrizze internetkapcsolatát és próbálja újra." -s "ERROR"
./trn-go-cli.exe naplo ir -b "Az alkalmazás használatához frissítse a licencét." -s "WARN"
./trn-go-cli.exe naplo ir -b "A tranzakció sikeresen megtörtént. Köszönjük a vásárlást" -s "INFO"
./trn-go-cli.exe naplo ir -b "Sikeresen kijelentkezett. Várjuk vissza" -s INFO
```

## Naplóbejegyzések olvasása a fájlból

### INFO

```bash
./trn-go-cli.exe naplo olvas -s "INFO"
```

### WARN

```bash
./trn-go-cli.exe naplo olvas -s "WARN"
```

### ERROR

```bash
./trn-go-cli.exe naplo olvas -s "ERROR"
```
