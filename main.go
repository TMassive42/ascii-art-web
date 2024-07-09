package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"Server/asciiart"
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
		log.Printf("Error parsing template: %v", err)
	}

	// Define the handler function for the root path
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		// Handle valid paths
		if r.URL.Path == "/" {
			asciiArtHandler(w, r)
			return
		}
	
		// Handle 404 for unregistered paths
		if !strings.HasPrefix(r.URL.Path, "/static/") {
			http.NotFound(w, r)
			return
		}
	
		// Serve static files (handled by FileServer)
		http.DefaultServeMux.ServeHTTP(w, r)
	})

	// Serve other static files (e.g., CSS, JS) using FileServer
	staticDir := http.Dir("static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func renderTemplate(w http.ResponseWriter, data *PageData) {
	if tmpl == nil {
		log.Println("Template file not found")
		http.Error(w, "Template file not found", http.StatusNotFound)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		log.Println("Template executed successfully")
	}
}

func handleError(w http.ResponseWriter, data *PageData, statusCode int, errMsg string, logMsg string) {
	data.Error = errMsg
	log.Println(logMsg)
	// Set the status code here
	w.WriteHeader(statusCode)
	// Render the template after setting the status code
	renderTemplate(w, data)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// If not a POST request, just render the form
		data := &PageData{}
		renderTemplate(w, data)
		return
	}

	input := r.FormValue("input")
	banner := "asciiart/banners/" + r.FormValue("banner")

	data := &PageData{}
	if input == "" || banner == "" {
		handleError(w, data, http.StatusBadRequest, "Both text input and banner selection are required.", "Error: Missing input or banner selection")
		return
	}

	art, err := asciiart.GenerateASCIIArt(input, banner)
	if err != nil {
		switch err {
		case asciiart.ErrNotFound:
			handleError(w, data, http.StatusNotFound, "The specified banner was not found.", fmt.Sprintf("Error: %v", err))
		case asciiart.ErrBadRequest:
			handleError(w, data, http.StatusBadRequest, "The request was incorrect. Please check your input.", fmt.Sprintf("Error: %v", err))
		default:
			handleError(w, data, http.StatusInternalServerError, "An internal error occurred. Please try again later.", fmt.Sprintf("Internal error: %v", err))
		}
		return
	}

	data.Art = art
	renderTemplate(w, data)
}
