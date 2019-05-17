package main

import (
	"net/http"
	"text/template"
)

////////////////////    Application related materials
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/navbar.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

///////////////////    Manuscript Handlers
func (app *application) ListManuscripts(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/manuscript.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/navbar.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
