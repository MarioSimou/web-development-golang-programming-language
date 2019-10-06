package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	var data string

	if r.Method == http.MethodPost {
		fmt.Println("Processing file...")
		f, fh, e := r.FormFile("file")
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		readable, e := ioutil.ReadAll(f)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		data = string(readable)

		// stores the file locally
		if nF, _ := os.Create(filepath.Join(".", "store", fh.Filename)); nF != nil {
			_, fe := nF.Write(readable)
			if fe != nil {
				http.Error(w, fe.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Error Processing the File", http.StatusInternalServerError)
			return
		}
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
