// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/book"
	"github.com/dkrasnovdev/heritage-api/ent/publisher"
)

// PublisherCreate is the builder for creating a Publisher entity.
type PublisherCreate struct {
	config
	mutation *PublisherMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PublisherCreate) SetCreatedAt(t time.Time) *PublisherCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableCreatedAt(t *time.Time) *PublisherCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *PublisherCreate) SetCreatedBy(s string) *PublisherCreate {
	pc.mutation.SetCreatedBy(s)
	return pc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableCreatedBy(s *string) *PublisherCreate {
	if s != nil {
		pc.SetCreatedBy(*s)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PublisherCreate) SetUpdatedAt(t time.Time) *PublisherCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableUpdatedAt(t *time.Time) *PublisherCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *PublisherCreate) SetUpdatedBy(s string) *PublisherCreate {
	pc.mutation.SetUpdatedBy(s)
	return pc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableUpdatedBy(s *string) *PublisherCreate {
	if s != nil {
		pc.SetUpdatedBy(*s)
	}
	return pc
}

// SetDisplayName sets the "display_name" field.
func (pc *PublisherCreate) SetDisplayName(s string) *PublisherCreate {
	pc.mutation.SetDisplayName(s)
	return pc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableDisplayName(s *string) *PublisherCreate {
	if s != nil {
		pc.SetDisplayName(*s)
	}
	return pc
}

// SetAbbreviation sets the "abbreviation" field.
func (pc *PublisherCreate) SetAbbreviation(s string) *PublisherCreate {
	pc.mutation.SetAbbreviation(s)
	return pc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableAbbreviation(s *string) *PublisherCreate {
	if s != nil {
		pc.SetAbbreviation(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *PublisherCreate) SetDescription(s string) *PublisherCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableDescription(s *string) *PublisherCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetExternalLinks sets the "external_links" field.
func (pc *PublisherCreate) SetExternalLinks(s []string) *PublisherCreate {
	pc.mutation.SetExternalLinks(s)
	return pc
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (pc *PublisherCreate) AddBookIDs(ids ...int) *PublisherCreate {
	pc.mutation.AddBookIDs(ids...)
	return pc
}

// AddBooks adds the "books" edges to the Book entity.
func (pc *PublisherCreate) AddBooks(b ...*Book) *PublisherCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBookIDs(ids...)
}

// Mutation returns the PublisherMutation object of the builder.
func (pc *PublisherCreate) Mutation() *PublisherMutation {
	return pc.mutation
}

// Save creates the Publisher in the database.
func (pc *PublisherCreate) Save(ctx context.Context) (*Publisher, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PublisherCreate) SaveX(ctx context.Context) *Publisher {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PublisherCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PublisherCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PublisherCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if publisher.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized publisher.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := publisher.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if publisher.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized publisher.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := publisher.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PublisherCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Publisher.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Publisher.updated_at"`)}
	}
	return nil
}

func (pc *PublisherCreate) sqlSave(ctx context.Context) (*Publisher, error) {
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

func (pc *PublisherCreate) createSpec() (*Publisher, *sqlgraph.CreateSpec) {
	var (
		_node = &Publisher{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(publisher.Table, sqlgraph.NewFieldSpec(publisher.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(publisher.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(publisher.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(publisher.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(publisher.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.DisplayName(); ok {
		_spec.SetField(publisher.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := pc.mutation.Abbreviation(); ok {
		_spec.SetField(publisher.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(publisher.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.ExternalLinks(); ok {
		_spec.SetField(publisher.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
	if nodes := pc.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   publisher.BooksTable,
			Columns: []string{publisher.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PublisherCreateBulk is the builder for creating many Publisher entities in bulk.
type PublisherCreateBulk struct {
	config
	builders []*PublisherCreate
}

// Save creates the Publisher entities in the database.
func (pcb *PublisherCreateBulk) Save(ctx context.Context) ([]*Publisher, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Publisher, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublisherMutation)
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
func (pcb *PublisherCreateBulk) SaveX(ctx context.Context) []*Publisher {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PublisherCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PublisherCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
