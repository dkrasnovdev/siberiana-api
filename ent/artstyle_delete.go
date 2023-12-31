// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/artstyle"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// ArtStyleDelete is the builder for deleting a ArtStyle entity.
type ArtStyleDelete struct {
	config
	hooks    []Hook
	mutation *ArtStyleMutation
}

// Where appends a list predicates to the ArtStyleDelete builder.
func (asd *ArtStyleDelete) Where(ps ...predicate.ArtStyle) *ArtStyleDelete {
	asd.mutation.Where(ps...)
	return asd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asd *ArtStyleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, asd.sqlExec, asd.mutation, asd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (asd *ArtStyleDelete) ExecX(ctx context.Context) int {
	n, err := asd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asd *ArtStyleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(artstyle.Table, sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt))
	if ps := asd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	asd.mutation.done = true
	return affected, err
}

// ArtStyleDeleteOne is the builder for deleting a single ArtStyle entity.
type ArtStyleDeleteOne struct {
	asd *ArtStyleDelete
}

// Where appends a list predicates to the ArtStyleDelete builder.
func (asdo *ArtStyleDeleteOne) Where(ps ...predicate.ArtStyle) *ArtStyleDeleteOne {
	asdo.asd.mutation.Where(ps...)
	return asdo
}

// Exec executes the deletion query.
func (asdo *ArtStyleDeleteOne) Exec(ctx context.Context) error {
	n, err := asdo.asd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{artstyle.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asdo *ArtStyleDeleteOne) ExecX(ctx context.Context) {
	if err := asdo.Exec(ctx); err != nil {
		panic(err)
	}
}
