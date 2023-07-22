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
	"github.com/dkrasnovdev/heritage-api/ent/organization"
	"github.com/dkrasnovdev/heritage-api/ent/organizationtype"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// OrganizationTypeUpdate is the builder for updating OrganizationType entities.
type OrganizationTypeUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationTypeMutation
}

// Where appends a list predicates to the OrganizationTypeUpdate builder.
func (otu *OrganizationTypeUpdate) Where(ps ...predicate.OrganizationType) *OrganizationTypeUpdate {
	otu.mutation.Where(ps...)
	return otu
}

// SetCreatedBy sets the "created_by" field.
func (otu *OrganizationTypeUpdate) SetCreatedBy(s string) *OrganizationTypeUpdate {
	otu.mutation.SetCreatedBy(s)
	return otu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (otu *OrganizationTypeUpdate) SetNillableCreatedBy(s *string) *OrganizationTypeUpdate {
	if s != nil {
		otu.SetCreatedBy(*s)
	}
	return otu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (otu *OrganizationTypeUpdate) ClearCreatedBy() *OrganizationTypeUpdate {
	otu.mutation.ClearCreatedBy()
	return otu
}

// SetUpdatedAt sets the "updated_at" field.
func (otu *OrganizationTypeUpdate) SetUpdatedAt(t time.Time) *OrganizationTypeUpdate {
	otu.mutation.SetUpdatedAt(t)
	return otu
}

// SetUpdatedBy sets the "updated_by" field.
func (otu *OrganizationTypeUpdate) SetUpdatedBy(s string) *OrganizationTypeUpdate {
	otu.mutation.SetUpdatedBy(s)
	return otu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (otu *OrganizationTypeUpdate) SetNillableUpdatedBy(s *string) *OrganizationTypeUpdate {
	if s != nil {
		otu.SetUpdatedBy(*s)
	}
	return otu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (otu *OrganizationTypeUpdate) ClearUpdatedBy() *OrganizationTypeUpdate {
	otu.mutation.ClearUpdatedBy()
	return otu
}

// SetDisplayName sets the "display_name" field.
func (otu *OrganizationTypeUpdate) SetDisplayName(s string) *OrganizationTypeUpdate {
	otu.mutation.SetDisplayName(s)
	return otu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (otu *OrganizationTypeUpdate) SetNillableDisplayName(s *string) *OrganizationTypeUpdate {
	if s != nil {
		otu.SetDisplayName(*s)
	}
	return otu
}

// ClearDisplayName clears the value of the "display_name" field.
func (otu *OrganizationTypeUpdate) ClearDisplayName() *OrganizationTypeUpdate {
	otu.mutation.ClearDisplayName()
	return otu
}

// SetDescription sets the "description" field.
func (otu *OrganizationTypeUpdate) SetDescription(s string) *OrganizationTypeUpdate {
	otu.mutation.SetDescription(s)
	return otu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (otu *OrganizationTypeUpdate) SetNillableDescription(s *string) *OrganizationTypeUpdate {
	if s != nil {
		otu.SetDescription(*s)
	}
	return otu
}

// ClearDescription clears the value of the "description" field.
func (otu *OrganizationTypeUpdate) ClearDescription() *OrganizationTypeUpdate {
	otu.mutation.ClearDescription()
	return otu
}

// SetExternalLinks sets the "external_links" field.
func (otu *OrganizationTypeUpdate) SetExternalLinks(s []string) *OrganizationTypeUpdate {
	otu.mutation.SetExternalLinks(s)
	return otu
}

// AppendExternalLinks appends s to the "external_links" field.
func (otu *OrganizationTypeUpdate) AppendExternalLinks(s []string) *OrganizationTypeUpdate {
	otu.mutation.AppendExternalLinks(s)
	return otu
}

// ClearExternalLinks clears the value of the "external_links" field.
func (otu *OrganizationTypeUpdate) ClearExternalLinks() *OrganizationTypeUpdate {
	otu.mutation.ClearExternalLinks()
	return otu
}

// AddOrganizationIDs adds the "organizations" edge to the Organization entity by IDs.
func (otu *OrganizationTypeUpdate) AddOrganizationIDs(ids ...int) *OrganizationTypeUpdate {
	otu.mutation.AddOrganizationIDs(ids...)
	return otu
}

// AddOrganizations adds the "organizations" edges to the Organization entity.
func (otu *OrganizationTypeUpdate) AddOrganizations(o ...*Organization) *OrganizationTypeUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otu.AddOrganizationIDs(ids...)
}

// Mutation returns the OrganizationTypeMutation object of the builder.
func (otu *OrganizationTypeUpdate) Mutation() *OrganizationTypeMutation {
	return otu.mutation
}

// ClearOrganizations clears all "organizations" edges to the Organization entity.
func (otu *OrganizationTypeUpdate) ClearOrganizations() *OrganizationTypeUpdate {
	otu.mutation.ClearOrganizations()
	return otu
}

// RemoveOrganizationIDs removes the "organizations" edge to Organization entities by IDs.
func (otu *OrganizationTypeUpdate) RemoveOrganizationIDs(ids ...int) *OrganizationTypeUpdate {
	otu.mutation.RemoveOrganizationIDs(ids...)
	return otu
}

// RemoveOrganizations removes "organizations" edges to Organization entities.
func (otu *OrganizationTypeUpdate) RemoveOrganizations(o ...*Organization) *OrganizationTypeUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otu.RemoveOrganizationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (otu *OrganizationTypeUpdate) Save(ctx context.Context) (int, error) {
	if err := otu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, otu.sqlSave, otu.mutation, otu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (otu *OrganizationTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := otu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (otu *OrganizationTypeUpdate) Exec(ctx context.Context) error {
	_, err := otu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otu *OrganizationTypeUpdate) ExecX(ctx context.Context) {
	if err := otu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (otu *OrganizationTypeUpdate) defaults() error {
	if _, ok := otu.mutation.UpdatedAt(); !ok {
		if organizationtype.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized organizationtype.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := organizationtype.UpdateDefaultUpdatedAt()
		otu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (otu *OrganizationTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(organizationtype.Table, organizationtype.Columns, sqlgraph.NewFieldSpec(organizationtype.FieldID, field.TypeInt))
	if ps := otu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otu.mutation.CreatedBy(); ok {
		_spec.SetField(organizationtype.FieldCreatedBy, field.TypeString, value)
	}
	if otu.mutation.CreatedByCleared() {
		_spec.ClearField(organizationtype.FieldCreatedBy, field.TypeString)
	}
	if value, ok := otu.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationtype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := otu.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationtype.FieldUpdatedBy, field.TypeString, value)
	}
	if otu.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationtype.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := otu.mutation.DisplayName(); ok {
		_spec.SetField(organizationtype.FieldDisplayName, field.TypeString, value)
	}
	if otu.mutation.DisplayNameCleared() {
		_spec.ClearField(organizationtype.FieldDisplayName, field.TypeString)
	}
	if value, ok := otu.mutation.Description(); ok {
		_spec.SetField(organizationtype.FieldDescription, field.TypeString, value)
	}
	if otu.mutation.DescriptionCleared() {
		_spec.ClearField(organizationtype.FieldDescription, field.TypeString)
	}
	if value, ok := otu.mutation.ExternalLinks(); ok {
		_spec.SetField(organizationtype.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := otu.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationtype.FieldExternalLinks, value)
		})
	}
	if otu.mutation.ExternalLinksCleared() {
		_spec.ClearField(organizationtype.FieldExternalLinks, field.TypeJSON)
	}
	if otu.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationtype.OrganizationsTable,
			Columns: []string{organizationtype.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otu.mutation.RemovedOrganizationsIDs(); len(nodes) > 0 && !otu.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationtype.OrganizationsTable,
			Columns: []string{organizationtype.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otu.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationtype.OrganizationsTable,
			Columns: []string{organizationtype.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, otu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	otu.mutation.done = true
	return n, nil
}

// OrganizationTypeUpdateOne is the builder for updating a single OrganizationType entity.
type OrganizationTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationTypeMutation
}

// SetCreatedBy sets the "created_by" field.
func (otuo *OrganizationTypeUpdateOne) SetCreatedBy(s string) *OrganizationTypeUpdateOne {
	otuo.mutation.SetCreatedBy(s)
	return otuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (otuo *OrganizationTypeUpdateOne) SetNillableCreatedBy(s *string) *OrganizationTypeUpdateOne {
	if s != nil {
		otuo.SetCreatedBy(*s)
	}
	return otuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (otuo *OrganizationTypeUpdateOne) ClearCreatedBy() *OrganizationTypeUpdateOne {
	otuo.mutation.ClearCreatedBy()
	return otuo
}

// SetUpdatedAt sets the "updated_at" field.
func (otuo *OrganizationTypeUpdateOne) SetUpdatedAt(t time.Time) *OrganizationTypeUpdateOne {
	otuo.mutation.SetUpdatedAt(t)
	return otuo
}

// SetUpdatedBy sets the "updated_by" field.
func (otuo *OrganizationTypeUpdateOne) SetUpdatedBy(s string) *OrganizationTypeUpdateOne {
	otuo.mutation.SetUpdatedBy(s)
	return otuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (otuo *OrganizationTypeUpdateOne) SetNillableUpdatedBy(s *string) *OrganizationTypeUpdateOne {
	if s != nil {
		otuo.SetUpdatedBy(*s)
	}
	return otuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (otuo *OrganizationTypeUpdateOne) ClearUpdatedBy() *OrganizationTypeUpdateOne {
	otuo.mutation.ClearUpdatedBy()
	return otuo
}

// SetDisplayName sets the "display_name" field.
func (otuo *OrganizationTypeUpdateOne) SetDisplayName(s string) *OrganizationTypeUpdateOne {
	otuo.mutation.SetDisplayName(s)
	return otuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (otuo *OrganizationTypeUpdateOne) SetNillableDisplayName(s *string) *OrganizationTypeUpdateOne {
	if s != nil {
		otuo.SetDisplayName(*s)
	}
	return otuo
}

// ClearDisplayName clears the value of the "display_name" field.
func (otuo *OrganizationTypeUpdateOne) ClearDisplayName() *OrganizationTypeUpdateOne {
	otuo.mutation.ClearDisplayName()
	return otuo
}

// SetDescription sets the "description" field.
func (otuo *OrganizationTypeUpdateOne) SetDescription(s string) *OrganizationTypeUpdateOne {
	otuo.mutation.SetDescription(s)
	return otuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (otuo *OrganizationTypeUpdateOne) SetNillableDescription(s *string) *OrganizationTypeUpdateOne {
	if s != nil {
		otuo.SetDescription(*s)
	}
	return otuo
}

// ClearDescription clears the value of the "description" field.
func (otuo *OrganizationTypeUpdateOne) ClearDescription() *OrganizationTypeUpdateOne {
	otuo.mutation.ClearDescription()
	return otuo
}

// SetExternalLinks sets the "external_links" field.
func (otuo *OrganizationTypeUpdateOne) SetExternalLinks(s []string) *OrganizationTypeUpdateOne {
	otuo.mutation.SetExternalLinks(s)
	return otuo
}

// AppendExternalLinks appends s to the "external_links" field.
func (otuo *OrganizationTypeUpdateOne) AppendExternalLinks(s []string) *OrganizationTypeUpdateOne {
	otuo.mutation.AppendExternalLinks(s)
	return otuo
}

// ClearExternalLinks clears the value of the "external_links" field.
func (otuo *OrganizationTypeUpdateOne) ClearExternalLinks() *OrganizationTypeUpdateOne {
	otuo.mutation.ClearExternalLinks()
	return otuo
}

// AddOrganizationIDs adds the "organizations" edge to the Organization entity by IDs.
func (otuo *OrganizationTypeUpdateOne) AddOrganizationIDs(ids ...int) *OrganizationTypeUpdateOne {
	otuo.mutation.AddOrganizationIDs(ids...)
	return otuo
}

// AddOrganizations adds the "organizations" edges to the Organization entity.
func (otuo *OrganizationTypeUpdateOne) AddOrganizations(o ...*Organization) *OrganizationTypeUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otuo.AddOrganizationIDs(ids...)
}

// Mutation returns the OrganizationTypeMutation object of the builder.
func (otuo *OrganizationTypeUpdateOne) Mutation() *OrganizationTypeMutation {
	return otuo.mutation
}

// ClearOrganizations clears all "organizations" edges to the Organization entity.
func (otuo *OrganizationTypeUpdateOne) ClearOrganizations() *OrganizationTypeUpdateOne {
	otuo.mutation.ClearOrganizations()
	return otuo
}

// RemoveOrganizationIDs removes the "organizations" edge to Organization entities by IDs.
func (otuo *OrganizationTypeUpdateOne) RemoveOrganizationIDs(ids ...int) *OrganizationTypeUpdateOne {
	otuo.mutation.RemoveOrganizationIDs(ids...)
	return otuo
}

// RemoveOrganizations removes "organizations" edges to Organization entities.
func (otuo *OrganizationTypeUpdateOne) RemoveOrganizations(o ...*Organization) *OrganizationTypeUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otuo.RemoveOrganizationIDs(ids...)
}

// Where appends a list predicates to the OrganizationTypeUpdate builder.
func (otuo *OrganizationTypeUpdateOne) Where(ps ...predicate.OrganizationType) *OrganizationTypeUpdateOne {
	otuo.mutation.Where(ps...)
	return otuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (otuo *OrganizationTypeUpdateOne) Select(field string, fields ...string) *OrganizationTypeUpdateOne {
	otuo.fields = append([]string{field}, fields...)
	return otuo
}

// Save executes the query and returns the updated OrganizationType entity.
func (otuo *OrganizationTypeUpdateOne) Save(ctx context.Context) (*OrganizationType, error) {
	if err := otuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, otuo.sqlSave, otuo.mutation, otuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (otuo *OrganizationTypeUpdateOne) SaveX(ctx context.Context) *OrganizationType {
	node, err := otuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (otuo *OrganizationTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := otuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otuo *OrganizationTypeUpdateOne) ExecX(ctx context.Context) {
	if err := otuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (otuo *OrganizationTypeUpdateOne) defaults() error {
	if _, ok := otuo.mutation.UpdatedAt(); !ok {
		if organizationtype.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized organizationtype.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := organizationtype.UpdateDefaultUpdatedAt()
		otuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (otuo *OrganizationTypeUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationType, err error) {
	_spec := sqlgraph.NewUpdateSpec(organizationtype.Table, organizationtype.Columns, sqlgraph.NewFieldSpec(organizationtype.FieldID, field.TypeInt))
	id, ok := otuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrganizationType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := otuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationtype.FieldID)
		for _, f := range fields {
			if !organizationtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organizationtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := otuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otuo.mutation.CreatedBy(); ok {
		_spec.SetField(organizationtype.FieldCreatedBy, field.TypeString, value)
	}
	if otuo.mutation.CreatedByCleared() {
		_spec.ClearField(organizationtype.FieldCreatedBy, field.TypeString)
	}
	if value, ok := otuo.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationtype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := otuo.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationtype.FieldUpdatedBy, field.TypeString, value)
	}
	if otuo.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationtype.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := otuo.mutation.DisplayName(); ok {
		_spec.SetField(organizationtype.FieldDisplayName, field.TypeString, value)
	}
	if otuo.mutation.DisplayNameCleared() {
		_spec.ClearField(organizationtype.FieldDisplayName, field.TypeString)
	}
	if value, ok := otuo.mutation.Description(); ok {
		_spec.SetField(organizationtype.FieldDescription, field.TypeString, value)
	}
	if otuo.mutation.DescriptionCleared() {
		_spec.ClearField(organizationtype.FieldDescription, field.TypeString)
	}
	if value, ok := otuo.mutation.ExternalLinks(); ok {
		_spec.SetField(organizationtype.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := otuo.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationtype.FieldExternalLinks, value)
		})
	}
	if otuo.mutation.ExternalLinksCleared() {
		_spec.ClearField(organizationtype.FieldExternalLinks, field.TypeJSON)
	}
	if otuo.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationtype.OrganizationsTable,
			Columns: []string{organizationtype.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otuo.mutation.RemovedOrganizationsIDs(); len(nodes) > 0 && !otuo.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationtype.OrganizationsTable,
			Columns: []string{organizationtype.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otuo.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationtype.OrganizationsTable,
			Columns: []string{organizationtype.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrganizationType{config: otuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, otuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	otuo.mutation.done = true
	return _node, nil
}