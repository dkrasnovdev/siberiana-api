package schema

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	e "github.com/dkrasnovdev/siberiana-api/ent"
	"github.com/dkrasnovdev/siberiana-api/ent/hook"
	"github.com/dkrasnovdev/siberiana-api/ent/personalcollection"
	"github.com/dkrasnovdev/siberiana-api/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/mixin"
	rule "github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
)

// PersonalCollection holds the schema definition for the PersonalCollection entity.
type PersonalCollection struct {
	ent.Schema
}

// Privacy policy of the PersonalCollection.
func (PersonalCollection) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

// Mixin of the PersonalCollection.
func (PersonalCollection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Annotations of the PersonalCollection.
func (PersonalCollection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Fields of the PersonalCollection.
func (PersonalCollection) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name").
			NotEmpty(),
		field.Bool("is_public").
			Default(false),
	}
}

// Edges of the PersonalCollection.
func (PersonalCollection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("art", Art.Type),
		edge.To("artifacts", Artifact.Type),
		edge.To("books", Book.Type),
		edge.To("dendrochronology", Dendrochronology.Type),
		edge.To("herbaria", Herbarium.Type),
		edge.To("petroglyphs", Petroglyph.Type),
		edge.To("protected_area_pictures", ProtectedAreaPicture.Type),
	}
}

func (PersonalCollection) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
				var usr string
				var isAdmin bool
				v := rule.FromContext(ctx)
				if v != nil {
					usr = v.GetPreferredUsername()
					isAdmin = v.IsAdministrator()
				}
				if q, ok := query.(*e.PersonalCollectionQuery); ok && !isAdmin {
					q.Where(personalcollection.Or(personalcollection.IsPublic(true), personalcollection.CreatedBy(usr)))
				}
				return next.Query(ctx, query)
			})
		}),
	}
}

func (PersonalCollection) Hooks() []ent.Hook {
	return []ent.Hook{
		OwnershipHook,
	}
}

func OwnershipHook(next ent.Mutator) ent.Mutator {
	return hook.PersonalCollectionFunc(func(ctx context.Context, m *e.PersonalCollectionMutation) (ent.Value, error) {
		v := rule.FromContext(ctx)
		if v == nil {
			return nil, fmt.Errorf("not authenticated")
		}
		usr := v.GetPreferredUsername()
		op := m.Op()
		created_by, _ := m.OldCreatedBy(ctx)
		if op.Is(ent.OpUpdateOne) && created_by != usr {
			return nil, fmt.Errorf("Operation is not allowed. User %s is not owner of the collection.", usr)
		}

		return next.Mutate(ctx, m)
	})
}
