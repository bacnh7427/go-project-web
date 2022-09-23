package handlers

import (
	"back-go-land/pkg/config"
	"back-go-land/pkg/render"
	"net/http"
)

// Repo the repository users by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creatates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func (m *Repository) New(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "new.page.tmpl")
}
