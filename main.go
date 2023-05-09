package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var db = initData()

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/list", listItemsHandler).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func initData() map[int]string {
	db := map[int]string{
		0: "Vinicius",
		1: "Naiara",
	}
	return db
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db)
}
