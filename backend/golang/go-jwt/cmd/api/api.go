package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/service/user"
)

/**
1. api struct holding listenAdd and DB
2. apiServer constructor function to initialize the api struct
3. apiServer method to start the server
*/

type APIServer struct {
	listenAddr string
	db         *sql.DB
}

func NewAPIServer(listenAddr string, db *sql.DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *APIServer) Run() error {
	// init router with api versioning
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//User Data-access layer
	userStore := user.NewStoreRepo(s.db)

	//Register User Handlers
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	srv := &http.Server{
		Addr:    s.listenAddr,
		Handler: subrouter,
	}

	log.Printf("Server started on port %s", s.listenAddr)
	return srv.ListenAndServe()
}
