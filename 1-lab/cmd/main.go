package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// type Guests

func main() {
	router := mux.NewRouter()

	// setting up endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/films", GetFilms).Methods("GET")
	router.HandleFunc("/films/Id/{Id}", GetFilmById).Methods("GET")
	router.HandleFunc("/films/Genre/{Genre}", GetFilmByGenre).Methods("GET")
	router.HandleFunc("/films/Search/{Name}", GetFilmsByName).Methods("GET")
	http.Handle("/", router)

	// starting listener
	http.ListenAndServe(":8080", router)
}
