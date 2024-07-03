package main

import (
	"Server/asciiart" // Import the asciiart package
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Art   string
	Error string
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
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			log.Println("Template executed successfully")
		}
		return
	}

	input := r.FormValue("input")
	banner := r.FormValue("banner")

	if input == "" || banner == "" {
		// If input or banner is missing, return a bad request
		w.WriteHeader(http.StatusBadRequest)
		data := &PageData{
			Error: "Both text input and banner selection are required.",
		}
		log.Println("Error: Missing input or banner selection")
		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	art, err := asciiart.GenerateASCIIArt(input, banner)
	data := &PageData{}
	if err != nil {
		switch err {
		case asciiart.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			data.Error = "The specified banner was not found."
			log.Printf("Error: %v", err)
		case asciiart.ErrBadRequest:
			w.WriteHeader(http.StatusBadRequest)
			data.Error = "The request was incorrect. Please check your input."
			log.Printf("Error: %v", err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			data.Error = "An internal error occurred. Please try again later."
			log.Printf("Internal error: %v", err)
		}
	} else {
		data.Art = art
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		log.Println("Template executed successfully")
	}
}
