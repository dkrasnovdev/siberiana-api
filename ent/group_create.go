// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/group"
	"github.com/dkrasnovdev/siberiana-api/ent/herbarium"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gc *GroupCreate) SetCreatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetCreatedBy sets the "created_by" field.
func (gc *GroupCreate) SetCreatedBy(s string) *GroupCreate {
	gc.mutation.SetCreatedBy(s)
	return gc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedBy(s *string) *GroupCreate {
	if s != nil {
		gc.SetCreatedBy(*s)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GroupCreate) SetUpdatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableUpdatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetUpdatedBy sets the "updated_by" field.
func (gc *GroupCreate) SetUpdatedBy(s string) *GroupCreate {
	gc.mutation.SetUpdatedBy(s)
	return gc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gc *GroupCreate) SetNillableUpdatedBy(s *string) *GroupCreate {
	if s != nil {
		gc.SetUpdatedBy(*s)
	}
	return gc
}

// SetDisplayName sets the "display_name" field.
func (gc *GroupCreate) SetDisplayName(s string) *GroupCreate {
	gc.mutation.SetDisplayName(s)
	return gc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDisplayName(s *string) *GroupCreate {
	if s != nil {
		gc.SetDisplayName(*s)
	}
	return gc
}

// SetAbbreviation sets the "abbreviation" field.
func (gc *GroupCreate) SetAbbreviation(s string) *GroupCreate {
	gc.mutation.SetAbbreviation(s)
	return gc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (gc *GroupCreate) SetNillableAbbreviation(s *string) *GroupCreate {
	if s != nil {
		gc.SetAbbreviation(*s)
	}
	return gc
}

// SetDescription sets the "description" field.
func (gc *GroupCreate) SetDescription(s string) *GroupCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDescription(s *string) *GroupCreate {
	if s != nil {
		gc.SetDescription(*s)
	}
	return gc
}

// SetExternalLink sets the "external_link" field.
func (gc *GroupCreate) SetExternalLink(s string) *GroupCreate {
	gc.mutation.SetExternalLink(s)
	return gc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (gc *GroupCreate) SetNillableExternalLink(s *string) *GroupCreate {
	if s != nil {
		gc.SetExternalLink(*s)
	}
	return gc
}

// SetPrimaryImageURL sets the "primary_image_url" field.
func (gc *GroupCreate) SetPrimaryImageURL(s string) *GroupCreate {
	gc.mutation.SetPrimaryImageURL(s)
	return gc
}

// SetNillablePrimaryImageURL sets the "primary_image_url" field if the given value is not nil.
func (gc *GroupCreate) SetNillablePrimaryImageURL(s *string) *GroupCreate {
	if s != nil {
		gc.SetPrimaryImageURL(*s)
	}
	return gc
}

// SetAdditionalImagesUrls sets the "additional_images_urls" field.
func (gc *GroupCreate) SetAdditionalImagesUrls(s []string) *GroupCreate {
	gc.mutation.SetAdditionalImagesUrls(s)
	return gc
}

// AddHerbariumIDs adds the "herbaria" edge to the Herbarium entity by IDs.
func (gc *GroupCreate) AddHerbariumIDs(ids ...int) *GroupCreate {
	gc.mutation.AddHerbariumIDs(ids...)
	return gc
}

// AddHerbaria adds the "herbaria" edges to the Herbarium entity.
func (gc *GroupCreate) AddHerbaria(h ...*Herbarium) *GroupCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return gc.AddHerbariumIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	if err := gc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		if group.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized group.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := group.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		if group.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized group.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := group.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Group.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Group.updated_at"`)}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(group.Table, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	)
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(group.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.CreatedBy(); ok {
		_spec.SetField(group.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.UpdatedBy(); ok {
		_spec.SetField(group.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := gc.mutation.DisplayName(); ok {
		_spec.SetField(group.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := gc.mutation.Abbreviation(); ok {
		_spec.SetField(group.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := gc.mutation.ExternalLink(); ok {
		_spec.SetField(group.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := gc.mutation.PrimaryImageURL(); ok {
		_spec.SetField(group.FieldPrimaryImageURL, field.TypeString, value)
		_node.PrimaryImageURL = value
	}
	if value, ok := gc.mutation.AdditionalImagesUrls(); ok {
		_spec.SetField(group.FieldAdditionalImagesUrls, field.TypeJSON, value)
		_node.AdditionalImagesUrls = value
	}
	if nodes := gc.mutation.HerbariaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.HerbariaTable,
			Columns: []string{group.HerbariaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(herbarium.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	err      error
	builders []*GroupCreate
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
