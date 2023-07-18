package schema

import "entgo.io/ent"

// Medium holds the schema definition for the Medium entity.
type Medium struct {
	ent.Schema
}

// Fields of the Medium.
func (Medium) Fields() []ent.Field {
	return nil
}

// Edges of the Medium.
func (Medium) Edges() []ent.Edge {
	return nil
}
