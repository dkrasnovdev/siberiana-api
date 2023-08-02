package minio

import (
	"fmt"
	"net/http"

	"github.com/minio/minio-go/v7"
)

func NewServer(client *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}
}
