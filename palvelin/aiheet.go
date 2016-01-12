package main

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type Aihe struct {
	Nimi string
	URL  template.URL
}

type AiheJaEdistyminen struct {
	Id string
	Aihe
	Tila uint8
}

var aiheet = lataaAiheet()

func lataaAiheet() (aiheet map[string]Aihe) {
	aiheet = make(map[string]Aihe)

	lataaAiheetKansiosta(aiheet, "../tehtävät/")
	lataaAiheetKansiosta(aiheet, "../ohjeet/")

	return
}

func lataaAiheetKansiosta(aiheet map[string]Aihe, kansio string) {

	markdownTiedostot, err := filepath.Glob(kansio + "*.md")
	if err != nil {
		panic(err)
	}

	for _, tiedostopolku := range markdownTiedostot {
		tiedosto, err := os.Open(tiedostopolku)
		if err != nil {
			panic(err)
		}
		defer tiedosto.Close()

		scanner := bufio.NewScanner(tiedosto)
		if scanner.Scan() && len(scanner.Text()) > 2 {
			tiedostonimi := filepath.Base(tiedostopolku)
			id := tiedostonimi[:len(tiedostonimi)-len(".md")]
			aiheet[id] = Aihe{
				Nimi: scanner.Text()[2:],
				URL:  template.URL(id),
			}
		} else {
			log.Println("Jotain pielessä tiedostossa", tiedostopolku)
		}
	}
}
