package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connString := "postgresql://msimou@localhost:5432/bookstore?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := db.Ping()
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("YOur are connected!")
}
