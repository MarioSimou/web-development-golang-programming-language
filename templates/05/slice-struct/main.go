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
	buddha := sage{Name: "Buddha", Motto: "The belief of no beliefs"}
	gandhi := sage{Name: "Gandhi", Motto: "Be the change"}
	mlk := sage{Name: "Martin Luther King", Motto: "Hatred never ceases with hatred but with love alone is healed"}
	jesus := sage{Name: "Jesus", Motto: "Love all"}
	sages := []sage{buddha, gandhi, mlk, jesus}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", sages)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
