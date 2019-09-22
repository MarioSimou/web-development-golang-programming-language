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
	nums := []float32{2.32, 1.22, 5.67, 8.5, 3.45}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nums)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
