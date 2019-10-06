package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", `text/html; charset="utf-8"`)
	fmt.Fprintln(w, `<img src="../../dog.jpeg"/>`)
}

func Dog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../dog.jpeg")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/dog", Dog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
