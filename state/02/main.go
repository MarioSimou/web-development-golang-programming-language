package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("name")

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `
	<form action="/" method="POST">
		<input type="text" name="name" placeholder="Your name" />
		<input type="submit" value="Submit" />
	</form><br>`+input)
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
