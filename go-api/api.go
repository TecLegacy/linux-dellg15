package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ErrorMessage struct {
	Error string
}

// Define Port Address
func NewApiServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (api *APIServer) Run() {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/account", makeHTTPHandleFunc(api.handleAccount))
	v1.HandleFunc("/account/{id}", makeHTTPHandleFunc(api.handleGetAccountByID))

	srv := http.Server{
		Addr:    api.listenAddr,
		Handler: r,

		// Enforce timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server Started on %s", api.listenAddr)
	log.Fatal(srv.ListenAndServe())
}

// Decorator Function to make http.HandleFunc
// Handle the errors from handlers
func makeHTTPHandleFunc(h apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			RespondWithJSON(w, http.StatusBadRequest, ErrorMessage{
				Error: err.Error(),
			})
		}
	}
}

// Method based Routing
func (api *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return api.handleGetAllAccounts(w, r)
	}

	if r.Method == "POST" {
		return api.handleCreateAccount(w, r)
	}

	return fmt.Errorf("method not allowed in this route : %s ", r.Method)
}

// Route: /v1/account/{id}
// Method: GET
// Description: This handler returns the account with the specified ID.
// Response: Returns a JSON object representing the account.
func (api *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("keshav", "kumar")

	id := mux.Vars(r)

	fmt.Print("Id", id)
	//TODO: query db to get account with specific id

	RespondWithJSON(w, http.StatusOK, account)

	return nil
}

// Route: /v1/account
// Method: GET
// Description: This handler returns all the accounts.
// Response: Returns a JSON array representing the accounts.
func (api *APIServer) handleGetAllAccounts(w http.ResponseWriter, _ *http.Request) error {

	accounts, err := api.store.GetAllAccounts()
	if err != nil {
		return err
	}

	RespondWithJSON(w, http.StatusOK, accounts)

	return nil
}

// Route: /v1/account
// Method: POST
// Description: This handler creates a new account.
// Request Body: Expects a JSON object representing the account details.
// Response: Returns a JSON object representing the created account.
func (api *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}

	account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)

	//Create account in PostgresDB -> DB request
	if err := api.store.CreateAccount(account); err != nil {
		return err
	}

	RespondWithJSON(w, http.StatusOK, account)

	return nil
}
