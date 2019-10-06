package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
func Dog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL.Path)
	fmt.Println("Method: ", r.Method)

	name := r.FormValue("name")
	fmt.Println("do some processing and then redirect ", name)

	// w.Header().Set("Location", "/doggy")
	// w.WriteHeader(http.StatusSeeOther)
	// or
	http.Redirect(w, r, "/doggy", http.StatusSeeOther)

}
func Doggy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL.Path)
	fmt.Println("Method: ", r.Method)

	fmt.Fprintln(w, "Thank you for your post")
}

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/dog", Dog)
	http.HandleFunc("/doggy", Doggy)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
