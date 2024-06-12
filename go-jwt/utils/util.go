package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Client Payload Validation
var Validator = validator.New()

// ParseJSON from request body
func ParseJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("no body provided in the request")
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJSON writes JSON response
func WriteJSON(w http.ResponseWriter, status int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}

// WriteError writes error response
func WriteError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errorPayload := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	WriteJSON(w, status, errorPayload)
}
