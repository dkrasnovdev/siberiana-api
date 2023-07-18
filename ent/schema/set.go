package schema

import "entgo.io/ent"

// Set holds the schema definition for the Set entity.
type Set struct {
	ent.Schema
}

// Fields of the Set.
func (Set) Fields() []ent.Field {
	return nil
}

// Edges of the Set.
func (Set) Edges() []ent.Edge {
	return nil
}
