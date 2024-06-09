package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/types"
	"github.com/teclegacy/golang-ecom/utils"
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
	/*
		route /api/v1/register POST
		{firstname , lastname , email , password}
		Check if user already Exists !Exists = create
	*/

	// Get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

}
