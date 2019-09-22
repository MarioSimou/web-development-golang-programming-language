package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// parse the template and return a pointer of template

	tpl := template.Must(template.ParseGlob("*.gohtml"))
	// or
	// tpl, err := template.ParseFiles("tpl.gohtml")

	// creates a writer stream of the file
	w, err := os.Create("tplReplica.gohtml")
	check(err)
	defer w.Close()

	// pass the writer stream and writes the template into it
	err = tpl.Execute(w, nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
