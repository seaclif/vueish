package vueish

import (
	"embed"
	"net/http"
)

//go:embed vueish.js
var Vueish embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/vueish.js"
	http.FileServer(http.FS(Vueish)).ServeHTTP(w, r)
}
