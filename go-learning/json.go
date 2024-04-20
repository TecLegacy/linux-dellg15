package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload
	byteData, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed Parsing payload %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(byteData)
}
