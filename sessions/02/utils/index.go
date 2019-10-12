package utils

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func ParseTemplates(root string) (*template.Template, error) {
	tpl := template.New("")

	cb := func(path string, info os.FileInfo, e error) error {
		if strings.Contains(path, ".gohtml") {
			if _, e := tpl.ParseFiles(path); e != nil {
				return e
			}
		}
		return e
	}

	if e := filepath.Walk(root, cb); e != nil {
		return tpl, e
	}

	return tpl, nil
}
