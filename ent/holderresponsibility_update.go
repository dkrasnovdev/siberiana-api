// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/holder"
	"github.com/dkrasnovdev/heritage-api/ent/holderresponsibility"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// HolderResponsibilityUpdate is the builder for updating HolderResponsibility entities.
type HolderResponsibilityUpdate struct {
	config
	hooks    []Hook
	mutation *HolderResponsibilityMutation
}

// Where appends a list predicates to the HolderResponsibilityUpdate builder.
func (hru *HolderResponsibilityUpdate) Where(ps ...predicate.HolderResponsibility) *HolderResponsibilityUpdate {
	hru.mutation.Where(ps...)
	return hru
}

// AddHolderIDs adds the "holder" edge to the Holder entity by IDs.
func (hru *HolderResponsibilityUpdate) AddHolderIDs(ids ...int) *HolderResponsibilityUpdate {
	hru.mutation.AddHolderIDs(ids...)
	return hru
}

// AddHolder adds the "holder" edges to the Holder entity.
func (hru *HolderResponsibilityUpdate) AddHolder(h ...*Holder) *HolderResponsibilityUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hru.AddHolderIDs(ids...)
}

// Mutation returns the HolderResponsibilityMutation object of the builder.
func (hru *HolderResponsibilityUpdate) Mutation() *HolderResponsibilityMutation {
	return hru.mutation
}

// ClearHolder clears all "holder" edges to the Holder entity.
func (hru *HolderResponsibilityUpdate) ClearHolder() *HolderResponsibilityUpdate {
	hru.mutation.ClearHolder()
	return hru
}

// RemoveHolderIDs removes the "holder" edge to Holder entities by IDs.
func (hru *HolderResponsibilityUpdate) RemoveHolderIDs(ids ...int) *HolderResponsibilityUpdate {
	hru.mutation.RemoveHolderIDs(ids...)
	return hru
}

// RemoveHolder removes "holder" edges to Holder entities.
func (hru *HolderResponsibilityUpdate) RemoveHolder(h ...*Holder) *HolderResponsibilityUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hru.RemoveHolderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hru *HolderResponsibilityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hru.sqlSave, hru.mutation, hru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hru *HolderResponsibilityUpdate) SaveX(ctx context.Context) int {
	affected, err := hru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hru *HolderResponsibilityUpdate) Exec(ctx context.Context) error {
	_, err := hru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hru *HolderResponsibilityUpdate) ExecX(ctx context.Context) {
	if err := hru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hru *HolderResponsibilityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(holderresponsibility.Table, holderresponsibility.Columns, sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt))
	if ps := hru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if hru.mutation.HolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hru.mutation.RemovedHolderIDs(); len(nodes) > 0 && !hru.mutation.HolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hru.mutation.HolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{holderresponsibility.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hru.mutation.done = true
	return n, nil
}

// HolderResponsibilityUpdateOne is the builder for updating a single HolderResponsibility entity.
type HolderResponsibilityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HolderResponsibilityMutation
}

// AddHolderIDs adds the "holder" edge to the Holder entity by IDs.
func (hruo *HolderResponsibilityUpdateOne) AddHolderIDs(ids ...int) *HolderResponsibilityUpdateOne {
	hruo.mutation.AddHolderIDs(ids...)
	return hruo
}

// AddHolder adds the "holder" edges to the Holder entity.
func (hruo *HolderResponsibilityUpdateOne) AddHolder(h ...*Holder) *HolderResponsibilityUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hruo.AddHolderIDs(ids...)
}

// Mutation returns the HolderResponsibilityMutation object of the builder.
func (hruo *HolderResponsibilityUpdateOne) Mutation() *HolderResponsibilityMutation {
	return hruo.mutation
}

// ClearHolder clears all "holder" edges to the Holder entity.
func (hruo *HolderResponsibilityUpdateOne) ClearHolder() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearHolder()
	return hruo
}

// RemoveHolderIDs removes the "holder" edge to Holder entities by IDs.
func (hruo *HolderResponsibilityUpdateOne) RemoveHolderIDs(ids ...int) *HolderResponsibilityUpdateOne {
	hruo.mutation.RemoveHolderIDs(ids...)
	return hruo
}

// RemoveHolder removes "holder" edges to Holder entities.
func (hruo *HolderResponsibilityUpdateOne) RemoveHolder(h ...*Holder) *HolderResponsibilityUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hruo.RemoveHolderIDs(ids...)
}

// Where appends a list predicates to the HolderResponsibilityUpdate builder.
func (hruo *HolderResponsibilityUpdateOne) Where(ps ...predicate.HolderResponsibility) *HolderResponsibilityUpdateOne {
	hruo.mutation.Where(ps...)
	return hruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hruo *HolderResponsibilityUpdateOne) Select(field string, fields ...string) *HolderResponsibilityUpdateOne {
	hruo.fields = append([]string{field}, fields...)
	return hruo
}

// Save executes the query and returns the updated HolderResponsibility entity.
func (hruo *HolderResponsibilityUpdateOne) Save(ctx context.Context) (*HolderResponsibility, error) {
	return withHooks(ctx, hruo.sqlSave, hruo.mutation, hruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hruo *HolderResponsibilityUpdateOne) SaveX(ctx context.Context) *HolderResponsibility {
	node, err := hruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hruo *HolderResponsibilityUpdateOne) Exec(ctx context.Context) error {
	_, err := hruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hruo *HolderResponsibilityUpdateOne) ExecX(ctx context.Context) {
	if err := hruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hruo *HolderResponsibilityUpdateOne) sqlSave(ctx context.Context) (_node *HolderResponsibility, err error) {
	_spec := sqlgraph.NewUpdateSpec(holderresponsibility.Table, holderresponsibility.Columns, sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt))
	id, ok := hruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HolderResponsibility.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, holderresponsibility.FieldID)
		for _, f := range fields {
			if !holderresponsibility.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != holderresponsibility.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if hruo.mutation.HolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hruo.mutation.RemovedHolderIDs(); len(nodes) > 0 && !hruo.mutation.HolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hruo.mutation.HolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HolderResponsibility{config: hruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{holderresponsibility.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hruo.mutation.done = true
	return _node, nil
}
