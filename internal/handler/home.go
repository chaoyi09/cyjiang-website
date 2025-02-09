package handler

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"../templates/home.html",
		"../templates/partials/navbar.html",
	))
	tmpl.Execute(w, nil)
}
