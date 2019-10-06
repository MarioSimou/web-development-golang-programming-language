package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func Foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo ran")
}

func Dog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func main() {
	http.HandleFunc("/", Foo)
	http.HandleFunc("/dog/", Dog)
	http.Handle("/images/dog.jpeg", http.StripPrefix("/images", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8080", nil)
}
