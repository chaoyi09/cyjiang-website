// internal/handler/about.go
package handler

import (
    "net/http"
    "html/template"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        "../templates/about.html",
        "../templates/partials/navbar.html",
    ))
    tmpl.Execute(w, nil)
}
