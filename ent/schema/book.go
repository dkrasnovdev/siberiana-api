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

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Privacy policy of the Book.
func (Book) Policy() ent.Policy {
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

// Mixin of the Book.
func (Book) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
		mixin.ImagesMixin{},
	}
}

// Annotations of the Book.
func (Book) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("files", []string{}).Optional(),
		field.Int("year").Max(2023).Optional(),
	}

}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("authors", Person.Type).Ref("books"),
		edge.From("book_genres", BookGenre.Type).Ref("books"),
		edge.From("collection", Collection.Type).Ref("books").Unique(),
		edge.From("holders", Holder.Type).Ref("books"),
		edge.From("publisher", Publisher.Type).Ref("books").Unique(),
		edge.From("license", License.Type).Ref("books").Unique(),
		edge.From("location", Location.Type).Ref("books").Unique(),
	}
}
