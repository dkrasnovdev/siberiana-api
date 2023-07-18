package schema

import "entgo.io/ent"

// Region holds the schema definition for the Region entity.
type Region struct {
	ent.Schema
}

// Fields of the Region.
func (Region) Fields() []ent.Field {
	return nil
}

// Edges of the Region.
func (Region) Edges() []ent.Edge {
	return nil
}
