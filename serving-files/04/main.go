package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir(".."))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
