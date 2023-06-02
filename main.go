package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
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
	router.HandleFunc("/delete/{id}", deleteItemHandler).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func initData() {
	items[1] = Item{
		"Vinicius",
	}
	items[2] = Item{
		"Nogueira",
	}
	items[3] = Item{
		"Costa",
	}
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	keys := make([]int, 0, len(items))
	for k, _ := range items {
		intKey := int(k)
		keys = append(keys, intKey)
	}
	sort.Ints(keys)

	lastIndex := len(keys) - 1
	lastIdInt64 := int64(keys[lastIndex])
	newIdInt64 := lastIdInt64 + 1

	name := r.FormValue("name")
	items[newIdInt64] = Item{Name: name}

	json.NewEncoder(w).Encode(items[newIdInt64])
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

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	idInt64, _ := strconv.ParseInt(id, 10, 64)

	delete(items, idInt64)
	json.NewEncoder(w).Encode(items[idInt64])
}
