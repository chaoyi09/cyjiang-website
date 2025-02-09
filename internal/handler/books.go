// internal/handler/books.go
package handler

import (
    "net/http"
    "html/template"
)

// Book structure
type Book struct {
    Title    string
    Author   string
    ISBN     string
    Summary  string
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
    // Define a list of books
    books := []Book{
        {
            Title:    "The Go Programming Language",
            Author:   "Alan A. A. Donovan and Brian W. Kernighan",
            ISBN:     "978-0134190440",
            Summary:  "A comprehensive guide to Go programming language",
        },
        {
            Title:    "Learning Go",
            Author:   "Jon Bodner",
            ISBN:     "978-1492077213",
            Summary:  "An excellent resource for learning Go programming",
        },
    }
    
    tmpl := template.Must(template.ParseFiles(
        "../templates/books.html",
        "../templates/partials/navbar.html",
    ))
    tmpl.Execute(w, books)
}

