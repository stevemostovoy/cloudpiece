package server

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"cloudpiece/model"
)

var hostname = flag.String("hostname", "localhost", "The hostname to use for routing.")

func Serve() {
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	app := r.Host(*hostname).Subrouter()
	app.HandleFunc("/", App)
	app.HandleFunc("/versions/{component:[a-zA-Z-]", Versions)

	blog := r.Host("blog." + *hostname).Subrouter()
	blog.HandleFunc("/", Blog)
	blog.HandleFunc("/posts", Posts)

	http.Handle("/", r)

	fmt.Println("Starting cloudpiece server on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func App(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/app.html")
}

func Versions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	component := vars["component"]
	model.ReadComponentVersions(component)
}

func Blog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/blog.html")
}

func Posts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, JSON(model.ReadPosts()))
}

func JSON(value interface{}) string {
	json, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	return string(json)
}
