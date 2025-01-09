package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	// debug.Stack() captures the current call stack as a string for debugging purposes.
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//ensure that the stack trace points to where the error was logged and not the location of the Output method itself
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
