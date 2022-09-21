package main

import (
	"back-go-land/pkg/handlers"
	"fmt"
	"net/http"
)

const PortNumber = ":8082"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
