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
	"github.com/dkrasnovdev/siberiana-api/ent/personal"
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
		field.Bool("is_public").
			Default(false),
	}
}

// Edges of the Personal.
func (Personal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artifacts", Artifact.Type),
		edge.To("petroglyphs", Petroglyph.Type),
		edge.To("books", Book.Type),
		edge.To("protected_area_pictures", ProtectedAreaPicture.Type),
	}
}

func (Personal) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
				var usr string
				v := rule.FromContext(ctx)
				if v != nil {
					usr = v.GetPreferredUsername()
				}
				if q, ok := query.(*e.PersonalQuery); ok {
					q.Where(personal.Or(personal.IsPublic(true), personal.CreatedBy(usr)))
				}
				return next.Query(ctx, query)
			})
		}),
	}
}

func (Personal) Hooks() []ent.Hook {
	return []ent.Hook{}
}

func OwnershipHook(next ent.Mutator) ent.Mutator {
	type Ownership interface {
		CreatedBy() string
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		mx, ok := m.(Ownership)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		v := rule.FromContext(ctx)
		if v == nil {
			return nil, fmt.Errorf("not authenticated")
		}
		usr := v.GetPreferredUsername()

		op := m.Op()

		if op.Is(ent.OpUpdateOne|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne) && mx.CreatedBy() != usr {
			return nil, fmt.Errorf("not allowed")
		}

		return next.Mutate(ctx, m)
	})
}
