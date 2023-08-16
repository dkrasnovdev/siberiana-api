// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/holder"
	"github.com/dkrasnovdev/siberiana-api/ent/holderresponsibility"
)

// HolderResponsibilityCreate is the builder for creating a HolderResponsibility entity.
type HolderResponsibilityCreate struct {
	config
	mutation *HolderResponsibilityMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (hrc *HolderResponsibilityCreate) SetCreatedAt(t time.Time) *HolderResponsibilityCreate {
	hrc.mutation.SetCreatedAt(t)
	return hrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableCreatedAt(t *time.Time) *HolderResponsibilityCreate {
	if t != nil {
		hrc.SetCreatedAt(*t)
	}
	return hrc
}

// SetCreatedBy sets the "created_by" field.
func (hrc *HolderResponsibilityCreate) SetCreatedBy(s string) *HolderResponsibilityCreate {
	hrc.mutation.SetCreatedBy(s)
	return hrc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableCreatedBy(s *string) *HolderResponsibilityCreate {
	if s != nil {
		hrc.SetCreatedBy(*s)
	}
	return hrc
}

// SetUpdatedAt sets the "updated_at" field.
func (hrc *HolderResponsibilityCreate) SetUpdatedAt(t time.Time) *HolderResponsibilityCreate {
	hrc.mutation.SetUpdatedAt(t)
	return hrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableUpdatedAt(t *time.Time) *HolderResponsibilityCreate {
	if t != nil {
		hrc.SetUpdatedAt(*t)
	}
	return hrc
}

// SetUpdatedBy sets the "updated_by" field.
func (hrc *HolderResponsibilityCreate) SetUpdatedBy(s string) *HolderResponsibilityCreate {
	hrc.mutation.SetUpdatedBy(s)
	return hrc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableUpdatedBy(s *string) *HolderResponsibilityCreate {
	if s != nil {
		hrc.SetUpdatedBy(*s)
	}
	return hrc
}

// SetDisplayName sets the "display_name" field.
func (hrc *HolderResponsibilityCreate) SetDisplayName(s string) *HolderResponsibilityCreate {
	hrc.mutation.SetDisplayName(s)
	return hrc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableDisplayName(s *string) *HolderResponsibilityCreate {
	if s != nil {
		hrc.SetDisplayName(*s)
	}
	return hrc
}

// SetAbbreviation sets the "abbreviation" field.
func (hrc *HolderResponsibilityCreate) SetAbbreviation(s string) *HolderResponsibilityCreate {
	hrc.mutation.SetAbbreviation(s)
	return hrc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableAbbreviation(s *string) *HolderResponsibilityCreate {
	if s != nil {
		hrc.SetAbbreviation(*s)
	}
	return hrc
}

// SetDescription sets the "description" field.
func (hrc *HolderResponsibilityCreate) SetDescription(s string) *HolderResponsibilityCreate {
	hrc.mutation.SetDescription(s)
	return hrc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableDescription(s *string) *HolderResponsibilityCreate {
	if s != nil {
		hrc.SetDescription(*s)
	}
	return hrc
}

// SetExternalLink sets the "external_link" field.
func (hrc *HolderResponsibilityCreate) SetExternalLink(s string) *HolderResponsibilityCreate {
	hrc.mutation.SetExternalLink(s)
	return hrc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableExternalLink(s *string) *HolderResponsibilityCreate {
	if s != nil {
		hrc.SetExternalLink(*s)
	}
	return hrc
}

// SetSlug sets the "slug" field.
func (hrc *HolderResponsibilityCreate) SetSlug(s string) *HolderResponsibilityCreate {
	hrc.mutation.SetSlug(s)
	return hrc
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (hrc *HolderResponsibilityCreate) SetNillableSlug(s *string) *HolderResponsibilityCreate {
	if s != nil {
		hrc.SetSlug(*s)
	}
	return hrc
}

// AddHolderIDs adds the "holder" edge to the Holder entity by IDs.
func (hrc *HolderResponsibilityCreate) AddHolderIDs(ids ...int) *HolderResponsibilityCreate {
	hrc.mutation.AddHolderIDs(ids...)
	return hrc
}

// AddHolder adds the "holder" edges to the Holder entity.
func (hrc *HolderResponsibilityCreate) AddHolder(h ...*Holder) *HolderResponsibilityCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hrc.AddHolderIDs(ids...)
}

// Mutation returns the HolderResponsibilityMutation object of the builder.
func (hrc *HolderResponsibilityCreate) Mutation() *HolderResponsibilityMutation {
	return hrc.mutation
}

// Save creates the HolderResponsibility in the database.
func (hrc *HolderResponsibilityCreate) Save(ctx context.Context) (*HolderResponsibility, error) {
	if err := hrc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, hrc.sqlSave, hrc.mutation, hrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hrc *HolderResponsibilityCreate) SaveX(ctx context.Context) *HolderResponsibility {
	v, err := hrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hrc *HolderResponsibilityCreate) Exec(ctx context.Context) error {
	_, err := hrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hrc *HolderResponsibilityCreate) ExecX(ctx context.Context) {
	if err := hrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hrc *HolderResponsibilityCreate) defaults() error {
	if _, ok := hrc.mutation.CreatedAt(); !ok {
		if holderresponsibility.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized holderresponsibility.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := holderresponsibility.DefaultCreatedAt()
		hrc.mutation.SetCreatedAt(v)
	}
	if _, ok := hrc.mutation.UpdatedAt(); !ok {
		if holderresponsibility.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized holderresponsibility.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := holderresponsibility.DefaultUpdatedAt()
		hrc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (hrc *HolderResponsibilityCreate) check() error {
	if _, ok := hrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HolderResponsibility.created_at"`)}
	}
	if _, ok := hrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "HolderResponsibility.updated_at"`)}
	}
	return nil
}

func (hrc *HolderResponsibilityCreate) sqlSave(ctx context.Context) (*HolderResponsibility, error) {
	if err := hrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hrc.mutation.id = &_node.ID
	hrc.mutation.done = true
	return _node, nil
}

func (hrc *HolderResponsibilityCreate) createSpec() (*HolderResponsibility, *sqlgraph.CreateSpec) {
	var (
		_node = &HolderResponsibility{config: hrc.config}
		_spec = sqlgraph.NewCreateSpec(holderresponsibility.Table, sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt))
	)
	if value, ok := hrc.mutation.CreatedAt(); ok {
		_spec.SetField(holderresponsibility.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hrc.mutation.CreatedBy(); ok {
		_spec.SetField(holderresponsibility.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := hrc.mutation.UpdatedAt(); ok {
		_spec.SetField(holderresponsibility.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := hrc.mutation.UpdatedBy(); ok {
		_spec.SetField(holderresponsibility.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := hrc.mutation.DisplayName(); ok {
		_spec.SetField(holderresponsibility.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := hrc.mutation.Abbreviation(); ok {
		_spec.SetField(holderresponsibility.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := hrc.mutation.Description(); ok {
		_spec.SetField(holderresponsibility.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := hrc.mutation.ExternalLink(); ok {
		_spec.SetField(holderresponsibility.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := hrc.mutation.Slug(); ok {
		_spec.SetField(holderresponsibility.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if nodes := hrc.mutation.HolderIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HolderResponsibilityCreateBulk is the builder for creating many HolderResponsibility entities in bulk.
type HolderResponsibilityCreateBulk struct {
	config
	builders []*HolderResponsibilityCreate
}

// Save creates the HolderResponsibility entities in the database.
func (hrcb *HolderResponsibilityCreateBulk) Save(ctx context.Context) ([]*HolderResponsibility, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hrcb.builders))
	nodes := make([]*HolderResponsibility, len(hrcb.builders))
	mutators := make([]Mutator, len(hrcb.builders))
	for i := range hrcb.builders {
		func(i int, root context.Context) {
			builder := hrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HolderResponsibilityMutation)
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
					_, err = mutators[i+1].Mutate(root, hrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hrcb *HolderResponsibilityCreateBulk) SaveX(ctx context.Context) []*HolderResponsibility {
	v, err := hrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hrcb *HolderResponsibilityCreateBulk) Exec(ctx context.Context) error {
	_, err := hrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hrcb *HolderResponsibilityCreateBulk) ExecX(ctx context.Context) {
	if err := hrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
