// internal/handler/experience.go
package handler

import (
    "net/http"
    "html/template"
)

func ExperienceHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        "../templates/experience.html",
        "../templates/partials/navbar.html",
    ))
    tmpl.Execute(w, nil)
}
