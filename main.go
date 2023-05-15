package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Item struct {
	Name string `json:"name"`
}

var items = make(map[int64]Item)

func main() {
	initData()

	router := mux.NewRouter()
	router.HandleFunc("/list", listItemsHandler).Methods("GET")
	router.HandleFunc("/create", createItemHandler).Methods("POST")
	router.HandleFunc("/update/{id}", updateItemHandler).Methods("PUT")

	http.ListenAndServe(":8000", router)
}

func initData() {
	items[0] = Item{
		"Vinicius",
	}
	items[1] = Item{
		"Nogueira",
	}
	items[2] = Item{
		"Costa",
	}
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.FormValue("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)

	nome := r.FormValue("name")
	items[idInt64] = Item{Name: nome}

	json.NewEncoder(w).Encode(items)
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	idInt64, _ := strconv.ParseInt(id, 10, 64)

	name := r.FormValue("name")
	items[idInt64] = Item{Name: name}

	json.NewEncoder(w).Encode(items[idInt64])
}
