package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", Hello)
	router.HandleFunc("/status", Status)

	http.ListenAndServe(":8000", router)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested %s\n", r.URL.Path)
}

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: %v\n", http.StatusOK)
}
