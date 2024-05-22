package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// Registering Routes
func (h *Handler) RegisteringRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}
