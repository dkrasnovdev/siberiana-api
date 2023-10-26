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

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Privacy policy of the Person.
func (Person) Policy() ent.Policy {
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

// Mixin of the Person.
func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.ContactInformationMixin{},
		mixin.DetailsMixin{},
		mixin.ImagesMixin{},
	}
}

// Annotations of the Person.
func (Person) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("given_name").
			Optional(),
		field.String("family_name").
			Optional(),
		field.String("patronymic_name").
			Optional(),
		field.Time("begin_data").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.Time("end_date").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.Enum("gender").
			Values("female", "male"),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("collections", Collection.Type),
		edge.To("artifacts", Artifact.Type),
		edge.To("books", Book.Type),
		edge.To("projects", Project.Type),
		edge.To("publications", Publication.Type),
		edge.From("affiliation", Organization.Type).Ref("people").Unique(),
	}
}
