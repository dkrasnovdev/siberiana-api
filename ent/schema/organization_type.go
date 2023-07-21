package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
)

// OrganizationType holds the schema definition for the OrganizationType entity.
type OrganizationType struct {
	ent.Schema
}

// Mixin of the OrganizationType.
func (OrganizationType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Edges of the OrganizationType.
func (OrganizationType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organizations", Organization.Type),
	}
}
