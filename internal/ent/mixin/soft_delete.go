package mixin

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	gen "github.com/dkrasnovdev/siberiana-api/ent"
	"github.com/dkrasnovdev/siberiana-api/ent/hook"
	"github.com/dkrasnovdev/siberiana-api/ent/intercept"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
)

// SoftDeleteMixin implements the soft delete pattern for schemas.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional(),
		field.String("deleted_by").
			Optional(),
	}
}

// softDeleteKey is a private key used to identify the context value for skipping soft-deletes.
type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the SoftDeleteMixin.
func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}
			d.P(q)
			return nil
		}),
	}
}

// Hooks of the SoftDeleteMixin.
func (d SoftDeleteMixin) Hooks() []ent.Hook {
	type SoftDelete interface {
		SetOp(ent.Op)
		Client() *gen.Client
		SetDeletedAt(time.Time)
		SetDeletedBy(string)
		WhereP(...func(*sql.Selector))
	}
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					// Skip soft-delete, means delete the entity permanently.
					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
						return next.Mutate(ctx, m)
					}

					// Retrieve the Viewer from the context.
					v := privacy.FromContext(ctx)

					// Check if the Viewer is present in the context.
					// If the Viewer is not found, it means that the user information is missing or not properly authenticated.
					if v == nil {
						// Return an error indicating that the Viewer is not available or the user is not authenticated.
						return nil, fmt.Errorf("not authenticated")
					}

					// Get the preferred username from the privacy policy.
					usr := v.GetPreferredUsername()

					mx, ok := m.(SoftDelete)
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}

					d.P(mx)
					mx.SetOp(ent.OpUpdate)
					mx.SetDeletedAt(time.Now())
					mx.SetDeletedBy(usr)
					return mx.Client().Mutate(ctx, m)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d SoftDeleteMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[0].Descriptor().Name),
	)
}
