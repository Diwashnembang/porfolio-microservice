package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, err error){
	slog.Error(fmt.Sprintf("%s",err.Error()))
	http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	slog.Error(fmt.Sprintf("CLIENT ERROR"))
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}