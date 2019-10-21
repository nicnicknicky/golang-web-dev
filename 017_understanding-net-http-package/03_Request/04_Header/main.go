package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		URL    *url.URL
		// url.Values
		// https://golang.org/pkg/net/url/#Values
		// type Values map[string][]string
		// ! keys are case-sensitive !
		Submissions map[string][]string
		// https://golang.org/pkg/net/http/#Header
		// key-value pairs in an HTTP header
		// ! HTTP defines that header names are case-insensitive !
		Header http.Header
	}{
		// https://golang.org/pkg/net/http/#Request
		req.Method,
		req.URL,
		req.Form,
		req.Header,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
