package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	io.WriteString(w, "foo ran!")
}
