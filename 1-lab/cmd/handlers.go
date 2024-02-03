package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	findfilm "tsis_1/pkg/find-film"

	"github.com/gorilla/mux"
)


func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// status code
	w.WriteHeader(http.StatusOK)

	// updating response writer
	w.Write([]byte("API is working"))
}


func GetFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(findfilm.Films)
}


func getFilmById(filmId int) (findfilm.Film, bool) {
	for i := range findfilm.Films {
		if findfilm.Films[i].ID == filmId {
			return findfilm.Films[i], true
		}
	}
	return findfilm.Films[0], false
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

// TODO: implement Levenstein distance
func GetFilmsByName(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	name, ok := argv["Name"]
	if !ok {
		http.Error(w, "Something wrong with Film Name", http.StatusBadRequest)
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