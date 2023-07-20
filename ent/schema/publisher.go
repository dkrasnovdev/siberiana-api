package schema

import "entgo.io/ent"

// Publisher holds the schema definition for the Publisher entity.
type Publisher struct {
	ent.Schema
}

// Fields of the Publisher.
func (Publisher) Fields() []ent.Field {
	return nil
}

// Edges of the Publisher.
func (Publisher) Edges() []ent.Edge {
	return nil
}
