package main

import (
	"fmt"
	"net/http"
	"reflect"
)

// https://golang.org/pkg/net/http/#Handler
// A Handler responds to a HTTP request
// Recommended Practices:
// 1. Read from Request.Body
// 2. Write to ReponseWriter

type hotdog int // demonstrate that we can use any underlying type

// by attaching the method ServeHTTP(w http.ResponseWriter, req *http.Request)
// implicitly implements <type Handler interface>
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// reply headers and data should be written to the ResponseWriter
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var testHotdog hotdog
	implementsHandler(testHotdog)
}

// TODO: BONUS - Check interface type implementation
func implementsHandler(i interface{}) {
	iVal := reflect.TypeOf(i)
	httpHandlerType := reflect.TypeOf((*http.Handler)(nil)).Elem()
	if iVal.Implements(httpHandlerType) {
		fmt.Println("YES!")
	}
}
