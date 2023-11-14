package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Privacy policy of the Collection.
func (Collection) Policy() ent.Policy {
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

// Mixin of the Collection.
func (Collection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
		mixin.ImagesMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Annotations of the Collection.
func (Collection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").
			Unique(),
		field.Enum("type").
			Values("art", "artifacts", "books", "protected_area_pictures", "petroglyphs").
			Optional().
			Immutable(),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("art", Art.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("artifacts", Artifact.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("petroglyphs", Petroglyph.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("books", Book.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("protected_area_pictures", ProtectedAreaPicture.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("category", Category.Type).Ref("collections").Unique().Required(),
		edge.From("authors", Person.Type).Ref("collections"),
	}
}
