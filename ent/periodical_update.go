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
	"github.com/dkrasnovdev/siberiana-api/ent/book"
	"github.com/dkrasnovdev/siberiana-api/ent/periodical"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// PeriodicalUpdate is the builder for updating Periodical entities.
type PeriodicalUpdate struct {
	config
	hooks    []Hook
	mutation *PeriodicalMutation
}

// Where appends a list predicates to the PeriodicalUpdate builder.
func (pu *PeriodicalUpdate) Where(ps ...predicate.Periodical) *PeriodicalUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *PeriodicalUpdate) SetCreatedBy(s string) *PeriodicalUpdate {
	pu.mutation.SetCreatedBy(s)
	return pu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pu *PeriodicalUpdate) SetNillableCreatedBy(s *string) *PeriodicalUpdate {
	if s != nil {
		pu.SetCreatedBy(*s)
	}
	return pu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (pu *PeriodicalUpdate) ClearCreatedBy() *PeriodicalUpdate {
	pu.mutation.ClearCreatedBy()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PeriodicalUpdate) SetUpdatedAt(t time.Time) *PeriodicalUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *PeriodicalUpdate) SetUpdatedBy(s string) *PeriodicalUpdate {
	pu.mutation.SetUpdatedBy(s)
	return pu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pu *PeriodicalUpdate) SetNillableUpdatedBy(s *string) *PeriodicalUpdate {
	if s != nil {
		pu.SetUpdatedBy(*s)
	}
	return pu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pu *PeriodicalUpdate) ClearUpdatedBy() *PeriodicalUpdate {
	pu.mutation.ClearUpdatedBy()
	return pu
}

// SetDisplayName sets the "display_name" field.
func (pu *PeriodicalUpdate) SetDisplayName(s string) *PeriodicalUpdate {
	pu.mutation.SetDisplayName(s)
	return pu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pu *PeriodicalUpdate) SetNillableDisplayName(s *string) *PeriodicalUpdate {
	if s != nil {
		pu.SetDisplayName(*s)
	}
	return pu
}

// ClearDisplayName clears the value of the "display_name" field.
func (pu *PeriodicalUpdate) ClearDisplayName() *PeriodicalUpdate {
	pu.mutation.ClearDisplayName()
	return pu
}

