package config

import (
	"CrudAPIWithGin/helper"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "khushbu"
	// dbName   = "ginDemoDB"
	dbName = "gindemodb"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s  dbname=%s",
		host, port, password, user, dbName,
	)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}

// * pgx
func DatabaseConnectionX() *pgx.Conn {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	return conn
}
func CreateTagsTable(conn *pgx.Conn) error {
	const createTagsTableSQL = `
	CREATE TABLE IF NOT EXISTS tagsSX (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		name TEXT NOT NULL,
	);`

	_, err := conn.Exec(context.Background(), createTagsTableSQL)
	if err != nil {
		return fmt.Errorf("failed to execute create table query: %w", err)
	}
	log.Print("created table")
	return nil
}

// * pgxPool
// func DatabaseConnectionPool() *pgxpool.Pool {
// 	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)

// 	config, err := pgxpool.ParseConfig(connString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	dbPool, err := pgxpool.ConnectConfig(context.Background(), config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return dbPool
// }
