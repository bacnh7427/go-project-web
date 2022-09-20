package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world")
		if err != nil {
			// fmt.Fprintf(err)
		}
		fmt.Println(fmt.Sprintf("NUmber of bytes writeen: %d ", n))
	})

	_ = http.ListenAndServe(":8082", nil)
}
