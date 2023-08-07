package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendErrorResponse sends a JSON error response with the specified status code and error message.
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]string{"error": message})
	if err != nil {
		// If encoding the JSON response fails, send a plain-text response with the error message.
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(statusCode)
		fmt.Fprint(w, message)
	}
}
