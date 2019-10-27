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
	meFuncForHandler := func(w http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", "NICHOLAS")
		if err != nil {
			log.Fatalln("error executing template", err)
		}
	}
	dogFuncForHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, req.URL.Path)
	}
	rootFuncForHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to /")
	}

	// routes that don't exist will be redirected to /
	http.HandleFunc("/", rootFuncForHandler)
	http.HandleFunc("/dog", dogFuncForHandler)
	http.HandleFunc("/me", meFuncForHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
