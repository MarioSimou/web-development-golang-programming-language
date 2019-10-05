package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	dog := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "dog dog...")
	}
	cat := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "cat cat...")
	}

	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	log.Fatal(http.ListenAndServe(":8080", nil)) // default router is used
}
