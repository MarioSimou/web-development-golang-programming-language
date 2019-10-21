package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pinger")
}
func Instance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "webserver-0002")
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/instance", Instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":80", nil))
}
