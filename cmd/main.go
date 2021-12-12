package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":10080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	password := r.URL.Query().Get("pass")

	if userId == "aaa" && password == "bbb" {
		w.Write([]byte("Success!"))
	} else {
		w.Write([]byte("Login Fail!"))
	}
}
