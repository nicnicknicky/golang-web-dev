package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	// Disambiguation - HandleFunc ( function ), HandlerFunc ( type )
	// https://golang.org/pkg/net/http/#HandlerFunc
	// type HandlerFunc is an adapter to allow ordinary functions
	// to be used as HTTP handlers. With the right signature,
	// HandlerFunc(f) is a Handler that calls the f
	// has method ServeHTTP which means it implements the Handler interface
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

// type casting / conversion
// https://play.golang.org/p/X2dlgVSIrd
// https://play.golang.org/p/YaUYR63b7L

// 3 Methods of Adding Handlers to DefaultServeMux
// - http.Handle: user-defined type with <method ServeHTTP> that implements the Handler interface
// - http.Handle + http.HandlerFunc:
//   ordinary function that takes arguments http.ResponseWriter, *http.Request
//   converted to http.Handler type with http.HandlerFunc
// - http.HandleFunc: directly pass in ordinary function that takes arguments http.ResponseWriter, *http.Request
