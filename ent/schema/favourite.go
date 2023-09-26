package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/mixin"
)

// Favourite holds the schema definition for the Favourite entity.
type Favourite struct {
	ent.Schema
}

// Mixin of the Favourite.
func (Favourite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.OwnerMixin{},
	}
}

// Edges of the Favourite.
func (Favourite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proxies", Proxy.Type),
	}
}
