package main

import (
	"log"
	"net/http"

	"github.com/dkrasnovdev/heritage-api/internal/http/chi"
)

func main() {
	r := chi.NewRouter()
	// Start the HTTP server and listen for incoming requests on port 4000.
	log.Println("listening on :4000")
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal("http server terminated", err)
	}
}
