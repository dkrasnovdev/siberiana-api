package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/dkrasnovdev/siberiana-api/ent"
	"github.com/dkrasnovdev/siberiana-api/internal/gql"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gql.NewExecutableSchema(gql.Config{
		Resolvers: &Resolver{client},
	})
}
