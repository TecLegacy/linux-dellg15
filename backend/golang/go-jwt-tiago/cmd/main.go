package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/teclegacy/golang-ecom/cmd/api"
	"github.com/teclegacy/golang-ecom/config"
	db_ "github.com/teclegacy/golang-ecom/db"
)

func main() {

	db, err := db_.NewMySqlStorage(mysql.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr:   config.Envs.DBAddress,
		DBName: config.Envs.DBName,

		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	//Establish connection with MYSQL DB
	// Ping it
	db_.InitStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
