package ent

import (
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/dkrasnovdev/siberiana-api/ent"
	"github.com/dkrasnovdev/siberiana-api/internal/gql/resolver"
)

// NewGraphQLServer creates a new GraphQL server with the provided ent client.
func NewGraphQLServer(client *ent.Client) *handler.Server {
	// Create a new GraphQL server using gqlgen's default server configuration.
	server := handler.NewDefaultServer(resolver.NewSchema(client))

	// Use the entgql.Transactioner middleware to automatically handle transactions for each GraphQL operation.
	server.Use(entgql.Transactioner{TxOpener: client})

	// Return the configured GraphQL server.
	return server
}
