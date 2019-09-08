package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	writeTemplateToFile(generateBaseHTML(1, os.Args...))
}

func writeTemplateToFile(tpl string) {
	f, err := os.Create("index.html")
	check(err)
	defer f.Close()

	// pass a string and return a reader stream
	r := strings.NewReader(tpl)
	check(err)

	// pipes the reader to the writer stream
	// written refers to the number of bytes that have been written
	written, err := io.Copy(f, r)

	fmt.Println(written)
	fmt.Printf("Wrote %d bytes\n", written)
}

func generateBaseHTML(index int, args ...string) string {
	var name string
	for i, v := range args {
		if i == index {
			name = v
		}
	}

	if name == "" {
		name = "Unknown"
	}

	return `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8"/>
			<meta name="Author" content="` + name + `" />
			<title>Intro</title>
		</head>
		<body>
			<h1>Dear ` + name + `</h1>
		</body>
	</html>
	`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
