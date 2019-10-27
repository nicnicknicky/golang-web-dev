package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	// From L19 <img src="/assets/toby.jpg">:
	// URL /assets/toby.jpg > match /assets/ pattern

	// [ Case 1 ] http.StripPrefix ( /toby.jpg ) > search for file ./assets/toby.jpg  [ DOG ]
	// http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets")))) // ! TODD ORIGINAL !
	http.Handle("/assets/", http.StripPrefix("/assets", wrapFileServerPrintURLPath("./assets")))

	// [ Case 2 ] NO http.StripPrefix, search for file ./assets/assets/toby.jpg [ CAT ]
	// http.Handle("/assets/", http.FileServer(http.Dir("./assets")))
	// http.Handle("/assets/", wrapFileServerPrintURLPath("./assets"))

	// [ Case 3 ] NO http.StripPrefix, search for file ./assets/toby.jpg [ DOG ]
	// http.Handle("/assets/", wrapFileServerPrintURLPath("."))

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<h1>Check the terminal!</h1>`)
	io.WriteString(w, `<img src="/assets/toby.jpg">`)
}

// Reference functions that return handlers
// https://golang.org/pkg/net/http/#Handler
func wrapFileServerPrintURLPath(dir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Searching in %s\n%s\n", dir, r.URL.Path)
		fsHandler := http.FileServer(http.Dir(dir))
		fsHandler.ServeHTTP(w, r)
	})
}
