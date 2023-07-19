package chi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/dkrasnovdev/heritage-api/config"
	"github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
)

// Authentication is a middleware function that performs authentication
// by validating the authorization token provided in the request header.
// If the token is valid, it retrieves user information from an external endpoint
// and adds it to the request context before passing the request to the next handler.
// If the token is invalid or an error occurs during the authentication process,
// it passes the request to the next handler without modifying the context.
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the authorization token from the request header
		token := r.Header.Get("Authorization")
		fields := strings.Fields(token)

		// Check if the token is in the correct format
		if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
			log.Println("Invalid authorization token format")
			// Token is invalid, pass the request to the next handler
			next.ServeHTTP(w, r)
			return
		}

		// Load configuration
		config, err := config.LoadConfig()
		if err != nil {
			log.Println("Error loading configuration:", err)
			// Error occurred while loading configuration, pass the request to the next handler
			next.ServeHTTP(w, r)
			return
		}

		client := &http.Client{}

		// Create a new HTTP GET request to retrieve user information from the external endpoint
		req, err := http.NewRequest("GET", config.OPENID_CONNECT_USERINFO_ENDPOINT, nil)
		if err != nil {
			log.Println("Error creating request:", err)
			// Error occurred while creating the request, pass the request to the next handler
			next.ServeHTTP(w, r)
			return
		}

		// Set the authorization token in the request header
		req.Header.Add("Authorization", token)

		// Send the request to the external endpoint
		res, err := client.Do(req)
		if err != nil {
			log.Println("Error sending request:", err)
			// Error occurred while sending the request, pass the request to the next handler
			next.ServeHTTP(w, r)
			return
		}

		defer res.Body.Close()

		// Read the response body
		ior, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			// Error occurred while reading the response body, pass the request to the next handler
			next.ServeHTTP(w, r)
			return
		}

		// Unmarshal the response body into a UserViewer struct
		var v *privacy.UserViewer
		err = json.Unmarshal(ior, &v)
		if err != nil {
			log.Println("Error unmarshaling response body:", err)
			// Error occurred while unmarshaling the response body, pass the request to the next handler
			next.ServeHTTP(w, r)
			return
		}

		// Create a new context with the UserViewer and add it to the request context
		ctx := privacy.NewContext(r.Context(), v)
		// Pass the modified request with the new context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

