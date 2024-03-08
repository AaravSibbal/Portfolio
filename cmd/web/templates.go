package main

import (
	"html/template"
	"path/filepath"

	"github.com/AaravSibbal/Portfolio/pkg/forms"
)

type templateData struct {
	Form  *forms.Form
	Flash string
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		fileName := filepath.Base(page)

		ts, err := template.New(fileName).ParseFiles(page)

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))

		if err != nil {
			return nil, err
		}

		cache[fileName] = ts

	}
	return cache, nil
}