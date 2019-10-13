package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Book struct {
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

type Books struct {
	Data []Book `json:"data"`
}

func BooksGet(w http.ResponseWriter, r *http.Request) {
	var books Books
	rows, e := db.Query("SELECT * FROM books;")
	if e != nil {
		log.Fatal(e)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if e := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price); e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		books.Data = append(books.Data, book)
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(books)
}

func main() {
	var e error
	db, e = sql.Open("postgres", "postgresql://msimou@localhost:5432/bookstore?sslmode=disable")
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()

	http.HandleFunc("/books", BooksGet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
