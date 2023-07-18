package schema

import "entgo.io/ent"

// Settlement holds the schema definition for the Settlement entity.
type Settlement struct {
	ent.Schema
}

// Fields of the Settlement.
func (Settlement) Fields() []ent.Field {
	return nil
}

// Edges of the Settlement.
func (Settlement) Edges() []ent.Edge {
	return nil
}
