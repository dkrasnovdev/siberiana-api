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
	"github.com/dkrasnovdev/heritage-api/ent/art"
	"github.com/dkrasnovdev/heritage-api/ent/artstyle"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// ArtStyleUpdate is the builder for updating ArtStyle entities.
type ArtStyleUpdate struct {
	config
	hooks    []Hook
	mutation *ArtStyleMutation
}

// Where appends a list predicates to the ArtStyleUpdate builder.
func (asu *ArtStyleUpdate) Where(ps ...predicate.ArtStyle) *ArtStyleUpdate {
	asu.mutation.Where(ps...)
	return asu
}

// SetCreatedBy sets the "created_by" field.
func (asu *ArtStyleUpdate) SetCreatedBy(s string) *ArtStyleUpdate {
	asu.mutation.SetCreatedBy(s)
	return asu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (asu *ArtStyleUpdate) SetNillableCreatedBy(s *string) *ArtStyleUpdate {
	if s != nil {
		asu.SetCreatedBy(*s)
	}
	return asu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (asu *ArtStyleUpdate) ClearCreatedBy() *ArtStyleUpdate {
	asu.mutation.ClearCreatedBy()
	return asu
}

// SetUpdatedAt sets the "updated_at" field.
func (asu *ArtStyleUpdate) SetUpdatedAt(t time.Time) *ArtStyleUpdate {
	asu.mutation.SetUpdatedAt(t)
	return asu
}

// SetUpdatedBy sets the "updated_by" field.
func (asu *ArtStyleUpdate) SetUpdatedBy(s string) *ArtStyleUpdate {
	asu.mutation.SetUpdatedBy(s)
	return asu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (asu *ArtStyleUpdate) SetNillableUpdatedBy(s *string) *ArtStyleUpdate {
	if s != nil {
		asu.SetUpdatedBy(*s)
	}
	return asu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (asu *ArtStyleUpdate) ClearUpdatedBy() *ArtStyleUpdate {
	asu.mutation.ClearUpdatedBy()
	return asu
}

// SetDisplayName sets the "display_name" field.
func (asu *ArtStyleUpdate) SetDisplayName(s string) *ArtStyleUpdate {
	asu.mutation.SetDisplayName(s)
	return asu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (asu *ArtStyleUpdate) SetNillableDisplayName(s *string) *ArtStyleUpdate {
	if s != nil {
		asu.SetDisplayName(*s)
	}
	return asu
}

// ClearDisplayName clears the value of the "display_name" field.
func (asu *ArtStyleUpdate) ClearDisplayName() *ArtStyleUpdate {
	asu.mutation.ClearDisplayName()
	return asu
}

// SetAbbreviation sets the "abbreviation" field.
func (asu *ArtStyleUpdate) SetAbbreviation(s string) *ArtStyleUpdate {
	asu.mutation.SetAbbreviation(s)
	return asu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (asu *ArtStyleUpdate) SetNillableAbbreviation(s *string) *ArtStyleUpdate {
	if s != nil {
		asu.SetAbbreviation(*s)
	}
	return asu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (asu *ArtStyleUpdate) ClearAbbreviation() *ArtStyleUpdate {
	asu.mutation.ClearAbbreviation()
	return asu
}

// SetDescription sets the "description" field.
func (asu *ArtStyleUpdate) SetDescription(s string) *ArtStyleUpdate {
	asu.mutation.SetDescription(s)
	return asu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (asu *ArtStyleUpdate) SetNillableDescription(s *string) *ArtStyleUpdate {
	if s != nil {
		asu.SetDescription(*s)
	}
	return asu
}

// ClearDescription clears the value of the "description" field.
func (asu *ArtStyleUpdate) ClearDescription() *ArtStyleUpdate {
	asu.mutation.ClearDescription()
	return asu
}

// SetExternalLinks sets the "external_links" field.
func (asu *ArtStyleUpdate) SetExternalLinks(s []string) *ArtStyleUpdate {
	asu.mutation.SetExternalLinks(s)
	return asu
}

// AppendExternalLinks appends s to the "external_links" field.
func (asu *ArtStyleUpdate) AppendExternalLinks(s []string) *ArtStyleUpdate {
	asu.mutation.AppendExternalLinks(s)
	return asu
}

// ClearExternalLinks clears the value of the "external_links" field.
func (asu *ArtStyleUpdate) ClearExternalLinks() *ArtStyleUpdate {
	asu.mutation.ClearExternalLinks()
	return asu
}

// AddArtIDs adds the "art" edge to the Art entity by IDs.
func (asu *ArtStyleUpdate) AddArtIDs(ids ...int) *ArtStyleUpdate {
	asu.mutation.AddArtIDs(ids...)
	return asu
}

// AddArt adds the "art" edges to the Art entity.
func (asu *ArtStyleUpdate) AddArt(a ...*Art) *ArtStyleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asu.AddArtIDs(ids...)
}

// Mutation returns the ArtStyleMutation object of the builder.
func (asu *ArtStyleUpdate) Mutation() *ArtStyleMutation {
	return asu.mutation
}

// ClearArt clears all "art" edges to the Art entity.
func (asu *ArtStyleUpdate) ClearArt() *ArtStyleUpdate {
	asu.mutation.ClearArt()
	return asu
}

// RemoveArtIDs removes the "art" edge to Art entities by IDs.
func (asu *ArtStyleUpdate) RemoveArtIDs(ids ...int) *ArtStyleUpdate {
	asu.mutation.RemoveArtIDs(ids...)
	return asu
}

// RemoveArt removes "art" edges to Art entities.
func (asu *ArtStyleUpdate) RemoveArt(a ...*Art) *ArtStyleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asu.RemoveArtIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *ArtStyleUpdate) Save(ctx context.Context) (int, error) {
	if err := asu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, asu.sqlSave, asu.mutation, asu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asu *ArtStyleUpdate) SaveX(ctx context.Context) int {
	affected, err := asu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (asu *ArtStyleUpdate) Exec(ctx context.Context) error {
	_, err := asu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asu *ArtStyleUpdate) ExecX(ctx context.Context) {
	if err := asu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asu *ArtStyleUpdate) defaults() error {
	if _, ok := asu.mutation.UpdatedAt(); !ok {
		if artstyle.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized artstyle.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := artstyle.UpdateDefaultUpdatedAt()
		asu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (asu *ArtStyleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(artstyle.Table, artstyle.Columns, sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt))
	if ps := asu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asu.mutation.CreatedBy(); ok {
		_spec.SetField(artstyle.FieldCreatedBy, field.TypeString, value)
	}
	if asu.mutation.CreatedByCleared() {
		_spec.ClearField(artstyle.FieldCreatedBy, field.TypeString)
	}
	if value, ok := asu.mutation.UpdatedAt(); ok {
		_spec.SetField(artstyle.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := asu.mutation.UpdatedBy(); ok {
		_spec.SetField(artstyle.FieldUpdatedBy, field.TypeString, value)
	}
	if asu.mutation.UpdatedByCleared() {
		_spec.ClearField(artstyle.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := asu.mutation.DisplayName(); ok {
		_spec.SetField(artstyle.FieldDisplayName, field.TypeString, value)
	}
	if asu.mutation.DisplayNameCleared() {
		_spec.ClearField(artstyle.FieldDisplayName, field.TypeString)
	}
	if value, ok := asu.mutation.Abbreviation(); ok {
		_spec.SetField(artstyle.FieldAbbreviation, field.TypeString, value)
	}
	if asu.mutation.AbbreviationCleared() {
		_spec.ClearField(artstyle.FieldAbbreviation, field.TypeString)
	}
	if value, ok := asu.mutation.Description(); ok {
		_spec.SetField(artstyle.FieldDescription, field.TypeString, value)
	}
	if asu.mutation.DescriptionCleared() {
		_spec.ClearField(artstyle.FieldDescription, field.TypeString)
	}
	if value, ok := asu.mutation.ExternalLinks(); ok {
		_spec.SetField(artstyle.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := asu.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, artstyle.FieldExternalLinks, value)
		})
	}
	if asu.mutation.ExternalLinksCleared() {
		_spec.ClearField(artstyle.FieldExternalLinks, field.TypeJSON)
	}
	if asu.mutation.ArtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artstyle.ArtTable,
			Columns: artstyle.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.RemovedArtIDs(); len(nodes) > 0 && !asu.mutation.ArtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artstyle.ArtTable,
			Columns: artstyle.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.ArtIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artstyle.ArtTable,
			Columns: artstyle.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, asu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artstyle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	asu.mutation.done = true
	return n, nil
}

// ArtStyleUpdateOne is the builder for updating a single ArtStyle entity.
type ArtStyleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtStyleMutation
}

// SetCreatedBy sets the "created_by" field.
func (asuo *ArtStyleUpdateOne) SetCreatedBy(s string) *ArtStyleUpdateOne {
	asuo.mutation.SetCreatedBy(s)
	return asuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (asuo *ArtStyleUpdateOne) SetNillableCreatedBy(s *string) *ArtStyleUpdateOne {
	if s != nil {
		asuo.SetCreatedBy(*s)
	}
	return asuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (asuo *ArtStyleUpdateOne) ClearCreatedBy() *ArtStyleUpdateOne {
	asuo.mutation.ClearCreatedBy()
	return asuo
}

// SetUpdatedAt sets the "updated_at" field.
func (asuo *ArtStyleUpdateOne) SetUpdatedAt(t time.Time) *ArtStyleUpdateOne {
	asuo.mutation.SetUpdatedAt(t)
	return asuo
}

// SetUpdatedBy sets the "updated_by" field.
func (asuo *ArtStyleUpdateOne) SetUpdatedBy(s string) *ArtStyleUpdateOne {
	asuo.mutation.SetUpdatedBy(s)
	return asuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (asuo *ArtStyleUpdateOne) SetNillableUpdatedBy(s *string) *ArtStyleUpdateOne {
	if s != nil {
		asuo.SetUpdatedBy(*s)
	}
	return asuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (asuo *ArtStyleUpdateOne) ClearUpdatedBy() *ArtStyleUpdateOne {
	asuo.mutation.ClearUpdatedBy()
	return asuo
}

// SetDisplayName sets the "display_name" field.
func (asuo *ArtStyleUpdateOne) SetDisplayName(s string) *ArtStyleUpdateOne {
	asuo.mutation.SetDisplayName(s)
	return asuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (asuo *ArtStyleUpdateOne) SetNillableDisplayName(s *string) *ArtStyleUpdateOne {
	if s != nil {
		asuo.SetDisplayName(*s)
	}
	return asuo
}

// ClearDisplayName clears the value of the "display_name" field.
func (asuo *ArtStyleUpdateOne) ClearDisplayName() *ArtStyleUpdateOne {
	asuo.mutation.ClearDisplayName()
	return asuo
}

// SetAbbreviation sets the "abbreviation" field.
func (asuo *ArtStyleUpdateOne) SetAbbreviation(s string) *ArtStyleUpdateOne {
	asuo.mutation.SetAbbreviation(s)
	return asuo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (asuo *ArtStyleUpdateOne) SetNillableAbbreviation(s *string) *ArtStyleUpdateOne {
	if s != nil {
		asuo.SetAbbreviation(*s)
	}
	return asuo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (asuo *ArtStyleUpdateOne) ClearAbbreviation() *ArtStyleUpdateOne {
	asuo.mutation.ClearAbbreviation()
	return asuo
}

// SetDescription sets the "description" field.
func (asuo *ArtStyleUpdateOne) SetDescription(s string) *ArtStyleUpdateOne {
	asuo.mutation.SetDescription(s)
	return asuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (asuo *ArtStyleUpdateOne) SetNillableDescription(s *string) *ArtStyleUpdateOne {
	if s != nil {
		asuo.SetDescription(*s)
	}
	return asuo
}

// ClearDescription clears the value of the "description" field.
func (asuo *ArtStyleUpdateOne) ClearDescription() *ArtStyleUpdateOne {
	asuo.mutation.ClearDescription()
	return asuo
}

// SetExternalLinks sets the "external_links" field.
func (asuo *ArtStyleUpdateOne) SetExternalLinks(s []string) *ArtStyleUpdateOne {
	asuo.mutation.SetExternalLinks(s)
	return asuo
}

// AppendExternalLinks appends s to the "external_links" field.
func (asuo *ArtStyleUpdateOne) AppendExternalLinks(s []string) *ArtStyleUpdateOne {
	asuo.mutation.AppendExternalLinks(s)
	return asuo
}

// ClearExternalLinks clears the value of the "external_links" field.
func (asuo *ArtStyleUpdateOne) ClearExternalLinks() *ArtStyleUpdateOne {
	asuo.mutation.ClearExternalLinks()
	return asuo
}

// AddArtIDs adds the "art" edge to the Art entity by IDs.
func (asuo *ArtStyleUpdateOne) AddArtIDs(ids ...int) *ArtStyleUpdateOne {
	asuo.mutation.AddArtIDs(ids...)
	return asuo
}

// AddArt adds the "art" edges to the Art entity.
func (asuo *ArtStyleUpdateOne) AddArt(a ...*Art) *ArtStyleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asuo.AddArtIDs(ids...)
}

// Mutation returns the ArtStyleMutation object of the builder.
func (asuo *ArtStyleUpdateOne) Mutation() *ArtStyleMutation {
	return asuo.mutation
}

// ClearArt clears all "art" edges to the Art entity.
func (asuo *ArtStyleUpdateOne) ClearArt() *ArtStyleUpdateOne {
	asuo.mutation.ClearArt()
	return asuo
}

// RemoveArtIDs removes the "art" edge to Art entities by IDs.
func (asuo *ArtStyleUpdateOne) RemoveArtIDs(ids ...int) *ArtStyleUpdateOne {
	asuo.mutation.RemoveArtIDs(ids...)
	return asuo
}

// RemoveArt removes "art" edges to Art entities.
func (asuo *ArtStyleUpdateOne) RemoveArt(a ...*Art) *ArtStyleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asuo.RemoveArtIDs(ids...)
}

// Where appends a list predicates to the ArtStyleUpdate builder.
func (asuo *ArtStyleUpdateOne) Where(ps ...predicate.ArtStyle) *ArtStyleUpdateOne {
	asuo.mutation.Where(ps...)
	return asuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (asuo *ArtStyleUpdateOne) Select(field string, fields ...string) *ArtStyleUpdateOne {
	asuo.fields = append([]string{field}, fields...)
	return asuo
}

// Save executes the query and returns the updated ArtStyle entity.
func (asuo *ArtStyleUpdateOne) Save(ctx context.Context) (*ArtStyle, error) {
	if err := asuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, asuo.sqlSave, asuo.mutation, asuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asuo *ArtStyleUpdateOne) SaveX(ctx context.Context) *ArtStyle {
	node, err := asuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asuo *ArtStyleUpdateOne) Exec(ctx context.Context) error {
	_, err := asuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asuo *ArtStyleUpdateOne) ExecX(ctx context.Context) {
	if err := asuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asuo *ArtStyleUpdateOne) defaults() error {
	if _, ok := asuo.mutation.UpdatedAt(); !ok {
		if artstyle.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized artstyle.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := artstyle.UpdateDefaultUpdatedAt()
		asuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (asuo *ArtStyleUpdateOne) sqlSave(ctx context.Context) (_node *ArtStyle, err error) {
	_spec := sqlgraph.NewUpdateSpec(artstyle.Table, artstyle.Columns, sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt))
	id, ok := asuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArtStyle.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := asuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artstyle.FieldID)
		for _, f := range fields {
			if !artstyle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != artstyle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := asuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asuo.mutation.CreatedBy(); ok {
		_spec.SetField(artstyle.FieldCreatedBy, field.TypeString, value)
	}
	if asuo.mutation.CreatedByCleared() {
		_spec.ClearField(artstyle.FieldCreatedBy, field.TypeString)
	}
	if value, ok := asuo.mutation.UpdatedAt(); ok {
		_spec.SetField(artstyle.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := asuo.mutation.UpdatedBy(); ok {
		_spec.SetField(artstyle.FieldUpdatedBy, field.TypeString, value)
	}
	if asuo.mutation.UpdatedByCleared() {
		_spec.ClearField(artstyle.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := asuo.mutation.DisplayName(); ok {
		_spec.SetField(artstyle.FieldDisplayName, field.TypeString, value)
	}
	if asuo.mutation.DisplayNameCleared() {
		_spec.ClearField(artstyle.FieldDisplayName, field.TypeString)
	}
	if value, ok := asuo.mutation.Abbreviation(); ok {
		_spec.SetField(artstyle.FieldAbbreviation, field.TypeString, value)
	}
	if asuo.mutation.AbbreviationCleared() {
		_spec.ClearField(artstyle.FieldAbbreviation, field.TypeString)
	}
	if value, ok := asuo.mutation.Description(); ok {
		_spec.SetField(artstyle.FieldDescription, field.TypeString, value)
	}
	if asuo.mutation.DescriptionCleared() {
		_spec.ClearField(artstyle.FieldDescription, field.TypeString)
	}
	if value, ok := asuo.mutation.ExternalLinks(); ok {
		_spec.SetField(artstyle.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := asuo.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, artstyle.FieldExternalLinks, value)
		})
	}
	if asuo.mutation.ExternalLinksCleared() {
		_spec.ClearField(artstyle.FieldExternalLinks, field.TypeJSON)
	}
	if asuo.mutation.ArtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artstyle.ArtTable,
			Columns: artstyle.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.RemovedArtIDs(); len(nodes) > 0 && !asuo.mutation.ArtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artstyle.ArtTable,
			Columns: artstyle.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.ArtIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artstyle.ArtTable,
			Columns: artstyle.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ArtStyle{config: asuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artstyle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	asuo.mutation.done = true
	return _node, nil
}
