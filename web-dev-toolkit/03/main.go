package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Fprintln(w, ctx)
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
