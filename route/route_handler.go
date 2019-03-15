package route

import (
	"encoding/json"
	"net/http"
	"simple-rest/models"

	"simple-rest/query"

	"github.com/gorilla/mux"
)

var data []models.PokemonModel

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-Control")
}

func SearchAll(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)

	result, err := query.QueryAll()
	if err != nil {
		http.Error(w, "Search Fail", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

func SearchByName(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)

	params := mux.Vars(r)
	result, err := query.Query(params["id"])
	if err != nil {
		http.Error(w, "Search Fail", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)

	var data models.PokemonModel
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := query.Insert(&data)
	if err != nil {
		http.Error(w, "Create Fail", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("Create Success")
	return

	http.Error(w, "Not Found", http.StatusNotFound)
	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)

	var data models.PokemonModel
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := query.Update(&data)
	if err != nil {
		http.Error(w, "Update Fail", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("Update Success")
	return

}

func Delete(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)

	params := mux.Vars(r)
	err := query.Delete(params["id"])
	if err != nil {
		http.Error(w, "Delete Fail", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode("Delete Success")
	return

}
