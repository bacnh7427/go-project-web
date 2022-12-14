package handlers

import (
	"back-go-land/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func New(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "new.page.tmpl")
}
