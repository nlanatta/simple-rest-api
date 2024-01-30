package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct represents a simple data model.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	// Initialize router
	router := mux.NewRouter()

	// Sample data
	books = []Book{
		{ID: "1", Title: "Golang Basics", Author: "John Doe"},
		{ID: "2", Title: "RESTful API Design", Author: "Jane Doe"},
	}

	// Define routes
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
