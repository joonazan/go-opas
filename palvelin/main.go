package main

import (
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"github.com/stretchr/signature"
	"html/template"
	"io"
	"net/http"
)

const PalvelimenOsoite = "http://localhost:8080"

func main() {

	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New("712450255947-o5uppfp82sc3sb2c1dum7es8k93ras4q.apps.googleusercontent.com", "TrFX1I5QJZ_unJ-P5UgYLF4N", PalvelimenOsoite+"/authcallback/google"),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.New("loginpage").Parse(`
			<!DOCTYPE html>
				<body>
				<h2>Log in with...</h2>
				<ul>
					<li>
					<a href="{{.}}">Google</a>
					</li>
				</ul>
				</body>
		`)

		provider, err := gomniauth.Provider("google")
		if err != nil {
			panic(err)
		}

		state := gomniauth.NewState("after", "success")
		authUrl, err := provider.GetBeginAuthURL(state, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		template.Execute(w, authUrl)
	})

	http.HandleFunc("/authcallback/google", callbackHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		panic(err)
	}

	omap, err := objx.FromURLQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	creds, err := provider.CompleteAuth(omap)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, userErr := provider.GetUser(creds)

	if userErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := fmt.Sprintf("%#v", user)
	io.WriteString(w, data)
}
