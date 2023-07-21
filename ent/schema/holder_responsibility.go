package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// HolderResponsibility holds the schema definition for the HolderResponsibility entity.
type HolderResponsibility struct {
	ent.Schema
}

// Fields of the HolderResponsibility.
func (HolderResponsibility) Fields() []ent.Field {
	return nil
}

// Edges of the HolderResponsibility.
func (HolderResponsibility) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("holder", Holder.Type).Ref("holder_responsibilities"),
	}
}
