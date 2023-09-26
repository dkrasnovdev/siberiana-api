// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/project"
	"github.com/dkrasnovdev/siberiana-api/ent/projecttype"
)

// ProjectTypeCreate is the builder for creating a ProjectType entity.
type ProjectTypeCreate struct {
	config
	mutation *ProjectTypeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ptc *ProjectTypeCreate) SetCreatedAt(t time.Time) *ProjectTypeCreate {
	ptc.mutation.SetCreatedAt(t)
	return ptc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptc *ProjectTypeCreate) SetNillableCreatedAt(t *time.Time) *ProjectTypeCreate {
	if t != nil {
		ptc.SetCreatedAt(*t)
	}
	return ptc
}

// SetCreatedBy sets the "created_by" field.
func (ptc *ProjectTypeCreate) SetCreatedBy(s string) *ProjectTypeCreate {
	ptc.mutation.SetCreatedBy(s)
	return ptc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ptc *ProjectTypeCreate) SetNillableCreatedBy(s *string) *ProjectTypeCreate {
	if s != nil {
		ptc.SetCreatedBy(*s)
	}
	return ptc
}

// SetUpdatedAt sets the "updated_at" field.
func (ptc *ProjectTypeCreate) SetUpdatedAt(t time.Time) *ProjectTypeCreate {
	ptc.mutation.SetUpdatedAt(t)
	return ptc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ptc *ProjectTypeCreate) SetNillableUpdatedAt(t *time.Time) *ProjectTypeCreate {
	if t != nil {
		ptc.SetUpdatedAt(*t)
	}
	return ptc
}

// SetUpdatedBy sets the "updated_by" field.
func (ptc *ProjectTypeCreate) SetUpdatedBy(s string) *ProjectTypeCreate {
	ptc.mutation.SetUpdatedBy(s)
	return ptc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ptc *ProjectTypeCreate) SetNillableUpdatedBy(s *string) *ProjectTypeCreate {
	if s != nil {
		ptc.SetUpdatedBy(*s)
	}
	return ptc
}

// SetDisplayName sets the "display_name" field.
func (ptc *ProjectTypeCreate) SetDisplayName(s string) *ProjectTypeCreate {
	ptc.mutation.SetDisplayName(s)
	return ptc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ptc *ProjectTypeCreate) SetNillableDisplayName(s *string) *ProjectTypeCreate {
	if s != nil {
		ptc.SetDisplayName(*s)
	}
	return ptc
}

// SetAbbreviation sets the "abbreviation" field.
func (ptc *ProjectTypeCreate) SetAbbreviation(s string) *ProjectTypeCreate {
	ptc.mutation.SetAbbreviation(s)
	return ptc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ptc *ProjectTypeCreate) SetNillableAbbreviation(s *string) *ProjectTypeCreate {
	if s != nil {
		ptc.SetAbbreviation(*s)
	}
	return ptc
}

// SetDescription sets the "description" field.
func (ptc *ProjectTypeCreate) SetDescription(s string) *ProjectTypeCreate {
	ptc.mutation.SetDescription(s)
	return ptc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ptc *ProjectTypeCreate) SetNillableDescription(s *string) *ProjectTypeCreate {
	if s != nil {
		ptc.SetDescription(*s)
	}
	return ptc
}

// SetExternalLink sets the "external_link" field.
func (ptc *ProjectTypeCreate) SetExternalLink(s string) *ProjectTypeCreate {
	ptc.mutation.SetExternalLink(s)
	return ptc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (ptc *ProjectTypeCreate) SetNillableExternalLink(s *string) *ProjectTypeCreate {
	if s != nil {
		ptc.SetExternalLink(*s)
	}
	return ptc
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (ptc *ProjectTypeCreate) AddProjectIDs(ids ...int) *ProjectTypeCreate {
	ptc.mutation.AddProjectIDs(ids...)
	return ptc
}

// AddProjects adds the "projects" edges to the Project entity.
func (ptc *ProjectTypeCreate) AddProjects(p ...*Project) *ProjectTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptc.AddProjectIDs(ids...)
}

// Mutation returns the ProjectTypeMutation object of the builder.
func (ptc *ProjectTypeCreate) Mutation() *ProjectTypeMutation {
	return ptc.mutation
}

// Save creates the ProjectType in the database.
func (ptc *ProjectTypeCreate) Save(ctx context.Context) (*ProjectType, error) {
	if err := ptc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ptc.sqlSave, ptc.mutation, ptc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *ProjectTypeCreate) SaveX(ctx context.Context) *ProjectType {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *ProjectTypeCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *ProjectTypeCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptc *ProjectTypeCreate) defaults() error {
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		if projecttype.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized projecttype.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := projecttype.DefaultCreatedAt()
		ptc.mutation.SetCreatedAt(v)
	}
	if _, ok := ptc.mutation.UpdatedAt(); !ok {
		if projecttype.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized projecttype.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := projecttype.DefaultUpdatedAt()
		ptc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ptc *ProjectTypeCreate) check() error {
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProjectType.created_at"`)}
	}
	if _, ok := ptc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProjectType.updated_at"`)}
	}
	return nil
}

func (ptc *ProjectTypeCreate) sqlSave(ctx context.Context) (*ProjectType, error) {
	if err := ptc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ptc.mutation.id = &_node.ID
	ptc.mutation.done = true
	return _node, nil
}

func (ptc *ProjectTypeCreate) createSpec() (*ProjectType, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectType{config: ptc.config}
		_spec = sqlgraph.NewCreateSpec(projecttype.Table, sqlgraph.NewFieldSpec(projecttype.FieldID, field.TypeInt))
	)
	if value, ok := ptc.mutation.CreatedAt(); ok {
		_spec.SetField(projecttype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ptc.mutation.CreatedBy(); ok {
		_spec.SetField(projecttype.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ptc.mutation.UpdatedAt(); ok {
		_spec.SetField(projecttype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ptc.mutation.UpdatedBy(); ok {
		_spec.SetField(projecttype.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ptc.mutation.DisplayName(); ok {
		_spec.SetField(projecttype.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := ptc.mutation.Abbreviation(); ok {
		_spec.SetField(projecttype.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := ptc.mutation.Description(); ok {
		_spec.SetField(projecttype.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ptc.mutation.ExternalLink(); ok {
		_spec.SetField(projecttype.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if nodes := ptc.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttype.ProjectsTable,
			Columns: []string{projecttype.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectTypeCreateBulk is the builder for creating many ProjectType entities in bulk.
type ProjectTypeCreateBulk struct {
	config
	err      error
	builders []*ProjectTypeCreate
}

// Save creates the ProjectType entities in the database.
func (ptcb *ProjectTypeCreateBulk) Save(ctx context.Context) ([]*ProjectType, error) {
	if ptcb.err != nil {
		return nil, ptcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*ProjectType, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *ProjectTypeCreateBulk) SaveX(ctx context.Context) []*ProjectType {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *ProjectTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *ProjectTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}
