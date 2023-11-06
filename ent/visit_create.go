// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/mound"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/visit"
)

// VisitCreate is the builder for creating a Visit entity.
type VisitCreate struct {
	config
	mutation *VisitMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (vc *VisitCreate) SetCreatedAt(t time.Time) *VisitCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VisitCreate) SetNillableCreatedAt(t *time.Time) *VisitCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetCreatedBy sets the "created_by" field.
func (vc *VisitCreate) SetCreatedBy(s string) *VisitCreate {
	vc.mutation.SetCreatedBy(s)
	return vc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vc *VisitCreate) SetNillableCreatedBy(s *string) *VisitCreate {
	if s != nil {
		vc.SetCreatedBy(*s)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VisitCreate) SetUpdatedAt(t time.Time) *VisitCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VisitCreate) SetNillableUpdatedAt(t *time.Time) *VisitCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetUpdatedBy sets the "updated_by" field.
func (vc *VisitCreate) SetUpdatedBy(s string) *VisitCreate {
	vc.mutation.SetUpdatedBy(s)
	return vc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vc *VisitCreate) SetNillableUpdatedBy(s *string) *VisitCreate {
	if s != nil {
		vc.SetUpdatedBy(*s)
	}
	return vc
}

// SetYear sets the "year" field.
func (vc *VisitCreate) SetYear(i int) *VisitCreate {
	vc.mutation.SetYear(i)
	return vc
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (vc *VisitCreate) SetNillableYear(i *int) *VisitCreate {
	if i != nil {
		vc.SetYear(*i)
	}
	return vc
}

// AddMoundIDs adds the "mounds" edge to the Mound entity by IDs.
func (vc *VisitCreate) AddMoundIDs(ids ...int) *VisitCreate {
	vc.mutation.AddMoundIDs(ids...)
	return vc
}

// AddMounds adds the "mounds" edges to the Mound entity.
func (vc *VisitCreate) AddMounds(m ...*Mound) *VisitCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return vc.AddMoundIDs(ids...)
}

// AddVisitorIDs adds the "visitors" edge to the Person entity by IDs.
func (vc *VisitCreate) AddVisitorIDs(ids ...int) *VisitCreate {
	vc.mutation.AddVisitorIDs(ids...)
	return vc
}

// AddVisitors adds the "visitors" edges to the Person entity.
func (vc *VisitCreate) AddVisitors(p ...*Person) *VisitCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vc.AddVisitorIDs(ids...)
}

// Mutation returns the VisitMutation object of the builder.
func (vc *VisitCreate) Mutation() *VisitMutation {
	return vc.mutation
}

// Save creates the Visit in the database.
func (vc *VisitCreate) Save(ctx context.Context) (*Visit, error) {
	if err := vc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VisitCreate) SaveX(ctx context.Context) *Visit {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VisitCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VisitCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VisitCreate) defaults() error {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		if visit.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized visit.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := visit.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		if visit.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized visit.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := visit.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vc *VisitCreate) check() error {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Visit.created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Visit.updated_at"`)}
	}
	if v, ok := vc.mutation.Year(); ok {
		if err := visit.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Visit.year": %w`, err)}
		}
	}
	return nil
}

func (vc *VisitCreate) sqlSave(ctx context.Context) (*Visit, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VisitCreate) createSpec() (*Visit, *sqlgraph.CreateSpec) {
	var (
		_node = &Visit{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(visit.Table, sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(visit.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.CreatedBy(); ok {
		_spec.SetField(visit.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(visit.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vc.mutation.UpdatedBy(); ok {
		_spec.SetField(visit.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := vc.mutation.Year(); ok {
		_spec.SetField(visit.FieldYear, field.TypeInt, value)
		_node.Year = value
	}
	if nodes := vc.mutation.MoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   visit.MoundsTable,
			Columns: visit.MoundsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mound.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.VisitorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   visit.VisitorsTable,
			Columns: visit.VisitorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VisitCreateBulk is the builder for creating many Visit entities in bulk.
type VisitCreateBulk struct {
	config
	err      error
	builders []*VisitCreate
}

// Save creates the Visit entities in the database.
func (vcb *VisitCreateBulk) Save(ctx context.Context) ([]*Visit, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Visit, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VisitMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VisitCreateBulk) SaveX(ctx context.Context) []*Visit {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VisitCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VisitCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
