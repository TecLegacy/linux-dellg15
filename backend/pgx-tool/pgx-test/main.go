package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"hello/config"
	"hello/model"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {
	// Establish database connection
	conn = config.DatabaseConnectionX()
	defer conn.Close(context.Background())

	// Create tags table
	if err := createTagsTable(conn); err != nil {
		log.Fatalf("Failed to create tags table: %v", err)
	}

	http.HandleFunc("/addTag", addTagHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTagsTable(conn *pgx.Conn) error {
	const createTagsTableSQL = `
	CREATE TABLE IF NOT EXISTS tags (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`

	_, err := conn.Exec(context.Background(), createTagsTableSQL)
	if err != nil {
		return fmt.Errorf("failed to execute create table query: %w", err)
	}

	return nil
}

func addTagHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var tag model.Tags
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = addTagToDB(conn, tag.Name)
	if err != nil {
		http.Error(w, "Failed to add tag to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Tag added successfully")
}

func addTagToDB(conn *pgx.Conn, name string) error {
	const insertTagSQL = `INSERT INTO tags (name) VALUES ($1)`
	_, err := conn.Exec(context.Background(), insertTagSQL, name)
	return err
}
