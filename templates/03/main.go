package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseGlob("./templates/*"))
	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "????")
}

func check(e error) {
	if e == nil {
		log.Fatalln(e)
	}
}
