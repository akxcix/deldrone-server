package main

import (
	"html/template"
	"path/filepath"

	"github.com/iamadarshk/deldrone-server/pkg/forms"
	"github.com/iamadarshk/deldrone-server/pkg/models"
)

type templateData struct {
	CSRFToken             string
	AuthenticatedCustomer int
	AuthenticatedVendor   int
	Form                  *forms.Form
	Flash                 string
	Listing               *models.Listing
	Listings              []*models.Listing
	Vendor                *models.Vendor
	Vendors               []*models.Vendor
	Carts                 []models.Cart
	Cart                  models.Cart
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
