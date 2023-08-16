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
	"github.com/dkrasnovdev/siberiana-api/ent/technique"
)

// TechniqueCreate is the builder for creating a Technique entity.
type TechniqueCreate struct {
	config
	mutation *TechniqueMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TechniqueCreate) SetCreatedAt(t time.Time) *TechniqueCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableCreatedAt(t *time.Time) *TechniqueCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetCreatedBy sets the "created_by" field.
func (tc *TechniqueCreate) SetCreatedBy(s string) *TechniqueCreate {
	tc.mutation.SetCreatedBy(s)
	return tc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableCreatedBy(s *string) *TechniqueCreate {
	if s != nil {
		tc.SetCreatedBy(*s)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TechniqueCreate) SetUpdatedAt(t time.Time) *TechniqueCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableUpdatedAt(t *time.Time) *TechniqueCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetUpdatedBy sets the "updated_by" field.
func (tc *TechniqueCreate) SetUpdatedBy(s string) *TechniqueCreate {
	tc.mutation.SetUpdatedBy(s)
	return tc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableUpdatedBy(s *string) *TechniqueCreate {
	if s != nil {
		tc.SetUpdatedBy(*s)
	}
	return tc
}

// SetDisplayName sets the "display_name" field.
func (tc *TechniqueCreate) SetDisplayName(s string) *TechniqueCreate {
	tc.mutation.SetDisplayName(s)
	return tc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableDisplayName(s *string) *TechniqueCreate {
	if s != nil {
		tc.SetDisplayName(*s)
	}
	return tc
}

// SetAbbreviation sets the "abbreviation" field.
func (tc *TechniqueCreate) SetAbbreviation(s string) *TechniqueCreate {
	tc.mutation.SetAbbreviation(s)
	return tc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableAbbreviation(s *string) *TechniqueCreate {
	if s != nil {
		tc.SetAbbreviation(*s)
	}
	return tc
}

// SetDescription sets the "description" field.
func (tc *TechniqueCreate) SetDescription(s string) *TechniqueCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableDescription(s *string) *TechniqueCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetExternalLink sets the "external_link" field.
func (tc *TechniqueCreate) SetExternalLink(s string) *TechniqueCreate {
	tc.mutation.SetExternalLink(s)
	return tc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableExternalLink(s *string) *TechniqueCreate {
	if s != nil {
		tc.SetExternalLink(*s)
	}
	return tc
}

// SetSlug sets the "slug" field.
func (tc *TechniqueCreate) SetSlug(s string) *TechniqueCreate {
	tc.mutation.SetSlug(s)
	return tc
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tc *TechniqueCreate) SetNillableSlug(s *string) *TechniqueCreate {
	if s != nil {
		tc.SetSlug(*s)
	}
	return tc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (tc *TechniqueCreate) AddArtifactIDs(ids ...int) *TechniqueCreate {
	tc.mutation.AddArtifactIDs(ids...)
	return tc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (tc *TechniqueCreate) AddArtifacts(a ...*Artifact) *TechniqueCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tc.AddArtifactIDs(ids...)
}

// Mutation returns the TechniqueMutation object of the builder.
func (tc *TechniqueCreate) Mutation() *TechniqueMutation {
	return tc.mutation
}

// Save creates the Technique in the database.
func (tc *TechniqueCreate) Save(ctx context.Context) (*Technique, error) {
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TechniqueCreate) SaveX(ctx context.Context) *Technique {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TechniqueCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TechniqueCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TechniqueCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if technique.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized technique.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := technique.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		if technique.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized technique.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := technique.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TechniqueCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Technique.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Technique.updated_at"`)}
	}
	return nil
}

func (tc *TechniqueCreate) sqlSave(ctx context.Context) (*Technique, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TechniqueCreate) createSpec() (*Technique, *sqlgraph.CreateSpec) {
	var (
		_node = &Technique{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(technique.Table, sqlgraph.NewFieldSpec(technique.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(technique.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.CreatedBy(); ok {
		_spec.SetField(technique.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(technique.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.UpdatedBy(); ok {
		_spec.SetField(technique.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := tc.mutation.DisplayName(); ok {
		_spec.SetField(technique.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := tc.mutation.Abbreviation(); ok {
		_spec.SetField(technique.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(technique.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.ExternalLink(); ok {
		_spec.SetField(technique.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := tc.mutation.Slug(); ok {
		_spec.SetField(technique.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if nodes := tc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   technique.ArtifactsTable,
			Columns: technique.ArtifactsPrimaryKey,
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

// TechniqueCreateBulk is the builder for creating many Technique entities in bulk.
type TechniqueCreateBulk struct {
	config
	builders []*TechniqueCreate
}

// Save creates the Technique entities in the database.
func (tcb *TechniqueCreateBulk) Save(ctx context.Context) ([]*Technique, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Technique, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TechniqueMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TechniqueCreateBulk) SaveX(ctx context.Context) []*Technique {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TechniqueCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TechniqueCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
