package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/page/*.gohtml"))
	tpl = template.Must(tpl.ParseFiles("./templates/index.gohtml"))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "John Doe")
	check(err)
}
