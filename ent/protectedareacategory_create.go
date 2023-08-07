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
	"github.com/dkrasnovdev/heritage-api/ent/protectedareacategory"
)

// ProtectedAreaCategoryCreate is the builder for creating a ProtectedAreaCategory entity.
type ProtectedAreaCategoryCreate struct {
	config
	mutation *ProtectedAreaCategoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pacc *ProtectedAreaCategoryCreate) SetCreatedAt(t time.Time) *ProtectedAreaCategoryCreate {
	pacc.mutation.SetCreatedAt(t)
	return pacc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pacc *ProtectedAreaCategoryCreate) SetNillableCreatedAt(t *time.Time) *ProtectedAreaCategoryCreate {
	if t != nil {
		pacc.SetCreatedAt(*t)
	}
	return pacc
}

// SetCreatedBy sets the "created_by" field.
func (pacc *ProtectedAreaCategoryCreate) SetCreatedBy(s string) *ProtectedAreaCategoryCreate {
	pacc.mutation.SetCreatedBy(s)
	return pacc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pacc *ProtectedAreaCategoryCreate) SetNillableCreatedBy(s *string) *ProtectedAreaCategoryCreate {
	if s != nil {
		pacc.SetCreatedBy(*s)
	}
	return pacc
}

// SetUpdatedAt sets the "updated_at" field.
func (pacc *ProtectedAreaCategoryCreate) SetUpdatedAt(t time.Time) *ProtectedAreaCategoryCreate {
	pacc.mutation.SetUpdatedAt(t)
	return pacc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pacc *ProtectedAreaCategoryCreate) SetNillableUpdatedAt(t *time.Time) *ProtectedAreaCategoryCreate {
	if t != nil {
		pacc.SetUpdatedAt(*t)
	}
	return pacc
}

// SetUpdatedBy sets the "updated_by" field.
func (pacc *ProtectedAreaCategoryCreate) SetUpdatedBy(s string) *ProtectedAreaCategoryCreate {
	pacc.mutation.SetUpdatedBy(s)
	return pacc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pacc *ProtectedAreaCategoryCreate) SetNillableUpdatedBy(s *string) *ProtectedAreaCategoryCreate {
	if s != nil {
		pacc.SetUpdatedBy(*s)
	}
	return pacc
}

// SetDisplayName sets the "display_name" field.
func (pacc *ProtectedAreaCategoryCreate) SetDisplayName(s string) *ProtectedAreaCategoryCreate {
	pacc.mutation.SetDisplayName(s)
	return pacc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pacc *ProtectedAreaCategoryCreate) SetNillableDisplayName(s *string) *ProtectedAreaCategoryCreate {
	if s != nil {
		pacc.SetDisplayName(*s)
	}
	return pacc
}

// SetAbbreviation sets the "abbreviation" field.
func (pacc *ProtectedAreaCategoryCreate) SetAbbreviation(s string) *ProtectedAreaCategoryCreate {
	pacc.mutation.SetAbbreviation(s)
	return pacc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pacc *ProtectedAreaCategoryCreate) SetNillableAbbreviation(s *string) *ProtectedAreaCategoryCreate {
	if s != nil {
		pacc.SetAbbreviation(*s)
	}
	return pacc
}

// SetDescription sets the "description" field.
func (pacc *ProtectedAreaCategoryCreate) SetDescription(s string) *ProtectedAreaCategoryCreate {
	pacc.mutation.SetDescription(s)
	return pacc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pacc *ProtectedAreaCategoryCreate) SetNillableDescription(s *string) *ProtectedAreaCategoryCreate {
	if s != nil {
		pacc.SetDescription(*s)
	}
	return pacc
}

// SetExternalLink sets the "external_link" field.
func (pacc *ProtectedAreaCategoryCreate) SetExternalLink(s string) *ProtectedAreaCategoryCreate {
	pacc.mutation.SetExternalLink(s)
	return pacc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (pacc *ProtectedAreaCategoryCreate) SetNillableExternalLink(s *string) *ProtectedAreaCategoryCreate {
	if s != nil {
		pacc.SetExternalLink(*s)
	}
	return pacc
}

// AddProtectedAreaIDs adds the "protected_areas" edge to the ProtectedArea entity by IDs.
func (pacc *ProtectedAreaCategoryCreate) AddProtectedAreaIDs(ids ...int) *ProtectedAreaCategoryCreate {
	pacc.mutation.AddProtectedAreaIDs(ids...)
	return pacc
}

// AddProtectedAreas adds the "protected_areas" edges to the ProtectedArea entity.
func (pacc *ProtectedAreaCategoryCreate) AddProtectedAreas(p ...*ProtectedArea) *ProtectedAreaCategoryCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pacc.AddProtectedAreaIDs(ids...)
}

