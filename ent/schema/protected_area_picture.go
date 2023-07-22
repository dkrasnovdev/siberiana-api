package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"github.com/dkrasnovdev/heritage-api/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
)

// ProtectedAreaPicture holds the schema definition for the ProtectedAreaPicture entity.
type ProtectedAreaPicture struct {
	ent.Schema
}

// Privacy policy of the ProtectedAreaPicture.
func (ProtectedAreaPicture) Policy() ent.Policy {
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

// Mixin of the ProtectedAreaPicture.
func (ProtectedAreaPicture) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Annotations of the ProtectedAreaPicture.
func (ProtectedAreaPicture) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the ProtectedAreaPicture.
func (ProtectedAreaPicture) Fields() []ent.Field {
	return nil
}

// Edges of the ProtectedAreaPicture.
func (ProtectedAreaPicture) Edges() []ent.Edge {
	return nil
}
