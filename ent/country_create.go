// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/country"
	"github.com/dkrasnovdev/siberiana-api/ent/location"
)

// CountryCreate is the builder for creating a Country entity.
type CountryCreate struct {
	config
	mutation *CountryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CountryCreate) SetCreatedAt(t time.Time) *CountryCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCreatedAt(t *time.Time) *CountryCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetCreatedBy sets the "created_by" field.
func (cc *CountryCreate) SetCreatedBy(s string) *CountryCreate {
	cc.mutation.SetCreatedBy(s)
	return cc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCreatedBy(s *string) *CountryCreate {
	if s != nil {
		cc.SetCreatedBy(*s)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CountryCreate) SetUpdatedAt(t time.Time) *CountryCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableUpdatedAt(t *time.Time) *CountryCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetUpdatedBy sets the "updated_by" field.
func (cc *CountryCreate) SetUpdatedBy(s string) *CountryCreate {
	cc.mutation.SetUpdatedBy(s)
	return cc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cc *CountryCreate) SetNillableUpdatedBy(s *string) *CountryCreate {
	if s != nil {
		cc.SetUpdatedBy(*s)
	}
	return cc
}

// SetDisplayName sets the "display_name" field.
func (cc *CountryCreate) SetDisplayName(s string) *CountryCreate {
	cc.mutation.SetDisplayName(s)
	return cc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cc *CountryCreate) SetNillableDisplayName(s *string) *CountryCreate {
	if s != nil {
		cc.SetDisplayName(*s)
	}
	return cc
}

// SetAbbreviation sets the "abbreviation" field.
func (cc *CountryCreate) SetAbbreviation(s string) *CountryCreate {
	cc.mutation.SetAbbreviation(s)
	return cc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (cc *CountryCreate) SetNillableAbbreviation(s *string) *CountryCreate {
	if s != nil {
		cc.SetAbbreviation(*s)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *CountryCreate) SetDescription(s string) *CountryCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CountryCreate) SetNillableDescription(s *string) *CountryCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetExternalLink sets the "external_link" field.
func (cc *CountryCreate) SetExternalLink(s string) *CountryCreate {
	cc.mutation.SetExternalLink(s)
	return cc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (cc *CountryCreate) SetNillableExternalLink(s *string) *CountryCreate {
	if s != nil {
		cc.SetExternalLink(*s)
	}
	return cc
}

// SetSlug sets the "slug" field.
func (cc *CountryCreate) SetSlug(s string) *CountryCreate {
	cc.mutation.SetSlug(s)
	return cc
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (cc *CountryCreate) SetNillableSlug(s *string) *CountryCreate {
	if s != nil {
		cc.SetSlug(*s)
	}
	return cc
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (cc *CountryCreate) SetLocationID(id int) *CountryCreate {
	cc.mutation.SetLocationID(id)
	return cc
}

// SetNillableLocationID sets the "location" edge to the Location entity by ID if the given value is not nil.
func (cc *CountryCreate) SetNillableLocationID(id *int) *CountryCreate {
	if id != nil {
		cc = cc.SetLocationID(*id)
	}
	return cc
}

// SetLocation sets the "location" edge to the Location entity.
func (cc *CountryCreate) SetLocation(l *Location) *CountryCreate {
	return cc.SetLocationID(l.ID)
}

// Mutation returns the CountryMutation object of the builder.
func (cc *CountryCreate) Mutation() *CountryMutation {
	return cc.mutation
}

// Save creates the Country in the database.
func (cc *CountryCreate) Save(ctx context.Context) (*Country, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CountryCreate) SaveX(ctx context.Context) *Country {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CountryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CountryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CountryCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if country.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized country.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := country.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if country.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized country.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := country.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CountryCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Country.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Country.updated_at"`)}
	}
	return nil
}

func (cc *CountryCreate) sqlSave(ctx context.Context) (*Country, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CountryCreate) createSpec() (*Country, *sqlgraph.CreateSpec) {
	var (
		_node = &Country{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(country.Table, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(country.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.CreatedBy(); ok {
		_spec.SetField(country.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(country.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.UpdatedBy(); ok {
		_spec.SetField(country.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := cc.mutation.DisplayName(); ok {
		_spec.SetField(country.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := cc.mutation.Abbreviation(); ok {
		_spec.SetField(country.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(country.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.ExternalLink(); ok {
		_spec.SetField(country.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := cc.mutation.Slug(); ok {
		_spec.SetField(country.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if nodes := cc.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   country.LocationTable,
			Columns: []string{country.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.location_country = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CountryCreateBulk is the builder for creating many Country entities in bulk.
type CountryCreateBulk struct {
	config
	builders []*CountryCreate
}

// Save creates the Country entities in the database.
func (ccb *CountryCreateBulk) Save(ctx context.Context) ([]*Country, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Country, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CountryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CountryCreateBulk) SaveX(ctx context.Context) []*Country {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CountryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CountryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
