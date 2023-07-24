// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
	"github.com/dkrasnovdev/heritage-api/ent/protectedarea"
)

// ProtectedAreaUpdate is the builder for updating ProtectedArea entities.
type ProtectedAreaUpdate struct {
	config
	hooks    []Hook
	mutation *ProtectedAreaMutation
}

// Where appends a list predicates to the ProtectedAreaUpdate builder.
func (pau *ProtectedAreaUpdate) Where(ps ...predicate.ProtectedArea) *ProtectedAreaUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// SetCreatedBy sets the "created_by" field.
func (pau *ProtectedAreaUpdate) SetCreatedBy(s string) *ProtectedAreaUpdate {
	pau.mutation.SetCreatedBy(s)
	return pau
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableCreatedBy(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetCreatedBy(*s)
	}
	return pau
}

// ClearCreatedBy clears the value of the "created_by" field.
func (pau *ProtectedAreaUpdate) ClearCreatedBy() *ProtectedAreaUpdate {
	pau.mutation.ClearCreatedBy()
	return pau
}

// SetUpdatedAt sets the "updated_at" field.
func (pau *ProtectedAreaUpdate) SetUpdatedAt(t time.Time) *ProtectedAreaUpdate {
	pau.mutation.SetUpdatedAt(t)
	return pau
}

// SetUpdatedBy sets the "updated_by" field.
func (pau *ProtectedAreaUpdate) SetUpdatedBy(s string) *ProtectedAreaUpdate {
	pau.mutation.SetUpdatedBy(s)
	return pau
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableUpdatedBy(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetUpdatedBy(*s)
	}
	return pau
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pau *ProtectedAreaUpdate) ClearUpdatedBy() *ProtectedAreaUpdate {
	pau.mutation.ClearUpdatedBy()
	return pau
}

// SetDisplayName sets the "display_name" field.
func (pau *ProtectedAreaUpdate) SetDisplayName(s string) *ProtectedAreaUpdate {
	pau.mutation.SetDisplayName(s)
	return pau
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableDisplayName(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetDisplayName(*s)
	}
	return pau
}

// ClearDisplayName clears the value of the "display_name" field.
func (pau *ProtectedAreaUpdate) ClearDisplayName() *ProtectedAreaUpdate {
	pau.mutation.ClearDisplayName()
	return pau
}

