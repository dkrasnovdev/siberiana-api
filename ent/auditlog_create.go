// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/auditlog"
)

// AuditLogCreate is the builder for creating a AuditLog entity.
type AuditLogCreate struct {
	config
	mutation *AuditLogMutation
	hooks    []Hook
}

// SetTable sets the "table" field.
func (alc *AuditLogCreate) SetTable(s string) *AuditLogCreate {
	alc.mutation.SetTable(s)
	return alc
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableTable(s *string) *AuditLogCreate {
	if s != nil {
		alc.SetTable(*s)
	}
	return alc
}

// SetRefID sets the "ref_id" field.
func (alc *AuditLogCreate) SetRefID(i int) *AuditLogCreate {
	alc.mutation.SetRefID(i)
	return alc
}

// SetNillableRefID sets the "ref_id" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableRefID(i *int) *AuditLogCreate {
	if i != nil {
		alc.SetRefID(*i)
	}
	return alc
}

// SetOperation sets the "operation" field.
func (alc *AuditLogCreate) SetOperation(s string) *AuditLogCreate {
	alc.mutation.SetOperation(s)
	return alc
}

// SetNillableOperation sets the "operation" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableOperation(s *string) *AuditLogCreate {
	if s != nil {
		alc.SetOperation(*s)
	}
	return alc
}

// SetChanges sets the "changes" field.
func (alc *AuditLogCreate) SetChanges(s []string) *AuditLogCreate {
	alc.mutation.SetChanges(s)
	return alc
}

// SetAddedIds sets the "added_ids" field.
func (alc *AuditLogCreate) SetAddedIds(s []string) *AuditLogCreate {
	alc.mutation.SetAddedIds(s)
	return alc
}

// SetRemovedIds sets the "removed_ids" field.
func (alc *AuditLogCreate) SetRemovedIds(s []string) *AuditLogCreate {
	alc.mutation.SetRemovedIds(s)
	return alc
}

// SetClearedEdges sets the "cleared_edges" field.
func (alc *AuditLogCreate) SetClearedEdges(s []string) *AuditLogCreate {
	alc.mutation.SetClearedEdges(s)
	return alc
}

// SetBlame sets the "blame" field.
func (alc *AuditLogCreate) SetBlame(s string) *AuditLogCreate {
	alc.mutation.SetBlame(s)
	return alc
}

// SetNillableBlame sets the "blame" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableBlame(s *string) *AuditLogCreate {
	if s != nil {
		alc.SetBlame(*s)
	}
	return alc
}

// SetCreatedAt sets the "created_at" field.
func (alc *AuditLogCreate) SetCreatedAt(t time.Time) *AuditLogCreate {
	alc.mutation.SetCreatedAt(t)
	return alc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableCreatedAt(t *time.Time) *AuditLogCreate {
	if t != nil {
		alc.SetCreatedAt(*t)
	}
	return alc
}

// Mutation returns the AuditLogMutation object of the builder.
func (alc *AuditLogCreate) Mutation() *AuditLogMutation {
	return alc.mutation
}

// Save creates the AuditLog in the database.
func (alc *AuditLogCreate) Save(ctx context.Context) (*AuditLog, error) {
	alc.defaults()
	return withHooks(ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *AuditLogCreate) SaveX(ctx context.Context) *AuditLog {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *AuditLogCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *AuditLogCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alc *AuditLogCreate) defaults() {
	if _, ok := alc.mutation.CreatedAt(); !ok {
		v := auditlog.DefaultCreatedAt()
		alc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *AuditLogCreate) check() error {
	if _, ok := alc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AuditLog.created_at"`)}
	}
	return nil
}

func (alc *AuditLogCreate) sqlSave(ctx context.Context) (*AuditLog, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *AuditLogCreate) createSpec() (*AuditLog, *sqlgraph.CreateSpec) {
	var (
		_node = &AuditLog{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(auditlog.Table, sqlgraph.NewFieldSpec(auditlog.FieldID, field.TypeInt))
	)
	if value, ok := alc.mutation.Table(); ok {
		_spec.SetField(auditlog.FieldTable, field.TypeString, value)
		_node.Table = value
	}
	if value, ok := alc.mutation.RefID(); ok {
		_spec.SetField(auditlog.FieldRefID, field.TypeInt, value)
		_node.RefID = value
	}
	if value, ok := alc.mutation.Operation(); ok {
		_spec.SetField(auditlog.FieldOperation, field.TypeString, value)
		_node.Operation = value
	}
	if value, ok := alc.mutation.Changes(); ok {
		_spec.SetField(auditlog.FieldChanges, field.TypeJSON, value)
		_node.Changes = value
	}
	if value, ok := alc.mutation.AddedIds(); ok {
		_spec.SetField(auditlog.FieldAddedIds, field.TypeJSON, value)
		_node.AddedIds = value
	}
	if value, ok := alc.mutation.RemovedIds(); ok {
		_spec.SetField(auditlog.FieldRemovedIds, field.TypeJSON, value)
		_node.RemovedIds = value
	}
	if value, ok := alc.mutation.GetClearedEdges(); ok {
		_spec.SetField(auditlog.FieldClearedEdges, field.TypeJSON, value)
		_node.ClearedEdges = value
	}
	if value, ok := alc.mutation.Blame(); ok {
		_spec.SetField(auditlog.FieldBlame, field.TypeString, value)
		_node.Blame = value
	}
	if value, ok := alc.mutation.CreatedAt(); ok {
		_spec.SetField(auditlog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// AuditLogCreateBulk is the builder for creating many AuditLog entities in bulk.
type AuditLogCreateBulk struct {
	config
	builders []*AuditLogCreate
}

// Save creates the AuditLog entities in the database.
func (alcb *AuditLogCreateBulk) Save(ctx context.Context) ([]*AuditLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*AuditLog, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuditLogMutation)
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
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *AuditLogCreateBulk) SaveX(ctx context.Context) []*AuditLog {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *AuditLogCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *AuditLogCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}
