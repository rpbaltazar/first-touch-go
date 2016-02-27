package main

import (
	"net/http"
	"html/template"
)

func main() {
	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Defined in route_main.go
	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// TODO: define in route_auth.go
	// GET - renders sign_in page
	mux.HandleFunc("/sign_in", sign_in)
	// POST - logs the user in with given username/pwd
	mux.HandleFunc("/authenticate", authenticate)
	// GET - renders sign_up page
	mux.HandleFunc("/sign_up", sign_up)
	// POST - registers the user with the given username/pwd
	mux.HandleFunc("/register", register)
	// GET - logs the user out
	mux.HandleFunc("/logout", logout)

	// TODO: define in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thred/create", createThread)
	mux.HandleFunc("/thred/post", postThread)
	mux.HandleFunc("/thred/read", readThread)



	// starting up the server
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
