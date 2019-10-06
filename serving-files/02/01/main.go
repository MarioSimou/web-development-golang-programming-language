package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<img src="../dog.jpeg" alt="dog" />`)
}

func Dog(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open("../dog.jpeg")
	if e != nil {
		http.Error(w, "File not found", 404)
		return
	}

	fi, e := f.Stat()
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/dog", Dog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
