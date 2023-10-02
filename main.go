package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vinicius-n4/golang-api-study/database"
	"net/http"
)

type Item struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
}

var respMessage = make(map[string]string)

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
	var data []Item
	database.DB.Find(&data)

	json.NewEncoder(w).Encode(data)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data = Item{
		Name:     r.FormValue("name"),
		Document: r.FormValue("document"),
	}
	database.DB.Create(&data)

	json.NewEncoder(w).Encode(data)
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var data Item

	err := database.DB.First(&data, id)
	if err.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		respMessage["message"] = http.StatusText(http.StatusBadRequest) +
			": ID " + id + " doesn't exist. Try to list items before update them."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	data.Name = r.FormValue("name")
	data.Document = r.FormValue("document")
	database.DB.Save(&data)

	json.NewEncoder(w).Encode(data)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var data Item

	err := database.DB.First(&data, id)
	if err.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		respMessage["message"] = http.StatusText(http.StatusBadRequest) +
			": ID " + id + " doesn't exist. Try to list items before delete them."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	database.DB.Delete(&data, id)

	json.NewEncoder(w).Encode(data)
}
