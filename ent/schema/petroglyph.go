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
	"github.com/dkrasnovdev/siberiana-api/internal/ent/types"
)

// Petroglyph holds the schema definition for the Petroglyph entity.
type Petroglyph struct {
	ent.Schema
}

// Privacy policy of the Petroglyph.
func (Petroglyph) Policy() ent.Policy {
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

// Mixin of the Petroglyph.
func (Petroglyph) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
		mixin.DraftMixin{},
		mixin.ImagesMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Annotations of the Petroglyph.
func (Petroglyph) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Petroglyph.
func (Petroglyph) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").Optional(),
		field.String("dating").Optional(),
		field.Int("dating_start").Optional(),
		field.Int("dating_end").Optional(),
		field.String("orientation").Optional(),
		field.String("position").Optional(),
		field.String("geometric_shape").Optional(),
		field.Float("height").Optional(),
		field.Float("width").Optional(),
		field.Float("length").Optional(),
		field.Float("depth").Optional(),
		field.Float("diameter").Optional(),
		field.String("weight").Optional(),
		field.String("dimensions").Optional(),
		field.String("plane_preservation").Optional(),
		field.String("photo_code").Optional(),
		field.String("accounting_documentation_information").Optional(),
		field.Time("accounting_documentation_date").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.Other("geometry", types.Geometry{}).
			Optional().
			Nillable().
			SchemaType(types.GeometrySchemaType()),
	}
}

// Edges of the Petroglyph.
func (Petroglyph) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cultural_affiliation", Culture.Type).Ref("petroglyphs").Unique(),
		edge.From("model", Model.Type).Ref("petroglyphs").Unique(),
		edge.From("mound", Mound.Type).Ref("petroglyphs").Unique(),
		edge.From("publications", Publication.Type).Ref("petroglyphs"),
		edge.From("techniques", Technique.Type).Ref("petroglyphs"),
		edge.From("region", Region.Type).Ref("petroglyphs").Unique(),
		edge.From("accounting_documentation_address", Location.Type).Ref("petroglyphs_accounting_documentation").Unique(),
		edge.From("accounting_documentation_author", Person.Type).Ref("petroglyphs_accounting_documentation").Unique(),
		edge.From("collection", Collection.Type).Ref("petroglyphs").Unique().Required(),
	}
}
