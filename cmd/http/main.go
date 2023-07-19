package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dkrasnovdev/heritage-api/config"
	"github.com/dkrasnovdev/heritage-api/internal/http/chi"
	"github.com/dkrasnovdev/heritage-api/internal/http/ent"
	"github.com/dkrasnovdev/heritage-api/pkg/shutdown"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// Load the configuration from the .env file.
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		exitCode = 1
		return
	}

	// Create a new internal chi router.
	r := chi.NewRouter()

	// Create a new ent client using the loaded configuration.
	client := ent.NewClient(config)
	defer client.Close()

	// Create a new GraphQL server using the ent client.
	gql := ent.NewGraphQLServer(client)

	// Mount the GraphQL server on the "/graphql" route of the chi router.
	r.Mount("/graphql", gql)

	// Mount the Apollo Sandbox on the "/" route of the chi router.
	r.Mount("/", playground.ApolloSandboxHandler("Heritage GraphQL API", "/graphql"))

	// Start the HTTP server and listen for incoming requests on port 4000.
	server := &http.Server{Addr: ":4000", Handler: r}
	go func() {
		log.Println("Listening on :4000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("HTTP server terminated with error:", err)
			exitCode = 1
			return
		}
	}()

	// Wait for a graceful shutdown signal (SIGINT or SIGTERM).
	shutdown.HandleGracefulShutdown()

	// Set a timeout for shutting down the server gracefully.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server within the given timeout.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error shutting down the server:", err)
		exitCode = 1
		return
	}
}
