package ent

import (
	// "context"

	"entgo.io/contrib/entgql"
	// "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/dkrasnovdev/siberiana-api/ent"
	// "github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/internal/gql/resolver"
)

// NewGraphQLServer creates a new GraphQL server with the provided ent client.
func NewGraphQLServer(client *ent.Client) *handler.Server {
	// Create a new GraphQL server using gqlgen's default server configuration.
	server := handler.NewDefaultServer(resolver.NewSchema(client))

	// // Add middleware to customize GraphQL operation behavior.
	// server.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	// 	// Get the privacy view from the context.
	// 	view := privacy.FromContext(ctx)
	//
	// 	// If the view is not provided, disable introspection to prevent schema introspection.
	// 	if view == nil {
	// 		graphql.GetOperationContext(ctx).DisableIntrospection = true
	// 	}
	//
	// 	// If the user is not an administrator, disable introspection to restrict access.
	// 	if !view.IsAdministrator() {
	// 		graphql.GetOperationContext(ctx).DisableIntrospection = true
	// 	}
	//
	// 	// Continue processing the GraphQL operation.
	// 	return next(ctx)
	// })

	// Use the entgql.Transactioner middleware to automatically handle transactions for each GraphQL operation.
	server.Use(entgql.Transactioner{TxOpener: client})

	// Return the configured GraphQL server.
	return server
}

