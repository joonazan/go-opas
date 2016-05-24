package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

const PalvelimenOsoite = "https://go-opas.herokuapp.com"

func main() {

	t, err := template.ParseFiles("data/aiheet.html")
	if err != nil {
		log.Fatal(err)
	}

	loginRequired("/", func(w http.ResponseWriter, r *http.Request, u User) {

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
	})

	loginRequired("/setstatus", func(w http.ResponseWriter, r *http.Request, u User) {
		if k, v, err := extractSetstatusArguments(r); err == nil {
			edistymiset.Edistyminen(u.Email).Set(k, v)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/logout", logout)

	edistymiset.TallennaVälein(time.Minute * 10)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func extractSetstatusArguments(r *http.Request) (aihe string, tila uint8, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	if len(r.Form) != 1 {
		err = errors.New("Wrong amount of form params")
		return
	}

	for k, v := range r.Form {
		if _, ok := aiheet[k]; !ok {
			err = errors.New("Aihe ei ole olemassa.")
			return
		}
		if len(v) != 1 {
			err = errors.New("???")
		}
		integer, interr := strconv.Atoi(v[0])
		if interr != nil {
			err = errors.New("Not an integer")
			return
		}
		if integer > TilojenMäärä {
			err = errors.New("Tuo ei ole yksikään tiloista!")
			return
		}

		return k, uint8(integer), nil
	}

	err = errors.New("Impossible!")
	return
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
