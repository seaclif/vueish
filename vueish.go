package vueish

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed vueish.js
var Vueish string

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/javascript")
	fmt.Fprintf(w, Vueish)
}
