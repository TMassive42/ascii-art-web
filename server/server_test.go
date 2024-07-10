package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"html/template"
)

func TestRenderTemplate(t *testing.T) {
	Tmpl = template.Must(template.New("test").Parse(`{{if .Error}}{{.Error}}{{else}}{{.Art}}{{end}}`))

	tests := []struct {
		data       *PageData
		expected   string
		statusCode int
	}{
		{&PageData{Art: "Art Content"}, "Art Content", http.StatusOK},
		{&PageData{Error: "Error Content"}, "Error Content", http.StatusOK},
	}

	for _, tt := range tests {
		rr := httptest.NewRecorder()
		renderTemplate(rr, tt.data)

		if status := rr.Code; status != tt.statusCode {
			t.Errorf("renderTemplate returned wrong status code: got %v want %v", status, tt.statusCode)
		}

		if !strings.Contains(rr.Body.String(), tt.expected) {
			t.Errorf("renderTemplate returned unexpected body: got %v want %v", rr.Body.String(), tt.expected)
		}
	}
}

func TestHandleError(t *testing.T) {
	Tmpl = template.Must(template.New("test").Parse(`{{if .Error}}{{.Error}}{{else}}{{.Art}}{{end}}`))

	tests := []struct {
		data       *PageData
		statusCode int
		errMsg     string
		logMsg     string
		expected   string
	}{
		{&PageData{}, http.StatusNotFound, "Error: Not Found", "Error: Not Found", "Error: Not Found"},
	}

	for _, tt := range tests {
		rr := httptest.NewRecorder()
		handleError(rr, tt.data, tt.statusCode, tt.errMsg, tt.logMsg)

		if status := rr.Code; status != tt.statusCode {
			t.Errorf("handleError returned wrong status code: got %v want %v", status, tt.statusCode)
		}

		if !strings.Contains(rr.Body.String(), tt.expected) {
			t.Errorf("handleError returned unexpected body: got %v want %v", rr.Body.String(), tt.expected)
		}
	}
}
