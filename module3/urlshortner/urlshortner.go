package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var (
	urls = make(map[string]string)
	mu   sync.Mutex
)

func shortenHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.URL.Query().Get("url")

	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Generate a simple short code
	shortCode := fmt.Sprintf("%d", len(urls)+1)

	mu.Lock()
	urls[shortCode] = longURL
	mu.Unlock()

	shortURL := "http://localhost:8080/" + shortCode

	fmt.Fprintln(w, shortURL)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

	shortCode := strings.TrimPrefix(r.URL.Path, "/")

	mu.Lock()
	longURL, exists := urls[shortCode]
	mu.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {

	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

// How it works

// Start the server:

// go run urlshortener.go

// Then send:

// POST http://localhost:8080/shorten?url=https://google.com

// It returns something like:

// http://localhost:8080/1

// Now open:

// http://localhost:8080/1

// The server looks up 1:

// 1 → https://google.com

// and redirects you to Google.
