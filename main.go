package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	server "asciiartserver/server"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}

	var err error

	// Parse the template file
	server.Tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
	}

	// Define the handler function for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handle valid paths
		if r.URL.Path == "/" {
			server.AsciiArtHandler(w, r)
			return
		}

		// Handle 404 for unregistered paths
		if !strings.HasPrefix(r.URL.Path, "/static/") {
			var PageNotfound *template.Template

			PageNotfound, err = template.ParseFiles("templates/error.html")
			if err != nil {
				log.Printf("Error parsing template: %v", err)
			}

			data := &server.PageData{
				Error: "Page Not Found",
			}
			w.WriteHeader(http.StatusNotFound)
			if err := PageNotfound.Execute(w, data); err != nil {
				log.Printf("Error executing template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			} else {
				log.Println("Template executed successfully")
			}

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
