// handlers.go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetBooks returns the list of all books.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(books); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error encoding response")
		return
	}
}

// GetBook returns a single book by ID.
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				respondWithError(w, http.StatusInternalServerError, "Error encoding response")
			}
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "Book not found")
}

// CreateBook adds a new book to the list.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	books = append(books, book)

	if err := json.NewEncoder(w).Encode(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error encoding response")
	}
}
