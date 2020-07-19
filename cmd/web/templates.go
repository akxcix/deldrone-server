package main

import (
	"html/template"
	"path/filepath"

	"github.com/deldrone/server/pkg/forms"
)

type templateData struct {
	CSRFToken             string
	AuthenticatedCustomer int
	AuthenticatedVendor   int
	Form                  *forms.Form
	Flash                 string
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
