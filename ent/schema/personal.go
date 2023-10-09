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

// Personal holds the schema definition for the Personal entity.
type Personal struct {
	ent.Schema
}

// Privacy policy of the Personal.
func (Personal) Policy() ent.Policy {
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

// Mixin of the Personal.
func (Personal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.OwnerMixin{},
	}
}

// Annotations of the Personal.
func (Personal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Personal.
func (Personal) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name").
			NotEmpty(),
	}
}

// Edges of the Personal.
func (Personal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proxies", Proxy.Type),
	}
}
