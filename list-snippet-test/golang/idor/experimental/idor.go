package main

import (
	"fmt"
	"net/http"
)

type User struct {
	ID       int
	Username string
	Password string
}

var users = []User{
	{ID: 1, Username: "user1", Password: "pass1"},
	{ID: 2, Username: "user2", Password: "pass2"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

	for _, user := range users {
		if fmt.Sprint(user.ID) == userID {
			fmt.Fprintf(w, "Username: %s, Password: %s", user.Username, user.Password)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
