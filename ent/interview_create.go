// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/interview"
)

// InterviewCreate is the builder for creating a Interview entity.
type InterviewCreate struct {
	config
	mutation *InterviewMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *InterviewCreate) SetCreatedAt(t time.Time) *InterviewCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableCreatedAt(t *time.Time) *InterviewCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetCreatedBy sets the "created_by" field.
func (ic *InterviewCreate) SetCreatedBy(s string) *InterviewCreate {
	ic.mutation.SetCreatedBy(s)
	return ic
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableCreatedBy(s *string) *InterviewCreate {
	if s != nil {
		ic.SetCreatedBy(*s)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *InterviewCreate) SetUpdatedAt(t time.Time) *InterviewCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableUpdatedAt(t *time.Time) *InterviewCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetUpdatedBy sets the "updated_by" field.
func (ic *InterviewCreate) SetUpdatedBy(s string) *InterviewCreate {
	ic.mutation.SetUpdatedBy(s)
	return ic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableUpdatedBy(s *string) *InterviewCreate {
	if s != nil {
		ic.SetUpdatedBy(*s)
	}
	return ic
}

// SetDisplayName sets the "display_name" field.
func (ic *InterviewCreate) SetDisplayName(s string) *InterviewCreate {
	ic.mutation.SetDisplayName(s)
	return ic
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableDisplayName(s *string) *InterviewCreate {
	if s != nil {
		ic.SetDisplayName(*s)
	}
	return ic
}

// SetAbbreviation sets the "abbreviation" field.
func (ic *InterviewCreate) SetAbbreviation(s string) *InterviewCreate {
	ic.mutation.SetAbbreviation(s)
	return ic
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableAbbreviation(s *string) *InterviewCreate {
	if s != nil {
		ic.SetAbbreviation(*s)
	}
	return ic
}

// SetDescription sets the "description" field.
func (ic *InterviewCreate) SetDescription(s string) *InterviewCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableDescription(s *string) *InterviewCreate {
	if s != nil {
		ic.SetDescription(*s)
	}
	return ic
}

// SetExternalLink sets the "external_link" field.
func (ic *InterviewCreate) SetExternalLink(s string) *InterviewCreate {
	ic.mutation.SetExternalLink(s)
	return ic
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableExternalLink(s *string) *InterviewCreate {
	if s != nil {
		ic.SetExternalLink(*s)
	}
	return ic
}

// SetDate sets the "date" field.
func (ic *InterviewCreate) SetDate(t time.Time) *InterviewCreate {
	ic.mutation.SetDate(t)
	return ic
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ic *InterviewCreate) SetNillableDate(t *time.Time) *InterviewCreate {
	if t != nil {
		ic.SetDate(*t)
	}
	return ic
}

// Mutation returns the InterviewMutation object of the builder.
func (ic *InterviewCreate) Mutation() *InterviewMutation {
	return ic.mutation
}

// Save creates the Interview in the database.
func (ic *InterviewCreate) Save(ctx context.Context) (*Interview, error) {
	if err := ic.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InterviewCreate) SaveX(ctx context.Context) *Interview {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InterviewCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InterviewCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InterviewCreate) defaults() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		if interview.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized interview.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := interview.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		if interview.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized interview.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := interview.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ic *InterviewCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Interview.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Interview.updated_at"`)}
	}
	return nil
}

func (ic *InterviewCreate) sqlSave(ctx context.Context) (*Interview, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InterviewCreate) createSpec() (*Interview, *sqlgraph.CreateSpec) {
	var (
		_node = &Interview{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(interview.Table, sqlgraph.NewFieldSpec(interview.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(interview.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.CreatedBy(); ok {
		_spec.SetField(interview.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(interview.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.UpdatedBy(); ok {
		_spec.SetField(interview.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ic.mutation.DisplayName(); ok {
		_spec.SetField(interview.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := ic.mutation.Abbreviation(); ok {
		_spec.SetField(interview.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(interview.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ic.mutation.ExternalLink(); ok {
		_spec.SetField(interview.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := ic.mutation.Date(); ok {
		_spec.SetField(interview.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	return _node, _spec
}

// InterviewCreateBulk is the builder for creating many Interview entities in bulk.
type InterviewCreateBulk struct {
	config
	err      error
	builders []*InterviewCreate
}

// Save creates the Interview entities in the database.
func (icb *InterviewCreateBulk) Save(ctx context.Context) ([]*Interview, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Interview, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InterviewMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InterviewCreateBulk) SaveX(ctx context.Context) []*Interview {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InterviewCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InterviewCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
