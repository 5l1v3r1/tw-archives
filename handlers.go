package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

// HandleNotFound is a handler for 404 not found.
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	template.Must(template.ParseFiles("templates/notfound.tmpl")).Execute(w, nil)
}

// IndexHandler provides top page.
func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl := template.Must(
		template.New("").ParseFiles(
			"templates/base.tmpl",
			"templates/index.tmpl",
		),
	)
	tmpl.ExecuteTemplate(w, "base", nil)
}

// HealthCheckHandler provides healthcheck page.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
}

// TwitterConnectHandler provides connecting to twitter.
func TwitterConnectHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
}
