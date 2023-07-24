package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"github.com/dkrasnovdev/heritage-api/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Privacy policy of the Location.
func (Location) Policy() ent.Policy {
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

// Mixin of the Location.
func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Annotations of the Location.
func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return nil
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artifacts", Artifact.Type),
		edge.To("books", Book.Type),
		edge.To("country", Country.Type).Unique(),
		edge.To("district", District.Type).Unique(),
		edge.To("settlement", Settlement.Type).Unique(),
		edge.To("region", Region.Type).Unique(),
	}
}
