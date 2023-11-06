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
	"github.com/dkrasnovdev/siberiana-api/ent/ethnos"
)

// EthnosCreate is the builder for creating a Ethnos entity.
type EthnosCreate struct {
	config
	mutation *EthnosMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ec *EthnosCreate) SetCreatedAt(t time.Time) *EthnosCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EthnosCreate) SetNillableCreatedAt(t *time.Time) *EthnosCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetCreatedBy sets the "created_by" field.
func (ec *EthnosCreate) SetCreatedBy(s string) *EthnosCreate {
	ec.mutation.SetCreatedBy(s)
	return ec
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ec *EthnosCreate) SetNillableCreatedBy(s *string) *EthnosCreate {
	if s != nil {
		ec.SetCreatedBy(*s)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EthnosCreate) SetUpdatedAt(t time.Time) *EthnosCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EthnosCreate) SetNillableUpdatedAt(t *time.Time) *EthnosCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetUpdatedBy sets the "updated_by" field.
func (ec *EthnosCreate) SetUpdatedBy(s string) *EthnosCreate {
	ec.mutation.SetUpdatedBy(s)
	return ec
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ec *EthnosCreate) SetNillableUpdatedBy(s *string) *EthnosCreate {
	if s != nil {
		ec.SetUpdatedBy(*s)
	}
	return ec
}

// SetDisplayName sets the "display_name" field.
func (ec *EthnosCreate) SetDisplayName(s string) *EthnosCreate {
	ec.mutation.SetDisplayName(s)
	return ec
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ec *EthnosCreate) SetNillableDisplayName(s *string) *EthnosCreate {
	if s != nil {
		ec.SetDisplayName(*s)
	}
	return ec
}

// SetAbbreviation sets the "abbreviation" field.
func (ec *EthnosCreate) SetAbbreviation(s string) *EthnosCreate {
	ec.mutation.SetAbbreviation(s)
	return ec
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ec *EthnosCreate) SetNillableAbbreviation(s *string) *EthnosCreate {
	if s != nil {
		ec.SetAbbreviation(*s)
	}
	return ec
}

// SetDescription sets the "description" field.
func (ec *EthnosCreate) SetDescription(s string) *EthnosCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ec *EthnosCreate) SetNillableDescription(s *string) *EthnosCreate {
	if s != nil {
		ec.SetDescription(*s)
	}
	return ec
}

// SetExternalLink sets the "external_link" field.
func (ec *EthnosCreate) SetExternalLink(s string) *EthnosCreate {
	ec.mutation.SetExternalLink(s)
	return ec
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (ec *EthnosCreate) SetNillableExternalLink(s *string) *EthnosCreate {
	if s != nil {
		ec.SetExternalLink(*s)
	}
	return ec
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (ec *EthnosCreate) AddArtifactIDs(ids ...int) *EthnosCreate {
	ec.mutation.AddArtifactIDs(ids...)
	return ec
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (ec *EthnosCreate) AddArtifacts(a ...*Artifact) *EthnosCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ec.AddArtifactIDs(ids...)
}

// Mutation returns the EthnosMutation object of the builder.
func (ec *EthnosCreate) Mutation() *EthnosMutation {
	return ec.mutation
}

// Save creates the Ethnos in the database.
func (ec *EthnosCreate) Save(ctx context.Context) (*Ethnos, error) {
	if err := ec.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EthnosCreate) SaveX(ctx context.Context) *Ethnos {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EthnosCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EthnosCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EthnosCreate) defaults() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		if ethnos.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized ethnos.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := ethnos.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		if ethnos.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized ethnos.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := ethnos.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ec *EthnosCreate) check() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Ethnos.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Ethnos.updated_at"`)}
	}
	return nil
}

func (ec *EthnosCreate) sqlSave(ctx context.Context) (*Ethnos, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EthnosCreate) createSpec() (*Ethnos, *sqlgraph.CreateSpec) {
	var (
		_node = &Ethnos{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(ethnos.Table, sqlgraph.NewFieldSpec(ethnos.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(ethnos.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.CreatedBy(); ok {
		_spec.SetField(ethnos.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(ethnos.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.UpdatedBy(); ok {
		_spec.SetField(ethnos.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ec.mutation.DisplayName(); ok {
		_spec.SetField(ethnos.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := ec.mutation.Abbreviation(); ok {
		_spec.SetField(ethnos.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.SetField(ethnos.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ec.mutation.ExternalLink(); ok {
		_spec.SetField(ethnos.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if nodes := ec.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ethnos.ArtifactsTable,
			Columns: []string{ethnos.ArtifactsColumn},
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

// EthnosCreateBulk is the builder for creating many Ethnos entities in bulk.
type EthnosCreateBulk struct {
	config
	err      error
	builders []*EthnosCreate
}

// Save creates the Ethnos entities in the database.
func (ecb *EthnosCreateBulk) Save(ctx context.Context) ([]*Ethnos, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Ethnos, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EthnosMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EthnosCreateBulk) SaveX(ctx context.Context) []*Ethnos {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EthnosCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EthnosCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}