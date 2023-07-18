package schema

import "entgo.io/ent"

// Holder holds the schema definition for the Holder entity.
type Holder struct {
	ent.Schema
}

// Fields of the Holder.
func (Holder) Fields() []ent.Field {
	return nil
}

// Edges of the Holder.
func (Holder) Edges() []ent.Edge {
	return nil
}
