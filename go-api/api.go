package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ErrorMessage struct {
	Error string
}

// Define Port Address
func NewApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (api *APIServer) Run() {
	r := mux.NewRouter()

	// Define handlers
	r.HandleFunc("/", makeHTTPHandleFunc(api.handleAccount))

	srv := http.Server{
		Addr:    api.listenAddr,
		Handler: r,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server Started on %s", api.listenAddr)
	log.Fatal(srv.ListenAndServe())

	// http.ListenAndServe(api.listenAddr, r)
}

// Decorator Function to make http.HandleFunc
func makeHTTPHandleFunc(h apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			RespondWithJSON(w, 200, ErrorMessage{
				Error: err.Error(),
			})
		}
	}
}

func (api *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		api.handleGetAccount(w, r)
	}

	return nil
}

func (api *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Gorilla!\n"))

	return nil
}
