package route

import (
	"encoding/json"
	"net/http"

	"simple-rest/query"
)

func SearchParticipant(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(query.QueryAll())
}
