package main

import (
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Error code %d", code)
	}

	type errorMessage struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorMessage{
		Error: msg,
	})
}
