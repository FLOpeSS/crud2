package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// json.NewDecoder is reading data from r.body
// r.body is data sent by the client
// Decode(*usr) is going to decode from the r.body to usr variable

// Insert new user to db

func inserUser(w http.ResponseWriter, r *http.Request) {
	var user Names
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := dbConn()
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	res, err := db.Exec("INSERT INTO names (name) VALUES(?)", user.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(res)

}

type Names struct {
	Id    int
	Name  string
	Email string
}

func dbConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "sneto:jms@tcp(127.0.0.1:3306)/crud")
	if err != nil {
		return nil, err
	}
	return db, nil
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
	db, err := dbConn()
	if err != nil {
		fmt.Errorf("Error on connecting db: ", err)
	}
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

// func showing_names(w http.ResponseWriter, r *http.Request) []Names {
// 	db := dbConn()
// 	defer db.Close()
// 	que, err := db.Query("SELECT * FROM names")
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 	}
//
// 	n := Names{}
// 	res := []Names{}
//
// 	for que.Next() {
// 		var name string
//
// 		err := que.Scan(&name)
// 		if err != nil {
// 			fmt.Println("que.Scan error: ", err)
// 		}
// 		n.Name = name
// 		res = append(res, n)
//
// 	}
// 	return res
// }

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
