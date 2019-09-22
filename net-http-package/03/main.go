package main

import (
	"html/template"
	"net/http"
	"net/url"
)

type myHandler string

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle http request
	data := struct {
		URL           *url.URL
		Method        string
		Headers       map[string][]string
		Host          string
		RequestUri    string
		ContentLength int64
	}{
		r.URL,
		r.Method,
		r.Header,
		r.Host,
		r.RequestURI,
		r.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var mh myHandler
	http.ListenAndServe(":8080", mh)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
