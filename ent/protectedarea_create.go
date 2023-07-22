// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/protectedarea"
)

// ProtectedAreaCreate is the builder for creating a ProtectedArea entity.
type ProtectedAreaCreate struct {
	config
	mutation *ProtectedAreaMutation
	hooks    []Hook
}

// Mutation returns the ProtectedAreaMutation object of the builder.
func (pac *ProtectedAreaCreate) Mutation() *ProtectedAreaMutation {
	return pac.mutation
}

// Save creates the ProtectedArea in the database.
func (pac *ProtectedAreaCreate) Save(ctx context.Context) (*ProtectedArea, error) {
	return withHooks(ctx, pac.sqlSave, pac.mutation, pac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pac *ProtectedAreaCreate) SaveX(ctx context.Context) *ProtectedArea {
	v, err := pac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pac *ProtectedAreaCreate) Exec(ctx context.Context) error {
	_, err := pac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pac *ProtectedAreaCreate) ExecX(ctx context.Context) {
	if err := pac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pac *ProtectedAreaCreate) check() error {
	return nil
}

func (pac *ProtectedAreaCreate) sqlSave(ctx context.Context) (*ProtectedArea, error) {
	if err := pac.check(); err != nil {
		return nil, err
	}
	_node, _spec := pac.createSpec()
	if err := sqlgraph.CreateNode(ctx, pac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pac.mutation.id = &_node.ID
	pac.mutation.done = true
	return _node, nil
}

func (pac *ProtectedAreaCreate) createSpec() (*ProtectedArea, *sqlgraph.CreateSpec) {
	var (
		_node = &ProtectedArea{config: pac.config}
		_spec = sqlgraph.NewCreateSpec(protectedarea.Table, sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// ProtectedAreaCreateBulk is the builder for creating many ProtectedArea entities in bulk.
type ProtectedAreaCreateBulk struct {
	config
	builders []*ProtectedAreaCreate
}

// Save creates the ProtectedArea entities in the database.
func (pacb *ProtectedAreaCreateBulk) Save(ctx context.Context) ([]*ProtectedArea, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pacb.builders))
	nodes := make([]*ProtectedArea, len(pacb.builders))
	mutators := make([]Mutator, len(pacb.builders))
	for i := range pacb.builders {
		func(i int, root context.Context) {
			builder := pacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProtectedAreaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pacb *ProtectedAreaCreateBulk) SaveX(ctx context.Context) []*ProtectedArea {
	v, err := pacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pacb *ProtectedAreaCreateBulk) Exec(ctx context.Context) error {
	_, err := pacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pacb *ProtectedAreaCreateBulk) ExecX(ctx context.Context) {
	if err := pacb.Exec(ctx); err != nil {
		panic(err)
	}
}