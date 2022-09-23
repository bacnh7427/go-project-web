package main

import (
	"back-go-land/pkg/config"
	"back-go-land/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/new", http.HandlerFunc(handlers.Repo.New))
	return mux
}
