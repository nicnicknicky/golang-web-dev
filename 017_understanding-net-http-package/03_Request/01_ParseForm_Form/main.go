package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// https://golang.org/pkg/net/http/#Request.ParseForm
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// https://golang.org/pkg/net/http/#Request
	// Form contains the parsed form data, including both the URL
	// field's query parameters via GET and the POST or PUT form data.
	// * This field is only available after ParseForm is called *
	// type Request struct {
	// 	...
	// 	Form url.Values
	// 	...
	// }
	// https://golang.org/pkg/net/url/#Values
	// - type Values map[string][]string

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
