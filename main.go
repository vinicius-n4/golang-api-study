package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/vinicius-n4/golang-api-study/database"
)

type Item struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
}

var respMessage = make(map[string]string)

func main() {
	loadEnv()
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

	for i := range data {
		data[i].Document = formatDocument(data[i].Document)
	}

	json.NewEncoder(w).Encode(data)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.FormValue("name") == "" || r.FormValue("document") == "" {
		w.WriteHeader(http.StatusBadRequest)
		respMessage["message"] = http.StatusText(http.StatusBadRequest) +
			": 'name' and 'document' fields mustn't be empty."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	if len(r.FormValue("document")) != 11 {
		w.WriteHeader(http.StatusBadRequest)
		respMessage["message"] = http.StatusText(http.StatusBadRequest) +
			": 'document' field must be 11 characters, instead of " +
			strconv.Itoa(len(r.FormValue("document"))) + "."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	var data = Item{
		Name:     r.FormValue("name"),
		Document: r.FormValue("document"),
	}
	database.DB.Create(&data)

	data.Document = formatDocument(data.Document)

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

	if r.FormValue("name") == "" && r.FormValue("document") == "" {
		w.WriteHeader(http.StatusBadRequest)
		respMessage["message"] = http.StatusText(http.StatusBadRequest) +
			": 'name' or 'document' field mustn't be empty."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	if r.FormValue("name") != "" {
		data.Name = r.FormValue("name")
	}

	if r.FormValue("document") != "" {
		data.Document = r.FormValue("document")

		if len(r.FormValue("document")) != 11 {
			w.WriteHeader(http.StatusBadRequest)
			respMessage["message"] = http.StatusText(http.StatusBadRequest) +
				": 'document' field must be 11 characters, instead of " +
				strconv.Itoa(len(r.FormValue("document"))) + "."

			json.NewEncoder(w).Encode(respMessage)
			return
		}
	}

	database.DB.Save(&data)

	data.Document = formatDocument(data.Document)

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

	data.Document = formatDocument(data.Document)

	json.NewEncoder(w).Encode(data)
}

func formatDocument(document string) string {
	return document[:3] + "." + document[3:6] + "." + document[6:9] + "-" + document[9:]
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
