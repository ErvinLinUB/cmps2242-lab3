package main

import (
	"fmt"
	"net/http"
	"time"
)

// Task 2: Implement Required Routes

// homeHandler handles GET / - returns "Welcome to the Shapes API"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Welcome to the Shapes API"))
}

// healthHandler handles GET /health - returns "Server is running"
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Server is running"))
}

// aboutHandler handles GET /about - returns your name only
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Ervin Lin"))
}

// timeHandler handles GET /time - returns current server time
func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(currentTime))
}

// Task 3: Implement Custom Route - Greeting

// greetingHandler handles GET /greeting - returns a personalized greeting message
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get name from query parameter, default to "Guest"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	message := fmt.Sprintf("Hello, %s! Welcome to the Shapes API!", name)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(message))
}

func main() {
	// Register all routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/greeting", greetingHandler) // Add custom route

	// Start server
	port := ":8080"
	fmt.Println("Starting Shapes API server on port", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /         - Welcome message")
	fmt.Println("  GET  /health   - Health check")
	fmt.Println("  GET  /about    - About information")
	fmt.Println("  GET  /time     - Current server time")
	fmt.Println("  GET  /greeting - Personalized greeting (optional: ?name=YourName)")

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
