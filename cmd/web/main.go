package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /gist/view", gistView)
	mux.HandleFunc("POST /gist/create", gistCreate)

	log.Println("listening to 4000...")

	server := &http.Server{
		Handler: mux,
		Addr:    ":4000",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
