package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// Note: Thinks tpl as a container of template, which means that more templates
	// can be added in a later stage
	tplName := "tpl.gohtml"
	tpl := template.Must(template.ParseFiles(tplName))
	err := tpl.ExecuteTemplate(os.Stdout, tplName, nil)
	// or (only for a single template)
	// err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
