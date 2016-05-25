package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type Tehtävä struct {
	kuvaus     template.HTML
	mallikoodi string
	tulosteet  []string
}

func init() {

	t, err := template.ParseFiles("data/koe.html")
	if err != nil {
		panic(err)
	}

	const (
		koeURL    = "/koe/"
		codeField = "koodi"

		koeKansio     = "../kokeet"
		syötetiedosto = "inputs"
	)

	kansio := "../kokeet/neliosumma"
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

	loginRequired(koeURL, func(w http.ResponseWriter, r *http.Request, u User) {

		tehtävä := tehtävät[0]

		var reply, code string
		if r.ParseForm(); len(r.Form) != 0 {
			code = r.FormValue(codeField)
			_, reply = grade(code, syötteet, tehtävä.tulosteet)
		}

		t.Execute(w, struct {
			FormAction, CodeName, Code, Reply string
			Description                       template.HTML
		}{
			FormAction:  koeURL,
			CodeName:    codeField,
			Description: tehtävä.kuvaus,
			Code:        code,
			Reply:       reply,
		})
	})
}

func paniccingReadFile(fn string) string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	return string(data)
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
			return false, "Ohjelmasi tulosti:\n" + output + "\nVäärin; pitäisi olla: \n" + outputs[i]
		}
	}

	return true, "Oikein!"
}
