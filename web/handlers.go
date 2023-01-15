package main

import (
	"net/http"
	"text/template"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// 404
		app.clientError(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Get", http.MethodGet)
		// 405
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		// 500
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		// 500
		app.serverError(w, err)
		return
	}
}

func (app *application) CreateAscii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		// 404
		app.notFound(w)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// 405
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		// 500
		app.serverError(w, err)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("bannerfile")
	text1, err := AsciiForWeb(text, banner)
	if err != nil {
		// 400
		app.clientError(w, http.StatusBadRequest)
		return
	}
	err = ts.Execute(w, text1)
	if err != nil {
		// 500
		app.serverError(w, err)
		return
	}
}
