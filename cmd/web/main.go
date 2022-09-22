package main

import (
	"back-go-land/pkg/config"
	"back-go-land/pkg/handlers"
	"back-go-land/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const PortNumber = ":8082"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/new", handlers.New)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
