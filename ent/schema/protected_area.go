package schema

import "entgo.io/ent"

// ProtectedArea holds the schema definition for the ProtectedArea entity.
type ProtectedArea struct {
	ent.Schema
}

// Fields of the ProtectedArea.
func (ProtectedArea) Fields() []ent.Field {
	return nil
}

// Edges of the ProtectedArea.
func (ProtectedArea) Edges() []ent.Edge {
	return nil
}
