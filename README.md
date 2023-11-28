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
mkdir cmd
mkdir -p pkg
```

2. Hozd létre a `main.go` fájlt a projekt gyökerében
3. A szerkesztőben a `main.go` fülön, ked el gépelni: `package`
4. A VS Code felajánl több lehetőséget is. Nekünk jelnleg a `package main` szükséges
5. Illeszd az alábbi kódot a `package main` alá:

```go
import (
	"./cmd"
)
```

6. Mentsd el. Hibákat hagyd figyelmen kívül.
7. Telepítsük a Cobra-t
```go
go get -u github.com/spf13/cobra@latest
```

8. Hozz létre egy `root.go` fájlt a `cmd` mappában

`cmd/root.go`:

```go
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
```