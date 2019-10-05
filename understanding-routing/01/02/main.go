package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// creates a new router, inside of it we register handlers
	mux := http.NewServeMux()

	cat := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "cat cat")
	}

	dog := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "dog dog")
	}

	mux.HandleFunc("/cat", cat)
	mux.HandleFunc("/dog", dog)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
