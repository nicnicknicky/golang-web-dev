package main

import (
	"log"
	"net/http"
)

func main() {
	// From package http documentation:
	// As a special case, the returned file server redirects any request ending in "/index.html" to the same path, without the final "index.html"
	// All other files in the same directory as index.html will not be visibly listed
	// [ TEST ]
	// - Use Chrome, rename index.html to ind.html
	// - Use FireFox, rename index.html to ind.html but reload in private browsing window
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
