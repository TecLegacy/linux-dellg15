package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define a struct to hold the JSON payload
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

// Define a struct to hold the JSON response
type Response struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {
	// Define the URL and payload
	url := "https://jsonplaceholder.typicode.com/posts"
	post := Post{
		Title:  "foo",
		Body:   "bar",
		UserID: 1,
	}

	// Convert the payload to JSON
	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to make POST request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Error: Status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Parse the JSON response
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Print the response data
	fmt.Printf("ID: %d\nTitle: %s\nBody: %s\nUserID: %d\n", response.ID, response.Title, response.Body, response.UserID)
}
