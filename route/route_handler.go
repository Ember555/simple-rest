package route

import (
	"encoding/json"
	"net/http"
	"simple-rest/models"

	"simple-rest/query"
)

var data []models.PokemonModel

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-Control")
}

func SearchByName(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		SetHeader(w)
		queryString := r.URL.Query()
		if len(queryString) < 1 {
			json.NewEncoder(w).Encode(query.QueryAll())
			return
		}
		json.NewEncoder(w).Encode(query.Query(queryString))
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data models.PokemonModel
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode("Create Success")
		query.Insert(&data)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
	return
}
