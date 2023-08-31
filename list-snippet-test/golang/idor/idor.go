package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// User struct
type User struct {
	Name     string
	Password string
	Email    string
}

// Mock users data
var users = map[int]User{
	1: {"Alice", "password1", "alice@example.com"},
	2: {"Bob", "password2", "bob@example.com"},
}

func main() {
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Path[len("/user/"):])
		if user, ok := users[id]; ok {
			fmt.Fprintf(w, "Name: %s, Email: %s\n", user.Name, user.Email)
		} else {
			http.Error(w, "User not found", http.StatusNotFound)
		}
	})

	http.ListenAndServe(":8080", nil)
}
