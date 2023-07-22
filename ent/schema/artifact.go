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

// Artifact holds the schema definition for the Artifact entity.
type Artifact struct {
	ent.Schema
}

// Privacy policy of the Artifact.
func (Artifact) Policy() ent.Policy {
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

// Mixin of the Artifact.
func (Artifact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.DetailsMixin{},
		mixin.ImagesMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Annotations of the Artifact.
func (Artifact) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the Artifact.
func (Artifact) Fields() []ent.Field {
	return []ent.Field{

		field.String("dimensions").Optional(),
		field.String("weight").Optional(),
		field.String("chemical_composition").Optional(),
		field.String("typology").Optional(),
		field.Time("admission_date").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
	}
}

// Edges of the Artifact.
func (Artifact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("authors", Person.Type).Ref("artifacts"),
		edge.From("mediums", Medium.Type).Ref("artifacts"),
		edge.From("techniques", Technique.Type).Ref("artifacts"),
		edge.From("projects", Project.Type).Ref("artifacts"),
		edge.From("publications", Publication.Type).Ref("artifacts"),
		edge.From("holders", Holder.Type).Ref("artifacts"),
		edge.From("cultural_affiliation", Culture.Type).Ref("artifacts").Unique(),
		edge.From("monument", Monument.Type).Ref("artifacts").Unique(),
		edge.From("model", Model.Type).Ref("artifacts").Unique(),
		edge.From("set", Set.Type).Ref("artifacts").Unique(),
		edge.From("period", Period.Type).Ref("artifacts").Unique(),
		edge.From("location", Location.Type).Ref("artifacts").Unique(),
		edge.From("collection", Collection.Type).Ref("artifacts").Unique(),
		edge.From("license", License.Type).Ref("artifacts").Unique(),
	}
}
