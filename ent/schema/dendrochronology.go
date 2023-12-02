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

// Dendrochronology holds the schema definition for the Dendrochronology entity.
type Dendrochronology struct {
	ent.Schema
}

// Privacy policy of the Dendrochronology.
func (Dendrochronology) Policy() ent.Policy {
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

// Mixin of the Dendrochronology.
func (Dendrochronology) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DatingMixin{},
		mixin.DetailsMixin{},
		mixin.DraftMixin{},
		mixin.ImagesMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Annotations of the Dendrochronology.
func (Dendrochronology) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Dendrochronology.
func (Dendrochronology) Fields() []ent.Field {
	return []ent.Field{
		field.String("analysis_data").Optional(),
		field.String("analysis_url").Optional(),
		field.String("data_url").Optional(),
		field.String("chart_url").Optional(),
	}
}

// Edges of the Dendrochronology.
func (Dendrochronology) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("analysis", DendrochronologicalAnalysis.Type).
			Ref("dendrochronology").
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
