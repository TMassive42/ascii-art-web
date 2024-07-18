package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"asciiartserver/asciiart" // Adjust the import path based on your project structure
)

var Tmpl *template.Template

type PageData struct {
	Art   string
	Error string
}

func RenderTemplate(w http.ResponseWriter, templateFile string, data *PageData) {
	var err error

	// Parse the template file
	Tmpl, err = template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
	}
	if Tmpl == nil {
		log.Println("Template file not found")
		http.Error(w, "Template file not found", http.StatusNotFound)
		return
	}
	if err := Tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		log.Println("Template executed successfully")
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// If not a POST request, just render the form
		data := &PageData{}
		RenderTemplate(w, "templates/index.html", data)
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
	RenderTemplate(w, "templates/index.html", data)
}

func handleError(w http.ResponseWriter, data *PageData, statusCode int, errMsg string, logMsg string) {
	data.Error = errMsg
	log.Println(logMsg)
	// Set the status code here
	w.WriteHeader(statusCode)
	// Render the appropriate template after setting the status code
	if statusCode == http.StatusNotFound {
		RenderTemplate(w, "templates/404.html", data)
	} else {
		RenderTemplate(w, "templates/index.html", data)
	}
}
