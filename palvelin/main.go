package main

import (
	"fmt"
	"net/http"
)

const PalvelimenOsoite = "http://localhost:8080"

func main() {

	http.HandleFunc("/", loginRequired(func(w http.ResponseWriter, r *http.Request, u User) {
		fmt.Fprint(w, edistymiset.Edistyminen(u.Email))
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

var edistymiset = make(Edistymiset)
