package main

import (
	"io"
	"net/http"
)

// Implements Handler interface with ServeHTTP method
type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

// Implements Handler interface with ServeHTTP method
type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	// Handler 1, 2
	var d hotdog
	var c hotcat
	// https://golang.org/pkg/net/http/#NewServeMux
	// ServeMux is an HTTP request multiplexer
	// Matches the URL of each incoming request against a list of registered patterns
	// and calls the handler for the pattern that most closely matches the URL

	// https://golang.org/pkg/net/http/#ServeMux.ServeHTTP
	// func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
	// *http.ServeMux implements Handler interface
	// http.ListenAndServe can accept *http.ServeMux

	// https://golang.org/pkg/net/http/#ServeMux.Handle
	mux := http.NewServeMux()
	mux.Handle("/dog/", d) // matches /dog/something
	mux.Handle("/cat", c)  // matches only /cat

	http.ListenAndServe(":8080", mux)
}
