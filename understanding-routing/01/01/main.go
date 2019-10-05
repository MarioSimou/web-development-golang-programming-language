package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler string

// routing middleware - single function
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/dog":
		fmt.Fprintln(w, "dog dog")
	case "/cat":
		fmt.Fprintln(w, "cat cat")
	default:
		fmt.Fprintln(w, "I don't know")
	}
	return
}

func main() {
	var h handler
	log.Fatal(http.ListenAndServe(":8080", h))
}
