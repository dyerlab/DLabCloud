package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"text/template"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

/**********    Server Error Reporting ***************/
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFoundError(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

/************   SERVER Routing *******************/
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/graph.page.tmpl",
		//		"./ui/html/base.layout.tmpl",
		//		"./ui/html/navbar.partial.tmpl",
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
