package main

import (
	"github.com/Vkanhan/code-bin/internal/models"
)

// templateData act as holds dynamic data that will be passed to html templates
type templateData struct {
	CurrentYear int
	Gist        *models.Gist
	Gists       []*models.Gist
}
