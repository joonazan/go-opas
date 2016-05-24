package main

import (
	"html/template"
	"net/http"
)

func init() {

	t, err := template.ParseFiles("data/koe.html")
	if err != nil {
		panic(err)
	}

	const (
		koeURL    = "/koe"
		codeField = "koodi"
	)

	loginRequired(koeURL, func(w http.ResponseWriter, r *http.Request, u User) {
		r.ParseForm()

		var reply, code string
		if len(r.Form) != 0 {
			code = r.FormValue(codeField)
			output, err := Run(code)
			if err != nil {
				reply = err.Error()
			} else {
				const toivottu = "Hello, world!\n"
				reply = "Ohjelmasi tulosti:\n" + output + "\n"
				if output == toivottu {
					reply += "Oikein!"
				} else {
					reply += "Väärin; pitäisi olla: \n" + toivottu
				}
			}
		}

		t.Execute(w, struct {
			FormAction, CodeName, Code, Reply string
		}{FormAction: koeURL,
			CodeName: codeField,
			Reply:    reply,
			Code:     code,
		})
	})
}