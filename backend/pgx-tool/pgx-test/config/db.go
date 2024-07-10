package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

// var (
//
//	user     = "your_db_user"
//	password = "your_db_password"
//	host     = "your_db_host"
//	port     = "your_db_port"
//	dbName   = "your_db_name"
//
// )
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "khushbu"
	// dbName   = "ginDemoDB"
	dbName = "gindemodb"
)

func DatabaseConnectionX() *pgx.Conn {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn
}
