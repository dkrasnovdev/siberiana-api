// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/bookgenre"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// BookGenreDelete is the builder for deleting a BookGenre entity.
type BookGenreDelete struct {
	config
	hooks    []Hook
	mutation *BookGenreMutation
}

// Where appends a list predicates to the BookGenreDelete builder.
func (bgd *BookGenreDelete) Where(ps ...predicate.BookGenre) *BookGenreDelete {
	bgd.mutation.Where(ps...)
	return bgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bgd *BookGenreDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bgd.sqlExec, bgd.mutation, bgd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bgd *BookGenreDelete) ExecX(ctx context.Context) int {
	n, err := bgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bgd *BookGenreDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bookgenre.Table, sqlgraph.NewFieldSpec(bookgenre.FieldID, field.TypeInt))
	if ps := bgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bgd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bgd.mutation.done = true
	return affected, err
}

// BookGenreDeleteOne is the builder for deleting a single BookGenre entity.
type BookGenreDeleteOne struct {
	bgd *BookGenreDelete
}

// Where appends a list predicates to the BookGenreDelete builder.
func (bgdo *BookGenreDeleteOne) Where(ps ...predicate.BookGenre) *BookGenreDeleteOne {
	bgdo.bgd.mutation.Where(ps...)
	return bgdo
}

// Exec executes the deletion query.
func (bgdo *BookGenreDeleteOne) Exec(ctx context.Context) error {
	n, err := bgdo.bgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bookgenre.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bgdo *BookGenreDeleteOne) ExecX(ctx context.Context) {
	if err := bgdo.Exec(ctx); err != nil {
		panic(err)
	}
}
