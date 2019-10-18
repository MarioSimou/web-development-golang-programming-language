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

	e = db.Ping()
	if e != nil {
		log.Fatal(e)
	}

	row := db.QueryRow("SELECT * FROM books WHERE isbn=$1", "978-1503261969")
	bk := Book{}
	err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No record found")
		return
	case err != nil:
		fmt.Printf(err.Error())
	}

	fmt.Printf("Record: %v, %v, %v, %v\n", bk.isbn, bk.title, bk.author, bk.price)
}
