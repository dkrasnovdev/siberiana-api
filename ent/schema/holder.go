package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
)

// Holder holds the schema definition for the Holder entity.
type Holder struct {
	ent.Schema
}

// Privacy policy of the Holder.
func (Holder) Policy() ent.Policy {
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

// Mixin of the Holder.
func (Holder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Annotations of the Holder.
func (Holder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Holder.
func (Holder) Fields() []ent.Field {
	return []ent.Field{
		field.Time("begin_date"),
		field.Time("end_date").Optional(),
	}
}

// Edges of the Holder.
func (Holder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artifacts", Artifact.Type),
		edge.To("holder_responsibilities", HolderResponsibility.Type),
		edge.To("person", Person.Type).Unique(),
		edge.To("organization", Organization.Type).Unique(),
	}
}
