package schema

import "entgo.io/ent"

// Keyword holds the schema definition for the Keyword entity.
type Keyword struct {
	ent.Schema
}

// Fields of the Keyword.
func (Keyword) Fields() []ent.Field {
	return nil
}

// Edges of the Keyword.
func (Keyword) Edges() []ent.Edge {
	return nil
}
