package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("starting server :8080")

	http.HandleFunc("/favicon.ico", favicon)

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func favicon(w http.ResponseWriter, r *http.Request) {}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
