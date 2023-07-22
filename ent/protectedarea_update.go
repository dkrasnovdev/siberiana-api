// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
	"github.com/dkrasnovdev/heritage-api/ent/protectedarea"
)

// ProtectedAreaUpdate is the builder for updating ProtectedArea entities.
type ProtectedAreaUpdate struct {
	config
	hooks    []Hook
	mutation *ProtectedAreaMutation
}

// Where appends a list predicates to the ProtectedAreaUpdate builder.
func (pau *ProtectedAreaUpdate) Where(ps ...predicate.ProtectedArea) *ProtectedAreaUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// Mutation returns the ProtectedAreaMutation object of the builder.
func (pau *ProtectedAreaUpdate) Mutation() *ProtectedAreaMutation {
	return pau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *ProtectedAreaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *ProtectedAreaUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *ProtectedAreaUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *ProtectedAreaUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pau *ProtectedAreaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(protectedarea.Table, protectedarea.Columns, sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{protectedarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// ProtectedAreaUpdateOne is the builder for updating a single ProtectedArea entity.
type ProtectedAreaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProtectedAreaMutation
}

// Mutation returns the ProtectedAreaMutation object of the builder.
func (pauo *ProtectedAreaUpdateOne) Mutation() *ProtectedAreaMutation {
	return pauo.mutation
}

// Where appends a list predicates to the ProtectedAreaUpdate builder.
func (pauo *ProtectedAreaUpdateOne) Where(ps ...predicate.ProtectedArea) *ProtectedAreaUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *ProtectedAreaUpdateOne) Select(field string, fields ...string) *ProtectedAreaUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated ProtectedArea entity.
func (pauo *ProtectedAreaUpdateOne) Save(ctx context.Context) (*ProtectedArea, error) {
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *ProtectedAreaUpdateOne) SaveX(ctx context.Context) *ProtectedArea {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *ProtectedAreaUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *ProtectedAreaUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pauo *ProtectedAreaUpdateOne) sqlSave(ctx context.Context) (_node *ProtectedArea, err error) {
	_spec := sqlgraph.NewUpdateSpec(protectedarea.Table, protectedarea.Columns, sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProtectedArea.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, protectedarea.FieldID)
		for _, f := range fields {
			if !protectedarea.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != protectedarea.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &ProtectedArea{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{protectedarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}