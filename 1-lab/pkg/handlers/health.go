package handlers

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// status code
	w.WriteHeader(http.StatusOK)

	// updating response writer
	fmt.Fprintf(w, "API is working")
}
