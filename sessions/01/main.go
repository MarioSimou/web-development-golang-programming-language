package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if id, e := uuid.NewV4(); e == nil {
		http.SetCookie(w, &http.Cookie{
			Name:     "sid",
			Value:    id.String(),
			MaxAge:   3600,
			HttpOnly: true,
		})
	} else {
		http.Error(w, "Error while creating cookie", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Welcome")
}

func main() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
