package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// https://golang.org/pkg/net/http/#ResponseWriter
	// https://golang.org/pkg/net/http/#Header
	// - ResponseWriter.Header()
	// returns the header map that will be sent by WriteHeader method

	// - ResponseWriter.Write([]byte) (int, error)
	// implements writer interface
	//  Write writes the data to the connection as part of an HTTP reply

	// https://golang.org/pkg/net/http/#Header.Set
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // view raw html use plain/html
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
