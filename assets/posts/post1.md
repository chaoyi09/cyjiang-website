// internal/handler/blog.go
package handler

import (
    "net/http"
    "html/template"
    "io/ioutil"
    "strings"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
    files, _ := ioutil.ReadDir("../assets/posts")
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
