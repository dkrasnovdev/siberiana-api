// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/culture"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// CultureDelete is the builder for deleting a Culture entity.
type CultureDelete struct {
	config
	hooks    []Hook
	mutation *CultureMutation
}

// Where appends a list predicates to the CultureDelete builder.
func (cd *CultureDelete) Where(ps ...predicate.Culture) *CultureDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CultureDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CultureDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CultureDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(culture.Table, sqlgraph.NewFieldSpec(culture.FieldID, field.TypeInt))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// CultureDeleteOne is the builder for deleting a single Culture entity.
type CultureDeleteOne struct {
	cd *CultureDelete
}

// Where appends a list predicates to the CultureDelete builder.
func (cdo *CultureDeleteOne) Where(ps ...predicate.Culture) *CultureDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *CultureDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{culture.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CultureDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
