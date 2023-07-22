package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"github.com/dkrasnovdev/heritage-api/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
)

// Publisher holds the schema definition for the Publisher entity.
type Publisher struct {
	ent.Schema
}

// Privacy policy of the Publisher.
func (Publisher) Policy() ent.Policy {
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

// Mixin of the Publisher.
func (Publisher) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Annotations of the Publisher.
func (Publisher) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Publisher.
func (Publisher) Fields() []ent.Field {
	return nil
}

// Edges of the Publisher.
func (Publisher) Edges() []ent.Edge {
	return nil
}