// SetAbbreviation sets the "abbreviation" field.
func (pu *PeriodicalUpdate) SetAbbreviation(s string) *PeriodicalUpdate {
	pu.mutation.SetAbbreviation(s)
	return pu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pu *PeriodicalUpdate) SetNillableAbbreviation(s *string) *PeriodicalUpdate {
	if s != nil {
		pu.SetAbbreviation(*s)
	}
	return pu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (pu *PeriodicalUpdate) ClearAbbreviation() *PeriodicalUpdate {
	pu.mutation.ClearAbbreviation()
	return pu
}

// SetDescription sets the "description" field.
func (pu *PeriodicalUpdate) SetDescription(s string) *PeriodicalUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PeriodicalUpdate) SetNillableDescription(s *string) *PeriodicalUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PeriodicalUpdate) ClearDescription() *PeriodicalUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetExternalLink sets the "external_link" field.
func (pu *PeriodicalUpdate) SetExternalLink(s string) *PeriodicalUpdate {
	pu.mutation.SetExternalLink(s)
	return pu
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (pu *PeriodicalUpdate) SetNillableExternalLink(s *string) *PeriodicalUpdate {
	if s != nil {
		pu.SetExternalLink(*s)
	}
	return pu
}

// ClearExternalLink clears the value of the "external_link" field.
func (pu *PeriodicalUpdate) ClearExternalLink() *PeriodicalUpdate {
	pu.mutation.ClearExternalLink()
	return pu
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (pu *PeriodicalUpdate) AddBookIDs(ids ...int) *PeriodicalUpdate {
	pu.mutation.AddBookIDs(ids...)
	return pu
}

// AddBooks adds the "books" edges to the Book entity.
func (pu *PeriodicalUpdate) AddBooks(b ...*Book) *PeriodicalUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddBookIDs(ids...)
}

// Mutation returns the PeriodicalMutation object of the builder.
func (pu *PeriodicalUpdate) Mutation() *PeriodicalMutation {
	return pu.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (pu *PeriodicalUpdate) ClearBooks() *PeriodicalUpdate {
	pu.mutation.ClearBooks()
	return pu
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (pu *PeriodicalUpdate) RemoveBookIDs(ids ...int) *PeriodicalUpdate {
	pu.mutation.RemoveBookIDs(ids...)
	return pu
}

// RemoveBooks removes "books" edges to Book entities.
func (pu *PeriodicalUpdate) RemoveBooks(b ...*Book) *PeriodicalUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveBookIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PeriodicalUpdate) Save(ctx context.Context) (int, error) {
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PeriodicalUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PeriodicalUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PeriodicalUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PeriodicalUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if periodical.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized periodical.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := periodical.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pu *PeriodicalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(periodical.Table, periodical.Columns, sqlgraph.NewFieldSpec(periodical.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.SetField(periodical.FieldCreatedBy, field.TypeString, value)
	}
	if pu.mutation.CreatedByCleared() {
		_spec.ClearField(periodical.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(periodical.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.SetField(periodical.FieldUpdatedBy, field.TypeString, value)
	}
	if pu.mutation.UpdatedByCleared() {
		_spec.ClearField(periodical.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.DisplayName(); ok {
		_spec.SetField(periodical.FieldDisplayName, field.TypeString, value)
	}
	if pu.mutation.DisplayNameCleared() {
		_spec.ClearField(periodical.FieldDisplayName, field.TypeString)
	}
	if value, ok := pu.mutation.Abbreviation(); ok {
		_spec.SetField(periodical.FieldAbbreviation, field.TypeString, value)
	}
	if pu.mutation.AbbreviationCleared() {
		_spec.ClearField(periodical.FieldAbbreviation, field.TypeString)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(periodical.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(periodical.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.ExternalLink(); ok {
		_spec.SetField(periodical.FieldExternalLink, field.TypeString, value)
	}
	if pu.mutation.ExternalLinkCleared() {
		_spec.ClearField(periodical.FieldExternalLink, field.TypeString)
	}
	if pu.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   periodical.BooksTable,
			Columns: []string{periodical.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedBooksIDs(); len(nodes) > 0 && !pu.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   periodical.BooksTable,
			Columns: []string{periodical.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   periodical.BooksTable,
			Columns: []string{periodical.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{periodical.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PeriodicalUpdateOne is the builder for updating a single Periodical entity.
type PeriodicalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PeriodicalMutation
}

// SetCreatedBy sets the "created_by" field.
func (puo *PeriodicalUpdateOne) SetCreatedBy(s string) *PeriodicalUpdateOne {
	puo.mutation.SetCreatedBy(s)
	return puo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (puo *PeriodicalUpdateOne) SetNillableCreatedBy(s *string) *PeriodicalUpdateOne {
	if s != nil {
		puo.SetCreatedBy(*s)
	}
	return puo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (puo *PeriodicalUpdateOne) ClearCreatedBy() *PeriodicalUpdateOne {
	puo.mutation.ClearCreatedBy()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PeriodicalUpdateOne) SetUpdatedAt(t time.Time) *PeriodicalUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *PeriodicalUpdateOne) SetUpdatedBy(s string) *PeriodicalUpdateOne {
	puo.mutation.SetUpdatedBy(s)
	return puo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (puo *PeriodicalUpdateOne) SetNillableUpdatedBy(s *string) *PeriodicalUpdateOne {
	if s != nil {
		puo.SetUpdatedBy(*s)
	}
	return puo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (puo *PeriodicalUpdateOne) ClearUpdatedBy() *PeriodicalUpdateOne {
	puo.mutation.ClearUpdatedBy()
	return puo
}

// SetDisplayName sets the "display_name" field.
func (puo *PeriodicalUpdateOne) SetDisplayName(s string) *PeriodicalUpdateOne {
	puo.mutation.SetDisplayName(s)
	return puo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (puo *PeriodicalUpdateOne) SetNillableDisplayName(s *string) *PeriodicalUpdateOne {
	if s != nil {
		puo.SetDisplayName(*s)
	}
	return puo
}

// ClearDisplayName clears the value of the "display_name" field.
func (puo *PeriodicalUpdateOne) ClearDisplayName() *PeriodicalUpdateOne {
	puo.mutation.ClearDisplayName()
	return puo
}

// SetAbbreviation sets the "abbreviation" field.
func (puo *PeriodicalUpdateOne) SetAbbreviation(s string) *PeriodicalUpdateOne {
	puo.mutation.SetAbbreviation(s)
	return puo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (puo *PeriodicalUpdateOne) SetNillableAbbreviation(s *string) *PeriodicalUpdateOne {
	if s != nil {
		puo.SetAbbreviation(*s)
	}
	return puo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (puo *PeriodicalUpdateOne) ClearAbbreviation() *PeriodicalUpdateOne {
	puo.mutation.ClearAbbreviation()
	return puo
}

// SetDescription sets the "description" field.
func (puo *PeriodicalUpdateOne) SetDescription(s string) *PeriodicalUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PeriodicalUpdateOne) SetNillableDescription(s *string) *PeriodicalUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PeriodicalUpdateOne) ClearDescription() *PeriodicalUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetExternalLink sets the "external_link" field.
func (puo *PeriodicalUpdateOne) SetExternalLink(s string) *PeriodicalUpdateOne {
	puo.mutation.SetExternalLink(s)
	return puo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (puo *PeriodicalUpdateOne) SetNillableExternalLink(s *string) *PeriodicalUpdateOne {
	if s != nil {
		puo.SetExternalLink(*s)
	}
	return puo
}

// ClearExternalLink clears the value of the "external_link" field.
func (puo *PeriodicalUpdateOne) ClearExternalLink() *PeriodicalUpdateOne {
	puo.mutation.ClearExternalLink()
	return puo
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (puo *PeriodicalUpdateOne) AddBookIDs(ids ...int) *PeriodicalUpdateOne {
	puo.mutation.AddBookIDs(ids...)
	return puo
}

// AddBooks adds the "books" edges to the Book entity.
func (puo *PeriodicalUpdateOne) AddBooks(b ...*Book) *PeriodicalUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddBookIDs(ids...)
}

// Mutation returns the PeriodicalMutation object of the builder.
func (puo *PeriodicalUpdateOne) Mutation() *PeriodicalMutation {
	return puo.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (puo *PeriodicalUpdateOne) ClearBooks() *PeriodicalUpdateOne {
	puo.mutation.ClearBooks()
	return puo
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (puo *PeriodicalUpdateOne) RemoveBookIDs(ids ...int) *PeriodicalUpdateOne {
	puo.mutation.RemoveBookIDs(ids...)
	return puo
}

// RemoveBooks removes "books" edges to Book entities.
func (puo *PeriodicalUpdateOne) RemoveBooks(b ...*Book) *PeriodicalUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveBookIDs(ids...)
}

// Where appends a list predicates to the PeriodicalUpdate builder.
func (puo *PeriodicalUpdateOne) Where(ps ...predicate.Periodical) *PeriodicalUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PeriodicalUpdateOne) Select(field string, fields ...string) *PeriodicalUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Periodical entity.
func (puo *PeriodicalUpdateOne) Save(ctx context.Context) (*Periodical, error) {
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PeriodicalUpdateOne) SaveX(ctx context.Context) *Periodical {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PeriodicalUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PeriodicalUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PeriodicalUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if periodical.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized periodical.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := periodical.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (puo *PeriodicalUpdateOne) sqlSave(ctx context.Context) (_node *Periodical, err error) {
	_spec := sqlgraph.NewUpdateSpec(periodical.Table, periodical.Columns, sqlgraph.NewFieldSpec(periodical.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Periodical.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, periodical.FieldID)
		for _, f := range fields {
			if !periodical.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != periodical.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedBy(); ok {
		_spec.SetField(periodical.FieldCreatedBy, field.TypeString, value)
	}
	if puo.mutation.CreatedByCleared() {
		_spec.ClearField(periodical.FieldCreatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(periodical.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.SetField(periodical.FieldUpdatedBy, field.TypeString, value)
	}
	if puo.mutation.UpdatedByCleared() {
		_spec.ClearField(periodical.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.DisplayName(); ok {
		_spec.SetField(periodical.FieldDisplayName, field.TypeString, value)
	}
	if puo.mutation.DisplayNameCleared() {
		_spec.ClearField(periodical.FieldDisplayName, field.TypeString)
	}
	if value, ok := puo.mutation.Abbreviation(); ok {
		_spec.SetField(periodical.FieldAbbreviation, field.TypeString, value)
	}
	if puo.mutation.AbbreviationCleared() {
		_spec.ClearField(periodical.FieldAbbreviation, field.TypeString)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(periodical.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(periodical.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.ExternalLink(); ok {
		_spec.SetField(periodical.FieldExternalLink, field.TypeString, value)
	}
	if puo.mutation.ExternalLinkCleared() {
		_spec.ClearField(periodical.FieldExternalLink, field.TypeString)
	}
	if puo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   periodical.BooksTable,
			Columns: []string{periodical.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedBooksIDs(); len(nodes) > 0 && !puo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   periodical.BooksTable,
			Columns: []string{periodical.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   periodical.BooksTable,
			Columns: []string{periodical.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Periodical{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{periodical.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
