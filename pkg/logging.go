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
			fmt.Print(s)
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
