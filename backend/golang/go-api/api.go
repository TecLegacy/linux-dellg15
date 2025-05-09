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
	Error string `json:"error"`
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
	v1.HandleFunc("/transfer", makeHTTPHandleFunc(api.handleTransfer))

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

	if r.Method == "GET" {

		id, err := getID(r)
		if err != nil {
			return err
		}

		account, err := api.store.GetAccountByID(id)
		if err != nil {
			return err
		}

		RespondWithJSON(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		return api.handleDeleteAccount(w, r)
	}

	return nil
}

// Route: /v1/account/{id}
// Method: DELETE
// Description: This handler deletes the account with the specified ID.
// Response: Returns a JSON object with a success message.
func (api *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	err = api.store.DeleteAccount(id)
	if err != nil {
		return err
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Account deleted successfully"})
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
	account, err := api.store.CreateAccount(account)
	if err != nil {
		return err
	}

	RespondWithJSON(w, http.StatusOK, account)

	return nil
}

// Route: /v1/transfer
// Method: POST
// Description: This handler performs a transfer between two accounts.
// Request Body: Expects a JSON object representing the transfer details.
// Response: Returns a JSON object representing the transfer result.
func (api *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {

	transferReq := new(TransferRequest)

	err := json.NewDecoder(r.Body).Decode(transferReq)
	if err != nil {
		return err
	}

	// TODO : transfer logic with db
	// Example: api.store.Transfer(transferReq.FromAccountID, transferReq.ToAccountID, transferReq.Amount)

	RespondWithJSON(w, http.StatusOK, transferReq)

	return nil
}
