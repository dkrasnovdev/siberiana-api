package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
)

// PersonRole holds the schema definition for the PersonRole entity.
type PersonRole struct {
	ent.Schema
}

// Mixin of the PersonRole.
func (PersonRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Edges of the PersonRole.
func (PersonRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("person", Person.Type).Ref("person_roles"),
	}
}
