// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/personrole"
)

// PersonRoleCreate is the builder for creating a PersonRole entity.
type PersonRoleCreate struct {
	config
	mutation *PersonRoleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (prc *PersonRoleCreate) SetCreatedAt(t time.Time) *PersonRoleCreate {
	prc.mutation.SetCreatedAt(t)
	return prc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (prc *PersonRoleCreate) SetNillableCreatedAt(t *time.Time) *PersonRoleCreate {
	if t != nil {
		prc.SetCreatedAt(*t)
	}
	return prc
}

// SetCreatedBy sets the "created_by" field.
func (prc *PersonRoleCreate) SetCreatedBy(s string) *PersonRoleCreate {
	prc.mutation.SetCreatedBy(s)
	return prc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (prc *PersonRoleCreate) SetNillableCreatedBy(s *string) *PersonRoleCreate {
	if s != nil {
		prc.SetCreatedBy(*s)
	}
	return prc
}

// SetUpdatedAt sets the "updated_at" field.
func (prc *PersonRoleCreate) SetUpdatedAt(t time.Time) *PersonRoleCreate {
	prc.mutation.SetUpdatedAt(t)
	return prc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (prc *PersonRoleCreate) SetNillableUpdatedAt(t *time.Time) *PersonRoleCreate {
	if t != nil {
		prc.SetUpdatedAt(*t)
	}
	return prc
}

// SetUpdatedBy sets the "updated_by" field.
func (prc *PersonRoleCreate) SetUpdatedBy(s string) *PersonRoleCreate {
	prc.mutation.SetUpdatedBy(s)
	return prc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (prc *PersonRoleCreate) SetNillableUpdatedBy(s *string) *PersonRoleCreate {
	if s != nil {
		prc.SetUpdatedBy(*s)
	}
	return prc
}

// SetDisplayName sets the "display_name" field.
func (prc *PersonRoleCreate) SetDisplayName(s string) *PersonRoleCreate {
	prc.mutation.SetDisplayName(s)
	return prc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (prc *PersonRoleCreate) SetNillableDisplayName(s *string) *PersonRoleCreate {
	if s != nil {
		prc.SetDisplayName(*s)
	}
	return prc
}

// SetAbbreviation sets the "abbreviation" field.
func (prc *PersonRoleCreate) SetAbbreviation(s string) *PersonRoleCreate {
	prc.mutation.SetAbbreviation(s)
	return prc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (prc *PersonRoleCreate) SetNillableAbbreviation(s *string) *PersonRoleCreate {
	if s != nil {
		prc.SetAbbreviation(*s)
	}
	return prc
}

// SetDescription sets the "description" field.
func (prc *PersonRoleCreate) SetDescription(s string) *PersonRoleCreate {
	prc.mutation.SetDescription(s)
	return prc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (prc *PersonRoleCreate) SetNillableDescription(s *string) *PersonRoleCreate {
	if s != nil {
		prc.SetDescription(*s)
	}
	return prc
}

// SetExternalLink sets the "external_link" field.
func (prc *PersonRoleCreate) SetExternalLink(s string) *PersonRoleCreate {
	prc.mutation.SetExternalLink(s)
	return prc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (prc *PersonRoleCreate) SetNillableExternalLink(s *string) *PersonRoleCreate {
	if s != nil {
		prc.SetExternalLink(*s)
	}
	return prc
}

// AddPersonIDs adds the "person" edge to the Person entity by IDs.
func (prc *PersonRoleCreate) AddPersonIDs(ids ...int) *PersonRoleCreate {
	prc.mutation.AddPersonIDs(ids...)
	return prc
}

// AddPerson adds the "person" edges to the Person entity.
func (prc *PersonRoleCreate) AddPerson(p ...*Person) *PersonRoleCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return prc.AddPersonIDs(ids...)
}

// Mutation returns the PersonRoleMutation object of the builder.
func (prc *PersonRoleCreate) Mutation() *PersonRoleMutation {
	return prc.mutation
}

// Save creates the PersonRole in the database.
func (prc *PersonRoleCreate) Save(ctx context.Context) (*PersonRole, error) {
	if err := prc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, prc.sqlSave, prc.mutation, prc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (prc *PersonRoleCreate) SaveX(ctx context.Context) *PersonRole {
	v, err := prc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prc *PersonRoleCreate) Exec(ctx context.Context) error {
	_, err := prc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prc *PersonRoleCreate) ExecX(ctx context.Context) {
	if err := prc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (prc *PersonRoleCreate) defaults() error {
	if _, ok := prc.mutation.CreatedAt(); !ok {
		if personrole.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized personrole.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := personrole.DefaultCreatedAt()
		prc.mutation.SetCreatedAt(v)
	}
	if _, ok := prc.mutation.UpdatedAt(); !ok {
		if personrole.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized personrole.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := personrole.DefaultUpdatedAt()
		prc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (prc *PersonRoleCreate) check() error {
	if _, ok := prc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PersonRole.created_at"`)}
	}
	if _, ok := prc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PersonRole.updated_at"`)}
	}
	return nil
}

func (prc *PersonRoleCreate) sqlSave(ctx context.Context) (*PersonRole, error) {
	if err := prc.check(); err != nil {
		return nil, err
	}
	_node, _spec := prc.createSpec()
	if err := sqlgraph.CreateNode(ctx, prc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	prc.mutation.id = &_node.ID
	prc.mutation.done = true
	return _node, nil
}

func (prc *PersonRoleCreate) createSpec() (*PersonRole, *sqlgraph.CreateSpec) {
	var (
		_node = &PersonRole{config: prc.config}
		_spec = sqlgraph.NewCreateSpec(personrole.Table, sqlgraph.NewFieldSpec(personrole.FieldID, field.TypeInt))
	)
	if value, ok := prc.mutation.CreatedAt(); ok {
		_spec.SetField(personrole.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := prc.mutation.CreatedBy(); ok {
		_spec.SetField(personrole.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := prc.mutation.UpdatedAt(); ok {
		_spec.SetField(personrole.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := prc.mutation.UpdatedBy(); ok {
		_spec.SetField(personrole.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := prc.mutation.DisplayName(); ok {
		_spec.SetField(personrole.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := prc.mutation.Abbreviation(); ok {
		_spec.SetField(personrole.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := prc.mutation.Description(); ok {
		_spec.SetField(personrole.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := prc.mutation.ExternalLink(); ok {
		_spec.SetField(personrole.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if nodes := prc.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   personrole.PersonTable,
			Columns: personrole.PersonPrimaryKey,
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

// PersonRoleCreateBulk is the builder for creating many PersonRole entities in bulk.
type PersonRoleCreateBulk struct {
	config
	builders []*PersonRoleCreate
}

// Save creates the PersonRole entities in the database.
func (prcb *PersonRoleCreateBulk) Save(ctx context.Context) ([]*PersonRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(prcb.builders))
	nodes := make([]*PersonRole, len(prcb.builders))
	mutators := make([]Mutator, len(prcb.builders))
	for i := range prcb.builders {
		func(i int, root context.Context) {
			builder := prcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, prcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, prcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, prcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (prcb *PersonRoleCreateBulk) SaveX(ctx context.Context) []*PersonRole {
	v, err := prcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prcb *PersonRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := prcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prcb *PersonRoleCreateBulk) ExecX(ctx context.Context) {
	if err := prcb.Exec(ctx); err != nil {
		panic(err)
	}
}
