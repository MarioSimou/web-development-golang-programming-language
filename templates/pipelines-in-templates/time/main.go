package main

import (
	"os"
	"text/template"
	"time"
)

var tpl *template.Template
var fm = template.FuncMap{
	"now": time.Now,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	t := time.Now().Format(time.RFC1123)
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", t)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
