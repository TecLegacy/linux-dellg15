package main

import (
	"log"

	"github.com/teclegacy/mysql-ecom/cmd/apis"
)

func main() {
	srv := apis.NewAPIServer(":8080", nil)

	err := srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
