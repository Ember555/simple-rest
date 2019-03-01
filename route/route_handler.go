package route

import (
	"encoding/json"
	"net/http"
	"simple-rest/models"

	"simple-rest/query"
)

var data []models.PokemonModel

func SearchAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(query.QueryAll())
}

func SearchByName(w http.ResponseWriter, r *http.Request) {
	name, ok := r.URL.Query()["name"]
	if !ok || len(name[0]) < 1 {
		json.NewEncoder(w).Encode("name requires")
		return
	}
	json.NewEncoder(w).Encode(query.Query(name[0]))
}
