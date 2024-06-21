package main

import (
	"html/template"
	"log"
	"net/http"
	"Server/asciiart" // Import the asciiart package
)

type PageData struct {
	Art string
}

var tmpl *template.Template

func main() {
	var err error
	// Parse the template file
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Define the handler function for the root path
	http.HandleFunc("/", asciiArtHandler)

	// Serve other static files (e.g., CSS, JS) using FileServer
	staticDir := http.Dir("static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// If not a POST request, just render the form
		data := &PageData{}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
		} else {
			log.Println("Template executed successfully")
		}
		return
	}

	input := r.FormValue("input")
	banner := r.FormValue("banner")

	art, err := asciiart.GenerateASCIIArt(input, banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &PageData{
		Art: art,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	} else {
		log.Println("Template executed successfully")
	}
}
