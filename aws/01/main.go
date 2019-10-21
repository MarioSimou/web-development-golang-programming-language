package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Oh yeah, I am running on AWS")
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":80", nil))
}
