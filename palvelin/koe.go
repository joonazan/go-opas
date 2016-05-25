package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type Tehtävä struct {
	kuvaus     template.HTML
	mallikoodi string
	tulosteet  []string
}

func init() {

	koeTemplate, err := template.ParseFiles("data/koe.html")
	if err != nil {
		panic(err)
	}

	trackitTemplate, err := template.ParseFiles("data/koetrackit.html")
	if err != nil {
		panic(err)
	}

	const (
		koeJuuri  = "/koe/"
		codeField = "koodi"

		koeKansio     = "../kokeet"
		syötetiedosto = "inputs"
	)

	kokeet, err := filepath.Glob(koeKansio + "/*")
	if err != nil {
		panic(err)
	}
	tehtävienPituudet := make([]int, len(kokeet))
	for i, polku := range kokeet {
		kokeet[i] = path.Base(polku)
	}

	trackSummariesFor := func(u User) []TrackSummary {
		dats := make([]TrackSummary, len(kokeet))

		for i, kokeenNimi := range kokeet {
			statukset := make([]int, tehtävienPituudet[i])
			for i := 0; i <= kokeenVaihe(u, kokeenNimi); i++ {
				if oikein(u, kokeenNimi, i) {
					statukset[i] = 1
				} else {
					statukset[i] = 2
				}
			}
			dats[i] = TrackSummary{Name: kokeenNimi, Statukset: statukset}
		}
		return dats
	}

	for moneskokoe, kokeenNimi := range kokeet {
		koeURL := koeJuuri + kokeenNimi
		ohitusURL := koeURL + "/skip"
		kansio := path.Join(koeKansio, kokeenNimi)
		syötteet, err := lueRivit(path.Join(kansio, syötetiedosto))
		if err != nil {
			panic("ei saatu syötteitä tiedostosta: " + err.Error()) // TODO: syötteettömät tehtävät
		}

		var tehtävät []Tehtävä
		for i := 0; ; i++ {
			fn := path.Join(kansio, fmt.Sprintf("%d.go", i))
			if _, err := os.Stat(fn); os.IsNotExist(err) {
				break
			}

			bin, err := CompileFile(fn)
			if err != nil {
				panic("virhe kääntäessä mallikoodia: " + err.Error())
			}
			defer os.Remove(bin)

			tehtävä := Tehtävä{
				kuvaus:     template.HTML(paniccingReadFile(path.Join(kansio, fmt.Sprintf("%d.html", i)))),
				mallikoodi: paniccingReadFile(fn),
				tulosteet:  make([]string, len(syötteet)),
			}
			for j, input := range syötteet {
				output, err := RunBinary(bin, input)
				if err != nil {
					panic("virhe ajaessa mallikoodia: " + err.Error())
				}
				tehtävä.tulosteet[j] = output
			}
			tehtävät = append(tehtävät, tehtävä)
		}

		tehtävienPituudet[moneskokoe] = len(tehtävät)

		previousSolution := func(u User, kokeenNimi string) string {
			edellinen := kokeenVaihe(u, kokeenNimi) - 1
			if edellinen == -1 {
				return ""
			}
			if oikein(u, kokeenNimi, edellinen) {
				return koodi(u, kokeenNimi, edellinen)
			} else {
				return tehtävät[edellinen].mallikoodi
			}
		}

		loginRequired(koeURL, func(w http.ResponseWriter, r *http.Request, u User) {

			vaihe := kokeenVaihe(u, kokeenNimi)
			last := int(vaihe) == len(tehtävät)-1
			tehtävä := tehtävät[vaihe]

			code := koodi(u, kokeenNimi, vaihe)

			var reply string
			if r.ParseForm(); len(r.Form) != 0 {
				code = r.FormValue(codeField)

				tallennaKoodi(u, kokeenNimi, vaihe, code)

				var correct bool
				correct, reply = grade(code, syötteet, tehtävä.tulosteet)
				if correct {
					redisClient.Set(oikeinAvain(kokelasAvain(u, kokeenNimi), vaihe), 1, 0)
					if !last {
						edistyKokeessa(u, kokeenNimi)

						vaihe++
						last = int(vaihe) == len(tehtävät)-1
						tehtävä = tehtävät[vaihe]
					}
				}
			}

			koeTemplate.Execute(w, struct {
				FormAction, SkipURL, CodeName, Code, Reply, PreviousSolution string
				Description                                                  template.HTML
				Last                                                         bool
				Phase                                                        int
			}{
				FormAction:       koeURL,
				SkipURL:          ohitusURL,
				CodeName:         codeField,
				Description:      tehtävä.kuvaus,
				Code:             code,
				Reply:            reply,
				PreviousSolution: previousSolution(u, kokeenNimi),
				Last:             last,
				Phase:            vaihe,
			})
		})

		loginRequired(ohitusURL, func(w http.ResponseWriter, r *http.Request, u User) {
			vaihe := kokeenVaihe(u, kokeenNimi)
			if int(vaihe) != len(tehtävät)-1 {
				edistyKokeessa(u, kokeenNimi)
			}
			http.Redirect(w, r, koeURL, http.StatusSeeOther)
		})
	}

	const kokeenTekijät = "kokeenTekijät"
	loginRequired(koeJuuri, func(w http.ResponseWriter, r *http.Request, u User) {
		redisClient.SAdd(kokeenTekijät, u.Email)

		trackitTemplate.Execute(w, []TestSummary{{u.Email, trackSummariesFor(u)}})
	})

	adminpage(koeJuuri+"admin", func(w http.ResponseWriter) {
		var tss []TestSummary
		emails, err := redisClient.SMembers(kokeenTekijät).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, email := range emails {
			ts := TestSummary{Email: email, Tracks: trackSummariesFor(User{Email: email})}
			tss = append(tss, ts)
		}
		trackitTemplate.Execute(w, tss)
	})
}

