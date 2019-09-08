package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", sages)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
