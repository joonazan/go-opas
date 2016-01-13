package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
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
		for id, aihe := range aiheet {
			a[i].Aihe = aihe
			a[i].Id = id
			a[i].Tila = oppilaanEdistyminen.Get(id)
			a[i].TäytetytVaatimukset, a[i].TäyttämättömätVaatimukset = luokitteleEsitiedot(aihe.Esitiedot, oppilaanEdistyminen)
			i++
		}

		Aiheluettelo(a).Lajittele()

		err := t.Execute(w, struct {
			Aiheet       []AiheJaEdistyminen
			TilojenNimet []string
			IdAiheeksi   map[string]Aihe
		}{
			Aiheet:       a,
			TilojenNimet: []string{"Aloittamaton", "Tehty!", "Aloitettu"},
			IdAiheeksi:   aiheet,
		})
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
			if integer > TilojenMäärä {
				http.Error(w, "Tuo ei ole yksikään tiloista!", http.StatusInternalServerError)
				return
			}
			edistymiset.Edistyminen(u.Email).Set(k, uint8(integer))
		}

		http.Error(w, "What did you expect?", http.StatusNoContent)
	}))

	http.HandleFunc("/authcallback/google", authentication)

	http.HandleFunc("/logout", logout)

	edistymiset.TallennaVälein(time.Minute * 10)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var edistymiset = lataaEdistymiset()

func luokitteleEsitiedot(esitiedot []string, edistyminen Edistyminen) (täytetty []string, täyttämätön []string) {
	for _, aiheId := range esitiedot {
		if edistyminen.Get(aiheId) == Tehty {
			täytetty = append(täytetty, aiheId)
		} else {
			täyttämätön = append(täyttämätön, aiheId)
		}
	}
	return
}
