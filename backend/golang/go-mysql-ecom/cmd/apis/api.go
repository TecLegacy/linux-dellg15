package apis

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teclegacy/mysql-ecom/service/users"
)

type APIServer struct {
	ListenAddr string
	db         *sql.DB
}

func NewAPIServer(listenAddr string, db *sql.DB) *APIServer {
	return &APIServer{
		ListenAddr: listenAddr,
		db:         db,
	}
}

func (api *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := users.NewHandler()
	userHandler.RegisteringRoutes(subRouter)

	log.Println("Listening to port", api.ListenAddr)

	return http.ListenAndServe(api.ListenAddr, router)
}
