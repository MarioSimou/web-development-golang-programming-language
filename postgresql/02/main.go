package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, e := sql.Open("postgres", "postgresql://msimou@localhost:5432/bookstore?sslmode=disable")
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()

	rows, e := db.Query("SELECT * FROM books;")
	if e != nil {
		log.Fatal(e)
	}
	defer rows.Close()

	for rows.Next() {
		var book Book

		if err := rows.Scan(&book.isbn, &book.title, &book.author, &book.price); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Row: %v, %v, %v, $%v\n", book.isbn, book.title, book.author, book.price)
	}
}
