package main

import (
	"io"
	"net/http"
)

type hotdog int

// [ Unelegant ] switch-case on URL path
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// https://golang.org/pkg/net/url/#URL
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
