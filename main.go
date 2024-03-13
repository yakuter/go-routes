package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Yakuter's Go Routes!")
	})

	// Route with path parameter (e.g. /users/123)
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Path[len("/users/"):]
		fmt.Fprintf(w, "User ID: %s", userID)
	})

	// Route handling different methods
	http.HandleFunc("/method", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.Method)
	})

	// Route handling different specific methods
	http.HandleFunc("GET /get-method", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Request Method is: %s", r.Method)
	})

	// Route with query parameters
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("keyword")
		fmt.Fprintf(w, "Search Keyword: %s", query)
	})

	// Route with wildcard
	http.HandleFunc("/products/{id}/details", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Product ID: %s", id)
	})

	// Route with wildcard using "..."
	// This route will match any path after "/files/"
	// For example: /files/images/logo.png or /files/docs/whitepaper.pdf
	http.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.PathValue("path")
		fmt.Fprintf(w, "File Path: %s", filePath)
	})

	// Route with exact match
	// This route will only match the exact path "/exact/match" and fail for "/exact/match/anything".
	http.HandleFunc("/exact/match/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Exact match")
	})

	// Route with static file server
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	fmt.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
