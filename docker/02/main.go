package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":80", nil))
}
