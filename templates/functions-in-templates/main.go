package main

import (
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

// declares a func map data type
var fm = template.FuncMap{
	"upper": strings.ToUpper,
	"lower": strings.ToLower,
}

func init() {
	// constructs a new template(returned from template.New("")), which allows to use the
	// receiver function Funcs and passes the fm variable to it. Since funcs returns a pointer to
	// a template, parseFiles is called to parse the template
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", struct {
		Fname string
		Lname string
	}{
		Fname: "John",
		Lname: "Doe",
	})
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
