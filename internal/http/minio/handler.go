package minio

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"

	"github.com/minio/minio-go/v7"
)

// MaxFileSize is the maximum allowed file size in bytes (2GB).
const MaxFileSize = 2 * 1024 * 1024 * 1024

// UploadResponse represents the JSON response for a successful upload.
type UploadResponse struct {
	URLs []string `json:"urls"`
}

// NewServer creates a new HTTP handler for handling file uploads using MinIO client.
func NewServer(client *minio.Client, baseURL, defaultBucket string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Display a message when a GET request is made to the server.
			fmt.Fprintf(w, "You've reached the file upload server. Use POST to upload files.")
		case http.MethodPost:
			// Parse the multipart form data for file uploads.
			err := r.ParseMultipartForm(MaxFileSize)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, "Failed to parse form data")
				return
			}

			// Get URL params for appending folders to the file names and selecting the S3 bucket.
			folder := r.URL.Query().Get("folder")
			bucket := r.URL.Query().Get("bucket")

			if bucket == "" {
				bucket = defaultBucket
			}

			// Check if the specified bucket exists.
			bucketExists, err := client.BucketExists(context.Background(), bucket)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, "Failed to check bucket existence")
				return
			}
			if !bucketExists {
				sendErrorResponse(w, http.StatusBadRequest, "Specified bucket does not exist")
				return
			}

			var wg sync.WaitGroup
			var mu sync.Mutex
			var URLs []string

			// Iterate through each file in the request and upload them concurrently.
			for _, fileHeaders := range r.MultipartForm.File {
				for _, handler := range fileHeaders {
					wg.Add(1)
					go func(handler *multipart.FileHeader) {
						defer wg.Done()

						// Open the file for reading.
						file, err := handler.Open()
						if err != nil {
							sendErrorResponse(w, http.StatusInternalServerError, "Failed to open file")
							return
						}
						defer file.Close()

						// Append folder to the file name if provided.
						objectName := handler.Filename
						if folder != "" {
							objectName = strings.TrimSuffix(folder, "/") + "/" + handler.Filename
						}

						// Perform the actual file upload.
						_, err = client.PutObject(context.Background(), bucket, objectName, file, handler.Size, minio.PutObjectOptions{
							ContentType: handler.Header.Get("Content-Type"),
						})
						if err != nil {
							sendErrorResponse(w, http.StatusInternalServerError, "Failed to upload file to MinIO")
							return
						}

						// Generate the public URL of the uploaded file.
						url := fmt.Sprintf("%s/%s/%s", baseURL, bucket, objectName)

						mu.Lock()
						URLs = append(URLs, url)
						mu.Unlock()
					}(handler)
				}
			}

			// Wait for all goroutines to finish.
			wg.Wait()

			// Send the slice of public URLs as a response.
			response := UploadResponse{
				URLs: URLs,
			}
			sendJSONResponse(w, http.StatusOK, response)
		default:
			// Return an error message for unsupported HTTP methods.
			sendErrorResponse(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		}
	}
}

// sendErrorResponse sends a JSON error response with the specified status code and error message.
func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
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

// sendJSONResponse sends a JSON response with the specified status code and payload.
func sendJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
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
