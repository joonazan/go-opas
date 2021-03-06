package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func adminpage(path string, page func(http.ResponseWriter)) {
	loginRequired(path, func(w http.ResponseWriter, r *http.Request, u User) {
		if u.Email == "joon.saar@gmail.com" {
			page(w)
		}
	})
}

func init() {

	adminpage("/progressdump", func(w http.ResponseWriter) {
		text, err := json.Marshal(edistymiset)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write(text)
	})

	adminpage("/progresssummary", func(w http.ResponseWriter) {

		edistymiset.RLock()
		defer edistymiset.RUnlock()

		for email, edistyminen := range edistymiset.data {

			tilalaskuri := make([]int, TilojenMäärä)

			edistyminen.RLock()
			for _, tila := range edistyminen.data {
				tilalaskuri[tila]++
			}
			edistyminen.RUnlock()

			line := fmt.Sprintf("%s tehty: %d aloitettu: %d\n", email, tilalaskuri[Tehty], tilalaskuri[Aloitettu])
			w.Write([]byte(line))
		}
	})
}
