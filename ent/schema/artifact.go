package schema

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	e "github.com/dkrasnovdev/siberiana-api/ent"
	"github.com/dkrasnovdev/siberiana-api/ent/hook"
	"github.com/dkrasnovdev/siberiana-api/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/pkg/transform"
	"github.com/minio/minio-go/v7"
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
		mixin.DraftMixin{},
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
		field.String("dating").Optional(),
		field.Int("dating_start").Optional(),
		field.Int("dating_end").Optional(),

		field.Float("height").Optional(),
		field.Float("width").Optional(),
		field.Float("length").Optional(),
		field.Float("depth").Optional(),
		field.Float("diameter").Optional(),
		field.String("weight").Optional(),
		field.String("dimensions").Optional(),

		field.String("chemical_composition").Optional(),

		field.String("kp_number").Optional(),
		field.String("goskatalog_number").Optional(),
		field.String("inventory_number").Optional(),

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
		edge.From("donor", Person.Type).Ref("donated_artifacts").Unique(),
		edge.From("mediums", Medium.Type).Ref("artifacts"),
		edge.From("techniques", Technique.Type).Ref("artifacts"),
		edge.From("projects", Project.Type).Ref("artifacts"),
		edge.From("publications", Publication.Type).Ref("artifacts"),
		edge.From("cultural_affiliation", Culture.Type).Ref("artifacts").Unique(),
		edge.From("ethnos", Ethnos.Type).Ref("artifacts").Unique(),
		edge.From("organization", Organization.Type).Ref("artifacts").Unique(),
		edge.From("monument", Monument.Type).Ref("artifacts").Unique(),
		edge.From("model", Model.Type).Ref("artifacts").Unique(),
		edge.From("set", Set.Type).Ref("artifacts").Unique(),
		edge.From("location", Location.Type).Ref("artifacts").Unique(),
		edge.From("collection", Collection.Type).Ref("artifacts").Unique().Required(),
		edge.From("license", License.Type).Ref("artifacts").Unique(),
		edge.From("country", Country.Type).Ref("artifacts").Unique(),
		edge.From("settlement", Settlement.Type).Ref("artifacts").Unique(),
		edge.From("district", District.Type).Ref("artifacts").Unique(),
		edge.From("region", Region.Type).Ref("artifacts").Unique(),
	}
}

// Hooks of the Artifact.
func (Artifact) Hooks() []ent.Hook {
	return []ent.Hook{
		DeleteOrphans(),
	}
}

func DeleteOrphans() ent.Hook {
	deleteOp := func(next ent.Mutator) ent.Mutator {
		return hook.ArtifactFunc(func(ctx context.Context, m *e.ArtifactMutation) (ent.Value, error) {
			id, exists := m.ID()
			if !exists {
				return nil, errors.New("id field is missing")
			}

			a, err := m.Client().Artifact.Get(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("getting deleted artifact: %w", err)
			}

			filename := transform.Filename(a.PrimaryImageURL)

			minioClient := m.Client().Minio

			err = minioClient.RemoveObject(ctx, "artifacts", filename, minio.RemoveObjectOptions{})

			if err != nil {
				return nil, fmt.Errorf("deleting image: %w", err)
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(deleteOp, ent.OpDeleteOne)
}
