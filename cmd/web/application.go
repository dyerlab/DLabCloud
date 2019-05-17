package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/dyerlab/DLabCloud/pkg/manuscript"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type application struct {
	DB       *gorm.DB
	errorLog *log.Logger
	infoLog  *log.Logger
	articles *manuscript.Model
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
