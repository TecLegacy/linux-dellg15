package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env via joho package
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port String not found")
	}

	// Chi router setup
	router := chi.NewRouter()

	//Cors Configuration
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{
				"https://*",
				"http://*",
			},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowCredentials: false,
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			MaxAge:           300,
		},
	))

	//Route handler
	v1Router := chi.NewRouter()

	v1Router.HandleFunc("/healthz", handlerReadiness)

	// Route versioning
	router.Mount("/v1", v1Router)

	// http server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Printf("Server Started on port %v", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("server stopped", err)
	}
}
