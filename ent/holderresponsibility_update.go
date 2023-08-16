// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/holder"
	"github.com/dkrasnovdev/siberiana-api/ent/holderresponsibility"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// HolderResponsibilityUpdate is the builder for updating HolderResponsibility entities.
type HolderResponsibilityUpdate struct {
	config
	hooks    []Hook
	mutation *HolderResponsibilityMutation
}

// Where appends a list predicates to the HolderResponsibilityUpdate builder.
func (hru *HolderResponsibilityUpdate) Where(ps ...predicate.HolderResponsibility) *HolderResponsibilityUpdate {
	hru.mutation.Where(ps...)
	return hru
}

// SetCreatedBy sets the "created_by" field.
func (hru *HolderResponsibilityUpdate) SetCreatedBy(s string) *HolderResponsibilityUpdate {
	hru.mutation.SetCreatedBy(s)
	return hru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (hru *HolderResponsibilityUpdate) SetNillableCreatedBy(s *string) *HolderResponsibilityUpdate {
	if s != nil {
		hru.SetCreatedBy(*s)
	}
	return hru
}

// ClearCreatedBy clears the value of the "created_by" field.
func (hru *HolderResponsibilityUpdate) ClearCreatedBy() *HolderResponsibilityUpdate {
	hru.mutation.ClearCreatedBy()
	return hru
}

// SetUpdatedAt sets the "updated_at" field.
func (hru *HolderResponsibilityUpdate) SetUpdatedAt(t time.Time) *HolderResponsibilityUpdate {
	hru.mutation.SetUpdatedAt(t)
	return hru
}

// SetUpdatedBy sets the "updated_by" field.
func (hru *HolderResponsibilityUpdate) SetUpdatedBy(s string) *HolderResponsibilityUpdate {
	hru.mutation.SetUpdatedBy(s)
	return hru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (hru *HolderResponsibilityUpdate) SetNillableUpdatedBy(s *string) *HolderResponsibilityUpdate {
	if s != nil {
		hru.SetUpdatedBy(*s)
	}
	return hru
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (hru *HolderResponsibilityUpdate) ClearUpdatedBy() *HolderResponsibilityUpdate {
	hru.mutation.ClearUpdatedBy()
	return hru
}

// SetDisplayName sets the "display_name" field.
func (hru *HolderResponsibilityUpdate) SetDisplayName(s string) *HolderResponsibilityUpdate {
	hru.mutation.SetDisplayName(s)
	return hru
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (hru *HolderResponsibilityUpdate) SetNillableDisplayName(s *string) *HolderResponsibilityUpdate {
	if s != nil {
		hru.SetDisplayName(*s)
	}
	return hru
}

// ClearDisplayName clears the value of the "display_name" field.
func (hru *HolderResponsibilityUpdate) ClearDisplayName() *HolderResponsibilityUpdate {
	hru.mutation.ClearDisplayName()
	return hru
}

// SetAbbreviation sets the "abbreviation" field.
func (hru *HolderResponsibilityUpdate) SetAbbreviation(s string) *HolderResponsibilityUpdate {
	hru.mutation.SetAbbreviation(s)
	return hru
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (hru *HolderResponsibilityUpdate) SetNillableAbbreviation(s *string) *HolderResponsibilityUpdate {
	if s != nil {
		hru.SetAbbreviation(*s)
	}
	return hru
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (hru *HolderResponsibilityUpdate) ClearAbbreviation() *HolderResponsibilityUpdate {
	hru.mutation.ClearAbbreviation()
	return hru
}

// SetDescription sets the "description" field.
func (hru *HolderResponsibilityUpdate) SetDescription(s string) *HolderResponsibilityUpdate {
	hru.mutation.SetDescription(s)
	return hru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (hru *HolderResponsibilityUpdate) SetNillableDescription(s *string) *HolderResponsibilityUpdate {
	if s != nil {
		hru.SetDescription(*s)
	}
	return hru
}

// ClearDescription clears the value of the "description" field.
func (hru *HolderResponsibilityUpdate) ClearDescription() *HolderResponsibilityUpdate {
	hru.mutation.ClearDescription()
	return hru
}

// SetExternalLink sets the "external_link" field.
func (hru *HolderResponsibilityUpdate) SetExternalLink(s string) *HolderResponsibilityUpdate {
	hru.mutation.SetExternalLink(s)
	return hru
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (hru *HolderResponsibilityUpdate) SetNillableExternalLink(s *string) *HolderResponsibilityUpdate {
	if s != nil {
		hru.SetExternalLink(*s)
	}
	return hru
}

// ClearExternalLink clears the value of the "external_link" field.
func (hru *HolderResponsibilityUpdate) ClearExternalLink() *HolderResponsibilityUpdate {
	hru.mutation.ClearExternalLink()
	return hru
}

// SetSlug sets the "slug" field.
func (hru *HolderResponsibilityUpdate) SetSlug(s string) *HolderResponsibilityUpdate {
	hru.mutation.SetSlug(s)
	return hru
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (hru *HolderResponsibilityUpdate) SetNillableSlug(s *string) *HolderResponsibilityUpdate {
	if s != nil {
		hru.SetSlug(*s)
	}
	return hru
}

// ClearSlug clears the value of the "slug" field.
func (hru *HolderResponsibilityUpdate) ClearSlug() *HolderResponsibilityUpdate {
	hru.mutation.ClearSlug()
	return hru
}

// AddHolderIDs adds the "holder" edge to the Holder entity by IDs.
func (hru *HolderResponsibilityUpdate) AddHolderIDs(ids ...int) *HolderResponsibilityUpdate {
	hru.mutation.AddHolderIDs(ids...)
	return hru
}

// AddHolder adds the "holder" edges to the Holder entity.
func (hru *HolderResponsibilityUpdate) AddHolder(h ...*Holder) *HolderResponsibilityUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hru.AddHolderIDs(ids...)
}

// Mutation returns the HolderResponsibilityMutation object of the builder.
func (hru *HolderResponsibilityUpdate) Mutation() *HolderResponsibilityMutation {
	return hru.mutation
}

// ClearHolder clears all "holder" edges to the Holder entity.
func (hru *HolderResponsibilityUpdate) ClearHolder() *HolderResponsibilityUpdate {
	hru.mutation.ClearHolder()
	return hru
}

// RemoveHolderIDs removes the "holder" edge to Holder entities by IDs.
func (hru *HolderResponsibilityUpdate) RemoveHolderIDs(ids ...int) *HolderResponsibilityUpdate {
	hru.mutation.RemoveHolderIDs(ids...)
	return hru
}

// RemoveHolder removes "holder" edges to Holder entities.
func (hru *HolderResponsibilityUpdate) RemoveHolder(h ...*Holder) *HolderResponsibilityUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hru.RemoveHolderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hru *HolderResponsibilityUpdate) Save(ctx context.Context) (int, error) {
	if err := hru.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, hru.sqlSave, hru.mutation, hru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hru *HolderResponsibilityUpdate) SaveX(ctx context.Context) int {
	affected, err := hru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hru *HolderResponsibilityUpdate) Exec(ctx context.Context) error {
	_, err := hru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hru *HolderResponsibilityUpdate) ExecX(ctx context.Context) {
	if err := hru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hru *HolderResponsibilityUpdate) defaults() error {
	if _, ok := hru.mutation.UpdatedAt(); !ok {
		if holderresponsibility.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized holderresponsibility.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := holderresponsibility.UpdateDefaultUpdatedAt()
		hru.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (hru *HolderResponsibilityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(holderresponsibility.Table, holderresponsibility.Columns, sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt))
	if ps := hru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hru.mutation.CreatedBy(); ok {
		_spec.SetField(holderresponsibility.FieldCreatedBy, field.TypeString, value)
	}
	if hru.mutation.CreatedByCleared() {
		_spec.ClearField(holderresponsibility.FieldCreatedBy, field.TypeString)
	}
	if value, ok := hru.mutation.UpdatedAt(); ok {
		_spec.SetField(holderresponsibility.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hru.mutation.UpdatedBy(); ok {
		_spec.SetField(holderresponsibility.FieldUpdatedBy, field.TypeString, value)
	}
	if hru.mutation.UpdatedByCleared() {
		_spec.ClearField(holderresponsibility.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := hru.mutation.DisplayName(); ok {
		_spec.SetField(holderresponsibility.FieldDisplayName, field.TypeString, value)
	}
	if hru.mutation.DisplayNameCleared() {
		_spec.ClearField(holderresponsibility.FieldDisplayName, field.TypeString)
	}
	if value, ok := hru.mutation.Abbreviation(); ok {
		_spec.SetField(holderresponsibility.FieldAbbreviation, field.TypeString, value)
	}
	if hru.mutation.AbbreviationCleared() {
		_spec.ClearField(holderresponsibility.FieldAbbreviation, field.TypeString)
	}
	if value, ok := hru.mutation.Description(); ok {
		_spec.SetField(holderresponsibility.FieldDescription, field.TypeString, value)
	}
	if hru.mutation.DescriptionCleared() {
		_spec.ClearField(holderresponsibility.FieldDescription, field.TypeString)
	}
	if value, ok := hru.mutation.ExternalLink(); ok {
		_spec.SetField(holderresponsibility.FieldExternalLink, field.TypeString, value)
	}
	if hru.mutation.ExternalLinkCleared() {
		_spec.ClearField(holderresponsibility.FieldExternalLink, field.TypeString)
	}
	if value, ok := hru.mutation.Slug(); ok {
		_spec.SetField(holderresponsibility.FieldSlug, field.TypeString, value)
	}
	if hru.mutation.SlugCleared() {
		_spec.ClearField(holderresponsibility.FieldSlug, field.TypeString)
	}
	if hru.mutation.HolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hru.mutation.RemovedHolderIDs(); len(nodes) > 0 && !hru.mutation.HolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hru.mutation.HolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{holderresponsibility.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hru.mutation.done = true
	return n, nil
}

// HolderResponsibilityUpdateOne is the builder for updating a single HolderResponsibility entity.
type HolderResponsibilityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HolderResponsibilityMutation
}

// SetCreatedBy sets the "created_by" field.
func (hruo *HolderResponsibilityUpdateOne) SetCreatedBy(s string) *HolderResponsibilityUpdateOne {
	hruo.mutation.SetCreatedBy(s)
	return hruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (hruo *HolderResponsibilityUpdateOne) SetNillableCreatedBy(s *string) *HolderResponsibilityUpdateOne {
	if s != nil {
		hruo.SetCreatedBy(*s)
	}
	return hruo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (hruo *HolderResponsibilityUpdateOne) ClearCreatedBy() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearCreatedBy()
	return hruo
}

// SetUpdatedAt sets the "updated_at" field.
func (hruo *HolderResponsibilityUpdateOne) SetUpdatedAt(t time.Time) *HolderResponsibilityUpdateOne {
	hruo.mutation.SetUpdatedAt(t)
	return hruo
}

// SetUpdatedBy sets the "updated_by" field.
func (hruo *HolderResponsibilityUpdateOne) SetUpdatedBy(s string) *HolderResponsibilityUpdateOne {
	hruo.mutation.SetUpdatedBy(s)
	return hruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (hruo *HolderResponsibilityUpdateOne) SetNillableUpdatedBy(s *string) *HolderResponsibilityUpdateOne {
	if s != nil {
		hruo.SetUpdatedBy(*s)
	}
	return hruo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (hruo *HolderResponsibilityUpdateOne) ClearUpdatedBy() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearUpdatedBy()
	return hruo
}

// SetDisplayName sets the "display_name" field.
func (hruo *HolderResponsibilityUpdateOne) SetDisplayName(s string) *HolderResponsibilityUpdateOne {
	hruo.mutation.SetDisplayName(s)
	return hruo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (hruo *HolderResponsibilityUpdateOne) SetNillableDisplayName(s *string) *HolderResponsibilityUpdateOne {
	if s != nil {
		hruo.SetDisplayName(*s)
	}
	return hruo
}

// ClearDisplayName clears the value of the "display_name" field.
func (hruo *HolderResponsibilityUpdateOne) ClearDisplayName() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearDisplayName()
	return hruo
}

// SetAbbreviation sets the "abbreviation" field.
func (hruo *HolderResponsibilityUpdateOne) SetAbbreviation(s string) *HolderResponsibilityUpdateOne {
	hruo.mutation.SetAbbreviation(s)
	return hruo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (hruo *HolderResponsibilityUpdateOne) SetNillableAbbreviation(s *string) *HolderResponsibilityUpdateOne {
	if s != nil {
		hruo.SetAbbreviation(*s)
	}
	return hruo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (hruo *HolderResponsibilityUpdateOne) ClearAbbreviation() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearAbbreviation()
	return hruo
}

// SetDescription sets the "description" field.
func (hruo *HolderResponsibilityUpdateOne) SetDescription(s string) *HolderResponsibilityUpdateOne {
	hruo.mutation.SetDescription(s)
	return hruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (hruo *HolderResponsibilityUpdateOne) SetNillableDescription(s *string) *HolderResponsibilityUpdateOne {
	if s != nil {
		hruo.SetDescription(*s)
	}
	return hruo
}

// ClearDescription clears the value of the "description" field.
func (hruo *HolderResponsibilityUpdateOne) ClearDescription() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearDescription()
	return hruo
}

// SetExternalLink sets the "external_link" field.
func (hruo *HolderResponsibilityUpdateOne) SetExternalLink(s string) *HolderResponsibilityUpdateOne {
	hruo.mutation.SetExternalLink(s)
	return hruo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (hruo *HolderResponsibilityUpdateOne) SetNillableExternalLink(s *string) *HolderResponsibilityUpdateOne {
	if s != nil {
		hruo.SetExternalLink(*s)
	}
	return hruo
}

// ClearExternalLink clears the value of the "external_link" field.
func (hruo *HolderResponsibilityUpdateOne) ClearExternalLink() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearExternalLink()
	return hruo
}

// SetSlug sets the "slug" field.
func (hruo *HolderResponsibilityUpdateOne) SetSlug(s string) *HolderResponsibilityUpdateOne {
	hruo.mutation.SetSlug(s)
	return hruo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (hruo *HolderResponsibilityUpdateOne) SetNillableSlug(s *string) *HolderResponsibilityUpdateOne {
	if s != nil {
		hruo.SetSlug(*s)
	}
	return hruo
}

// ClearSlug clears the value of the "slug" field.
func (hruo *HolderResponsibilityUpdateOne) ClearSlug() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearSlug()
	return hruo
}

// AddHolderIDs adds the "holder" edge to the Holder entity by IDs.
func (hruo *HolderResponsibilityUpdateOne) AddHolderIDs(ids ...int) *HolderResponsibilityUpdateOne {
	hruo.mutation.AddHolderIDs(ids...)
	return hruo
}

// AddHolder adds the "holder" edges to the Holder entity.
func (hruo *HolderResponsibilityUpdateOne) AddHolder(h ...*Holder) *HolderResponsibilityUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hruo.AddHolderIDs(ids...)
}

// Mutation returns the HolderResponsibilityMutation object of the builder.
func (hruo *HolderResponsibilityUpdateOne) Mutation() *HolderResponsibilityMutation {
	return hruo.mutation
}

// ClearHolder clears all "holder" edges to the Holder entity.
func (hruo *HolderResponsibilityUpdateOne) ClearHolder() *HolderResponsibilityUpdateOne {
	hruo.mutation.ClearHolder()
	return hruo
}

// RemoveHolderIDs removes the "holder" edge to Holder entities by IDs.
func (hruo *HolderResponsibilityUpdateOne) RemoveHolderIDs(ids ...int) *HolderResponsibilityUpdateOne {
	hruo.mutation.RemoveHolderIDs(ids...)
	return hruo
}

// RemoveHolder removes "holder" edges to Holder entities.
func (hruo *HolderResponsibilityUpdateOne) RemoveHolder(h ...*Holder) *HolderResponsibilityUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hruo.RemoveHolderIDs(ids...)
}

// Where appends a list predicates to the HolderResponsibilityUpdate builder.
func (hruo *HolderResponsibilityUpdateOne) Where(ps ...predicate.HolderResponsibility) *HolderResponsibilityUpdateOne {
	hruo.mutation.Where(ps...)
	return hruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hruo *HolderResponsibilityUpdateOne) Select(field string, fields ...string) *HolderResponsibilityUpdateOne {
	hruo.fields = append([]string{field}, fields...)
	return hruo
}

// Save executes the query and returns the updated HolderResponsibility entity.
func (hruo *HolderResponsibilityUpdateOne) Save(ctx context.Context) (*HolderResponsibility, error) {
	if err := hruo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, hruo.sqlSave, hruo.mutation, hruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hruo *HolderResponsibilityUpdateOne) SaveX(ctx context.Context) *HolderResponsibility {
	node, err := hruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hruo *HolderResponsibilityUpdateOne) Exec(ctx context.Context) error {
	_, err := hruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hruo *HolderResponsibilityUpdateOne) ExecX(ctx context.Context) {
	if err := hruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hruo *HolderResponsibilityUpdateOne) defaults() error {
	if _, ok := hruo.mutation.UpdatedAt(); !ok {
		if holderresponsibility.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized holderresponsibility.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := holderresponsibility.UpdateDefaultUpdatedAt()
		hruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (hruo *HolderResponsibilityUpdateOne) sqlSave(ctx context.Context) (_node *HolderResponsibility, err error) {
	_spec := sqlgraph.NewUpdateSpec(holderresponsibility.Table, holderresponsibility.Columns, sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt))
	id, ok := hruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HolderResponsibility.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, holderresponsibility.FieldID)
		for _, f := range fields {
			if !holderresponsibility.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != holderresponsibility.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hruo.mutation.CreatedBy(); ok {
		_spec.SetField(holderresponsibility.FieldCreatedBy, field.TypeString, value)
	}
	if hruo.mutation.CreatedByCleared() {
		_spec.ClearField(holderresponsibility.FieldCreatedBy, field.TypeString)
	}
	if value, ok := hruo.mutation.UpdatedAt(); ok {
		_spec.SetField(holderresponsibility.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hruo.mutation.UpdatedBy(); ok {
		_spec.SetField(holderresponsibility.FieldUpdatedBy, field.TypeString, value)
	}
	if hruo.mutation.UpdatedByCleared() {
		_spec.ClearField(holderresponsibility.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := hruo.mutation.DisplayName(); ok {
		_spec.SetField(holderresponsibility.FieldDisplayName, field.TypeString, value)
	}
	if hruo.mutation.DisplayNameCleared() {
		_spec.ClearField(holderresponsibility.FieldDisplayName, field.TypeString)
	}
	if value, ok := hruo.mutation.Abbreviation(); ok {
		_spec.SetField(holderresponsibility.FieldAbbreviation, field.TypeString, value)
	}
	if hruo.mutation.AbbreviationCleared() {
		_spec.ClearField(holderresponsibility.FieldAbbreviation, field.TypeString)
	}
	if value, ok := hruo.mutation.Description(); ok {
		_spec.SetField(holderresponsibility.FieldDescription, field.TypeString, value)
	}
	if hruo.mutation.DescriptionCleared() {
		_spec.ClearField(holderresponsibility.FieldDescription, field.TypeString)
	}
	if value, ok := hruo.mutation.ExternalLink(); ok {
		_spec.SetField(holderresponsibility.FieldExternalLink, field.TypeString, value)
	}
	if hruo.mutation.ExternalLinkCleared() {
		_spec.ClearField(holderresponsibility.FieldExternalLink, field.TypeString)
	}
	if value, ok := hruo.mutation.Slug(); ok {
		_spec.SetField(holderresponsibility.FieldSlug, field.TypeString, value)
	}
	if hruo.mutation.SlugCleared() {
		_spec.ClearField(holderresponsibility.FieldSlug, field.TypeString)
	}
	if hruo.mutation.HolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hruo.mutation.RemovedHolderIDs(); len(nodes) > 0 && !hruo.mutation.HolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hruo.mutation.HolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   holderresponsibility.HolderTable,
			Columns: holderresponsibility.HolderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HolderResponsibility{config: hruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{holderresponsibility.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hruo.mutation.done = true
	return _node, nil
}
