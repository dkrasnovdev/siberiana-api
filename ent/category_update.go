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
	"github.com/dkrasnovdev/heritage-api/ent/category"
	"github.com/dkrasnovdev/heritage-api/ent/collection"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedBy sets the "created_by" field.
func (cu *CategoryUpdate) SetCreatedBy(s string) *CategoryUpdate {
	cu.mutation.SetCreatedBy(s)
	return cu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableCreatedBy(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetCreatedBy(*s)
	}
	return cu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (cu *CategoryUpdate) ClearCreatedBy() *CategoryUpdate {
	cu.mutation.ClearCreatedBy()
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CategoryUpdate) SetUpdatedAt(t time.Time) *CategoryUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetUpdatedBy sets the "updated_by" field.
func (cu *CategoryUpdate) SetUpdatedBy(s string) *CategoryUpdate {
	cu.mutation.SetUpdatedBy(s)
	return cu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableUpdatedBy(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetUpdatedBy(*s)
	}
	return cu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (cu *CategoryUpdate) ClearUpdatedBy() *CategoryUpdate {
	cu.mutation.ClearUpdatedBy()
	return cu
}

// SetDisplayName sets the "display_name" field.
func (cu *CategoryUpdate) SetDisplayName(s string) *CategoryUpdate {
	cu.mutation.SetDisplayName(s)
	return cu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDisplayName(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetDisplayName(*s)
	}
	return cu
}

// ClearDisplayName clears the value of the "display_name" field.
func (cu *CategoryUpdate) ClearDisplayName() *CategoryUpdate {
	cu.mutation.ClearDisplayName()
	return cu
}

// SetDescription sets the "description" field.
func (cu *CategoryUpdate) SetDescription(s string) *CategoryUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDescription(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CategoryUpdate) ClearDescription() *CategoryUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetExternalLinks sets the "external_links" field.
func (cu *CategoryUpdate) SetExternalLinks(s []string) *CategoryUpdate {
	cu.mutation.SetExternalLinks(s)
	return cu
}

// AppendExternalLinks appends s to the "external_links" field.
func (cu *CategoryUpdate) AppendExternalLinks(s []string) *CategoryUpdate {
	cu.mutation.AppendExternalLinks(s)
	return cu
}

// ClearExternalLinks clears the value of the "external_links" field.
func (cu *CategoryUpdate) ClearExternalLinks() *CategoryUpdate {
	cu.mutation.ClearExternalLinks()
	return cu
}

// AddCollectionIDs adds the "collections" edge to the Collection entity by IDs.
func (cu *CategoryUpdate) AddCollectionIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddCollectionIDs(ids...)
	return cu
}

// AddCollections adds the "collections" edges to the Collection entity.
func (cu *CategoryUpdate) AddCollections(c ...*Collection) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCollectionIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearCollections clears all "collections" edges to the Collection entity.
func (cu *CategoryUpdate) ClearCollections() *CategoryUpdate {
	cu.mutation.ClearCollections()
	return cu
}

// RemoveCollectionIDs removes the "collections" edge to Collection entities by IDs.
func (cu *CategoryUpdate) RemoveCollectionIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveCollectionIDs(ids...)
	return cu
}

