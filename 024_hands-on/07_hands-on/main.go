package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", wrapFileServerPrintURLPath("public")))
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, res *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
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
