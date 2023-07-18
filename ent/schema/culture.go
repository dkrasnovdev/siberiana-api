package schema

import "entgo.io/ent"

// Culture holds the schema definition for the Culture entity.
type Culture struct {
	ent.Schema
}

// Fields of the Culture.
func (Culture) Fields() []ent.Field {
	return nil
}

// Edges of the Culture.
func (Culture) Edges() []ent.Edge {
	return nil
}
