package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("GET /gist/view", app.gistView)
	mux.HandleFunc("POST /gist/create", app.gistCreate)
	mux.HandleFunc("GET /user/signup", app.userSignup)
	mux.HandleFunc("POST /user/signup", app.registerUser)
	mux.HandleFunc("POST /user/login", app.userLogin)
	mux.HandleFunc("POST /user/logout", app.userLogout)

	return app.logRequests(app.enableCORS(mux))
}
