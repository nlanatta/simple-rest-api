// main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetBooks(t *testing.T) {
	r := setup()

	// Reset the books slice to ensure a clean state for testing
	books = []Book{
		{ID: "1", Title: "Golang Basics", Author: "John Doe"},
		{ID: "2", Title: "RESTful API Design", Author: "Jane Doe"},
	}

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected HTTP status %v but got %v", http.StatusOK, status)
	}

	expected := `[{"id":"1","title":"Golang Basics","author":"John Doe"},{"id":"2","title":"RESTful API Design","author":"Jane Doe"}]
`
	got := rr.Body.String()
	if got != expected {
		t.Errorf("Expected response body %v but got %v", expected, rr.Body.String())
	}
}

func TestGetBook(t *testing.T) {
	r := setup()

	// Reset the books slice to ensure a clean state for testing
	books = []Book{
		{ID: "1", Title: "Golang Basics", Author: "John Doe"},
		{ID: "2", Title: "RESTful API Design", Author: "Jane Doe"},
	}

	tests := []struct {
		id       string
		expected string
		status   int
	}{
		{"1", `{"id":"1","title":"Golang Basics","author":"John Doe"}
`, http.StatusOK},
		{"3", `{"error":"Book not found"}
`, http.StatusNotFound},
	}

	for _, tt := range tests {
		req, err := http.NewRequest("GET", "/books/"+tt.id, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != tt.status {
			t.Errorf("For ID %v, expected HTTP status %v but got %v", tt.id, tt.status, status)
		}

		got := rr.Body.String()
		if got != tt.expected {
			t.Errorf("For ID %v, expected response body %v but got %v", tt.id, tt.expected, rr.Body.String())
		}
	}
}

func TestCreateBook(t *testing.T) {
	r := setup()

	// Reset the books slice to ensure a clean state for testing
	books = []Book{
		{ID: "1", Title: "Golang Basics", Author: "John Doe"},
		{ID: "2", Title: "RESTful API Design", Author: "Jane Doe"},
	}

	testBook := Book{ID: "3", Title: "New Book", Author: "Author Name"}

	payload, err := json.Marshal(testBook)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected HTTP status %v but got %v", http.StatusOK, status)
	}

	expected := `{"id":"3","title":"New Book","author":"Author Name"}
`
	got := rr.Body.String()
	if got != expected {
		t.Errorf("Expected response body %v but got %v", expected, rr.Body.String())
	}
}

func setup() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	return r
}
