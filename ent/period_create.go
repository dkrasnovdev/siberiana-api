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
	"github.com/dkrasnovdev/siberiana-api/ent/period"
)

// PeriodCreate is the builder for creating a Period entity.
type PeriodCreate struct {
	config
	mutation *PeriodMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PeriodCreate) SetCreatedAt(t time.Time) *PeriodCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PeriodCreate) SetNillableCreatedAt(t *time.Time) *PeriodCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *PeriodCreate) SetCreatedBy(s string) *PeriodCreate {
	pc.mutation.SetCreatedBy(s)
	return pc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pc *PeriodCreate) SetNillableCreatedBy(s *string) *PeriodCreate {
	if s != nil {
		pc.SetCreatedBy(*s)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PeriodCreate) SetUpdatedAt(t time.Time) *PeriodCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PeriodCreate) SetNillableUpdatedAt(t *time.Time) *PeriodCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *PeriodCreate) SetUpdatedBy(s string) *PeriodCreate {
	pc.mutation.SetUpdatedBy(s)
	return pc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pc *PeriodCreate) SetNillableUpdatedBy(s *string) *PeriodCreate {
	if s != nil {
		pc.SetUpdatedBy(*s)
	}
	return pc
}

// SetDisplayName sets the "display_name" field.
func (pc *PeriodCreate) SetDisplayName(s string) *PeriodCreate {
	pc.mutation.SetDisplayName(s)
	return pc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pc *PeriodCreate) SetNillableDisplayName(s *string) *PeriodCreate {
	if s != nil {
		pc.SetDisplayName(*s)
	}
	return pc
}

// SetAbbreviation sets the "abbreviation" field.
func (pc *PeriodCreate) SetAbbreviation(s string) *PeriodCreate {
	pc.mutation.SetAbbreviation(s)
	return pc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pc *PeriodCreate) SetNillableAbbreviation(s *string) *PeriodCreate {
	if s != nil {
		pc.SetAbbreviation(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *PeriodCreate) SetDescription(s string) *PeriodCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PeriodCreate) SetNillableDescription(s *string) *PeriodCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetExternalLink sets the "external_link" field.
func (pc *PeriodCreate) SetExternalLink(s string) *PeriodCreate {
	pc.mutation.SetExternalLink(s)
	return pc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (pc *PeriodCreate) SetNillableExternalLink(s *string) *PeriodCreate {
	if s != nil {
		pc.SetExternalLink(*s)
	}
	return pc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (pc *PeriodCreate) AddArtifactIDs(ids ...int) *PeriodCreate {
	pc.mutation.AddArtifactIDs(ids...)
	return pc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (pc *PeriodCreate) AddArtifacts(a ...*Artifact) *PeriodCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddArtifactIDs(ids...)
}

// Mutation returns the PeriodMutation object of the builder.
func (pc *PeriodCreate) Mutation() *PeriodMutation {
	return pc.mutation
}

// Save creates the Period in the database.
func (pc *PeriodCreate) Save(ctx context.Context) (*Period, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PeriodCreate) SaveX(ctx context.Context) *Period {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PeriodCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PeriodCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PeriodCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if period.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized period.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := period.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if period.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized period.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := period.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PeriodCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Period.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Period.updated_at"`)}
	}
	return nil
}

func (pc *PeriodCreate) sqlSave(ctx context.Context) (*Period, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PeriodCreate) createSpec() (*Period, *sqlgraph.CreateSpec) {
	var (
		_node = &Period{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(period.Table, sqlgraph.NewFieldSpec(period.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(period.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(period.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(period.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(period.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.DisplayName(); ok {
		_spec.SetField(period.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := pc.mutation.Abbreviation(); ok {
		_spec.SetField(period.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(period.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.ExternalLink(); ok {
		_spec.SetField(period.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if nodes := pc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   period.ArtifactsTable,
			Columns: []string{period.ArtifactsColumn},
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

// PeriodCreateBulk is the builder for creating many Period entities in bulk.
type PeriodCreateBulk struct {
	config
	err      error
	builders []*PeriodCreate
}

// Save creates the Period entities in the database.
func (pcb *PeriodCreateBulk) Save(ctx context.Context) ([]*Period, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Period, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PeriodMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PeriodCreateBulk) SaveX(ctx context.Context) []*Period {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PeriodCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PeriodCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
