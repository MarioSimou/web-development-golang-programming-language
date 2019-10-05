package main

import (
	"html/template"
	"log"
	"net/http"
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
func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/dog", Dog)
	http.HandleFunc("/me/", Me)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
