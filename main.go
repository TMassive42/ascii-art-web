package main

import (
	"log"
	"net/http"
	"strings"
	server "asciiartserver/server"
)

func main() {
	// Define the handler function for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handle valid paths
		if r.URL.Path == "/" {
			server.AsciiArtHandler(w, r)
			return
		}

		// Handle 404 for unregistered paths
		if !strings.HasPrefix(r.URL.Path, "/static/") {
			
			data := &server.PageData{
				Error: "Page Not Found",
			}
			w.WriteHeader(http.StatusNotFound)
			server.RenderTemplate(w, "templates/404.html", data)
			
			return
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

