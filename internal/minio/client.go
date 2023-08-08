package minio

import (
	"fmt"
	"log"

	"github.com/dkrasnovdev/siberiana-api/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// NewClient creates a new MinIO client based on the provided environment configuration.
func NewClient(env config.Config) *minio.Client {
	// Determine if SSL/TLS should be used for the MinIO connection based on the configuration.
	useSSL := env.MINIO_USE_SSL == "true"

	// Create a new MinIO client with the specified options.
	client, err := minio.New(fmt.Sprintf("%s:%s", env.MINIO_ENDPOINT, env.MINIO_PORT), &minio.Options{
		Creds:  credentials.NewStaticV4(env.MINIO_ACCESS_KEY_ID, env.MINIO_SECRET_ACCESS_KEY, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err) // If an error occurs while creating the MinIO client, log and terminate the application.
	}
	return client
}
