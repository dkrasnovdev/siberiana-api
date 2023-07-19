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
	"github.com/dkrasnovdev/heritage-api/ent/model"
)

// ModelCreate is the builder for creating a Model entity.
type ModelCreate struct {
	config
	mutation *ModelMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *ModelCreate) SetCreatedAt(t time.Time) *ModelCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *ModelCreate) SetNillableCreatedAt(t *time.Time) *ModelCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetCreatedBy sets the "created_by" field.
func (mc *ModelCreate) SetCreatedBy(s string) *ModelCreate {
	mc.mutation.SetCreatedBy(s)
	return mc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mc *ModelCreate) SetNillableCreatedBy(s *string) *ModelCreate {
	if s != nil {
		mc.SetCreatedBy(*s)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *ModelCreate) SetUpdatedAt(t time.Time) *ModelCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *ModelCreate) SetNillableUpdatedAt(t *time.Time) *ModelCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetUpdatedBy sets the "updated_by" field.
func (mc *ModelCreate) SetUpdatedBy(s string) *ModelCreate {
	mc.mutation.SetUpdatedBy(s)
	return mc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mc *ModelCreate) SetNillableUpdatedBy(s *string) *ModelCreate {
	if s != nil {
		mc.SetUpdatedBy(*s)
	}
	return mc
}

// SetDisplayName sets the "display_name" field.
func (mc *ModelCreate) SetDisplayName(s string) *ModelCreate {
	mc.mutation.SetDisplayName(s)
	return mc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (mc *ModelCreate) SetNillableDisplayName(s *string) *ModelCreate {
	if s != nil {
		mc.SetDisplayName(*s)
	}
	return mc
}

// SetDescription sets the "description" field.
func (mc *ModelCreate) SetDescription(s string) *ModelCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *ModelCreate) SetNillableDescription(s *string) *ModelCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (mc *ModelCreate) AddArtifactIDs(ids ...int) *ModelCreate {
	mc.mutation.AddArtifactIDs(ids...)
	return mc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (mc *ModelCreate) AddArtifacts(a ...*Artifact) *ModelCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mc.AddArtifactIDs(ids...)
}

// Mutation returns the ModelMutation object of the builder.
func (mc *ModelCreate) Mutation() *ModelMutation {
	return mc.mutation
}

// Save creates the Model in the database.
func (mc *ModelCreate) Save(ctx context.Context) (*Model, error) {
	if err := mc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *ModelCreate) SaveX(ctx context.Context) *Model {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *ModelCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *ModelCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *ModelCreate) defaults() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		if model.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized model.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := model.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		if model.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized model.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := model.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mc *ModelCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Model.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Model.updated_at"`)}
	}
	return nil
}

func (mc *ModelCreate) sqlSave(ctx context.Context) (*Model, error) {
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

func (mc *ModelCreate) createSpec() (*Model, *sqlgraph.CreateSpec) {
	var (
		_node = &Model{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(model.Table, sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(model.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.CreatedBy(); ok {
		_spec.SetField(model.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(model.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.UpdatedBy(); ok {
		_spec.SetField(model.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := mc.mutation.DisplayName(); ok {
		_spec.SetField(model.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(model.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := mc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ArtifactsTable,
			Columns: []string{model.ArtifactsColumn},
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

// ModelCreateBulk is the builder for creating many Model entities in bulk.
type ModelCreateBulk struct {
	config
	builders []*ModelCreate
}

// Save creates the Model entities in the database.
func (mcb *ModelCreateBulk) Save(ctx context.Context) ([]*Model, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Model, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ModelMutation)
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
func (mcb *ModelCreateBulk) SaveX(ctx context.Context) []*Model {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *ModelCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *ModelCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
