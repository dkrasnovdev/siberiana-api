// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/artifact"
	"github.com/dkrasnovdev/heritage-api/ent/culture"
)

// CultureCreate is the builder for creating a Culture entity.
type CultureCreate struct {
	config
	mutation *CultureMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CultureCreate) SetCreatedAt(t time.Time) *CultureCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CultureCreate) SetNillableCreatedAt(t *time.Time) *CultureCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetCreatedBy sets the "created_by" field.
func (cc *CultureCreate) SetCreatedBy(s string) *CultureCreate {
	cc.mutation.SetCreatedBy(s)
	return cc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cc *CultureCreate) SetNillableCreatedBy(s *string) *CultureCreate {
	if s != nil {
		cc.SetCreatedBy(*s)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CultureCreate) SetUpdatedAt(t time.Time) *CultureCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CultureCreate) SetNillableUpdatedAt(t *time.Time) *CultureCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetUpdatedBy sets the "updated_by" field.
func (cc *CultureCreate) SetUpdatedBy(s string) *CultureCreate {
	cc.mutation.SetUpdatedBy(s)
	return cc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cc *CultureCreate) SetNillableUpdatedBy(s *string) *CultureCreate {
	if s != nil {
		cc.SetUpdatedBy(*s)
	}
	return cc
}

// SetDisplayName sets the "display_name" field.
func (cc *CultureCreate) SetDisplayName(s string) *CultureCreate {
	cc.mutation.SetDisplayName(s)
	return cc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cc *CultureCreate) SetNillableDisplayName(s *string) *CultureCreate {
	if s != nil {
		cc.SetDisplayName(*s)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *CultureCreate) SetDescription(s string) *CultureCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CultureCreate) SetNillableDescription(s *string) *CultureCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetExternalLinks sets the "external_links" field.
func (cc *CultureCreate) SetExternalLinks(s []string) *CultureCreate {
	cc.mutation.SetExternalLinks(s)
	return cc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (cc *CultureCreate) AddArtifactIDs(ids ...int) *CultureCreate {
	cc.mutation.AddArtifactIDs(ids...)
	return cc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (cc *CultureCreate) AddArtifacts(a ...*Artifact) *CultureCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddArtifactIDs(ids...)
}

// Mutation returns the CultureMutation object of the builder.
func (cc *CultureCreate) Mutation() *CultureMutation {
	return cc.mutation
}

// Save creates the Culture in the database.
func (cc *CultureCreate) Save(ctx context.Context) (*Culture, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CultureCreate) SaveX(ctx context.Context) *Culture {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CultureCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CultureCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CultureCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if culture.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized culture.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := culture.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if culture.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized culture.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := culture.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CultureCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Culture.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Culture.updated_at"`)}
	}
	return nil
}

func (cc *CultureCreate) sqlSave(ctx context.Context) (*Culture, error) {
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

func (cc *CultureCreate) createSpec() (*Culture, *sqlgraph.CreateSpec) {
	var (
		_node = &Culture{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(culture.Table, sqlgraph.NewFieldSpec(culture.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(culture.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.CreatedBy(); ok {
		_spec.SetField(culture.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(culture.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.UpdatedBy(); ok {
		_spec.SetField(culture.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := cc.mutation.DisplayName(); ok {
		_spec.SetField(culture.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(culture.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.ExternalLinks(); ok {
		_spec.SetField(culture.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
	if nodes := cc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   culture.ArtifactsTable,
			Columns: []string{culture.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CultureCreateBulk is the builder for creating many Culture entities in bulk.
type CultureCreateBulk struct {
	config
	builders []*CultureCreate
}

// Save creates the Culture entities in the database.
func (ccb *CultureCreateBulk) Save(ctx context.Context) ([]*Culture, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Culture, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CultureMutation)
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
func (ccb *CultureCreateBulk) SaveX(ctx context.Context) []*Culture {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CultureCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CultureCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
