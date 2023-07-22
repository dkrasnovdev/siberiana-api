package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
)

// Period holds the schema definition for the Period entity.
type Period struct {
	ent.Schema
}

// Mixin of the Period.
func (Period) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Edges of the Period.
func (Period) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artifacts", Artifact.Type),
	}
}
