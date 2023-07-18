package schema

import "entgo.io/ent"

// District holds the schema definition for the District entity.
type District struct {
	ent.Schema
}

// Fields of the District.
func (District) Fields() []ent.Field {
	return nil
}

// Edges of the District.
func (District) Edges() []ent.Edge {
	return nil
}
