package main

import (
	"encoding/json"
	"net/http"
)

func init() {
	http.HandleFunc("/progressdump", loginRequired(func(w http.ResponseWriter, r *http.Request, u User) {
		if u.Email == "joon.saar@gmail.com" {
			text, err := json.Marshal(edistymiset)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write(text)
		}
	}))
}
