package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func logOlvaso(naploFajl string, sullyossag string) {
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

func logIro(naploFajl string, sullyossag string, bejegyzes string) {
	f, err := os.Open(naploFajl)
	if err != nil {
		log.Fatal(err)
	}

	l, err := f.WriteString("Hello World")
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
