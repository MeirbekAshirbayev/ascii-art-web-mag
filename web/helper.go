package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"text/template"
)

type Msg struct {
	TextErr string
	NumErr  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./ui/html/error.html"))
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	var Msg Msg
	Msg.NumErr = 500
	Msg.TextErr = http.StatusText(500)
	w.WriteHeader(500)

	err = tpl.Execute(w, Msg)
	if err != nil {
		http.Error(w, "500 | Internal Server Error!", http.StatusInternalServerError)
		return
	}
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	var Msg Msg
	Msg.NumErr = status
	Msg.TextErr = http.StatusText(status)
	w.WriteHeader(status)

	err := tpl.Execute(w, Msg)
	if err != nil {
		http.Error(w, "500 | Internal Server Error!", http.StatusInternalServerError)
		return
	}
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
