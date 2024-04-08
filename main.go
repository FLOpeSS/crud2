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
type Names struct {
	// Id    int
	Name string
	// Email string
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
	r.HandleFunc("/new", insertUser)
	r.HandleFunc("/delete", deleteUser)

	// r.HandleFunc("/show", show)

	fmt.Println("Server ruuning at port: 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	db, err := dbConn()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer db.Close()

	que, err := db.Query("SELECT * FROM names ORDER BY id DESC")
	if err != nil {
		fmt.Println("err: ", err)
	}
	defer que.Close()

	var n Names
	var res []Names
	// n := Names{}
	// res := []Names{}

	for que.Next() {
		// var id int
		var name string
		err = que.Scan(&name)
		if err != nil {
			fmt.Println("Error: ", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		// n.Id = id
		// n.Email = email

		n.Name = name
		res = append(res, n)
	}
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

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

	response := map[string]string{"message": "User inserted"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	w.Write(jsonResponse)
}