// RemoveCollections removes "collections" edges to Collection entities.
func (cu *CategoryUpdate) RemoveCollections(c ...*Collection) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCollectionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CategoryUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if category.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized category.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := category.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedBy(); ok {
		_spec.SetField(category.FieldCreatedBy, field.TypeString, value)
	}
	if cu.mutation.CreatedByCleared() {
		_spec.ClearField(category.FieldCreatedBy, field.TypeString)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedBy(); ok {
		_spec.SetField(category.FieldUpdatedBy, field.TypeString, value)
	}
	if cu.mutation.UpdatedByCleared() {
		_spec.ClearField(category.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := cu.mutation.DisplayName(); ok {
		_spec.SetField(category.FieldDisplayName, field.TypeString, value)
	}
	if cu.mutation.DisplayNameCleared() {
		_spec.ClearField(category.FieldDisplayName, field.TypeString)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(category.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.ExternalLinks(); ok {
		_spec.SetField(category.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, category.FieldExternalLinks, value)
		})
	}
	if cu.mutation.ExternalLinksCleared() {
		_spec.ClearField(category.FieldExternalLinks, field.TypeJSON)
	}
	if cu.mutation.CollectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.CollectionsTable,
			Columns: []string{category.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCollectionsIDs(); len(nodes) > 0 && !cu.mutation.CollectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.CollectionsTable,
			Columns: []string{category.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CollectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.CollectionsTable,
			Columns: []string{category.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetCreatedBy sets the "created_by" field.
func (cuo *CategoryUpdateOne) SetCreatedBy(s string) *CategoryUpdateOne {
	cuo.mutation.SetCreatedBy(s)
	return cuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableCreatedBy(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetCreatedBy(*s)
	}
	return cuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (cuo *CategoryUpdateOne) ClearCreatedBy() *CategoryUpdateOne {
	cuo.mutation.ClearCreatedBy()
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CategoryUpdateOne) SetUpdatedAt(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetUpdatedBy sets the "updated_by" field.
func (cuo *CategoryUpdateOne) SetUpdatedBy(s string) *CategoryUpdateOne {
	cuo.mutation.SetUpdatedBy(s)
	return cuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableUpdatedBy(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetUpdatedBy(*s)
	}
	return cuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (cuo *CategoryUpdateOne) ClearUpdatedBy() *CategoryUpdateOne {
	cuo.mutation.ClearUpdatedBy()
	return cuo
}

// SetDisplayName sets the "display_name" field.
func (cuo *CategoryUpdateOne) SetDisplayName(s string) *CategoryUpdateOne {
	cuo.mutation.SetDisplayName(s)
	return cuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDisplayName(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetDisplayName(*s)
	}
	return cuo
}

// ClearDisplayName clears the value of the "display_name" field.
func (cuo *CategoryUpdateOne) ClearDisplayName() *CategoryUpdateOne {
	cuo.mutation.ClearDisplayName()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CategoryUpdateOne) SetDescription(s string) *CategoryUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDescription(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CategoryUpdateOne) ClearDescription() *CategoryUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetExternalLinks sets the "external_links" field.
func (cuo *CategoryUpdateOne) SetExternalLinks(s []string) *CategoryUpdateOne {
	cuo.mutation.SetExternalLinks(s)
	return cuo
}

// AppendExternalLinks appends s to the "external_links" field.
func (cuo *CategoryUpdateOne) AppendExternalLinks(s []string) *CategoryUpdateOne {
	cuo.mutation.AppendExternalLinks(s)
	return cuo
}

// ClearExternalLinks clears the value of the "external_links" field.
func (cuo *CategoryUpdateOne) ClearExternalLinks() *CategoryUpdateOne {
	cuo.mutation.ClearExternalLinks()
	return cuo
}

// AddCollectionIDs adds the "collections" edge to the Collection entity by IDs.
func (cuo *CategoryUpdateOne) AddCollectionIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddCollectionIDs(ids...)
	return cuo
}

// AddCollections adds the "collections" edges to the Collection entity.
func (cuo *CategoryUpdateOne) AddCollections(c ...*Collection) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCollectionIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearCollections clears all "collections" edges to the Collection entity.
func (cuo *CategoryUpdateOne) ClearCollections() *CategoryUpdateOne {
	cuo.mutation.ClearCollections()
	return cuo
}

// RemoveCollectionIDs removes the "collections" edge to Collection entities by IDs.
func (cuo *CategoryUpdateOne) RemoveCollectionIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveCollectionIDs(ids...)
	return cuo
}

// RemoveCollections removes "collections" edges to Collection entities.
func (cuo *CategoryUpdateOne) RemoveCollections(c ...*Collection) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCollectionIDs(ids...)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CategoryUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if category.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized category.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := category.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedBy(); ok {
		_spec.SetField(category.FieldCreatedBy, field.TypeString, value)
	}
	if cuo.mutation.CreatedByCleared() {
		_spec.ClearField(category.FieldCreatedBy, field.TypeString)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedBy(); ok {
		_spec.SetField(category.FieldUpdatedBy, field.TypeString, value)
	}
	if cuo.mutation.UpdatedByCleared() {
		_spec.ClearField(category.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := cuo.mutation.DisplayName(); ok {
		_spec.SetField(category.FieldDisplayName, field.TypeString, value)
	}
	if cuo.mutation.DisplayNameCleared() {
		_spec.ClearField(category.FieldDisplayName, field.TypeString)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(category.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.ExternalLinks(); ok {
		_spec.SetField(category.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, category.FieldExternalLinks, value)
		})
	}
	if cuo.mutation.ExternalLinksCleared() {
		_spec.ClearField(category.FieldExternalLinks, field.TypeJSON)
	}
	if cuo.mutation.CollectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.CollectionsTable,
			Columns: []string{category.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCollectionsIDs(); len(nodes) > 0 && !cuo.mutation.CollectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.CollectionsTable,
			Columns: []string{category.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CollectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.CollectionsTable,
			Columns: []string{category.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
