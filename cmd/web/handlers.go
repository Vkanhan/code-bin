package main

import (
	"errors"
	"fmt"
	"html/template"
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

	for gist := range gists {
		fmt.Fprintf(w, "%+v\n", gist)
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	//holds dynamic gist data
	data := &templateData{
		Gists: gists,
	}

	for _, file := range files {
		app.renderTemplate(w, file, data)
	}

}

func (app *application) gistCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Einstein"
	content := "Photonics"
	expires := 5

	err := app.gists.Insert(title, content, expires)
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

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	data := &templateData{
		Gist: gist,
	}

	for _, file := range files {
		app.renderTemplate(w, file, data)
	}

}

func (app *application) renderTemplate(w http.ResponseWriter, templatePath string, data any) {
	templ, err := template.ParseFiles(templatePath)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if err := templ.Execute(w, data); err != nil {
		app.serverError(w, err)
	}
}
