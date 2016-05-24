package main

import (
	"bytes"
	"fmt"
	gfm "github.com/shurcooL/github_flavored_markdown"
	"io/ioutil"
)

func markdownHTMLläksi(md []byte) []byte {

	return rinnastuksetHTMLläksi(md)
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

func rinnastuksetHTMLläksi(md []byte) []byte {
	var (
		alku  = []byte("??rinnan")
		väli  = []byte("??v")
		loppu = []byte("??l")
	)

	palat := make([][]byte, 0)

	for {
		alkupaikka := bytes.Index(md, alku)
		if alkupaikka == -1 {
			break
		}

		palat = append(palat,
			gfm.Markdown(lisääKoodi(md[:alkupaikka])),
			[]byte("<div class=comparison><div class=compcell>"),
		)

		md = md[alkupaikka+len(alku):]

		loppupaikka := bytes.Index(md, loppu)

		for seuraava := bytes.Index(md, väli); seuraava != -1 && seuraava < loppupaikka; seuraava = bytes.Index(md, väli) {
			palat = append(palat,
				gfm.Markdown(lisääKoodi(md[:seuraava])),
				[]byte("</div><div class=compcell>"),
			)
			uusiAlku := seuraava + len(väli)
			md = md[uusiAlku:]
			loppupaikka -= uusiAlku
		}
		palat = append(palat,
			gfm.Markdown(lisääKoodi(md[:loppupaikka])),
			[]byte("</div></div>"),
		)
		md = md[loppupaikka+len(loppu):]
	}

	palat = append(palat, md)

	return bytes.Join(palat, []byte{})
}
