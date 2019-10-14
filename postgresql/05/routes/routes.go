package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"../models"
	"github.com/julienschmidt/httprouter"
)

func GetBooks(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var bks models.Books
		rows, e := db.Query("SELECT * FROM books")
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
		defer rows.Close()

		for rows.Next() {
			var bk models.Book
			if e := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price); e != nil {
				http.Error(w, e.Error(), http.StatusInternalServerError)
			}
			bks = append(bks, bk)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(bks)
	}
}

func GetBook(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var bk models.Book
		isbn := p.ByName("isbn")

		row := db.QueryRow("SELECT * FROM books WHERE isbn=$1", isbn)
		if e := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price); e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(bk)
	}
}

func PostBook(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var bk models.Book
		json.NewDecoder(r.Body).Decode(&bk)

		_, e := db.Exec("INSERT INTO books(isbn,title,author,price) VALUES($1,$2,$3,$4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Location", strings.Join([]string{r.URL.Path, bk.Isbn}, "/"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(bk)
	}
}

func DeleteBook(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		isbn := p.ByName("isbn")

		_, e := db.Exec("DELETE FROM books WHERE isbn=$1", isbn)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(204)
	}
}

func PutBook(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var bk models.Book
		isbn := p.ByName("isbn")

		json.NewDecoder(r.Body).Decode(&bk)

		_, e := db.Exec("UPDATE books SET isbn=$1, title=$2,author=$3,price=$4 WHERE isbn=$5", bk.Isbn, bk.Title, bk.Author, bk.Price, isbn)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bk)
	}
}
