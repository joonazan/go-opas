package main

import (
	"github.com/zond/gosafe"
	"net/http"
)

func init() {

	compiler := gosafe.NewCompiler()
	for _, allowed := range []string{"fmt", "math"} {
		compiler.Allow(allowed)
	}

	loginRequired("/koe", func(w http.ResponseWriter, r *http.Request, u User) {})
}
