package schema

import "entgo.io/ent"

// PersonRole holds the schema definition for the PersonRole entity.
type PersonRole struct {
	ent.Schema
}

// Fields of the PersonRole.
func (PersonRole) Fields() []ent.Field {
	return nil
}

// Edges of the PersonRole.
func (PersonRole) Edges() []ent.Edge {
	return nil
}
