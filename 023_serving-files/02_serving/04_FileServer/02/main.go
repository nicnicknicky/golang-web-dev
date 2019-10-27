package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	// https://golang.org/pkg/net/http/#StripPrefix
	// http.StripPrefix can modify the request
	// URL's path before the FileServer sees it
	// allows a directory on disk to be served under an alternate URL path
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

/*

./assets/toby.jpg

*/
