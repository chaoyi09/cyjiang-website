package main

import (
	"log"
	"net/http"
	"github.com/cyjiang-website/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	// Static file service
	fs := http.FileServer(http.Dir("../assets/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page routing
	mux.HandleFunc("/", handler.HomeHandler)

	// Other page routing
    mux.HandleFunc("/experience", handler.ExperienceHandler)
    mux.HandleFunc("/about", handler.AboutHandler)
    mux.HandleFunc("/books", handler.BookHandler)
    mux.HandleFunc("/blog", handler.BlogHandler)
    mux.HandleFunc("/contact", handler.ContactHandler)

	log.Println("Server starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
