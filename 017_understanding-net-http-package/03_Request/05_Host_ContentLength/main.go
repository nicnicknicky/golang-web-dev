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
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		// https://golang.org/pkg/net/http/#Request
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		// ContentLength records the length of the associated content
		// Value == -1 indicates that the length is unknown
		// Values >= 0 indicate that the given number of bytes may
		// be read from Body
		req.ContentLength,
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
