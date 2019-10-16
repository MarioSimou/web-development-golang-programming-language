package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
}

var user User // global scope

func marshal(w http.ResponseWriter, r *http.Request) {
	j, e := json.Marshal(user)

	if e != nil {
		http.Error(w, e.Error(), http.StatusUnprocessableEntity)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintln(w, string(j))

}

func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}

func main() {
	user = User{1, "john", ""}

	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
