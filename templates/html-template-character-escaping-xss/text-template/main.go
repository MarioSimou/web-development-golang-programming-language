package main

import (
	"os"
	"text/template"
)

type page struct {
	Header string
	Title  string
	Input  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func main() {
	h := page{Header: "Cross-Site Scripting", Title: "Cross Site Scripting", Input: "<script>alert('malicious code')</script>"}
	e := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", h)
	check(e)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
