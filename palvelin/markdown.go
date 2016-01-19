package main

import (
	"bytes"
	"fmt"
	gfm "github.com/shurcooL/github_flavored_markdown"
	"io/ioutil"
)

func markdownHTMLläksi(md []byte) []byte {

	return gfm.Markdown(lisääKoodi(md))
}

func lisääKoodi(md []byte) []byte {

	var koodimerkki = []byte("$$$")
	const koodikansio = "../esimerkit/"

	palat := make([][]byte, 0)

	for {
		alku := bytes.Index(md, koodimerkki)
		if alku == -1 {
			break
		}

		polunAlku := alku + len(koodimerkki)
		polunLoppu := bytes.Index(md[polunAlku:], koodimerkki) + polunAlku

		polku := koodikansio + string(md[polunAlku:polunLoppu])

		palat = append(palat, md[:alku],
			[]byte("```Go\n"),
			haeKoodi(polku),
			[]byte(fmt.Sprintf("```\n*%s*\n", polku)),
		)
		md = md[polunLoppu+len(koodimerkki):]
	}

	palat = append(palat, md)

	return bytes.Join(palat, []byte{})
}

func haeKoodi(polku string) []byte {
	tavut, virhe := ioutil.ReadFile(polku)
	if virhe != nil {
		return []byte("Kooditiedostoa ei löytynyt")
	} else {
		return tavut
	}
}
