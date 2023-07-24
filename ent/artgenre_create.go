// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/artgenre"
)

// ArtGenreCreate is the builder for creating a ArtGenre entity.
type ArtGenreCreate struct {
	config
	mutation *ArtGenreMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (agc *ArtGenreCreate) SetCreatedAt(t time.Time) *ArtGenreCreate {
	agc.mutation.SetCreatedAt(t)
	return agc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (agc *ArtGenreCreate) SetNillableCreatedAt(t *time.Time) *ArtGenreCreate {
	if t != nil {
		agc.SetCreatedAt(*t)
	}
	return agc
}

// SetCreatedBy sets the "created_by" field.
func (agc *ArtGenreCreate) SetCreatedBy(s string) *ArtGenreCreate {
	agc.mutation.SetCreatedBy(s)
	return agc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (agc *ArtGenreCreate) SetNillableCreatedBy(s *string) *ArtGenreCreate {
	if s != nil {
		agc.SetCreatedBy(*s)
	}
	return agc
}

// SetUpdatedAt sets the "updated_at" field.
func (agc *ArtGenreCreate) SetUpdatedAt(t time.Time) *ArtGenreCreate {
	agc.mutation.SetUpdatedAt(t)
	return agc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (agc *ArtGenreCreate) SetNillableUpdatedAt(t *time.Time) *ArtGenreCreate {
	if t != nil {
		agc.SetUpdatedAt(*t)
	}
	return agc
}

// SetUpdatedBy sets the "updated_by" field.
func (agc *ArtGenreCreate) SetUpdatedBy(s string) *ArtGenreCreate {
	agc.mutation.SetUpdatedBy(s)
	return agc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (agc *ArtGenreCreate) SetNillableUpdatedBy(s *string) *ArtGenreCreate {
	if s != nil {
		agc.SetUpdatedBy(*s)
	}
	return agc
}

// SetDisplayName sets the "display_name" field.
func (agc *ArtGenreCreate) SetDisplayName(s string) *ArtGenreCreate {
	agc.mutation.SetDisplayName(s)
	return agc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (agc *ArtGenreCreate) SetNillableDisplayName(s *string) *ArtGenreCreate {
	if s != nil {
		agc.SetDisplayName(*s)
	}
	return agc
}

// SetAbbreviation sets the "abbreviation" field.
func (agc *ArtGenreCreate) SetAbbreviation(s string) *ArtGenreCreate {
	agc.mutation.SetAbbreviation(s)
	return agc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (agc *ArtGenreCreate) SetNillableAbbreviation(s *string) *ArtGenreCreate {
	if s != nil {
		agc.SetAbbreviation(*s)
	}
	return agc
}

// SetDescription sets the "description" field.
func (agc *ArtGenreCreate) SetDescription(s string) *ArtGenreCreate {
	agc.mutation.SetDescription(s)
	return agc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (agc *ArtGenreCreate) SetNillableDescription(s *string) *ArtGenreCreate {
	if s != nil {
		agc.SetDescription(*s)
	}
	return agc
}

// SetExternalLinks sets the "external_links" field.
func (agc *ArtGenreCreate) SetExternalLinks(s []string) *ArtGenreCreate {
	agc.mutation.SetExternalLinks(s)
	return agc
}

// Mutation returns the ArtGenreMutation object of the builder.
func (agc *ArtGenreCreate) Mutation() *ArtGenreMutation {
	return agc.mutation
}

// Save creates the ArtGenre in the database.
func (agc *ArtGenreCreate) Save(ctx context.Context) (*ArtGenre, error) {
	if err := agc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, agc.sqlSave, agc.mutation, agc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (agc *ArtGenreCreate) SaveX(ctx context.Context) *ArtGenre {
	v, err := agc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agc *ArtGenreCreate) Exec(ctx context.Context) error {
	_, err := agc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agc *ArtGenreCreate) ExecX(ctx context.Context) {
	if err := agc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agc *ArtGenreCreate) defaults() error {
	if _, ok := agc.mutation.CreatedAt(); !ok {
		if artgenre.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized artgenre.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := artgenre.DefaultCreatedAt()
		agc.mutation.SetCreatedAt(v)
	}
	if _, ok := agc.mutation.UpdatedAt(); !ok {
		if artgenre.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized artgenre.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := artgenre.DefaultUpdatedAt()
		agc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (agc *ArtGenreCreate) check() error {
	if _, ok := agc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ArtGenre.created_at"`)}
	}
	if _, ok := agc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ArtGenre.updated_at"`)}
	}
	return nil
}

func (agc *ArtGenreCreate) sqlSave(ctx context.Context) (*ArtGenre, error) {
	if err := agc.check(); err != nil {
		return nil, err
	}
	_node, _spec := agc.createSpec()
	if err := sqlgraph.CreateNode(ctx, agc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	agc.mutation.id = &_node.ID
	agc.mutation.done = true
	return _node, nil
}

func (agc *ArtGenreCreate) createSpec() (*ArtGenre, *sqlgraph.CreateSpec) {
	var (
		_node = &ArtGenre{config: agc.config}
		_spec = sqlgraph.NewCreateSpec(artgenre.Table, sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt))
	)
	if value, ok := agc.mutation.CreatedAt(); ok {
		_spec.SetField(artgenre.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := agc.mutation.CreatedBy(); ok {
		_spec.SetField(artgenre.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := agc.mutation.UpdatedAt(); ok {
		_spec.SetField(artgenre.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := agc.mutation.UpdatedBy(); ok {
		_spec.SetField(artgenre.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := agc.mutation.DisplayName(); ok {
		_spec.SetField(artgenre.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := agc.mutation.Abbreviation(); ok {
		_spec.SetField(artgenre.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := agc.mutation.Description(); ok {
		_spec.SetField(artgenre.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := agc.mutation.ExternalLinks(); ok {
		_spec.SetField(artgenre.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
	return _node, _spec
}

// ArtGenreCreateBulk is the builder for creating many ArtGenre entities in bulk.
type ArtGenreCreateBulk struct {
	config
	builders []*ArtGenreCreate
}

// Save creates the ArtGenre entities in the database.
func (agcb *ArtGenreCreateBulk) Save(ctx context.Context) ([]*ArtGenre, error) {
	specs := make([]*sqlgraph.CreateSpec, len(agcb.builders))
	nodes := make([]*ArtGenre, len(agcb.builders))
	mutators := make([]Mutator, len(agcb.builders))
	for i := range agcb.builders {
		func(i int, root context.Context) {
			builder := agcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtGenreMutation)
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
					_, err = mutators[i+1].Mutate(root, agcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, agcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, agcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (agcb *ArtGenreCreateBulk) SaveX(ctx context.Context) []*ArtGenre {
	v, err := agcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agcb *ArtGenreCreateBulk) Exec(ctx context.Context) error {
	_, err := agcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agcb *ArtGenreCreateBulk) ExecX(ctx context.Context) {
	if err := agcb.Exec(ctx); err != nil {
		panic(err)
	}
}
