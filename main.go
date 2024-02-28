package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mvrilo/go-cpf"
	"github.com/vinicius-n4/golang-api-study/database"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	dbErr := database.DB.Find(&data)
	if dbErr.Error != nil {
		w.WriteHeader(http.StatusNoContent)
		respMessage["message"] = http.StatusText(http.StatusNoContent) +
			": Database table is empty."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	for i := range data {
		data[i].Document = formatDocument(data[i].Document)
	}

	sortedList := sortNames(data)

	json.NewEncoder(w).Encode(sortedList)
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

	_, err := cpf.Valid(r.FormValue("document"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respMessage["message"] = http.StatusText(http.StatusBadRequest) +
			": 'document' validation: " + err.Error()

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	titleCaser := cases.Title(language.English)
	var data = Item{
		Name:     titleCaser.String(r.FormValue("name")),
		Document: r.FormValue("document"),
	}

	dbErr := database.DB.Create(&data)
	if dbErr.Error != nil {
		w.WriteHeader(http.StatusFailedDependency)
		respMessage["message"] = http.StatusText(http.StatusFailedDependency) +
			": Error writing data on table."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	data.Document = formatDocument(data.Document)

	json.NewEncoder(w).Encode(data)
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var data Item

	dbErr := database.DB.First(&data, id)
	if dbErr.Error != nil {
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
		titleCaser := cases.Title(language.English)
		data.Name = titleCaser.String(r.FormValue("name"))
	}

	if r.FormValue("document") != "" {
		_, err := cpf.Valid(r.FormValue("document"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			respMessage["message"] = http.StatusText(http.StatusBadRequest) +
				": 'document' validation: " + err.Error()

			json.NewEncoder(w).Encode(respMessage)
			return
		}
		data.Document = r.FormValue("document")
	}

	dbErr = database.DB.Updates(&data)
	if dbErr.Error != nil {
		w.WriteHeader(http.StatusFailedDependency)
		respMessage["message"] = http.StatusText(http.StatusFailedDependency) +
			": Error updating data on table."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

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

	dbErr := database.DB.Delete(&data, id)
	if dbErr.Error != nil {
		w.WriteHeader(http.StatusFailedDependency)
		respMessage["message"] = http.StatusText(http.StatusFailedDependency) +
			": Error deleting data on table."

		json.NewEncoder(w).Encode(respMessage)
		return
	}

	data.Document = formatDocument(data.Document)

	json.NewEncoder(w).Encode(data)
}

func formatDocument(document string) string {
	return document[:3] + "." + document[3:6] + "." + document[6:9] + "-" + document[9:]
}

func sortNames(items []Item) []Item {
	n := len(items)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if items[j].Name > items[j+1].Name {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}
