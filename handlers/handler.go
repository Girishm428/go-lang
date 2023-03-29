package handlers

import (
	"net/http"

	"github.com/girishm428/go-lang/renders"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemp(w, "home.page.tmpl")
}

// dont know where to find the erro
func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemp(w, "about.page.tmpl")
}
