package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type file struct {
	Header []string
	Rows   []row
}
type row struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	AdjClose float64
}

// global
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	d, e := ioutil.ReadFile("./table.csv")
	check(e)

	var rows []row
	var header []string
	for i, v := range strings.Split(string(d), "\n")[0:100] {
		if i == 0 {
			header = strings.Split(v, ",")
			continue
		}

		r := createRow(v, ",")
		rows = append(rows, r)
	}

	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", file{
		Header: header,
		Rows:   rows,
	})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createRow(s string, sep string) row {
	sa := strings.Split(s, sep)
	open, _ := strconv.ParseFloat(sa[1], 64)
	high, _ := strconv.ParseFloat(sa[2], 64)
	low, _ := strconv.ParseFloat(sa[3], 64)
	close, _ := strconv.ParseFloat(sa[4], 64)
	volume, _ := strconv.ParseFloat(sa[5], 64)
	adjClose, _ := strconv.ParseFloat(sa[6], 64)

	return row{
		Date:     sa[0],
		Open:     open,
		High:     high,
		Low:      low,
		Close:    close,
		Volume:   volume,
		AdjClose: adjClose,
	}
}
