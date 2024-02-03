package handlers

import (
	"encoding/json"
	"net/http"
	"tsis_1/api"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(api.Films)
}
