package schema

import "entgo.io/ent"

// Publication holds the schema definition for the Publication entity.
type Publication struct {
	ent.Schema
}

// Fields of the Publication.
func (Publication) Fields() []ent.Field {
	return nil
}

// Edges of the Publication.
func (Publication) Edges() []ent.Edge {
	return nil
}
