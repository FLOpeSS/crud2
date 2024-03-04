package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/FLOpeSS/crud2/internal/book"
	"github.com/gorilla/mux"
)

var bk book.Book

type App struct {
	DB *sql.DB
}

func (app *App) GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer rows.Close()

	books := []book.Book{}

	for rows.Next() {
		var book book.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			fmt.Println("error occurred: ", err)
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func (app *App) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("error ocurred: ", err)
		log.Fatal(err)
	}

	var book book.Book
	row := app.DB.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id).Scan(&book.Id, &book.Title, &book.Author)
	fmt.Println(row)
}

func (app *App) CreateBook(w http.ResponseWriter, r *http.Request) {

}

func (app *App) UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func (app *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
}
