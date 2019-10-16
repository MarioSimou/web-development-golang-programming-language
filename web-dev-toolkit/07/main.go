package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServeTLS(":8080", "./cert.pem", "./key.pem", nil))
}

// generate key with: go run /usr/local/go/src/crypto/tls/generate_cert.go
