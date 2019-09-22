package main

import (
	"fmt"
	"net/http"
)

type myHandler string

func main() {
	var mh myHandler
	http.ListenAndServe(":8080", mh)
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Anyc ode you want to write")
}
