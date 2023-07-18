package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/dkrasnovdev/heritage-api/ent"
	gql "github.com/dkrasnovdev/heritage-api/internal/graphql"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gql.NewExecutableSchema(gql.Config{
		Resolvers: &Resolver{client},
	})
}
