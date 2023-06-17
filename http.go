package main

import (
	"fmt"
	"log"
	"net/http"
)

// Implementation of seperate HandleFunc.
func _HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server listening, you've requested: %s\n", r.URL.Path)
}

// Working hello world http server.  Browse to localhost:3000.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server listening, you've requested: %s\n", r.URL.Path)
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
