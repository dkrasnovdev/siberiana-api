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

// Proxy holds the schema definition for the Proxy entity.
type Proxy struct {
	ent.Schema
}

// Mixin of the Proxy.
func (Proxy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Privacy policy of the Proxy.
func (Proxy) Policy() ent.Policy {
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

// Annotations of the Proxy.
func (Proxy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Proxy.
func (Proxy) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("artifacts", "books", "protected_area_pictures").
			Immutable(),
		field.String("ref_id"),
		field.String("url"),
	}
}

// Edges of the Proxy.
func (Proxy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("favourite", Favourite.Type).
			Ref("proxies").
			Unique(),
		edge.From("personal", Personal.Type).
			Ref("proxies").
			Unique(),
	}
}
