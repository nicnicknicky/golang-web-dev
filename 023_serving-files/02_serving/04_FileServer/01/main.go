package main

import (
	"io"
	"net/http"
)

func main() {
	// https://golang.org/pkg/net/http/#FileServer
	// listing of all served files is viewable at the served path

	// https: //golang.org/pkg/net/http/#Dir
	// https://golang.org/pkg/net/http/#FileSystem
	// http.Dir has method Open which implements http.Filesystem
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`)
}
