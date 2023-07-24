// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/location"
	"github.com/dkrasnovdev/heritage-api/ent/settlement"
)

// SettlementCreate is the builder for creating a Settlement entity.
type SettlementCreate struct {
	config
	mutation *SettlementMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SettlementCreate) SetCreatedAt(t time.Time) *SettlementCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SettlementCreate) SetNillableCreatedAt(t *time.Time) *SettlementCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *SettlementCreate) SetCreatedBy(s string) *SettlementCreate {
	sc.mutation.SetCreatedBy(s)
	return sc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sc *SettlementCreate) SetNillableCreatedBy(s *string) *SettlementCreate {
	if s != nil {
		sc.SetCreatedBy(*s)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SettlementCreate) SetUpdatedAt(t time.Time) *SettlementCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SettlementCreate) SetNillableUpdatedAt(t *time.Time) *SettlementCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *SettlementCreate) SetUpdatedBy(s string) *SettlementCreate {
	sc.mutation.SetUpdatedBy(s)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *SettlementCreate) SetNillableUpdatedBy(s *string) *SettlementCreate {
	if s != nil {
		sc.SetUpdatedBy(*s)
	}
	return sc
}

// SetAbbreviation sets the "abbreviation" field.
func (sc *SettlementCreate) SetAbbreviation(s string) *SettlementCreate {
	sc.mutation.SetAbbreviation(s)
	return sc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (sc *SettlementCreate) SetNillableAbbreviation(s *string) *SettlementCreate {
	if s != nil {
		sc.SetAbbreviation(*s)
	}
	return sc
}

// SetDisplayName sets the "display_name" field.
func (sc *SettlementCreate) SetDisplayName(s string) *SettlementCreate {
	sc.mutation.SetDisplayName(s)
	return sc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (sc *SettlementCreate) SetNillableDisplayName(s *string) *SettlementCreate {
	if s != nil {
		sc.SetDisplayName(*s)
	}
	return sc
}

// SetDescription sets the "description" field.
func (sc *SettlementCreate) SetDescription(s string) *SettlementCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *SettlementCreate) SetNillableDescription(s *string) *SettlementCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetExternalLinks sets the "external_links" field.
func (sc *SettlementCreate) SetExternalLinks(s []string) *SettlementCreate {
	sc.mutation.SetExternalLinks(s)
	return sc
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (sc *SettlementCreate) SetLocationID(id int) *SettlementCreate {
	sc.mutation.SetLocationID(id)
	return sc
}

// SetNillableLocationID sets the "location" edge to the Location entity by ID if the given value is not nil.
func (sc *SettlementCreate) SetNillableLocationID(id *int) *SettlementCreate {
	if id != nil {
		sc = sc.SetLocationID(*id)
	}
	return sc
}

// SetLocation sets the "location" edge to the Location entity.
func (sc *SettlementCreate) SetLocation(l *Location) *SettlementCreate {
	return sc.SetLocationID(l.ID)
}

// Mutation returns the SettlementMutation object of the builder.
func (sc *SettlementCreate) Mutation() *SettlementMutation {
	return sc.mutation
}

// Save creates the Settlement in the database.
func (sc *SettlementCreate) Save(ctx context.Context) (*Settlement, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettlementCreate) SaveX(ctx context.Context) *Settlement {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettlementCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettlementCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SettlementCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if settlement.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized settlement.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := settlement.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if settlement.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized settlement.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := settlement.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettlementCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Settlement.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Settlement.updated_at"`)}
	}
	return nil
}

func (sc *SettlementCreate) sqlSave(ctx context.Context) (*Settlement, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettlementCreate) createSpec() (*Settlement, *sqlgraph.CreateSpec) {
	var (
		_node = &Settlement{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(settlement.Table, sqlgraph.NewFieldSpec(settlement.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(settlement.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.SetField(settlement.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(settlement.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.SetField(settlement.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.Abbreviation(); ok {
		_spec.SetField(settlement.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := sc.mutation.DisplayName(); ok {
		_spec.SetField(settlement.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(settlement.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.ExternalLinks(); ok {
		_spec.SetField(settlement.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
	if nodes := sc.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   settlement.LocationTable,
			Columns: []string{settlement.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.location_settlement = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SettlementCreateBulk is the builder for creating many Settlement entities in bulk.
type SettlementCreateBulk struct {
	config
	builders []*SettlementCreate
}

// Save creates the Settlement entities in the database.
func (scb *SettlementCreateBulk) Save(ctx context.Context) ([]*Settlement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Settlement, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettlementMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettlementCreateBulk) SaveX(ctx context.Context) []*Settlement {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettlementCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettlementCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
