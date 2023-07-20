package schema

import "entgo.io/ent"

// Art holds the schema definition for the Art entity.
type Art struct {
	ent.Schema
}

// Fields of the Art.
func (Art) Fields() []ent.Field {
	return nil
}

// Edges of the Art.
func (Art) Edges() []ent.Edge {
	return nil
}
