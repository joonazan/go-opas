package main

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Aihe struct {
	Nimi      string
	URL       template.URL
	Esitiedot []string
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
				Nimi:      scanner.Text()[2:],
				URL:       template.URL(tiedostonimi),
				Esitiedot: haeEsitiedot(id),
			}
			lisääOpetussivu(tiedostonimi, tiedostopolku)
		} else {
			log.Println("Jotain pielessä tiedostossa", tiedostopolku)
		}
	}

	kuvaTiedostot, err := filepath.Glob(kansio + "*.png")
	if err != nil {
		panic(err)
	}

	for _, kuvapolku := range kuvaTiedostot {
		tiedostonimi := filepath.Base(kuvapolku)
		lisääKuva(tiedostonimi, kuvapolku)
	}
}

// jos esitietotiedostoa ei löydy, palauttaa yhden esitiedon id:llä "määrittelemättömät esitiedot"
func haeEsitiedot(id string) []string {
	tiedosto, err := os.Open("../esitiedot/" + id)
	if err != nil {
		return []string{"määrittelemättömät esitiedot"}
	}
	defer tiedosto.Close()

	esitiedot := make([]string, 0)
	scanner := bufio.NewScanner(tiedosto)
	for scanner.Scan() {
		rivi := strings.TrimSpace(scanner.Text())
		if rivi != "" {
			esitiedot = append(esitiedot, rivi)
		}
	}

	return esitiedot
}
