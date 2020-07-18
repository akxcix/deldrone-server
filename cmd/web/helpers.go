package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) (*templateData, error) {
	if td == nil {
		td = &templateData{}
	}
	session, err := app.sessionStore.Get(r, "session-name")
	if err != nil {
		return nil, err
	}
	flashes := session.Flashes()
	if len(flashes) > 0 {
		td.Flash = flashes[0].(string)
	}

	return td, nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("Template not found: %s", name))
		return
	}

	buf := new(bytes.Buffer)
	td, err := app.addDefaultData(td, r)
	if err != nil {
		app.serverError(w, err)
	}
	err = ts.Execute(buf, td)
	if err != nil {
		app.serverError(w, err)
		return
	}
	buf.WriteTo(w)
}
