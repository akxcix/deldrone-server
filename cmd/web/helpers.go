package main

import (
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w)
		return
	}
	err := ts.Execute(w, nil)
	if err != nil {
		app.serverError(w)
		return
	}

}
