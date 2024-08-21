package common

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, payload any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(payload)

}

func ReadJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorJSON(w http.ResponseWriter, status int, message error) {
	WriteJSON(w, status, ErrorResponse{Error: message.Error()})
}
