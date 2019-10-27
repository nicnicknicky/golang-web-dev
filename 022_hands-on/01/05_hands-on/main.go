package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	// HandlerFunc >
	http.Handle("/", http.HandlerFunc(rootFuncForHandler))
	http.Handle("/dog", http.HandlerFunc(dogFuncForHandler))
	http.Handle("/me", http.HandlerFunc(meFuncForHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func meFuncForHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "NICHOLAS")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
func dogFuncForHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, req.URL.Path)
}
func rootFuncForHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to \\")
}