// SetAbbreviation sets the "abbreviation" field.
func (pau *ProtectedAreaUpdate) SetAbbreviation(s string) *ProtectedAreaUpdate {
	pau.mutation.SetAbbreviation(s)
	return pau
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableAbbreviation(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetAbbreviation(*s)
	}
	return pau
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (pau *ProtectedAreaUpdate) ClearAbbreviation() *ProtectedAreaUpdate {
	pau.mutation.ClearAbbreviation()
	return pau
}

// SetDescription sets the "description" field.
func (pau *ProtectedAreaUpdate) SetDescription(s string) *ProtectedAreaUpdate {
	pau.mutation.SetDescription(s)
	return pau
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableDescription(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetDescription(*s)
	}
	return pau
}

// ClearDescription clears the value of the "description" field.
func (pau *ProtectedAreaUpdate) ClearDescription() *ProtectedAreaUpdate {
	pau.mutation.ClearDescription()
	return pau
}

// SetExternalLinks sets the "external_links" field.
func (pau *ProtectedAreaUpdate) SetExternalLinks(s []string) *ProtectedAreaUpdate {
	pau.mutation.SetExternalLinks(s)
	return pau
}

// AppendExternalLinks appends s to the "external_links" field.
func (pau *ProtectedAreaUpdate) AppendExternalLinks(s []string) *ProtectedAreaUpdate {
	pau.mutation.AppendExternalLinks(s)
	return pau
}

// ClearExternalLinks clears the value of the "external_links" field.
func (pau *ProtectedAreaUpdate) ClearExternalLinks() *ProtectedAreaUpdate {
	pau.mutation.ClearExternalLinks()
	return pau
}

// Mutation returns the ProtectedAreaMutation object of the builder.
func (pau *ProtectedAreaUpdate) Mutation() *ProtectedAreaMutation {
	return pau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *ProtectedAreaUpdate) Save(ctx context.Context) (int, error) {
	if err := pau.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *ProtectedAreaUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *ProtectedAreaUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *ProtectedAreaUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pau *ProtectedAreaUpdate) defaults() error {
	if _, ok := pau.mutation.UpdatedAt(); !ok {
		if protectedarea.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedarea.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := protectedarea.UpdateDefaultUpdatedAt()
		pau.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pau *ProtectedAreaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(protectedarea.Table, protectedarea.Columns, sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.CreatedBy(); ok {
		_spec.SetField(protectedarea.FieldCreatedBy, field.TypeString, value)
	}
	if pau.mutation.CreatedByCleared() {
		_spec.ClearField(protectedarea.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pau.mutation.UpdatedAt(); ok {
		_spec.SetField(protectedarea.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pau.mutation.UpdatedBy(); ok {
		_spec.SetField(protectedarea.FieldUpdatedBy, field.TypeString, value)
	}
	if pau.mutation.UpdatedByCleared() {
		_spec.ClearField(protectedarea.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pau.mutation.DisplayName(); ok {
		_spec.SetField(protectedarea.FieldDisplayName, field.TypeString, value)
	}
	if pau.mutation.DisplayNameCleared() {
		_spec.ClearField(protectedarea.FieldDisplayName, field.TypeString)
	}
	if value, ok := pau.mutation.Abbreviation(); ok {
		_spec.SetField(protectedarea.FieldAbbreviation, field.TypeString, value)
	}
	if pau.mutation.AbbreviationCleared() {
		_spec.ClearField(protectedarea.FieldAbbreviation, field.TypeString)
	}
	if value, ok := pau.mutation.Description(); ok {
		_spec.SetField(protectedarea.FieldDescription, field.TypeString, value)
	}
	if pau.mutation.DescriptionCleared() {
		_spec.ClearField(protectedarea.FieldDescription, field.TypeString)
	}
	if value, ok := pau.mutation.ExternalLinks(); ok {
		_spec.SetField(protectedarea.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := pau.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, protectedarea.FieldExternalLinks, value)
		})
	}
	if pau.mutation.ExternalLinksCleared() {
		_spec.ClearField(protectedarea.FieldExternalLinks, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{protectedarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// ProtectedAreaUpdateOne is the builder for updating a single ProtectedArea entity.
type ProtectedAreaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProtectedAreaMutation
}

// SetCreatedBy sets the "created_by" field.
func (pauo *ProtectedAreaUpdateOne) SetCreatedBy(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetCreatedBy(s)
	return pauo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableCreatedBy(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetCreatedBy(*s)
	}
	return pauo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (pauo *ProtectedAreaUpdateOne) ClearCreatedBy() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearCreatedBy()
	return pauo
}

// SetUpdatedAt sets the "updated_at" field.
func (pauo *ProtectedAreaUpdateOne) SetUpdatedAt(t time.Time) *ProtectedAreaUpdateOne {
	pauo.mutation.SetUpdatedAt(t)
	return pauo
}

// SetUpdatedBy sets the "updated_by" field.
func (pauo *ProtectedAreaUpdateOne) SetUpdatedBy(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetUpdatedBy(s)
	return pauo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableUpdatedBy(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetUpdatedBy(*s)
	}
	return pauo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pauo *ProtectedAreaUpdateOne) ClearUpdatedBy() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearUpdatedBy()
	return pauo
}

// SetDisplayName sets the "display_name" field.
func (pauo *ProtectedAreaUpdateOne) SetDisplayName(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetDisplayName(s)
	return pauo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableDisplayName(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetDisplayName(*s)
	}
	return pauo
}

// ClearDisplayName clears the value of the "display_name" field.
func (pauo *ProtectedAreaUpdateOne) ClearDisplayName() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearDisplayName()
	return pauo
}

// SetAbbreviation sets the "abbreviation" field.
func (pauo *ProtectedAreaUpdateOne) SetAbbreviation(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetAbbreviation(s)
	return pauo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableAbbreviation(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetAbbreviation(*s)
	}
	return pauo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (pauo *ProtectedAreaUpdateOne) ClearAbbreviation() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearAbbreviation()
	return pauo
}

// SetDescription sets the "description" field.
func (pauo *ProtectedAreaUpdateOne) SetDescription(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetDescription(s)
	return pauo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableDescription(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetDescription(*s)
	}
	return pauo
}

// ClearDescription clears the value of the "description" field.
func (pauo *ProtectedAreaUpdateOne) ClearDescription() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearDescription()
	return pauo
}

// SetExternalLinks sets the "external_links" field.
func (pauo *ProtectedAreaUpdateOne) SetExternalLinks(s []string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetExternalLinks(s)
	return pauo
}

// AppendExternalLinks appends s to the "external_links" field.
func (pauo *ProtectedAreaUpdateOne) AppendExternalLinks(s []string) *ProtectedAreaUpdateOne {
	pauo.mutation.AppendExternalLinks(s)
	return pauo
}

// ClearExternalLinks clears the value of the "external_links" field.
func (pauo *ProtectedAreaUpdateOne) ClearExternalLinks() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearExternalLinks()
	return pauo
}

// Mutation returns the ProtectedAreaMutation object of the builder.
func (pauo *ProtectedAreaUpdateOne) Mutation() *ProtectedAreaMutation {
	return pauo.mutation
}

// Where appends a list predicates to the ProtectedAreaUpdate builder.
func (pauo *ProtectedAreaUpdateOne) Where(ps ...predicate.ProtectedArea) *ProtectedAreaUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *ProtectedAreaUpdateOne) Select(field string, fields ...string) *ProtectedAreaUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated ProtectedArea entity.
func (pauo *ProtectedAreaUpdateOne) Save(ctx context.Context) (*ProtectedArea, error) {
	if err := pauo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *ProtectedAreaUpdateOne) SaveX(ctx context.Context) *ProtectedArea {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *ProtectedAreaUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *ProtectedAreaUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pauo *ProtectedAreaUpdateOne) defaults() error {
	if _, ok := pauo.mutation.UpdatedAt(); !ok {
		if protectedarea.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedarea.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := protectedarea.UpdateDefaultUpdatedAt()
		pauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pauo *ProtectedAreaUpdateOne) sqlSave(ctx context.Context) (_node *ProtectedArea, err error) {
	_spec := sqlgraph.NewUpdateSpec(protectedarea.Table, protectedarea.Columns, sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProtectedArea.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, protectedarea.FieldID)
		for _, f := range fields {
			if !protectedarea.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != protectedarea.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.CreatedBy(); ok {
		_spec.SetField(protectedarea.FieldCreatedBy, field.TypeString, value)
	}
	if pauo.mutation.CreatedByCleared() {
		_spec.ClearField(protectedarea.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pauo.mutation.UpdatedAt(); ok {
		_spec.SetField(protectedarea.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pauo.mutation.UpdatedBy(); ok {
		_spec.SetField(protectedarea.FieldUpdatedBy, field.TypeString, value)
	}
	if pauo.mutation.UpdatedByCleared() {
		_spec.ClearField(protectedarea.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pauo.mutation.DisplayName(); ok {
		_spec.SetField(protectedarea.FieldDisplayName, field.TypeString, value)
	}
	if pauo.mutation.DisplayNameCleared() {
		_spec.ClearField(protectedarea.FieldDisplayName, field.TypeString)
	}
	if value, ok := pauo.mutation.Abbreviation(); ok {
		_spec.SetField(protectedarea.FieldAbbreviation, field.TypeString, value)
	}
	if pauo.mutation.AbbreviationCleared() {
		_spec.ClearField(protectedarea.FieldAbbreviation, field.TypeString)
	}
	if value, ok := pauo.mutation.Description(); ok {
		_spec.SetField(protectedarea.FieldDescription, field.TypeString, value)
	}
	if pauo.mutation.DescriptionCleared() {
		_spec.ClearField(protectedarea.FieldDescription, field.TypeString)
	}
	if value, ok := pauo.mutation.ExternalLinks(); ok {
		_spec.SetField(protectedarea.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := pauo.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, protectedarea.FieldExternalLinks, value)
		})
	}
	if pauo.mutation.ExternalLinksCleared() {
		_spec.ClearField(protectedarea.FieldExternalLinks, field.TypeJSON)
	}
	_node = &ProtectedArea{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{protectedarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}
