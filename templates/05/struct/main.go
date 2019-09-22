package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	sages := sage{Name: "Buddha", Motto: "The belief of no beliefs"}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", sages)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
