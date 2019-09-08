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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
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

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        4,
	}

	t := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{buddha, gandhi, mlk, jesus}
	cars := []car{f, t}

	data := items{Wisdom: sages, Transport: cars}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
