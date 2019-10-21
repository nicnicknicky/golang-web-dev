package main

import (
	"io"
	"net/http"
)

// Implements Handler with ServeHTTP method
type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

// ! DOES NOT IMPLEMENT HANDLER !
type hotcat int

// Function with required signature of http.HandleFunc handler argument
func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	// Handler 1
	var d hotdog
	// https://golang.org/pkg/net/http/#Handle
	// ! For DefaultServeMux !
	// http.Handle takes a Handler as input
	http.Handle("/dog", d)

	// https://golang.org/pkg/net/http/#HandleFunc
	// ! For DefaultServeMux !
	// http.HandleFunc takes a func(ResponseWriter, *Request) as input
	http.HandleFunc("/cat", c)

	// https://golang.org/pkg/net/http/#ListenAndServe
	// The handler is typically nil, in which case the DefaultServeMux is used
	// - Both Handle and HandleFunc can be used to add handlers to DefaultServeMux
	http.ListenAndServe(":8080", nil)
}
