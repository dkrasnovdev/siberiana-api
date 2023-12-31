package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"github.com/dkrasnovdev/siberiana-api/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
)

// Settlement holds the schema definition for the Settlement entity.
type Settlement struct {
	ent.Schema
}

// Privacy policy of the Settlement.
func (Settlement) Policy() ent.Policy {
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

// Mixin of the Settlement.
func (Settlement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Annotations of the Settlement.
func (Settlement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Edges of the Settlement.
func (Settlement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("art", Art.Type),
		edge.To("artifacts", Artifact.Type),
		edge.To("books", Book.Type),
		edge.To("herbaria", Herbarium.Type),
		edge.To("protected_area_pictures", ProtectedAreaPicture.Type),
		edge.From("locations", Location.Type).Ref("settlement"),
		edge.From("region", Region.Type).Ref("settlements").Unique(),
		edge.From("district", District.Type).Ref("settlements").Unique(),
		edge.To("known_as_before", Settlement.Type).
			From("known_as_after"),
	}
}
