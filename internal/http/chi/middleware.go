package chi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/dkrasnovdev/heritage-api/config"
	"github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/pkg/response"
)

// CacheTTL specifies the time-to-live for cached user information.
const CacheTTL = 5 * time.Minute

// userCacheEntry represents an entry in the user cache with an expiration time.
type userCacheEntry struct {
	value      *privacy.UserViewer
	expiration time.Time
}

var userCache sync.Map // Global cache to store user information

// Authentication is a middleware function that performs authentication
// by validating the authorization token provided in the request header.
// If the token is in valid JWT format, it retrieves user information from an external endpoint
// and adds it to the request context before passing the request to the next handler.
// If the token is not in valid JWT format or an error occurs during the authentication process,
// it returns an error response and does not proceed with further checks.
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the authorization token from the request header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// No token provided, proceed to the next handler
			next.ServeHTTP(w, r)
			return
		}

		// Check if the token is in the correct format (Bearer <token>) and parse the token
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || parts[0] != "Bearer" {
			// Invalid header format, return an error response
			response.SendErrorResponse(w, http.StatusUnauthorized, "Invalid Authorization Header Format")
			return
		}
		token := parts[1]

		// Check if user information is already cached for the token
		if entry, ok := userCache.Load(token); ok {
			cacheEntry := entry.(*userCacheEntry)
			// Check if the cached entry has not expired
			if time.Now().Before(cacheEntry.expiration) {
				// User information found in the cache, create a new context with it
				ctx := privacy.NewContext(r.Context(), cacheEntry.value)
				// Pass the modified request with the cached context to the next handler
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			// Entry has expired, remove it from the cache
			userCache.Delete(token)
		}

		// Load the configuration from the .env file.
		config, err := config.LoadConfig()
		if err != nil {
			log.Println("Error loading .env file:", err)
			response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// If user information is not cached or has expired, retrieve it from the external endpoint
		client := &http.Client{}
		req, err := http.NewRequest("GET", config.OIDC_USERINFO_ENDPOINT, nil)
		if err != nil {
			log.Println("Error creating request:", err)
			response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Set the authorization token in the request header
		req.Header.Add("Authorization", authHeader)

		// Send the request to the external endpoint
		res, err := client.Do(req)
		if err != nil {
			log.Println("Error sending request:", err)
			response.SendErrorResponse(w, http.StatusInternalServerError, "Authentication Failed")
			return
		}
		defer res.Body.Close()

		// If the response status code indicates an error
		if res.StatusCode < 200 || res.StatusCode >= 300 {
			log.Println("External endpoint returned an error:", res.Status)
			response.SendErrorResponse(w, http.StatusUnauthorized, "Authentication Failed")
			return
		}

		// Read the response body
		ior, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Unmarshal the response body into a UserViewer struct
		var v *privacy.UserViewer
		err = json.Unmarshal(ior, &v)
		if err != nil {
			log.Println("Error unmarshaling response body:", err)
			response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Cache the user information with the current time as expiration
		cacheEntry := &userCacheEntry{
			value:      v,
			expiration: time.Now().Add(CacheTTL),
		}
		userCache.Store(token, cacheEntry)

		// Create a new context with the UserViewer and add it to the request context
		ctx := privacy.NewContext(r.Context(), v)
		// Pass the modified request with the new context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
