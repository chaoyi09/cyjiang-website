package main

import (
	"log"
	"net/http"
	"github.com/cyjiang-website/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	// Static file service
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page routing
	mux.HandleFunc("/", handler.HomeHandler)

	log.Println("Server starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
