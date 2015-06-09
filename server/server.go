package server

import (
	"fmt"
	"net/http"
	"flag"

	"github.com/gorilla/mux"
)

var hostname = flag.String("hostname", "localhost", "The hostname - defaults to localhost")

func Serve() {
	fmt.Println("Starting webserver on :80")
	fmt.Println(*hostname)
	r := mux.NewRouter()
	s := r.Host(*hostname).Subrouter()
	s.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	s = r.Host("blog." + *hostname).Subrouter()
	s.PathPrefix("/").Handler(http.FileServer(http.Dir("./blog/static/")))
	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}