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

// District holds the schema definition for the District entity.
type District struct {
	ent.Schema
}

// Privacy policy of the District.
func (District) Policy() ent.Policy {
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

// Mixin of the District.
func (District) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Annotations of the District.
func (District) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Edges of the District.
func (District) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("art", Art.Type),
		edge.To("artifacts", Artifact.Type),
		edge.To("books", Book.Type),
		edge.To("protected_area_pictures", ProtectedAreaPicture.Type),
		edge.To("settlements", Settlement.Type),
		edge.From("locations", Location.Type).Ref("district"),
		edge.From("region", Region.Type).Ref("districts").Unique(),
	}
}
