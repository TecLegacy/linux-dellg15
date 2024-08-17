package main

import (
	"log"
	"net/http"
	"teclegacy/oms/common"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	mux := http.NewServeMux()

	handler := NewHandler()
	handler.RegisterRoutes(mux)

	log.Printf("Server started on port %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server")
	}

}
