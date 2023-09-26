package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Proxy holds the schema definition for the Proxy entity.
type Proxy struct {
	ent.Schema
}

// Fields of the Proxy.
func (Proxy) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("artifacts", "books", "protected_area_pictures").
			Immutable(),
		field.String("ref_id"),
		field.String("url"),
	}
}

// Edges of the Proxy.
func (Proxy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("favourite", Favourite.Type).
			Ref("proxies").
			Unique(),
		edge.From("personal", Personal.Type).
			Ref("proxies").
			Unique(),
	}
}
