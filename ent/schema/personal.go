package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/mixin"
)

// Personal holds the schema definition for the Personal entity.
type Personal struct {
	ent.Schema
}

// Mixin of the Personal.
func (Personal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.OwnerMixin{},
	}
}

// Fields of the Personal.
func (Personal) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name").
			NotEmpty(),
	}
}

// Edges of the Personal.
func (Personal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proxies", Proxy.Type),
	}
}
