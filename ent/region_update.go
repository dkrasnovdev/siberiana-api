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
	"github.com/dkrasnovdev/siberiana-api/ent/location"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
	"github.com/dkrasnovdev/siberiana-api/ent/region"
)

// RegionUpdate is the builder for updating Region entities.
type RegionUpdate struct {
	config
	hooks    []Hook
	mutation *RegionMutation
}

// Where appends a list predicates to the RegionUpdate builder.
func (ru *RegionUpdate) Where(ps ...predicate.Region) *RegionUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedBy sets the "created_by" field.
func (ru *RegionUpdate) SetCreatedBy(s string) *RegionUpdate {
	ru.mutation.SetCreatedBy(s)
	return ru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableCreatedBy(s *string) *RegionUpdate {
	if s != nil {
		ru.SetCreatedBy(*s)
	}
	return ru
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ru *RegionUpdate) ClearCreatedBy() *RegionUpdate {
	ru.mutation.ClearCreatedBy()
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RegionUpdate) SetUpdatedAt(t time.Time) *RegionUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetUpdatedBy sets the "updated_by" field.
func (ru *RegionUpdate) SetUpdatedBy(s string) *RegionUpdate {
	ru.mutation.SetUpdatedBy(s)
	return ru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableUpdatedBy(s *string) *RegionUpdate {
	if s != nil {
		ru.SetUpdatedBy(*s)
	}
	return ru
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ru *RegionUpdate) ClearUpdatedBy() *RegionUpdate {
	ru.mutation.ClearUpdatedBy()
	return ru
}

// SetDisplayName sets the "display_name" field.
func (ru *RegionUpdate) SetDisplayName(s string) *RegionUpdate {
	ru.mutation.SetDisplayName(s)
	return ru
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableDisplayName(s *string) *RegionUpdate {
	if s != nil {
		ru.SetDisplayName(*s)
	}
	return ru
}

// ClearDisplayName clears the value of the "display_name" field.
func (ru *RegionUpdate) ClearDisplayName() *RegionUpdate {
	ru.mutation.ClearDisplayName()
	return ru
}

// SetAbbreviation sets the "abbreviation" field.
func (ru *RegionUpdate) SetAbbreviation(s string) *RegionUpdate {
	ru.mutation.SetAbbreviation(s)
	return ru
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableAbbreviation(s *string) *RegionUpdate {
	if s != nil {
		ru.SetAbbreviation(*s)
	}
	return ru
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (ru *RegionUpdate) ClearAbbreviation() *RegionUpdate {
	ru.mutation.ClearAbbreviation()
	return ru
}

// SetDescription sets the "description" field.
func (ru *RegionUpdate) SetDescription(s string) *RegionUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableDescription(s *string) *RegionUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// ClearDescription clears the value of the "description" field.
func (ru *RegionUpdate) ClearDescription() *RegionUpdate {
	ru.mutation.ClearDescription()
	return ru
}

// SetExternalLink sets the "external_link" field.
func (ru *RegionUpdate) SetExternalLink(s string) *RegionUpdate {
	ru.mutation.SetExternalLink(s)
	return ru
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableExternalLink(s *string) *RegionUpdate {
	if s != nil {
		ru.SetExternalLink(*s)
	}
	return ru
}

// ClearExternalLink clears the value of the "external_link" field.
func (ru *RegionUpdate) ClearExternalLink() *RegionUpdate {
	ru.mutation.ClearExternalLink()
	return ru
}

// SetSlug sets the "slug" field.
func (ru *RegionUpdate) SetSlug(s string) *RegionUpdate {
	ru.mutation.SetSlug(s)
	return ru
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableSlug(s *string) *RegionUpdate {
	if s != nil {
		ru.SetSlug(*s)
	}
	return ru
}

// ClearSlug clears the value of the "slug" field.
func (ru *RegionUpdate) ClearSlug() *RegionUpdate {
	ru.mutation.ClearSlug()
	return ru
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (ru *RegionUpdate) SetLocationID(id int) *RegionUpdate {
	ru.mutation.SetLocationID(id)
	return ru
}

// SetNillableLocationID sets the "location" edge to the Location entity by ID if the given value is not nil.
func (ru *RegionUpdate) SetNillableLocationID(id *int) *RegionUpdate {
	if id != nil {
		ru = ru.SetLocationID(*id)
	}
	return ru
}

// SetLocation sets the "location" edge to the Location entity.
func (ru *RegionUpdate) SetLocation(l *Location) *RegionUpdate {
	return ru.SetLocationID(l.ID)
}

// Mutation returns the RegionMutation object of the builder.
func (ru *RegionUpdate) Mutation() *RegionMutation {
	return ru.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (ru *RegionUpdate) ClearLocation() *RegionUpdate {
	ru.mutation.ClearLocation()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RegionUpdate) Save(ctx context.Context) (int, error) {
	if err := ru.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RegionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RegionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RegionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RegionUpdate) defaults() error {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		if region.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized region.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := region.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ru *RegionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(region.Table, region.Columns, sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CreatedBy(); ok {
		_spec.SetField(region.FieldCreatedBy, field.TypeString, value)
	}
	if ru.mutation.CreatedByCleared() {
		_spec.ClearField(region.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(region.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedBy(); ok {
		_spec.SetField(region.FieldUpdatedBy, field.TypeString, value)
	}
	if ru.mutation.UpdatedByCleared() {
		_spec.ClearField(region.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ru.mutation.DisplayName(); ok {
		_spec.SetField(region.FieldDisplayName, field.TypeString, value)
	}
	if ru.mutation.DisplayNameCleared() {
		_spec.ClearField(region.FieldDisplayName, field.TypeString)
	}
	if value, ok := ru.mutation.Abbreviation(); ok {
		_spec.SetField(region.FieldAbbreviation, field.TypeString, value)
	}
	if ru.mutation.AbbreviationCleared() {
		_spec.ClearField(region.FieldAbbreviation, field.TypeString)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(region.FieldDescription, field.TypeString, value)
	}
	if ru.mutation.DescriptionCleared() {
		_spec.ClearField(region.FieldDescription, field.TypeString)
	}
	if value, ok := ru.mutation.ExternalLink(); ok {
		_spec.SetField(region.FieldExternalLink, field.TypeString, value)
	}
	if ru.mutation.ExternalLinkCleared() {
		_spec.ClearField(region.FieldExternalLink, field.TypeString)
	}
	if value, ok := ru.mutation.Slug(); ok {
		_spec.SetField(region.FieldSlug, field.TypeString, value)
	}
	if ru.mutation.SlugCleared() {
		_spec.ClearField(region.FieldSlug, field.TypeString)
	}
	if ru.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   region.LocationTable,
			Columns: []string{region.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   region.LocationTable,
			Columns: []string{region.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{region.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RegionUpdateOne is the builder for updating a single Region entity.
type RegionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RegionMutation
}

// SetCreatedBy sets the "created_by" field.
func (ruo *RegionUpdateOne) SetCreatedBy(s string) *RegionUpdateOne {
	ruo.mutation.SetCreatedBy(s)
	return ruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableCreatedBy(s *string) *RegionUpdateOne {
	if s != nil {
		ruo.SetCreatedBy(*s)
	}
	return ruo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ruo *RegionUpdateOne) ClearCreatedBy() *RegionUpdateOne {
	ruo.mutation.ClearCreatedBy()
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RegionUpdateOne) SetUpdatedAt(t time.Time) *RegionUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetUpdatedBy sets the "updated_by" field.
func (ruo *RegionUpdateOne) SetUpdatedBy(s string) *RegionUpdateOne {
	ruo.mutation.SetUpdatedBy(s)
	return ruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableUpdatedBy(s *string) *RegionUpdateOne {
	if s != nil {
		ruo.SetUpdatedBy(*s)
	}
	return ruo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ruo *RegionUpdateOne) ClearUpdatedBy() *RegionUpdateOne {
	ruo.mutation.ClearUpdatedBy()
	return ruo
}

// SetDisplayName sets the "display_name" field.
func (ruo *RegionUpdateOne) SetDisplayName(s string) *RegionUpdateOne {
	ruo.mutation.SetDisplayName(s)
	return ruo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableDisplayName(s *string) *RegionUpdateOne {
	if s != nil {
		ruo.SetDisplayName(*s)
	}
	return ruo
}

// ClearDisplayName clears the value of the "display_name" field.
func (ruo *RegionUpdateOne) ClearDisplayName() *RegionUpdateOne {
	ruo.mutation.ClearDisplayName()
	return ruo
}

// SetAbbreviation sets the "abbreviation" field.
func (ruo *RegionUpdateOne) SetAbbreviation(s string) *RegionUpdateOne {
	ruo.mutation.SetAbbreviation(s)
	return ruo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableAbbreviation(s *string) *RegionUpdateOne {
	if s != nil {
		ruo.SetAbbreviation(*s)
	}
	return ruo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (ruo *RegionUpdateOne) ClearAbbreviation() *RegionUpdateOne {
	ruo.mutation.ClearAbbreviation()
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RegionUpdateOne) SetDescription(s string) *RegionUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableDescription(s *string) *RegionUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// ClearDescription clears the value of the "description" field.
func (ruo *RegionUpdateOne) ClearDescription() *RegionUpdateOne {
	ruo.mutation.ClearDescription()
	return ruo
}

// SetExternalLink sets the "external_link" field.
func (ruo *RegionUpdateOne) SetExternalLink(s string) *RegionUpdateOne {
	ruo.mutation.SetExternalLink(s)
	return ruo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableExternalLink(s *string) *RegionUpdateOne {
	if s != nil {
		ruo.SetExternalLink(*s)
	}
	return ruo
}

// ClearExternalLink clears the value of the "external_link" field.
func (ruo *RegionUpdateOne) ClearExternalLink() *RegionUpdateOne {
	ruo.mutation.ClearExternalLink()
	return ruo
}

// SetSlug sets the "slug" field.
func (ruo *RegionUpdateOne) SetSlug(s string) *RegionUpdateOne {
	ruo.mutation.SetSlug(s)
	return ruo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableSlug(s *string) *RegionUpdateOne {
	if s != nil {
		ruo.SetSlug(*s)
	}
	return ruo
}

// ClearSlug clears the value of the "slug" field.
func (ruo *RegionUpdateOne) ClearSlug() *RegionUpdateOne {
	ruo.mutation.ClearSlug()
	return ruo
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (ruo *RegionUpdateOne) SetLocationID(id int) *RegionUpdateOne {
	ruo.mutation.SetLocationID(id)
	return ruo
}

// SetNillableLocationID sets the "location" edge to the Location entity by ID if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableLocationID(id *int) *RegionUpdateOne {
	if id != nil {
		ruo = ruo.SetLocationID(*id)
	}
	return ruo
}

// SetLocation sets the "location" edge to the Location entity.
func (ruo *RegionUpdateOne) SetLocation(l *Location) *RegionUpdateOne {
	return ruo.SetLocationID(l.ID)
}

// Mutation returns the RegionMutation object of the builder.
func (ruo *RegionUpdateOne) Mutation() *RegionMutation {
	return ruo.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (ruo *RegionUpdateOne) ClearLocation() *RegionUpdateOne {
	ruo.mutation.ClearLocation()
	return ruo
}

// Where appends a list predicates to the RegionUpdate builder.
func (ruo *RegionUpdateOne) Where(ps ...predicate.Region) *RegionUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RegionUpdateOne) Select(field string, fields ...string) *RegionUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Region entity.
func (ruo *RegionUpdateOne) Save(ctx context.Context) (*Region, error) {
	if err := ruo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RegionUpdateOne) SaveX(ctx context.Context) *Region {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RegionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RegionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RegionUpdateOne) defaults() error {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		if region.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized region.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := region.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ruo *RegionUpdateOne) sqlSave(ctx context.Context) (_node *Region, err error) {
	_spec := sqlgraph.NewUpdateSpec(region.Table, region.Columns, sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Region.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, region.FieldID)
		for _, f := range fields {
			if !region.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != region.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.CreatedBy(); ok {
		_spec.SetField(region.FieldCreatedBy, field.TypeString, value)
	}
	if ruo.mutation.CreatedByCleared() {
		_spec.ClearField(region.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(region.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedBy(); ok {
		_spec.SetField(region.FieldUpdatedBy, field.TypeString, value)
	}
	if ruo.mutation.UpdatedByCleared() {
		_spec.ClearField(region.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ruo.mutation.DisplayName(); ok {
		_spec.SetField(region.FieldDisplayName, field.TypeString, value)
	}
	if ruo.mutation.DisplayNameCleared() {
		_spec.ClearField(region.FieldDisplayName, field.TypeString)
	}
	if value, ok := ruo.mutation.Abbreviation(); ok {
		_spec.SetField(region.FieldAbbreviation, field.TypeString, value)
	}
	if ruo.mutation.AbbreviationCleared() {
		_spec.ClearField(region.FieldAbbreviation, field.TypeString)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(region.FieldDescription, field.TypeString, value)
	}
	if ruo.mutation.DescriptionCleared() {
		_spec.ClearField(region.FieldDescription, field.TypeString)
	}
	if value, ok := ruo.mutation.ExternalLink(); ok {
		_spec.SetField(region.FieldExternalLink, field.TypeString, value)
	}
	if ruo.mutation.ExternalLinkCleared() {
		_spec.ClearField(region.FieldExternalLink, field.TypeString)
	}
	if value, ok := ruo.mutation.Slug(); ok {
		_spec.SetField(region.FieldSlug, field.TypeString, value)
	}
	if ruo.mutation.SlugCleared() {
		_spec.ClearField(region.FieldSlug, field.TypeString)
	}
	if ruo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   region.LocationTable,
			Columns: []string{region.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   region.LocationTable,
			Columns: []string{region.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Region{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{region.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
