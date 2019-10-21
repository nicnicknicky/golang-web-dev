package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// hotdog implements <type Handler interface>
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
	// https://golang.org/pkg/net/http/#ListenAndServe
	// ListenAndServe listens on the TCP network address {addr} and then calls
	// Serve with {handler} to handle requests on incoming connections.

	// [ INVESTIGATE ] http.ListenAndServe
	// https://golang.org/src/net/http/server.go?s=96333:96388#L3068
	// - Create http.*Server
	// https://golang.org/pkg/net/http/#Server
	// - Returns with *Server method ListenAndServe call
	// [ INVESTIGATE ] *Server.ListenAndServe
	// https://golang.org/pkg/net/http/#Server.ListenAndServe
	// https://golang.org/src/net/http/server.go?#L2805
	// - Create net.Listener with net.Listen
	// https://golang.org/pkg/net/#Listen
	// - Returns with *Server.Serve(Listener)
	// [ INVESTIGATE ] *Server.Serve
	// https://golang.org/src/net/http/server.go?#L2856
	// Serve accepts incoming connections on the Listener l, creating a
	// new service goroutine for each. The service goroutines read requests and
	// then call *Server.Handler to reply to them.

	// [ Recall: TCP Server Basics! ]
	// which calls net.Listen on the TCP network address to get a Listener
	// returns with *Server method Serve call
	// that accepts incoming connections on the Listener with a new
	// service goroutine for each that
	// reads requests and calls the Handler to reply
}
