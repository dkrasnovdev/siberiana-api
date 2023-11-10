package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
)

// Art holds the schema definition for the Art entity.
type Art struct {
	ent.Schema
}

// Privacy policy of the Art.
func (Art) Policy() ent.Policy {
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

// Mixin of the Art.
func (Art) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
		mixin.ImagesMixin{},
	}
}

// Annotations of the Art.
func (Art) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Art.
func (Art) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").Optional(),
		field.String("dating").Optional(),
		field.String("dimensions").Optional(),
	}
}

// Edges of the Art.
func (Art) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Person.Type).Ref("art").Unique(),
		edge.From("art_genre", ArtGenre.Type).Ref("art"),
		edge.From("art_style", ArtStyle.Type).Ref("art"),
		edge.From("techniques", Technique.Type).Ref("art"),
		edge.From("collection", Collection.Type).Ref("art").Unique().Required(),
		edge.From("country", Country.Type).Ref("art").Unique(),
		edge.From("settlement", Settlement.Type).Ref("art").Unique(),
		edge.From("district", District.Type).Ref("art").Unique(),
		edge.From("region", Region.Type).Ref("art").Unique(),
	}
}
