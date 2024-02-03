package main

import (
	"net/http"
	handlers "tsis_1/pkg/handlers"

	"github.com/gorilla/mux"
)

// type Guests

func main() {
	router := mux.NewRouter()

	// setting up endpoints
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/films", handlers.GetFilms).Methods("GET")
	router.HandleFunc("/films/Id/{Id}", handlers.GetFilmById).Methods("GET")
	router.HandleFunc("/films/Genre/{Genre}", handlers.GetFilmByGenre).Methods("GET")
	router.HandleFunc("/films/Search/{Name}", handlers.GetFilmsByName).Methods("GET")
	http.Handle("/", router)

	// starting listener
	http.ListenAndServe(":8080", router)
}