package schema

import "entgo.io/ent"

// Technique holds the schema definition for the Technique entity.
type Technique struct {
	ent.Schema
}

// Fields of the Technique.
func (Technique) Fields() []ent.Field {
	return nil
}

// Edges of the Technique.
func (Technique) Edges() []ent.Edge {
	return nil
}
