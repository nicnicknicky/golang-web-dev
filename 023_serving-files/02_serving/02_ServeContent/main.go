package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	// https://golang.org/pkg/os/#example_OpenFile
	// https://golang.org/pkg/os/#FileInfo
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	// https: //golang.org/pkg/net/http/#ServeContent
	// os.File has method Read and Seek
	// implements io.Readseeker interface
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
