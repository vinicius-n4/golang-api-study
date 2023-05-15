package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{0, "Vinicius"},
	{1, "Naiara"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/list", listItemsHandler).Methods("GET")
	router.HandleFunc("/create", createItemHandler).Methods("POST")
	router.HandleFunc("/update/{id}", updateItemHandler).Methods("PUT")
	http.ListenAndServe(":8000", router)
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)

	var newItem Item
	json.Unmarshal(reqBody, &newItem)
	items = append(items, newItem)
	json.NewEncoder(w).Encode(items)
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, item := range items {
		i, _ := strconv.Atoi(id)
		if item.Id == i {
			items = append(items[:index], items[index+1:]...)

			var updateItem Item

			json.NewDecoder(r.Body).Decode(&updateItem)
			items = append(items, updateItem)
			json.NewEncoder(w).Encode(updateItem)
			return
		}
	}
}
