package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
)

// HolderResponsibility holds the schema definition for the HolderResponsibility entity.
type HolderResponsibility struct {
	ent.Schema
}

// Mixin of the HolderResponsibility.
func (HolderResponsibility) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Edges of the HolderResponsibility.
func (HolderResponsibility) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("holder", Holder.Type).Ref("holder_responsibilities"),
	}
}
