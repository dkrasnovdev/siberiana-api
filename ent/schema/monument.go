package schema

import "entgo.io/ent"

// Monument holds the schema definition for the Monument entity.
type Monument struct {
	ent.Schema
}

// Fields of the Monument.
func (Monument) Fields() []ent.Field {
	return nil
}

// Edges of the Monument.
func (Monument) Edges() []ent.Edge {
	return nil
}
