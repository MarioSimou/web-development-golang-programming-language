package main

import (
	"database/sql"
	"log"
	"net/http"

	"./routes"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var e error
	db, e = sql.Open("postgres", "postgresql://msimou@localhost:5432/bookstore?sslmode=disable")
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()

	router := httprouter.New()
	router.GET("/books", routes.GetBooks(db))
	router.GET("/books/:isbn", routes.GetBook(db))
	router.POST("/books", routes.PostBook(db))
	router.PUT("/books/:isbn", routes.PutBook(db))
	router.DELETE("/books/:isbn", routes.DeleteBook(db))

	log.Fatal(http.ListenAndServe(":8080", router))
}
