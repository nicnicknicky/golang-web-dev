package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// wants toby.jpg in the root directory
	io.WriteString(w, `
	<!--image doesn't serve-->
	<img src="/toby.jpg">
	`)
}
