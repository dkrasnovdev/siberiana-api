package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"github.com/dkrasnovdev/heritage-api/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
)

// ProtectedAreaCategory holds the schema definition for the ProtectedAreaCategory entity.
type ProtectedAreaCategory struct {
	ent.Schema
}

// Privacy policy of the ProtectedAreaCategory.
func (ProtectedAreaCategory) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdministrator(),
			rule.AllowIfModerator(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

// Mixin of the ProtectedAreaCategory.
func (ProtectedAreaCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Annotations of the ProtectedAreaCategory.
func (ProtectedAreaCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the ProtectedAreaCategory.
func (ProtectedAreaCategory) Fields() []ent.Field {
	return nil
}

// Edges of the ProtectedAreaCategory.
func (ProtectedAreaCategory) Edges() []ent.Edge {
	return nil
}
