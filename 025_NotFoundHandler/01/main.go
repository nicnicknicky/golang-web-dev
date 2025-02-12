package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	// https://golang.org/pkg/net/http/#NotFoundHandler
	// https://golang.org/pkg/net/http/#Error
	// HTTP Status Codes
	// https://golang.org/pkg/net/http/#pkg-constants
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go look at your terminal")
}