// Mutation returns the ProtectedAreaCategoryMutation object of the builder.
func (pacc *ProtectedAreaCategoryCreate) Mutation() *ProtectedAreaCategoryMutation {
	return pacc.mutation
}

// Save creates the ProtectedAreaCategory in the database.
func (pacc *ProtectedAreaCategoryCreate) Save(ctx context.Context) (*ProtectedAreaCategory, error) {
	if err := pacc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pacc.sqlSave, pacc.mutation, pacc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pacc *ProtectedAreaCategoryCreate) SaveX(ctx context.Context) *ProtectedAreaCategory {
	v, err := pacc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pacc *ProtectedAreaCategoryCreate) Exec(ctx context.Context) error {
	_, err := pacc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pacc *ProtectedAreaCategoryCreate) ExecX(ctx context.Context) {
	if err := pacc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pacc *ProtectedAreaCategoryCreate) defaults() error {
	if _, ok := pacc.mutation.CreatedAt(); !ok {
		if protectedareacategory.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedareacategory.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := protectedareacategory.DefaultCreatedAt()
		pacc.mutation.SetCreatedAt(v)
	}
	if _, ok := pacc.mutation.UpdatedAt(); !ok {
		if protectedareacategory.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedareacategory.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := protectedareacategory.DefaultUpdatedAt()
		pacc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pacc *ProtectedAreaCategoryCreate) check() error {
	if _, ok := pacc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProtectedAreaCategory.created_at"`)}
	}
	if _, ok := pacc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProtectedAreaCategory.updated_at"`)}
	}
	return nil
}

func (pacc *ProtectedAreaCategoryCreate) sqlSave(ctx context.Context) (*ProtectedAreaCategory, error) {
	if err := pacc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pacc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pacc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pacc.mutation.id = &_node.ID
	pacc.mutation.done = true
	return _node, nil
}

func (pacc *ProtectedAreaCategoryCreate) createSpec() (*ProtectedAreaCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &ProtectedAreaCategory{config: pacc.config}
		_spec = sqlgraph.NewCreateSpec(protectedareacategory.Table, sqlgraph.NewFieldSpec(protectedareacategory.FieldID, field.TypeInt))
	)
	if value, ok := pacc.mutation.CreatedAt(); ok {
		_spec.SetField(protectedareacategory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pacc.mutation.CreatedBy(); ok {
		_spec.SetField(protectedareacategory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pacc.mutation.UpdatedAt(); ok {
		_spec.SetField(protectedareacategory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pacc.mutation.UpdatedBy(); ok {
		_spec.SetField(protectedareacategory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pacc.mutation.DisplayName(); ok {
		_spec.SetField(protectedareacategory.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := pacc.mutation.Abbreviation(); ok {
		_spec.SetField(protectedareacategory.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := pacc.mutation.Description(); ok {
		_spec.SetField(protectedareacategory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pacc.mutation.ExternalLink(); ok {
		_spec.SetField(protectedareacategory.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if nodes := pacc.mutation.ProtectedAreasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   protectedareacategory.ProtectedAreasTable,
			Columns: []string{protectedareacategory.ProtectedAreasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProtectedAreaCategoryCreateBulk is the builder for creating many ProtectedAreaCategory entities in bulk.
type ProtectedAreaCategoryCreateBulk struct {
	config
	builders []*ProtectedAreaCategoryCreate
}

// Save creates the ProtectedAreaCategory entities in the database.
func (paccb *ProtectedAreaCategoryCreateBulk) Save(ctx context.Context) ([]*ProtectedAreaCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(paccb.builders))
	nodes := make([]*ProtectedAreaCategory, len(paccb.builders))
	mutators := make([]Mutator, len(paccb.builders))
	for i := range paccb.builders {
		func(i int, root context.Context) {
			builder := paccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProtectedAreaCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, paccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, paccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, paccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (paccb *ProtectedAreaCategoryCreateBulk) SaveX(ctx context.Context) []*ProtectedAreaCategory {
	v, err := paccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (paccb *ProtectedAreaCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := paccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (paccb *ProtectedAreaCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := paccb.Exec(ctx); err != nil {
		panic(err)
	}
}
