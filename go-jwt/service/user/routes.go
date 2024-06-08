package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Define Handler struct DI
// Constructor for Handler
// RegisterRoutes on Handler

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	response := []byte("hello")
	w.Write(response)
}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}
