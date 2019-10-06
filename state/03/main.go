package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	user := Person{
		FirstName:  r.FormValue("fname"),
		LastName:   r.FormValue("lname"),
		Subscribed: r.FormValue("subscribed") == "on",
	}
	w.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
