// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

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

// SetCreatedAt sets the "created_at" field.
func (papc *ProtectedAreaPictureCreate) SetCreatedAt(t time.Time) *ProtectedAreaPictureCreate {
	papc.mutation.SetCreatedAt(t)
	return papc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (papc *ProtectedAreaPictureCreate) SetNillableCreatedAt(t *time.Time) *ProtectedAreaPictureCreate {
	if t != nil {
		papc.SetCreatedAt(*t)
	}
	return papc
}

// SetCreatedBy sets the "created_by" field.
func (papc *ProtectedAreaPictureCreate) SetCreatedBy(s string) *ProtectedAreaPictureCreate {
	papc.mutation.SetCreatedBy(s)
	return papc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (papc *ProtectedAreaPictureCreate) SetNillableCreatedBy(s *string) *ProtectedAreaPictureCreate {
	if s != nil {
		papc.SetCreatedBy(*s)
	}
	return papc
}

// SetUpdatedAt sets the "updated_at" field.
func (papc *ProtectedAreaPictureCreate) SetUpdatedAt(t time.Time) *ProtectedAreaPictureCreate {
	papc.mutation.SetUpdatedAt(t)
	return papc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (papc *ProtectedAreaPictureCreate) SetNillableUpdatedAt(t *time.Time) *ProtectedAreaPictureCreate {
	if t != nil {
		papc.SetUpdatedAt(*t)
	}
	return papc
}

// SetUpdatedBy sets the "updated_by" field.
func (papc *ProtectedAreaPictureCreate) SetUpdatedBy(s string) *ProtectedAreaPictureCreate {
	papc.mutation.SetUpdatedBy(s)
	return papc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (papc *ProtectedAreaPictureCreate) SetNillableUpdatedBy(s *string) *ProtectedAreaPictureCreate {
	if s != nil {
		papc.SetUpdatedBy(*s)
	}
	return papc
}

// SetDisplayName sets the "display_name" field.
func (papc *ProtectedAreaPictureCreate) SetDisplayName(s string) *ProtectedAreaPictureCreate {
	papc.mutation.SetDisplayName(s)
	return papc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (papc *ProtectedAreaPictureCreate) SetNillableDisplayName(s *string) *ProtectedAreaPictureCreate {
	if s != nil {
		papc.SetDisplayName(*s)
	}
	return papc
}

// SetDescription sets the "description" field.
func (papc *ProtectedAreaPictureCreate) SetDescription(s string) *ProtectedAreaPictureCreate {
	papc.mutation.SetDescription(s)
	return papc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (papc *ProtectedAreaPictureCreate) SetNillableDescription(s *string) *ProtectedAreaPictureCreate {
	if s != nil {
		papc.SetDescription(*s)
	}
	return papc
}

// SetExternalLinks sets the "external_links" field.
func (papc *ProtectedAreaPictureCreate) SetExternalLinks(s []string) *ProtectedAreaPictureCreate {
	papc.mutation.SetExternalLinks(s)
	return papc
}

// Mutation returns the ProtectedAreaPictureMutation object of the builder.
func (papc *ProtectedAreaPictureCreate) Mutation() *ProtectedAreaPictureMutation {
	return papc.mutation
}

// Save creates the ProtectedAreaPicture in the database.
func (papc *ProtectedAreaPictureCreate) Save(ctx context.Context) (*ProtectedAreaPicture, error) {
	if err := papc.defaults(); err != nil {
		return nil, err
	}
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

// defaults sets the default values of the builder before save.
func (papc *ProtectedAreaPictureCreate) defaults() error {
	if _, ok := papc.mutation.CreatedAt(); !ok {
		if protectedareapicture.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedareapicture.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := protectedareapicture.DefaultCreatedAt()
		papc.mutation.SetCreatedAt(v)
	}
	if _, ok := papc.mutation.UpdatedAt(); !ok {
		if protectedareapicture.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedareapicture.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := protectedareapicture.DefaultUpdatedAt()
		papc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (papc *ProtectedAreaPictureCreate) check() error {
	if _, ok := papc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProtectedAreaPicture.created_at"`)}
	}
	if _, ok := papc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProtectedAreaPicture.updated_at"`)}
	}
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
	if value, ok := papc.mutation.CreatedAt(); ok {
		_spec.SetField(protectedareapicture.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := papc.mutation.CreatedBy(); ok {
		_spec.SetField(protectedareapicture.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := papc.mutation.UpdatedAt(); ok {
		_spec.SetField(protectedareapicture.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := papc.mutation.UpdatedBy(); ok {
		_spec.SetField(protectedareapicture.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := papc.mutation.DisplayName(); ok {
		_spec.SetField(protectedareapicture.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := papc.mutation.Description(); ok {
		_spec.SetField(protectedareapicture.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := papc.mutation.ExternalLinks(); ok {
		_spec.SetField(protectedareapicture.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
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
			builder.defaults()
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
