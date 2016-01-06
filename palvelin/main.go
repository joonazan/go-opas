package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const PalvelimenOsoite = "http://localhost:8080"

func main() {

	t, err := template.ParseFiles("data/aiheet.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", loginRequired(func(w http.ResponseWriter, r *http.Request, u User) {

		oppilaanEdistyminen := edistymiset.Edistyminen(u.Email)
		a := make([]AiheJaEdistyminen, len(aiheet))
		i := 0
		for k, v := range aiheet {
			a[i].Aihe = v
			a[i].Id = k
			a[i].Tila = oppilaanEdistyminen[k]
			i++
		}
		err := t.Execute(w, struct{ Aiheet []AiheJaEdistyminen }{Aiheet: a})
		if err != nil {
			log.Println(err)
		}
	}))

	http.HandleFunc("/setstatus", loginRequired(func(w http.ResponseWriter, r *http.Request, u User) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(r.Form) != 1 {
			http.Error(w, "Wrong amount of form params", http.StatusInternalServerError)
			return
		}

		for k, v := range r.Form {
			if _, ok := aiheet[k]; !ok {
				http.Error(w, "Aihe ei ole olemassa.", http.StatusInternalServerError)
				return
			}
			if len(v) != 1 {
				http.Error(w, "???", http.StatusInternalServerError)
			}
			integer, err := strconv.Atoi(v[0])
			if err != nil {
				http.Error(w, "Not an integer", http.StatusInternalServerError)
				return
			}
			edistymiset.Edistyminen(u.Email)[k] = uint8(integer)
		}

		http.Error(w, "What did you expect?", http.StatusNoContent)
	}))

	http.HandleFunc("/authcallback/google", authentication)

	log.Fatal(http.ListenAndServe(":8080", nil))
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
