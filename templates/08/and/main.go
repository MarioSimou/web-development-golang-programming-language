package main

import (
	"os"
	"text/template"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	u1 := user{
		Name:  "Budha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}
	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}
	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: false,
	}

	users := []user{u1, u2, u3}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", users)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
