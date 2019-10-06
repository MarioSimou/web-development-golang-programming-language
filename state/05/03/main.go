package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl *template.Template

type handler string

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	readable, e := ioutil.ReadAll(r.Body)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	fmt.Println(string(readable))
	tpl.ExecuteTemplate(w, "index.gohtml", string(readable))
}

func main() {
	var h handler
	log.Fatal(http.ListenAndServe(":8080", h))
}
