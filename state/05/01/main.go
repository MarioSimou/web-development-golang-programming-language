package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type handler string

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body string
	if r.Method == http.MethodPost {
		bf := make([]byte, r.ContentLength)
		r.Body.Read(bf)
		body = string(bf)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", body)
}

func main() {
	var h handler
	log.Fatal(http.ListenAndServe(":8080", h))
}
