package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/service/auth"
	"github.com/teclegacy/golang-ecom/types"
	"github.com/teclegacy/golang-ecom/utils"
)

// Define Handler struct DI for Data access layer
// Constructor for Handler
// RegisterRoutes on Handler

type Handler struct {
	Store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		Store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	response := []byte("hello")
	w.Write(response)
}

// route /api/v1/register POST
// {firstname , lastname , email , password}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Check if user already exists
	_, err := h.Store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// Hash password
	hashPass, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Create user
	user := types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashPass,
	}

	err = h.Store.CreateUser(user)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Return success response
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}
