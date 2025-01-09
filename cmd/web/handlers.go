package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Vkanhan/code-bin/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	gists, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData()
	data.Gists = gists

	app.renderTemplate(w, data, "./ui/html/pages/home.html")
}

func (app *application) gistCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Expires int    `json:"expires"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if requestBody.Title == "" || requestBody.Content == "" || requestBody.Expires <= 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.gists.Insert(requestBody.Title, requestBody.Content, requestBody.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/gist/view", http.StatusSeeOther)
}

func (app *application) gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	gist, err := app.gists.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecords) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	//to get the templateData struct
	data := app.newTemplateData()
	data.Gist = gist

	app.renderTemplate(w, data, "./ui/html/pages/home.html")
}
