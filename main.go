package main

import (
	"log"
	"net/http"

	"github.com/rohitmj/goshorty/handlers"
)

func main() {
	// Serve static files from the /static path (for future use: CSS, JS, images)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route for homepage or redirect
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If path is "/", serve index.html
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./static/index.html")
			return
		}
		// Otherwise, handle as redirect (short URL)
		handlers.ResolveURL(w, r)
	})

	// API endpoint to shorten URLs
	http.HandleFunc("/shorten", handlers.ShortenURL)

	log.Println("🚀 Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
