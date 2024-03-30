package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type Names struct {
	Id    int
	Name  string
	Email string
}

func dbConn() *sql.DB {
	db, err := sql.Open("mysql", "sneto:jms@tcp(127.0.0.1:3306)/crud")
	if err != nil {
		fmt.Println("err: ", err)
	}
	return db
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homepage).Methods("GET")
	r.HandleFunc("/new", newUser)
	r.HandleFunc("/edit", edit)

	// r.HandleFunc("/show", show)

	fmt.Println("Server ruuning at port: 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()
	que, err := db.Query("SELECT * FROM names ORDER BY id DESC")
	if err != nil {
		fmt.Println("err: ", err)
	}
	n := Names{}
	res := []Names{}

	for que.Next() {
		var id int
		var name, email string

		err = que.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		n.Id = id
		n.Name = name
		n.Email = email

		res = append(res, n)

	}
	fmt.Println(res)
}

//	func show(w http.ResponseWriter, r *http.Request) {
//		db := dbConn()
//
//		nId := r.URL.Query().Get("id")
//		que, err := db.Query("SELECT * FROM names WHERE id=?", nId)
//		if err != nil {
//			fmt.Println("err: ", err)
//		}
//	}
func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request from newUser")
}

func edit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request from edit")
}
