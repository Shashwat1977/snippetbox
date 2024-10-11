package main

import "github.com.Shashwat1977.snippetbox/internal/models"

type templateData struct {
	Snippets []*models.Snippet
	Snippet  *models.Snippet
}
