package main

import (
	"fmt"
	"net/http"
)

func handleUsers(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Fprintln(w, "Welcome to our page "+name)
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8080", nil)
}
