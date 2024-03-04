package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/FLOpeSS/crud2/internal/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var app handler.App

func main() {

	db, err := sql.Open("mysql", "sneto:jms@tcp(127.0.0.1:3306)/crud")
	if err != nil {
		fmt.Println("error ocurred: ", err)
	}
	defer db.Close()
	//
	r := mux.NewRouter()
	r.HandleFunc("/books", app.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", app.GetBook).Methods("GET")
	r.HandleFunc("/books", app.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", app.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", app.DeleteBook).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
