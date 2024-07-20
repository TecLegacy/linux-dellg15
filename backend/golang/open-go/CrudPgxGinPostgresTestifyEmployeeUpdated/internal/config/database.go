package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	// password = "khushbu"
	// dbName   = "ginDemoDB"
	password = "password"  //KK docker
	dbName   = "gindemodb" //KK docker
)

func ConnectDatabase() {
	var err error
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)

	// connString := "postgres://postgres:khushbu:5432/ginDemoDB"
	DB, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Connected to database")
}
