package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}
func Dog(w http.ResponseWriter, r *http.Request) {
	// read image and serve it
	f, e := os.Open("../dog.jpeg")
	if e != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/dog", Dog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
