package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "home")
}
func Dog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", "dog")
}
func Me(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", "me")
}

type handler string

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		tpl.ExecuteTemplate(w, "index.gohtml", "home")
	case "/dog":
		tpl.ExecuteTemplate(w, "dog.gohtml", "dog")
	case "/me/":
		tpl.ExecuteTemplate(w, "me.gohtml", "me")
	}
}

func init() {
	fm := template.FuncMap{
		"uc": strings.ToUpper,
	}

	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("./templates/*.gohtml"))
}

func main() {
	var h handler
	http.Handle("/", h)
	http.Handle("/dog", h)
	http.Handle("/me/", h)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
