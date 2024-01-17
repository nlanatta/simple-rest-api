// utils.go
package main

import (
	"encoding/json"
	"net/http"
)

// respondWithError sends an HTTP error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
