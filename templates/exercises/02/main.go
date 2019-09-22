package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type region struct {
	Name string
}
type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int32
	Region  string
}
type hotels []hotel

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	region := region{}
	hotels := hotels{
		hotel{
			Name:    "Hotel 1",
			Address: "Address 1",
			City:    "City 1",
			Zip:     54321,
			Region:  region.Central(),
		},
		hotel{
			Name:    "Hotel 2",
			Address: "Address 2",
			City:    "City 2",
			Zip:     43213,
			Region:  region.Nothern(),
		},
		hotel{
			Name:    "Hotel 3",
			Address: "Address 3",
			City:    "City 3",
			Zip:     35754,
			Region:  region.Southern(),
		},
	}

	e := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", hotels)
	check(e)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (r region) Nothern() string {
	return "Nothern"
}

func (r region) Central() string {
	return "Central"
}

func (r region) Southern() string {
	return "Southern"
}
