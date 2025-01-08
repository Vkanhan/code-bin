package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
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

	//holds dynamic gist data
	data := &templateData{
		Gists: gists,
	}

	app.renderTemplate(w, data, "./ui/html/pages/home.html")

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

	data := &templateData{
		Gist: gist,
	}

	app.renderTemplate(w, data, "./ui/html/pages/home.html")

}

func (app *application) renderTemplate(w http.ResponseWriter, data any, pageTemplate string) {
	commonFiles, err := filepath.Glob("./ui/html/{base,partials/nav}.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	files, err := filepath.Glob(pageTemplate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	files = append(commonFiles, files...)

	templ, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if err := templ.Execute(w, data); err != nil {
		app.serverError(w, err)
	}
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Signing up")
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logging in")
}

func (app *application) userLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logging out")
}