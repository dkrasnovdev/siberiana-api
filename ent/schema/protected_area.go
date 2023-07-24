package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
)

// ProtectedArea holds the schema definition for the ProtectedArea entity.
type ProtectedArea struct {
	ent.Schema
}

// Privacy policy of the ProtectedArea.
func (ProtectedArea) Policy() ent.Policy {
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

// Mixin of the ProtectedArea.
func (ProtectedArea) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
	}
}

// Annotations of the ProtectedArea.
func (ProtectedArea) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the ProtectedArea.
func (ProtectedArea) Fields() []ent.Field {
	return []ent.Field{
		field.String("area").Optional(),
		field.Time("establishment_date").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
	}
}

// Edges of the ProtectedArea.
func (ProtectedArea) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("protected_area_pictures", ProtectedAreaPicture.Type),
		edge.From("protected_area_category", ProtectedAreaCategory.Type).Ref("protected_areas").Unique(),
	}
}
