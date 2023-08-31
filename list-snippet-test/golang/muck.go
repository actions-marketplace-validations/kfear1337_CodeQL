package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:password@/blogdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/search", searchBlog).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func searchBlog(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	var blog Blog
	result, err := db.Query("SELECT * FROM blogs WHERE title LIKE '%" + query + "%'")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&blog.ID, &blog.Title, &blog.Content)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(blog)
}
