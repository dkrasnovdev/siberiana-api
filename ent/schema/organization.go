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

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Privacy policy of the Organization.
func (Organization) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfModerator(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

// Mixin of the Organization.
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.ContactInformationMixin{},
		mixin.DetailsMixin{},
		mixin.ImagesMixin{},
	}
}

// Annotations of the Organization.
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		// entgql.MultiOrder(),
	}
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("previous_names", []string{}).
			Optional(),
		field.Bool("is_in_a_consortium").
			Optional(),
		field.String("consortium_document_url").
			Optional(),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("people", Person.Type),
		edge.From("holder", Holder.Type).Ref("organization").Unique(),
		edge.From("organization_type", OrganizationType.Type).Ref("organizations").Unique(),
	}
}
