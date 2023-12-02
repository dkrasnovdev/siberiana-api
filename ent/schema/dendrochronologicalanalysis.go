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

// DendrochronologicalAnalysis holds the schema definition for the DendrochronologicalAnalysis entity.
type DendrochronologicalAnalysis struct {
	ent.Schema
}

// Privacy policy of the DendrochronologicalAnalysis.
func (DendrochronologicalAnalysis) Policy() ent.Policy {
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

// Mixin of the DendrochronologicalAnalysis.
func (DendrochronologicalAnalysis) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Annotations of the DendrochronologicalAnalysis.
func (DendrochronologicalAnalysis) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the DendrochronologicalAnalysis.
func (DendrochronologicalAnalysis) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name"),
		field.Int("start_year"),
		field.Int("end_year"),
		field.Int("number_of_rings"),
		field.Float("coefficient_correlation"),
		field.Float("standard_deviation"),
		field.Float("sensitivity"),
		field.String("sampling_location"),
	}
}

// Edges of the DendrochronologicalAnalysis.
func (DendrochronologicalAnalysis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dendrochronology", Dendrochronology.Type).
			Unique().
			Required(),
	}
}
