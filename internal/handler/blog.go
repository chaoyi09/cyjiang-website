// internal/handler/blog.go
package handler

import (
    "net/http"
    "html/template"
    "os"
    "strings"
)

type Post struct {
    Title string
    Content string
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
    files, _ := os.ReadDir("../assets/posts")
    var posts []string
    for _, file := range files {
        posts = append(posts, strings.TrimSuffix(file.Name(), ".md"))
    }
    
    tmpl := template.Must(template.ParseFiles(
        "../templates/blog.html",
        "../templates/partials/navbar.html",
    ))
    tmpl.Execute(w, posts)
}
