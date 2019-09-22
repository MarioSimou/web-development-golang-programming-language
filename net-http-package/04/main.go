package main

import (
	"fmt"
	"net/http"
)

type myHandler string

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := "<h1>Hello World</h1>"
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache")
	fmt.Fprintln(w, s)
}

func main() {
	var mh myHandler
	http.ListenAndServe(":8080", mh)
}
