package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define a struct to match the JSON response
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	// FUNCTION TO DO GET REQUEST IN GOLANG

	// Define the API endpoint
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Parse the JSON response
	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Print the parsed data
	fmt.Printf("UserID: %d\nID: %d\nTitle: %s\nBody: %s\n", post.UserID, post.ID, post.Title, post.Body)
}
