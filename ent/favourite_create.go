// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/favourite"
	"github.com/dkrasnovdev/siberiana-api/ent/proxy"
)

// FavouriteCreate is the builder for creating a Favourite entity.
type FavouriteCreate struct {
	config
	mutation *FavouriteMutation
	hooks    []Hook
}

// SetOwnerID sets the "owner_id" field.
func (fc *FavouriteCreate) SetOwnerID(s string) *FavouriteCreate {
	fc.mutation.SetOwnerID(s)
	return fc
}

// AddProxyIDs adds the "proxies" edge to the Proxy entity by IDs.
func (fc *FavouriteCreate) AddProxyIDs(ids ...int) *FavouriteCreate {
	fc.mutation.AddProxyIDs(ids...)
	return fc
}

// AddProxies adds the "proxies" edges to the Proxy entity.
func (fc *FavouriteCreate) AddProxies(p ...*Proxy) *FavouriteCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return fc.AddProxyIDs(ids...)
}

// Mutation returns the FavouriteMutation object of the builder.
func (fc *FavouriteCreate) Mutation() *FavouriteMutation {
	return fc.mutation
}

// Save creates the Favourite in the database.
func (fc *FavouriteCreate) Save(ctx context.Context) (*Favourite, error) {
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FavouriteCreate) SaveX(ctx context.Context) *Favourite {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FavouriteCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FavouriteCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FavouriteCreate) check() error {
	if _, ok := fc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Favourite.owner_id"`)}
	}
	if v, ok := fc.mutation.OwnerID(); ok {
		if err := favourite.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`ent: validator failed for field "Favourite.owner_id": %w`, err)}
		}
	}
	return nil
}

func (fc *FavouriteCreate) sqlSave(ctx context.Context) (*Favourite, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FavouriteCreate) createSpec() (*Favourite, *sqlgraph.CreateSpec) {
	var (
		_node = &Favourite{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(favourite.Table, sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeInt))
	)
	if value, ok := fc.mutation.OwnerID(); ok {
		_spec.SetField(favourite.FieldOwnerID, field.TypeString, value)
		_node.OwnerID = value
	}
	if nodes := fc.mutation.ProxiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   favourite.ProxiesTable,
			Columns: []string{favourite.ProxiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proxy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FavouriteCreateBulk is the builder for creating many Favourite entities in bulk.
type FavouriteCreateBulk struct {
	config
	err      error
	builders []*FavouriteCreate
}

// Save creates the Favourite entities in the database.
func (fcb *FavouriteCreateBulk) Save(ctx context.Context) ([]*Favourite, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Favourite, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FavouriteMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FavouriteCreateBulk) SaveX(ctx context.Context) []*Favourite {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FavouriteCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FavouriteCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
