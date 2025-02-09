// internal/handler/contact.go
package handler

import (
    "net/http"
    "html/template"
)

// ContactInfo structure
type ContactInfo struct {
    Email    string
    GitHub   string
    LinkedIn string
    Location string
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
    contact := ContactInfo{
        Email:    "your.email@example.com",
        GitHub:   "https://github.com/yourusername",
        LinkedIn: "https://linkedin.com/in/yourprofile",
        Location: "Your Location",
    }
    
    tmpl := template.Must(template.ParseFiles(
        "../templates/contact.html",
        "../templates/partials/navbar.html",
    ))
    tmpl.Execute(w, contact)
}