type TestSummary struct {
	Email  string
	Tracks []TrackSummary
}

type TrackSummary struct {
	Name      string
	Statukset []int
}

func paniccingReadFile(fn string) string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func kokeenVaihe(u User, kokeenNimi string) int {
	return lueLuku(kokelasAvain(u, kokeenNimi))
}

func edistyKokeessa(u User, kokeenNimi string) {
	avain := kokelasAvain(u, kokeenNimi)
	vaihe := kokeenVaihe(u, kokeenNimi)
	redisClient.Set(avain, vaihe+1, 0)
}

func oikein(u User, kokeenNimi string, vaihe int) bool {
	i, _ := redisClient.Get(oikeinAvain(kokelasAvain(u, kokeenNimi), vaihe)).Int64()
	if i == 1 {
		return true
	}
	return false
}

func kokelasAvain(u User, kokeenNimi string) string {
	return u.Email + kokeenNimi
}

func oikeinAvain(x string, y int) string {
	return fmt.Sprintf("OV%d", y) + x
}

func lueLuku(key string) int {
	resp, err := redisClient.Get(key).Int64()
	if err != nil {
		return 0
	}
	return int(resp)
}

func koodi(u User, kokeenNimi string, tehtävä int) string {
	b, _ := redisClient.Get(tehtäväAvain(u, kokeenNimi, tehtävä)).Bytes()
	return string(b)
}

func tallennaKoodi(u User, kokeenNimi string, tehtävä int, koodi string) {
	redisClient.Set(tehtäväAvain(u, kokeenNimi, tehtävä), koodi, 0)
}

func tehtäväAvain(u User, kokeenNimi string, tehtävä int) string {
	return fmt.Sprintf("%d", tehtävä) + kokelasAvain(u, kokeenNimi)
}

func grade(code string, inputs, outputs []string) (bool, string) {

	binary, err := Compile(code)
	if err != nil {
		return false, err.Error()
	}
	defer os.Remove(binary)

	for i, input := range inputs {
		output, err := RunBinary(binary, input)
		if err != nil {
			return false, err.Error()
		}
		if output != outputs[i] {
			return false, "Syötteellä '" + input + "' ohjelmasi tulosti:\n" + output + "\nVäärin; pitäisi olla: \n" + outputs[i]
		}
	}

	return true, "Oikein!"
}
