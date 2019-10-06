package main

import (
	"fmt"
	"log"
	"net/http"
)

func Dog(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doggy", http.StatusMovedPermanently) // Memoised the redirection to /doggy
}
func Doggy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to doggy")
}

func main() {
	http.HandleFunc("/dog", Dog)
	http.HandleFunc("/doggy", Doggy)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
