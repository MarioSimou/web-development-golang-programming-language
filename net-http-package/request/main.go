package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type myHandler string

var tpl *template.Template

func init() {
	// parses all templates and pass them to tpl variable. Tpl is a container of
	// all templates
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	var mh myHandler
	log.Fatal(http.ListenAndServe(":8080", mh))
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// - a proper handling would request a filtering of all routes
	// - in that case we assume that all users will request /

	switch r.Method {
	case "GET":
		// render data without data
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
		break
	case "POST":
		fmt.Println(r.Header["Content-Type"])
		// we can check the Content-Type header to inspect for
		// - application/x-www-form-urlencoded
		// - multipart/form-data

		// get request data
		e := r.ParseForm()
		check(e)

		// render template with data
		fmt.Println(r.PostForm)
		tpl.ExecuteTemplate(w, "index.gohtml", r.PostForm)
		break
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
