package main

import (
	"html/template"
	"path/filepath"

	"github.com.Shashwat1977.snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippets    []*models.Snippet
	Snippet     *models.Snippet
	Form        any
	Flash       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
