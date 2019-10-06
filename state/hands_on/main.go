package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Cookie"))
	if counter, _ := r.Cookie("counter"); counter != nil {
		fmt.Fprintln(w, renderHTML("home", counter.Value))
	} else {
		http.Error(w, "Error Retrieving the cookie", http.StatusInternalServerError)
		return
	}
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Cookie"))
	if counter, _ := r.Cookie("counter"); counter != nil {
		fmt.Fprintln(w, renderHTML("About", counter.Value))
	} else {
		http.Error(w, "Error Retrieving the cookie", http.StatusInternalServerError)
		return
	}
}
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Cookie"))
	if counter, _ := r.Cookie("counter"); counter != nil {
		fmt.Fprintln(w, renderHTML("Contact", counter.Value))
	} else {
		http.Error(w, "Error Retrieving the cookie", http.StatusInternalServerError)
		return
	}
}
func CookieMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if counter, _ := r.Cookie("counter"); counter != nil {
			if v, _ := strconv.Atoi(counter.Value); v != 0 {
				http.SetCookie(w, &http.Cookie{
					Name:   "counter",
					Value:  strconv.Itoa(v + 1),
					MaxAge: 3600,
				})
			} else {
				counter.Value = "1"
			}
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:   "counter",
				Value:  "1",
				MaxAge: 3600,
			})
		}

		next.ServeHTTP(w, r)
	})
}

func renderHTML(site string, counter string) string {
	return `
	<div>
		<h1>` + strings.ToUpper(site) + `</h1>
		<p>Counter: ` + counter + `</p>
	</div>`
}

func main() {
	http.Handle("/", CookieMiddleware(Index))
	http.Handle("/about", CookieMiddleware(About))
	http.Handle("/contact", CookieMiddleware(Contact))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
