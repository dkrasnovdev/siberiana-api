// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
	"github.com/dkrasnovdev/siberiana-api/ent/projecttype"
)

// ProjectTypeDelete is the builder for deleting a ProjectType entity.
type ProjectTypeDelete struct {
	config
	hooks    []Hook
	mutation *ProjectTypeMutation
}

// Where appends a list predicates to the ProjectTypeDelete builder.
func (ptd *ProjectTypeDelete) Where(ps ...predicate.ProjectType) *ProjectTypeDelete {
	ptd.mutation.Where(ps...)
	return ptd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptd *ProjectTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ptd.sqlExec, ptd.mutation, ptd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ptd *ProjectTypeDelete) ExecX(ctx context.Context) int {
	n, err := ptd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptd *ProjectTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(projecttype.Table, sqlgraph.NewFieldSpec(projecttype.FieldID, field.TypeInt))
	if ps := ptd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ptd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ptd.mutation.done = true
	return affected, err
}

// ProjectTypeDeleteOne is the builder for deleting a single ProjectType entity.
type ProjectTypeDeleteOne struct {
	ptd *ProjectTypeDelete
}

// Where appends a list predicates to the ProjectTypeDelete builder.
func (ptdo *ProjectTypeDeleteOne) Where(ps ...predicate.ProjectType) *ProjectTypeDeleteOne {
	ptdo.ptd.mutation.Where(ps...)
	return ptdo
}

// Exec executes the deletion query.
func (ptdo *ProjectTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ptdo.ptd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{projecttype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdo *ProjectTypeDeleteOne) ExecX(ctx context.Context) {
	if err := ptdo.Exec(ctx); err != nil {
		panic(err)
	}
}
