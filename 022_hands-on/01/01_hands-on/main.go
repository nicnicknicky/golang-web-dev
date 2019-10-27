package main

import (
	"io"
	"log"
	"net/http"
)

// Guidelines
// - ListenAndServe with DefaultServeMux
// - Use HandleFunc to add routes:
// 	- /
//	- /dog/
// 	- /me/

func main() {
	meFuncForHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "MY NAME IS...\n")
	}
	dogFuncForHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, req.URL.Path)
	}
	rootFuncForHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, req.URL.Path)
	}

	http.HandleFunc("/", rootFuncForHandler)
	http.HandleFunc("/dog/", dogFuncForHandler)
	http.HandleFunc("/me/", meFuncForHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
