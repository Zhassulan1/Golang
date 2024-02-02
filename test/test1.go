package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// status code
	w.WriteHeader(http.StatusOK)

	// updating response writer
	fmt.Fprintf(w, "API is working")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	var response Response

	persons := PrepareResponse()
	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func PrepareResponse() []Person {
	var personlist []Person

	var person Person
	// setting up Person-1
	person.Id = 1
	person.FirstName = "Zh"
	person.LastName = "K"
	personlist = append(personlist, person)

	// setting up Person-2
	person.Id = 2
	person.FirstName = "Zh2"
	person.LastName = "K2"
	personlist = append(personlist, person)

	return personlist

}

func main() {
	// creating new router
	router := mux.NewRouter()

	// endpoints and handlers
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	http.Handle("/", router)

	// starting listener
	http.ListenAndServe(":8080", router)
}
