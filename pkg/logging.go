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
	f, err := os.Open(naploFajl)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, sullyossag) {
			fmt.Println(s)
		}
	}
}

func NaploIro(naploFajl string, sulyossag string, bejegyzes string) {

	f, err := os.OpenFile(naploFajl,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	naploBejegyzes := (time.Now().Format(time.DateTime)) + " - " + sulyossag + " - " + bejegyzes + "\n"
	fmt.Println(naploBejegyzes)
	l, err := f.WriteString(naploBejegyzes)
	if err != nil {
		log.Fatalln(err)
		defer f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
		defer f.Close()
		return
	}
	defer f.Close()
}
