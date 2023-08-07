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
	"github.com/dkrasnovdev/heritage-api/ent/monument"
	"github.com/dkrasnovdev/heritage-api/ent/set"
)

// SetCreate is the builder for creating a Set entity.
type SetCreate struct {
	config
	mutation *SetMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SetCreate) SetCreatedAt(t time.Time) *SetCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SetCreate) SetNillableCreatedAt(t *time.Time) *SetCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *SetCreate) SetCreatedBy(s string) *SetCreate {
	sc.mutation.SetCreatedBy(s)
	return sc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sc *SetCreate) SetNillableCreatedBy(s *string) *SetCreate {
	if s != nil {
		sc.SetCreatedBy(*s)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SetCreate) SetUpdatedAt(t time.Time) *SetCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SetCreate) SetNillableUpdatedAt(t *time.Time) *SetCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *SetCreate) SetUpdatedBy(s string) *SetCreate {
	sc.mutation.SetUpdatedBy(s)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *SetCreate) SetNillableUpdatedBy(s *string) *SetCreate {
	if s != nil {
		sc.SetUpdatedBy(*s)
	}
	return sc
}

// SetDisplayName sets the "display_name" field.
func (sc *SetCreate) SetDisplayName(s string) *SetCreate {
	sc.mutation.SetDisplayName(s)
	return sc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (sc *SetCreate) SetNillableDisplayName(s *string) *SetCreate {
	if s != nil {
		sc.SetDisplayName(*s)
	}
	return sc
}

// SetAbbreviation sets the "abbreviation" field.
func (sc *SetCreate) SetAbbreviation(s string) *SetCreate {
	sc.mutation.SetAbbreviation(s)
	return sc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (sc *SetCreate) SetNillableAbbreviation(s *string) *SetCreate {
	if s != nil {
		sc.SetAbbreviation(*s)
	}
	return sc
}

// SetDescription sets the "description" field.
func (sc *SetCreate) SetDescription(s string) *SetCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *SetCreate) SetNillableDescription(s *string) *SetCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetExternalLink sets the "external_link" field.
func (sc *SetCreate) SetExternalLink(s string) *SetCreate {
	sc.mutation.SetExternalLink(s)
	return sc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (sc *SetCreate) SetNillableExternalLink(s *string) *SetCreate {
	if s != nil {
		sc.SetExternalLink(*s)
	}
	return sc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (sc *SetCreate) AddArtifactIDs(ids ...int) *SetCreate {
	sc.mutation.AddArtifactIDs(ids...)
	return sc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (sc *SetCreate) AddArtifacts(a ...*Artifact) *SetCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddArtifactIDs(ids...)
}

// AddMonumentIDs adds the "monuments" edge to the Monument entity by IDs.
func (sc *SetCreate) AddMonumentIDs(ids ...int) *SetCreate {
	sc.mutation.AddMonumentIDs(ids...)
	return sc
}

// AddMonuments adds the "monuments" edges to the Monument entity.
func (sc *SetCreate) AddMonuments(m ...*Monument) *SetCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return sc.AddMonumentIDs(ids...)
}

// Mutation returns the SetMutation object of the builder.
func (sc *SetCreate) Mutation() *SetMutation {
	return sc.mutation
}

// Save creates the Set in the database.
func (sc *SetCreate) Save(ctx context.Context) (*Set, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SetCreate) SaveX(ctx context.Context) *Set {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SetCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SetCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SetCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if set.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized set.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := set.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if set.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized set.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := set.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SetCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Set.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Set.updated_at"`)}
	}
	return nil
}

func (sc *SetCreate) sqlSave(ctx context.Context) (*Set, error) {
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

func (sc *SetCreate) createSpec() (*Set, *sqlgraph.CreateSpec) {
	var (
		_node = &Set{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(set.Table, sqlgraph.NewFieldSpec(set.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(set.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.SetField(set.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(set.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.SetField(set.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.DisplayName(); ok {
		_spec.SetField(set.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := sc.mutation.Abbreviation(); ok {
		_spec.SetField(set.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(set.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.ExternalLink(); ok {
		_spec.SetField(set.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if nodes := sc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   set.ArtifactsTable,
			Columns: []string{set.ArtifactsColumn},
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
	if nodes := sc.mutation.MonumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   set.MonumentsTable,
			Columns: set.MonumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetCreateBulk is the builder for creating many Set entities in bulk.
type SetCreateBulk struct {
	config
	builders []*SetCreate
}

// Save creates the Set entities in the database.
func (scb *SetCreateBulk) Save(ctx context.Context) ([]*Set, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Set, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SetMutation)
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
func (scb *SetCreateBulk) SaveX(ctx context.Context) []*Set {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SetCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SetCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
