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
	"github.com/dkrasnovdev/heritage-api/ent/bookgenre"
)

// BookGenreCreate is the builder for creating a BookGenre entity.
type BookGenreCreate struct {
	config
	mutation *BookGenreMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bgc *BookGenreCreate) SetCreatedAt(t time.Time) *BookGenreCreate {
	bgc.mutation.SetCreatedAt(t)
	return bgc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bgc *BookGenreCreate) SetNillableCreatedAt(t *time.Time) *BookGenreCreate {
	if t != nil {
		bgc.SetCreatedAt(*t)
	}
	return bgc
}

// SetCreatedBy sets the "created_by" field.
func (bgc *BookGenreCreate) SetCreatedBy(s string) *BookGenreCreate {
	bgc.mutation.SetCreatedBy(s)
	return bgc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (bgc *BookGenreCreate) SetNillableCreatedBy(s *string) *BookGenreCreate {
	if s != nil {
		bgc.SetCreatedBy(*s)
	}
	return bgc
}

// SetUpdatedAt sets the "updated_at" field.
func (bgc *BookGenreCreate) SetUpdatedAt(t time.Time) *BookGenreCreate {
	bgc.mutation.SetUpdatedAt(t)
	return bgc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bgc *BookGenreCreate) SetNillableUpdatedAt(t *time.Time) *BookGenreCreate {
	if t != nil {
		bgc.SetUpdatedAt(*t)
	}
	return bgc
}

// SetUpdatedBy sets the "updated_by" field.
func (bgc *BookGenreCreate) SetUpdatedBy(s string) *BookGenreCreate {
	bgc.mutation.SetUpdatedBy(s)
	return bgc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (bgc *BookGenreCreate) SetNillableUpdatedBy(s *string) *BookGenreCreate {
	if s != nil {
		bgc.SetUpdatedBy(*s)
	}
	return bgc
}

// SetDisplayName sets the "display_name" field.
func (bgc *BookGenreCreate) SetDisplayName(s string) *BookGenreCreate {
	bgc.mutation.SetDisplayName(s)
	return bgc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (bgc *BookGenreCreate) SetNillableDisplayName(s *string) *BookGenreCreate {
	if s != nil {
		bgc.SetDisplayName(*s)
	}
	return bgc
}

// SetAbbreviation sets the "abbreviation" field.
func (bgc *BookGenreCreate) SetAbbreviation(s string) *BookGenreCreate {
	bgc.mutation.SetAbbreviation(s)
	return bgc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (bgc *BookGenreCreate) SetNillableAbbreviation(s *string) *BookGenreCreate {
	if s != nil {
		bgc.SetAbbreviation(*s)
	}
	return bgc
}

// SetDescription sets the "description" field.
func (bgc *BookGenreCreate) SetDescription(s string) *BookGenreCreate {
	bgc.mutation.SetDescription(s)
	return bgc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bgc *BookGenreCreate) SetNillableDescription(s *string) *BookGenreCreate {
	if s != nil {
		bgc.SetDescription(*s)
	}
	return bgc
}

// SetExternalLink sets the "external_link" field.
func (bgc *BookGenreCreate) SetExternalLink(s string) *BookGenreCreate {
	bgc.mutation.SetExternalLink(s)
	return bgc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (bgc *BookGenreCreate) SetNillableExternalLink(s *string) *BookGenreCreate {
	if s != nil {
		bgc.SetExternalLink(*s)
	}
	return bgc
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (bgc *BookGenreCreate) AddBookIDs(ids ...int) *BookGenreCreate {
	bgc.mutation.AddBookIDs(ids...)
	return bgc
}

// AddBooks adds the "books" edges to the Book entity.
func (bgc *BookGenreCreate) AddBooks(b ...*Book) *BookGenreCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bgc.AddBookIDs(ids...)
}

// Mutation returns the BookGenreMutation object of the builder.
func (bgc *BookGenreCreate) Mutation() *BookGenreMutation {
	return bgc.mutation
}

// Save creates the BookGenre in the database.
func (bgc *BookGenreCreate) Save(ctx context.Context) (*BookGenre, error) {
	if err := bgc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, bgc.sqlSave, bgc.mutation, bgc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bgc *BookGenreCreate) SaveX(ctx context.Context) *BookGenre {
	v, err := bgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bgc *BookGenreCreate) Exec(ctx context.Context) error {
	_, err := bgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bgc *BookGenreCreate) ExecX(ctx context.Context) {
	if err := bgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bgc *BookGenreCreate) defaults() error {
	if _, ok := bgc.mutation.CreatedAt(); !ok {
		if bookgenre.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized bookgenre.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := bookgenre.DefaultCreatedAt()
		bgc.mutation.SetCreatedAt(v)
	}
	if _, ok := bgc.mutation.UpdatedAt(); !ok {
		if bookgenre.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized bookgenre.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := bookgenre.DefaultUpdatedAt()
		bgc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (bgc *BookGenreCreate) check() error {
	if _, ok := bgc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BookGenre.created_at"`)}
	}
	if _, ok := bgc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BookGenre.updated_at"`)}
	}
	return nil
}

func (bgc *BookGenreCreate) sqlSave(ctx context.Context) (*BookGenre, error) {
	if err := bgc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bgc.mutation.id = &_node.ID
	bgc.mutation.done = true
	return _node, nil
}

func (bgc *BookGenreCreate) createSpec() (*BookGenre, *sqlgraph.CreateSpec) {
	var (
		_node = &BookGenre{config: bgc.config}
		_spec = sqlgraph.NewCreateSpec(bookgenre.Table, sqlgraph.NewFieldSpec(bookgenre.FieldID, field.TypeInt))
	)
	if value, ok := bgc.mutation.CreatedAt(); ok {
		_spec.SetField(bookgenre.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bgc.mutation.CreatedBy(); ok {
		_spec.SetField(bookgenre.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := bgc.mutation.UpdatedAt(); ok {
		_spec.SetField(bookgenre.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bgc.mutation.UpdatedBy(); ok {
		_spec.SetField(bookgenre.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := bgc.mutation.DisplayName(); ok {
		_spec.SetField(bookgenre.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := bgc.mutation.Abbreviation(); ok {
		_spec.SetField(bookgenre.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := bgc.mutation.Description(); ok {
		_spec.SetField(bookgenre.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := bgc.mutation.ExternalLink(); ok {
		_spec.SetField(bookgenre.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if nodes := bgc.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookgenre.BooksTable,
			Columns: bookgenre.BooksPrimaryKey,
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

// BookGenreCreateBulk is the builder for creating many BookGenre entities in bulk.
type BookGenreCreateBulk struct {
	config
	builders []*BookGenreCreate
}

// Save creates the BookGenre entities in the database.
func (bgcb *BookGenreCreateBulk) Save(ctx context.Context) ([]*BookGenre, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bgcb.builders))
	nodes := make([]*BookGenre, len(bgcb.builders))
	mutators := make([]Mutator, len(bgcb.builders))
	for i := range bgcb.builders {
		func(i int, root context.Context) {
			builder := bgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookGenreMutation)
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
					_, err = mutators[i+1].Mutate(root, bgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bgcb *BookGenreCreateBulk) SaveX(ctx context.Context) []*BookGenre {
	v, err := bgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bgcb *BookGenreCreateBulk) Exec(ctx context.Context) error {
	_, err := bgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bgcb *BookGenreCreateBulk) ExecX(ctx context.Context) {
	if err := bgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
