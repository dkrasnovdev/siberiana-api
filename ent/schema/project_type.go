package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
)

// ProjectType holds the schema definition for the ProjectType entity.
type ProjectType struct {
	ent.Schema
}

// Mixin of the ProjectType.
func (ProjectType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Edges of the ProjectType.
func (ProjectType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type),
	}
}
