package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", []string{
		"John",
		"Foo",
		"Bar",
	})
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
