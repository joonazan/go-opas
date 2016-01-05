package main

import (
	"fmt"
	"net/http"
)

const PalvelimenOsoite = "http://localhost:8080"

func main() {

	http.HandleFunc("/", loginRequired(func(w http.ResponseWriter, r *http.Request, u User) {
		fmt.Fprintf(w, "%#v", u)
	}))

	http.HandleFunc("/authcallback/google", authentication)

	panic(http.ListenAndServe(":8080", nil))
}
