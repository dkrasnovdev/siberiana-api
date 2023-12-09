package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
)

// Herbarium holds the schema definition for the Herbarium entity.
type Herbarium struct {
	ent.Schema
}

// Privacy policy of the Herbarium.
func (Herbarium) Policy() ent.Policy {
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

// Mixin of the Herbarium.
func (Herbarium) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
		mixin.DraftMixin{},
		mixin.ImagesMixin{},
	}
}

// Annotations of the Herbarium.
func (Herbarium) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Herbarium.
func (Herbarium) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.String("location").Optional(),
	}
}

// Edges of the Herbarium.
func (Herbarium) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Person.Type).Ref("herbaria").Unique(),
		edge.From("familia", Familia.Type).Ref("herbaria").Unique(),
		edge.From("genus", Genus.Type).Ref("herbaria").Unique(),
		edge.From("group", Group.Type).Ref("herbaria").Unique(),
		edge.From("species", Species.Type).Ref("herbaria").Unique(),
		edge.From("collection", Collection.Type).Ref("herbaria").Unique().Required(),
		edge.From("country", Country.Type).Ref("herbaria").Unique(),
		edge.From("settlement", Settlement.Type).Ref("herbaria").Unique(),
		edge.From("district", District.Type).Ref("herbaria").Unique(),
		edge.From("region", Region.Type).Ref("herbaria").Unique(),
		edge.From("personal_collection", PersonalCollection.Type).Ref("herbaria"),
	}
}
