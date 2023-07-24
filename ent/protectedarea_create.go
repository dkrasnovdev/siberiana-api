// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

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

// SetCreatedAt sets the "created_at" field.
func (pac *ProtectedAreaCreate) SetCreatedAt(t time.Time) *ProtectedAreaCreate {
	pac.mutation.SetCreatedAt(t)
	return pac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pac *ProtectedAreaCreate) SetNillableCreatedAt(t *time.Time) *ProtectedAreaCreate {
	if t != nil {
		pac.SetCreatedAt(*t)
	}
	return pac
}

// SetCreatedBy sets the "created_by" field.
func (pac *ProtectedAreaCreate) SetCreatedBy(s string) *ProtectedAreaCreate {
	pac.mutation.SetCreatedBy(s)
	return pac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pac *ProtectedAreaCreate) SetNillableCreatedBy(s *string) *ProtectedAreaCreate {
	if s != nil {
		pac.SetCreatedBy(*s)
	}
	return pac
}

// SetUpdatedAt sets the "updated_at" field.
func (pac *ProtectedAreaCreate) SetUpdatedAt(t time.Time) *ProtectedAreaCreate {
	pac.mutation.SetUpdatedAt(t)
	return pac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pac *ProtectedAreaCreate) SetNillableUpdatedAt(t *time.Time) *ProtectedAreaCreate {
	if t != nil {
		pac.SetUpdatedAt(*t)
	}
	return pac
}

// SetUpdatedBy sets the "updated_by" field.
func (pac *ProtectedAreaCreate) SetUpdatedBy(s string) *ProtectedAreaCreate {
	pac.mutation.SetUpdatedBy(s)
	return pac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pac *ProtectedAreaCreate) SetNillableUpdatedBy(s *string) *ProtectedAreaCreate {
	if s != nil {
		pac.SetUpdatedBy(*s)
	}
	return pac
}

// SetDisplayName sets the "display_name" field.
func (pac *ProtectedAreaCreate) SetDisplayName(s string) *ProtectedAreaCreate {
	pac.mutation.SetDisplayName(s)
	return pac
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pac *ProtectedAreaCreate) SetNillableDisplayName(s *string) *ProtectedAreaCreate {
	if s != nil {
		pac.SetDisplayName(*s)
	}
	return pac
}

// SetAbbreviation sets the "abbreviation" field.
func (pac *ProtectedAreaCreate) SetAbbreviation(s string) *ProtectedAreaCreate {
	pac.mutation.SetAbbreviation(s)
	return pac
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pac *ProtectedAreaCreate) SetNillableAbbreviation(s *string) *ProtectedAreaCreate {
	if s != nil {
		pac.SetAbbreviation(*s)
	}
	return pac
}

// SetDescription sets the "description" field.
func (pac *ProtectedAreaCreate) SetDescription(s string) *ProtectedAreaCreate {
	pac.mutation.SetDescription(s)
	return pac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pac *ProtectedAreaCreate) SetNillableDescription(s *string) *ProtectedAreaCreate {
	if s != nil {
		pac.SetDescription(*s)
	}
	return pac
}

// SetExternalLinks sets the "external_links" field.
func (pac *ProtectedAreaCreate) SetExternalLinks(s []string) *ProtectedAreaCreate {
	pac.mutation.SetExternalLinks(s)
	return pac
}

// Mutation returns the ProtectedAreaMutation object of the builder.
func (pac *ProtectedAreaCreate) Mutation() *ProtectedAreaMutation {
	return pac.mutation
}

// Save creates the ProtectedArea in the database.
func (pac *ProtectedAreaCreate) Save(ctx context.Context) (*ProtectedArea, error) {
	if err := pac.defaults(); err != nil {
		return nil, err
	}
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

// defaults sets the default values of the builder before save.
func (pac *ProtectedAreaCreate) defaults() error {
	if _, ok := pac.mutation.CreatedAt(); !ok {
		if protectedarea.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedarea.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := protectedarea.DefaultCreatedAt()
		pac.mutation.SetCreatedAt(v)
	}
	if _, ok := pac.mutation.UpdatedAt(); !ok {
		if protectedarea.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedarea.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := protectedarea.DefaultUpdatedAt()
		pac.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pac *ProtectedAreaCreate) check() error {
	if _, ok := pac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProtectedArea.created_at"`)}
	}
	if _, ok := pac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProtectedArea.updated_at"`)}
	}
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
	if value, ok := pac.mutation.CreatedAt(); ok {
		_spec.SetField(protectedarea.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pac.mutation.CreatedBy(); ok {
		_spec.SetField(protectedarea.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pac.mutation.UpdatedAt(); ok {
		_spec.SetField(protectedarea.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pac.mutation.UpdatedBy(); ok {
		_spec.SetField(protectedarea.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pac.mutation.DisplayName(); ok {
		_spec.SetField(protectedarea.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := pac.mutation.Abbreviation(); ok {
		_spec.SetField(protectedarea.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := pac.mutation.Description(); ok {
		_spec.SetField(protectedarea.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pac.mutation.ExternalLinks(); ok {
		_spec.SetField(protectedarea.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
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
			builder.defaults()
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
