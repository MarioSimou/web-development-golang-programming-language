package main

import (
	"math"
	"os"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"sqrt":  math.Sqrt,
	"pow":   math.Pow,
	"round": math.Round,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	var x = 4.00
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", x)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
