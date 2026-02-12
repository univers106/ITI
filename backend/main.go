// Package main starts a simple HTTP server.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register a handler function for the default path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, "Hello, Docker!"); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
