package main

import (
	"fmt"
	"log"
)

func main() {

	store, err := NewPostgesStore()
	if err != nil {
		log.Fatalf("Database connection failed %v", err)
	}

	fmt.Printf("posgres store running %+v \n", store)

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewApiServer(":3000", store)
	server.Run()

}
