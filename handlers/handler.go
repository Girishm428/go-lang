package handlers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemp(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemp(w, "about.page.tmpl")
}
