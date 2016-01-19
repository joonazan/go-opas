package main

import (
	"io/ioutil"
	"net/http"
)

var (
	markdownsivunAlku  = []byte(`<!DOCTYPE html><head><meta charset="utf-8"><link href="../static/gfm.css" media="all" rel="stylesheet" type="text/css" /></head><body>`)
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
	tavut, err := ioutil.ReadFile("data/github-markdown.css")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/static/gfm.css", func(w http.ResponseWriter, r *http.Request) {
		w.Write(tavut)
	})
}
