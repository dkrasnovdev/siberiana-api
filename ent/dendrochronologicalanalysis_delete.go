// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/dendrochronologicalanalysis"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// DendrochronologicalAnalysisDelete is the builder for deleting a DendrochronologicalAnalysis entity.
type DendrochronologicalAnalysisDelete struct {
	config
	hooks    []Hook
	mutation *DendrochronologicalAnalysisMutation
}

// Where appends a list predicates to the DendrochronologicalAnalysisDelete builder.
func (dad *DendrochronologicalAnalysisDelete) Where(ps ...predicate.DendrochronologicalAnalysis) *DendrochronologicalAnalysisDelete {
	dad.mutation.Where(ps...)
	return dad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dad *DendrochronologicalAnalysisDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dad.sqlExec, dad.mutation, dad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dad *DendrochronologicalAnalysisDelete) ExecX(ctx context.Context) int {
	n, err := dad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dad *DendrochronologicalAnalysisDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(dendrochronologicalanalysis.Table, sqlgraph.NewFieldSpec(dendrochronologicalanalysis.FieldID, field.TypeInt))
	if ps := dad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dad.mutation.done = true
	return affected, err
}

// DendrochronologicalAnalysisDeleteOne is the builder for deleting a single DendrochronologicalAnalysis entity.
type DendrochronologicalAnalysisDeleteOne struct {
	dad *DendrochronologicalAnalysisDelete
}

// Where appends a list predicates to the DendrochronologicalAnalysisDelete builder.
func (dado *DendrochronologicalAnalysisDeleteOne) Where(ps ...predicate.DendrochronologicalAnalysis) *DendrochronologicalAnalysisDeleteOne {
	dado.dad.mutation.Where(ps...)
	return dado
}

// Exec executes the deletion query.
func (dado *DendrochronologicalAnalysisDeleteOne) Exec(ctx context.Context) error {
	n, err := dado.dad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dendrochronologicalanalysis.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dado *DendrochronologicalAnalysisDeleteOne) ExecX(ctx context.Context) {
	if err := dado.Exec(ctx); err != nil {
		panic(err)
	}
}
