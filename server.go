package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Define a struct to hold the data to be passed to the template
	type PageData struct {
		Title string
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Define the handler function for the template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create an instance of PageData with the data to be passed to the template
		data := &PageData{
			Title: "ASCII Art Generator",
		}

		// Execute the template with the data
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
		} else {
			log.Println("Template executed successfully")
		}
	})

	// Serve other static files (e.g., CSS, JS) using FileServer
	staticDir := http.Dir("static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
