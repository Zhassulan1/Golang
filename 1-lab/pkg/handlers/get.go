package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tsis_1/api"

	"github.com/gorilla/mux"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(api.Films)
}

func getFilmById(filmId int) (api.Film, bool) {
	for i := range api.Films {
		if api.Films[i].ID == filmId {
			return api.Films[i], true
		}
	}
	return api.Films[0], false
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	
	filmIdStr, ok := argv["Id"]
	if !ok {
		http.Error(w, "Something wrong with Film id", http.StatusBadRequest)
		return
	}
	
	filmId, err := strconv.Atoi(filmIdStr)
	if err != nil {
		http.Error(w, "Wrong Film id", http.StatusBadRequest)
		return
	}
	
	filmData, ok := getFilmById(filmId)
	if !ok {
		http.Error(w, "Wrong Film id", http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filmData)
}

func getFilmByGenre(genre string) ([]api.Film, bool) {
	var films []api.Film
	for i := range api.Films {
		if api.Films[i].Genre == genre {
			films = append(films, api.Films[i])
		}
	}
	return films, len(films) > 0

}

func GetFilmByGenre(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	genre, ok := argv["Genre"]
	if !ok {
		http.Error(w, "Something wrong with Film Genre", http.StatusBadRequest)
		return
	}

	films, ok := getFilmByGenre(genre)
	if !ok {
		http.Error(w, "Something wrong with Film Genre", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(films)
}

func GetFilmsByName(w http.ResponseWriter, r *http.Request) {
	
}