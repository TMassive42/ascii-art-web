package main

import (
	"html/template"
	"log"
	"net/http"
	"asciiartserver/server"
)

func main() {

	var err error

	// Parse the template file
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
	}

	// Define the handler function for the root path
	http.HandleFunc("/", server.AsciiArtHandler)

	// Serve other static files (e.g., CSS, JS) using FileServer
	staticDir := http.Dir("static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

var tmpl *template.Template
