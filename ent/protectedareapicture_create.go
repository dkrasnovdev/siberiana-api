// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/protectedareapicture"
)

// ProtectedAreaPictureCreate is the builder for creating a ProtectedAreaPicture entity.
type ProtectedAreaPictureCreate struct {
	config
	mutation *ProtectedAreaPictureMutation
	hooks    []Hook
}

// Mutation returns the ProtectedAreaPictureMutation object of the builder.
func (papc *ProtectedAreaPictureCreate) Mutation() *ProtectedAreaPictureMutation {
	return papc.mutation
}

// Save creates the ProtectedAreaPicture in the database.
func (papc *ProtectedAreaPictureCreate) Save(ctx context.Context) (*ProtectedAreaPicture, error) {
	return withHooks(ctx, papc.sqlSave, papc.mutation, papc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (papc *ProtectedAreaPictureCreate) SaveX(ctx context.Context) *ProtectedAreaPicture {
	v, err := papc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (papc *ProtectedAreaPictureCreate) Exec(ctx context.Context) error {
	_, err := papc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (papc *ProtectedAreaPictureCreate) ExecX(ctx context.Context) {
	if err := papc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (papc *ProtectedAreaPictureCreate) check() error {
	return nil
}

func (papc *ProtectedAreaPictureCreate) sqlSave(ctx context.Context) (*ProtectedAreaPicture, error) {
	if err := papc.check(); err != nil {
		return nil, err
	}
	_node, _spec := papc.createSpec()
	if err := sqlgraph.CreateNode(ctx, papc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	papc.mutation.id = &_node.ID
	papc.mutation.done = true
	return _node, nil
}

func (papc *ProtectedAreaPictureCreate) createSpec() (*ProtectedAreaPicture, *sqlgraph.CreateSpec) {
	var (
		_node = &ProtectedAreaPicture{config: papc.config}
		_spec = sqlgraph.NewCreateSpec(protectedareapicture.Table, sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// ProtectedAreaPictureCreateBulk is the builder for creating many ProtectedAreaPicture entities in bulk.
type ProtectedAreaPictureCreateBulk struct {
	config
	builders []*ProtectedAreaPictureCreate
}

// Save creates the ProtectedAreaPicture entities in the database.
func (papcb *ProtectedAreaPictureCreateBulk) Save(ctx context.Context) ([]*ProtectedAreaPicture, error) {
	specs := make([]*sqlgraph.CreateSpec, len(papcb.builders))
	nodes := make([]*ProtectedAreaPicture, len(papcb.builders))
	mutators := make([]Mutator, len(papcb.builders))
	for i := range papcb.builders {
		func(i int, root context.Context) {
			builder := papcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProtectedAreaPictureMutation)
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
					_, err = mutators[i+1].Mutate(root, papcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, papcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, papcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (papcb *ProtectedAreaPictureCreateBulk) SaveX(ctx context.Context) []*ProtectedAreaPicture {
	v, err := papcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (papcb *ProtectedAreaPictureCreateBulk) Exec(ctx context.Context) error {
	_, err := papcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (papcb *ProtectedAreaPictureCreateBulk) ExecX(ctx context.Context) {
	if err := papcb.Exec(ctx); err != nil {
		panic(err)
	}
}