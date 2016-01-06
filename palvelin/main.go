package main

import (
	"html/template"
	"net/http"
)

const PalvelimenOsoite = "http://localhost:8080"

func main() {

	t, err := template.ParseFiles("data/aiheet.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", loginRequired(func(w http.ResponseWriter, r *http.Request, u User) {

		oppilaanEdistyminen := edistymiset.Edistyminen(u.Email)
		a := make([]AiheJaEdistyminen, len(aiheet))
		i := 0
		for k, v := range aiheet {
			a[i].Aihe = v
			a[i].Tila = oppilaanEdistyminen[k]
			i++
		}
		t.Execute(w, struct{ Aiheet []AiheJaEdistyminen }{Aiheet: a})
	}))

	http.HandleFunc("/authcallback/google", authentication)

	panic(http.ListenAndServe(":8080", nil))
}

const (
	EiTehty = iota
	Tehty
	Aloitettu
)

type Edistyminen map[string]uint8

type Edistymiset map[string]Edistyminen

func (e Edistymiset) Edistyminen(oppilas string) Edistyminen {
	if o, ok := e[oppilas]; ok {
		return o
	}
	e[oppilas] = make(Edistyminen)
	return e[oppilas]
}

func (e Edistymiset) Tallenna() {

}

var (
	edistymiset = make(Edistymiset)
	aiheet      = make(map[string]Aihe)
)

func init() {
	aiheet = map[string]Aihe{
		"ohjelma": {
			Nimi: "Mikä on ohjelma?",
		},
		"setup": {Nimi: "Go:n asentaminen"},
	}
}
