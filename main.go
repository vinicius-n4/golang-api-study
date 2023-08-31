package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vinicius-n4/golang-api-study/database"
	"net/http"
)

type Item struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// TODO: error message: var respMessage = make(map[string]string)

func main() {
	database.ConnectToDatabase()

	router := mux.NewRouter()
	router.HandleFunc("/list", listItemsHandler).Methods("GET")
	router.HandleFunc("/create", createItemHandler).Methods("POST")
	router.HandleFunc("/update/{id}", updateItemHandler).Methods("PUT")
	router.HandleFunc("/delete/{id}", deleteItemHandler).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var name []Item
	database.DB.Find(&name)

	json.NewEncoder(w).Encode(name)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var name = Item{Name: r.FormValue("name")}
	database.DB.Create(&name)

	json.NewEncoder(w).Encode(name)
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var name Item
	database.DB.First(&name, id)
	name.Name = r.FormValue("name")
	database.DB.Save(&name)
	// TODO: implement inexistent id validation (record not found)

	json.NewEncoder(w).Encode(name)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var name Item
	database.DB.Delete(&name, id)
	// TODO: implement inexistent id validation (record not found)

	json.NewEncoder(w).Encode(name)
}
