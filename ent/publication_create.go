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
	"github.com/dkrasnovdev/heritage-api/ent/person"
	"github.com/dkrasnovdev/heritage-api/ent/publication"
)

// PublicationCreate is the builder for creating a Publication entity.
type PublicationCreate struct {
	config
	mutation *PublicationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PublicationCreate) SetCreatedAt(t time.Time) *PublicationCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PublicationCreate) SetNillableCreatedAt(t *time.Time) *PublicationCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *PublicationCreate) SetCreatedBy(s string) *PublicationCreate {
	pc.mutation.SetCreatedBy(s)
	return pc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pc *PublicationCreate) SetNillableCreatedBy(s *string) *PublicationCreate {
	if s != nil {
		pc.SetCreatedBy(*s)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PublicationCreate) SetUpdatedAt(t time.Time) *PublicationCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PublicationCreate) SetNillableUpdatedAt(t *time.Time) *PublicationCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *PublicationCreate) SetUpdatedBy(s string) *PublicationCreate {
	pc.mutation.SetUpdatedBy(s)
	return pc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pc *PublicationCreate) SetNillableUpdatedBy(s *string) *PublicationCreate {
	if s != nil {
		pc.SetUpdatedBy(*s)
	}
	return pc
}

// SetDisplayName sets the "display_name" field.
func (pc *PublicationCreate) SetDisplayName(s string) *PublicationCreate {
	pc.mutation.SetDisplayName(s)
	return pc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pc *PublicationCreate) SetNillableDisplayName(s *string) *PublicationCreate {
	if s != nil {
		pc.SetDisplayName(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *PublicationCreate) SetDescription(s string) *PublicationCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PublicationCreate) SetNillableDescription(s *string) *PublicationCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (pc *PublicationCreate) AddArtifactIDs(ids ...int) *PublicationCreate {
	pc.mutation.AddArtifactIDs(ids...)
	return pc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (pc *PublicationCreate) AddArtifacts(a ...*Artifact) *PublicationCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddArtifactIDs(ids...)
}

// AddAuthorIDs adds the "authors" edge to the Person entity by IDs.
func (pc *PublicationCreate) AddAuthorIDs(ids ...int) *PublicationCreate {
	pc.mutation.AddAuthorIDs(ids...)
	return pc
}

// AddAuthors adds the "authors" edges to the Person entity.
func (pc *PublicationCreate) AddAuthors(p ...*Person) *PublicationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddAuthorIDs(ids...)
}

// Mutation returns the PublicationMutation object of the builder.
func (pc *PublicationCreate) Mutation() *PublicationMutation {
	return pc.mutation
}

// Save creates the Publication in the database.
func (pc *PublicationCreate) Save(ctx context.Context) (*Publication, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PublicationCreate) SaveX(ctx context.Context) *Publication {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PublicationCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PublicationCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PublicationCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if publication.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized publication.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := publication.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if publication.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized publication.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := publication.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PublicationCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Publication.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Publication.updated_at"`)}
	}
	return nil
}

func (pc *PublicationCreate) sqlSave(ctx context.Context) (*Publication, error) {
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

func (pc *PublicationCreate) createSpec() (*Publication, *sqlgraph.CreateSpec) {
	var (
		_node = &Publication{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(publication.Table, sqlgraph.NewFieldSpec(publication.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(publication.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(publication.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(publication.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(publication.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.DisplayName(); ok {
		_spec.SetField(publication.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(publication.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := pc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   publication.ArtifactsTable,
			Columns: publication.ArtifactsPrimaryKey,
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
	if nodes := pc.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publication.AuthorsTable,
			Columns: publication.AuthorsPrimaryKey,
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

// PublicationCreateBulk is the builder for creating many Publication entities in bulk.
type PublicationCreateBulk struct {
	config
	builders []*PublicationCreate
}

// Save creates the Publication entities in the database.
func (pcb *PublicationCreateBulk) Save(ctx context.Context) ([]*Publication, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Publication, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublicationMutation)
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
func (pcb *PublicationCreateBulk) SaveX(ctx context.Context) []*Publication {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PublicationCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PublicationCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
