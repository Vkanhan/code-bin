package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Vkanhan/code-bin/internal/models"
)

// templateData act as holds dynamic data that will be passed to html templates
type templateData struct {
	CurrentYear int
	Gist        *models.Gist
	Gists       []*models.Gist
}

func (app *application) newTemplateData() *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
		Gists:       make([]*models.Gist, 0),
	}
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
