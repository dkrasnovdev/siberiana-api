// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/artifact"
	"github.com/dkrasnovdev/siberiana-api/ent/monument"
	"github.com/dkrasnovdev/siberiana-api/ent/set"
)

// MonumentCreate is the builder for creating a Monument entity.
type MonumentCreate struct {
	config
	mutation *MonumentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MonumentCreate) SetCreatedAt(t time.Time) *MonumentCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MonumentCreate) SetNillableCreatedAt(t *time.Time) *MonumentCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetCreatedBy sets the "created_by" field.
func (mc *MonumentCreate) SetCreatedBy(s string) *MonumentCreate {
	mc.mutation.SetCreatedBy(s)
	return mc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mc *MonumentCreate) SetNillableCreatedBy(s *string) *MonumentCreate {
	if s != nil {
		mc.SetCreatedBy(*s)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MonumentCreate) SetUpdatedAt(t time.Time) *MonumentCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MonumentCreate) SetNillableUpdatedAt(t *time.Time) *MonumentCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetUpdatedBy sets the "updated_by" field.
func (mc *MonumentCreate) SetUpdatedBy(s string) *MonumentCreate {
	mc.mutation.SetUpdatedBy(s)
	return mc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mc *MonumentCreate) SetNillableUpdatedBy(s *string) *MonumentCreate {
	if s != nil {
		mc.SetUpdatedBy(*s)
	}
	return mc
}

// SetDisplayName sets the "display_name" field.
func (mc *MonumentCreate) SetDisplayName(s string) *MonumentCreate {
	mc.mutation.SetDisplayName(s)
	return mc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (mc *MonumentCreate) SetNillableDisplayName(s *string) *MonumentCreate {
	if s != nil {
		mc.SetDisplayName(*s)
	}
	return mc
}

// SetAbbreviation sets the "abbreviation" field.
func (mc *MonumentCreate) SetAbbreviation(s string) *MonumentCreate {
	mc.mutation.SetAbbreviation(s)
	return mc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (mc *MonumentCreate) SetNillableAbbreviation(s *string) *MonumentCreate {
	if s != nil {
		mc.SetAbbreviation(*s)
	}
	return mc
}

// SetDescription sets the "description" field.
func (mc *MonumentCreate) SetDescription(s string) *MonumentCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *MonumentCreate) SetNillableDescription(s *string) *MonumentCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// SetExternalLink sets the "external_link" field.
func (mc *MonumentCreate) SetExternalLink(s string) *MonumentCreate {
	mc.mutation.SetExternalLink(s)
	return mc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (mc *MonumentCreate) SetNillableExternalLink(s *string) *MonumentCreate {
	if s != nil {
		mc.SetExternalLink(*s)
	}
	return mc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (mc *MonumentCreate) AddArtifactIDs(ids ...int) *MonumentCreate {
	mc.mutation.AddArtifactIDs(ids...)
	return mc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (mc *MonumentCreate) AddArtifacts(a ...*Artifact) *MonumentCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mc.AddArtifactIDs(ids...)
}

// AddSetIDs adds the "sets" edge to the Set entity by IDs.
func (mc *MonumentCreate) AddSetIDs(ids ...int) *MonumentCreate {
	mc.mutation.AddSetIDs(ids...)
	return mc
}

// AddSets adds the "sets" edges to the Set entity.
func (mc *MonumentCreate) AddSets(s ...*Set) *MonumentCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mc.AddSetIDs(ids...)
}

// Mutation returns the MonumentMutation object of the builder.
func (mc *MonumentCreate) Mutation() *MonumentMutation {
	return mc.mutation
}

// Save creates the Monument in the database.
func (mc *MonumentCreate) Save(ctx context.Context) (*Monument, error) {
	if err := mc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MonumentCreate) SaveX(ctx context.Context) *Monument {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MonumentCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MonumentCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MonumentCreate) defaults() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		if monument.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized monument.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := monument.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		if monument.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized monument.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := monument.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mc *MonumentCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Monument.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Monument.updated_at"`)}
	}
	return nil
}

func (mc *MonumentCreate) sqlSave(ctx context.Context) (*Monument, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MonumentCreate) createSpec() (*Monument, *sqlgraph.CreateSpec) {
	var (
		_node = &Monument{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(monument.Table, sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(monument.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.CreatedBy(); ok {
		_spec.SetField(monument.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(monument.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.UpdatedBy(); ok {
		_spec.SetField(monument.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := mc.mutation.DisplayName(); ok {
		_spec.SetField(monument.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := mc.mutation.Abbreviation(); ok {
		_spec.SetField(monument.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(monument.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mc.mutation.ExternalLink(); ok {
		_spec.SetField(monument.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if nodes := mc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   monument.ArtifactsTable,
			Columns: []string{monument.ArtifactsColumn},
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
	if nodes := mc.mutation.SetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   monument.SetsTable,
			Columns: monument.SetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(set.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MonumentCreateBulk is the builder for creating many Monument entities in bulk.
type MonumentCreateBulk struct {
	config
	err      error
	builders []*MonumentCreate
}

// Save creates the Monument entities in the database.
func (mcb *MonumentCreateBulk) Save(ctx context.Context) ([]*Monument, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Monument, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MonumentMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MonumentCreateBulk) SaveX(ctx context.Context) []*Monument {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MonumentCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MonumentCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
