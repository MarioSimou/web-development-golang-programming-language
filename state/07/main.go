package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "sid",
		Value:    "1234",
		MaxAge:   86400,
		HttpOnly: true,
	}

	// writes a Set-Cookie HTTP Header to the response
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Welcome to our page</h1>")

	// set cookie and returns a response
}

func About(w http.ResponseWriter, r *http.Request) {
	// reads the cookie and returns a response
	// fmt.Println(r.Cookie("sid"))
	// fmt.Println(r.Cookies())
	cookie := r.Header.Get("Cookie")
	fmt.Println("Cookie: ", cookie)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Welcome to about section</h1>")
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/about", About)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
