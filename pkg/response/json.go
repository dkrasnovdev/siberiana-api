package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendJSONResponse sends a JSON response with the specified status code and payload.
func SendJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		// If encoding the JSON response fails, send a plain-text response with the error message.
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to encode JSON response")
	}
}
