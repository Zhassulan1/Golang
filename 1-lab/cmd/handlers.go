package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	findfilm "tsis_1/pkg/find-film"
	searchengine "tsis_1/pkg/search-engine"

	"github.com/gorilla/mux"
)

const NOT_FOUND = "Wrong request or film not found"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is working"))
}

func GetFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(findfilm.Films)
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)

	filmIdStr, ok := argv["Id"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}
	filmId, err := strconv.Atoi(filmIdStr)
	if err != nil {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}
	if filmId < 0 || filmId >= len(findfilm.Films) {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}
	filmData := findfilm.Films[filmId]
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filmData)
}

func getFilmByGenre(genre string) ([]findfilm.Film, bool) {
	var films []findfilm.Film
	for i := range findfilm.Films {
		if findfilm.Films[i].Genre == genre {
			films = append(films, findfilm.Films[i])
		}
	}
	return films, len(films) > 0

}

func GetFilmByGenre(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	genre, ok := argv["Genre"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	films, ok := getFilmByGenre(genre)
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(films)
}

func GetFilmsByName(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	name, ok := argv["Name"]
	if !ok {
		http.Error(w, "Could not find any film", http.StatusBadRequest)
		return
	}

	var films []findfilm.Film
	for i := range findfilm.Films {
		if findfilm.Films[i].Name == name {
			films = append(films, findfilm.Films[i])
		}
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(films)
}

func SearchFilms(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	searchText, ok := argv["SearchText"]
	if !ok {
		http.Error(w, "Could not find any film", http.StatusBadRequest)
		return
	}

	result := searchengine.Search(searchText)
	if result == nil {
		http.Error(w, "Could not find any film", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
