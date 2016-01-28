package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var (
	markdownsivunAlku  = []byte(`<!DOCTYPE html><head><meta charset="utf-8"><link rel="stylesheet" type="text/css" href="/static/ohje.css"></head><body>`)
	markdownsivunLoppu = []byte(`</body>`)
)

func lisääOpetussivu(nimi string, tiedostopolku string) {
	tavut, err := ioutil.ReadFile(tiedostopolku)
	if err != nil {
		panic(err)
	}

	html := markdownHTMLläksi(tavut)

	http.HandleFunc("/materiaali/"+nimi, func(w http.ResponseWriter, r *http.Request) {
		w.Write(markdownsivunAlku)
		w.Write(html)
		w.Write(markdownsivunLoppu)
	})
}

func lisääKuva(nimi string, tiedostopolku string) {
	tavut, err := ioutil.ReadFile(tiedostopolku)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/materiaali/"+nimi, func(w http.ResponseWriter, r *http.Request) {
		w.Write(tavut)
	})
}

func init() {
	const stylefolder = "data/"
	polut, err := filepath.Glob(stylefolder + "*.css")
	if err != nil {
		panic(err)
	}
	for _, polku := range polut {
		tavut, err := ioutil.ReadFile(polku)
		if err != nil {
			panic(err)
		}

		http.HandleFunc("/static/"+filepath.Base(polku), func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/css")
			w.Write(tavut)
		})
	}
}
