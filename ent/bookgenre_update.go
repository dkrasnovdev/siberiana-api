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
	"github.com/dkrasnovdev/siberiana-api/ent/bookgenre"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// BookGenreUpdate is the builder for updating BookGenre entities.
type BookGenreUpdate struct {
	config
	hooks    []Hook
	mutation *BookGenreMutation
}

// Where appends a list predicates to the BookGenreUpdate builder.
func (bgu *BookGenreUpdate) Where(ps ...predicate.BookGenre) *BookGenreUpdate {
	bgu.mutation.Where(ps...)
	return bgu
}

// SetCreatedBy sets the "created_by" field.
func (bgu *BookGenreUpdate) SetCreatedBy(s string) *BookGenreUpdate {
	bgu.mutation.SetCreatedBy(s)
	return bgu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (bgu *BookGenreUpdate) SetNillableCreatedBy(s *string) *BookGenreUpdate {
	if s != nil {
		bgu.SetCreatedBy(*s)
	}
	return bgu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (bgu *BookGenreUpdate) ClearCreatedBy() *BookGenreUpdate {
	bgu.mutation.ClearCreatedBy()
	return bgu
}

// SetUpdatedAt sets the "updated_at" field.
func (bgu *BookGenreUpdate) SetUpdatedAt(t time.Time) *BookGenreUpdate {
	bgu.mutation.SetUpdatedAt(t)
	return bgu
}

// SetUpdatedBy sets the "updated_by" field.
func (bgu *BookGenreUpdate) SetUpdatedBy(s string) *BookGenreUpdate {
	bgu.mutation.SetUpdatedBy(s)
	return bgu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (bgu *BookGenreUpdate) SetNillableUpdatedBy(s *string) *BookGenreUpdate {
	if s != nil {
		bgu.SetUpdatedBy(*s)
	}
	return bgu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (bgu *BookGenreUpdate) ClearUpdatedBy() *BookGenreUpdate {
	bgu.mutation.ClearUpdatedBy()
	return bgu
}

// SetDisplayName sets the "display_name" field.
func (bgu *BookGenreUpdate) SetDisplayName(s string) *BookGenreUpdate {
	bgu.mutation.SetDisplayName(s)
	return bgu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (bgu *BookGenreUpdate) SetNillableDisplayName(s *string) *BookGenreUpdate {
	if s != nil {
		bgu.SetDisplayName(*s)
	}
	return bgu
}

// ClearDisplayName clears the value of the "display_name" field.
func (bgu *BookGenreUpdate) ClearDisplayName() *BookGenreUpdate {
	bgu.mutation.ClearDisplayName()
	return bgu
}

// SetAbbreviation sets the "abbreviation" field.
func (bgu *BookGenreUpdate) SetAbbreviation(s string) *BookGenreUpdate {
	bgu.mutation.SetAbbreviation(s)
	return bgu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (bgu *BookGenreUpdate) SetNillableAbbreviation(s *string) *BookGenreUpdate {
	if s != nil {
		bgu.SetAbbreviation(*s)
	}
	return bgu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (bgu *BookGenreUpdate) ClearAbbreviation() *BookGenreUpdate {
	bgu.mutation.ClearAbbreviation()
	return bgu
}

// SetDescription sets the "description" field.
func (bgu *BookGenreUpdate) SetDescription(s string) *BookGenreUpdate {
	bgu.mutation.SetDescription(s)
	return bgu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bgu *BookGenreUpdate) SetNillableDescription(s *string) *BookGenreUpdate {
	if s != nil {
		bgu.SetDescription(*s)
	}
	return bgu
}

// ClearDescription clears the value of the "description" field.
func (bgu *BookGenreUpdate) ClearDescription() *BookGenreUpdate {
	bgu.mutation.ClearDescription()
	return bgu
}

// SetExternalLink sets the "external_link" field.
func (bgu *BookGenreUpdate) SetExternalLink(s string) *BookGenreUpdate {
	bgu.mutation.SetExternalLink(s)
	return bgu
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (bgu *BookGenreUpdate) SetNillableExternalLink(s *string) *BookGenreUpdate {
	if s != nil {
		bgu.SetExternalLink(*s)
	}
	return bgu
}

// ClearExternalLink clears the value of the "external_link" field.
func (bgu *BookGenreUpdate) ClearExternalLink() *BookGenreUpdate {
	bgu.mutation.ClearExternalLink()
	return bgu
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (bgu *BookGenreUpdate) AddBookIDs(ids ...int) *BookGenreUpdate {
	bgu.mutation.AddBookIDs(ids...)
	return bgu
}

// AddBooks adds the "books" edges to the Book entity.
func (bgu *BookGenreUpdate) AddBooks(b ...*Book) *BookGenreUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bgu.AddBookIDs(ids...)
}

// Mutation returns the BookGenreMutation object of the builder.
func (bgu *BookGenreUpdate) Mutation() *BookGenreMutation {
	return bgu.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (bgu *BookGenreUpdate) ClearBooks() *BookGenreUpdate {
	bgu.mutation.ClearBooks()
	return bgu
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (bgu *BookGenreUpdate) RemoveBookIDs(ids ...int) *BookGenreUpdate {
	bgu.mutation.RemoveBookIDs(ids...)
	return bgu
}

// RemoveBooks removes "books" edges to Book entities.
func (bgu *BookGenreUpdate) RemoveBooks(b ...*Book) *BookGenreUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bgu.RemoveBookIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bgu *BookGenreUpdate) Save(ctx context.Context) (int, error) {
	if err := bgu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, bgu.sqlSave, bgu.mutation, bgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bgu *BookGenreUpdate) SaveX(ctx context.Context) int {
	affected, err := bgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bgu *BookGenreUpdate) Exec(ctx context.Context) error {
	_, err := bgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bgu *BookGenreUpdate) ExecX(ctx context.Context) {
	if err := bgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bgu *BookGenreUpdate) defaults() error {
	if _, ok := bgu.mutation.UpdatedAt(); !ok {
		if bookgenre.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized bookgenre.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := bookgenre.UpdateDefaultUpdatedAt()
		bgu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (bgu *BookGenreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookgenre.Table, bookgenre.Columns, sqlgraph.NewFieldSpec(bookgenre.FieldID, field.TypeInt))
	if ps := bgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bgu.mutation.CreatedBy(); ok {
		_spec.SetField(bookgenre.FieldCreatedBy, field.TypeString, value)
	}
	if bgu.mutation.CreatedByCleared() {
		_spec.ClearField(bookgenre.FieldCreatedBy, field.TypeString)
	}
	if value, ok := bgu.mutation.UpdatedAt(); ok {
		_spec.SetField(bookgenre.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bgu.mutation.UpdatedBy(); ok {
		_spec.SetField(bookgenre.FieldUpdatedBy, field.TypeString, value)
	}
	if bgu.mutation.UpdatedByCleared() {
		_spec.ClearField(bookgenre.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := bgu.mutation.DisplayName(); ok {
		_spec.SetField(bookgenre.FieldDisplayName, field.TypeString, value)
	}
	if bgu.mutation.DisplayNameCleared() {
		_spec.ClearField(bookgenre.FieldDisplayName, field.TypeString)
	}
	if value, ok := bgu.mutation.Abbreviation(); ok {
		_spec.SetField(bookgenre.FieldAbbreviation, field.TypeString, value)
	}
	if bgu.mutation.AbbreviationCleared() {
		_spec.ClearField(bookgenre.FieldAbbreviation, field.TypeString)
	}
	if value, ok := bgu.mutation.Description(); ok {
		_spec.SetField(bookgenre.FieldDescription, field.TypeString, value)
	}
	if bgu.mutation.DescriptionCleared() {
		_spec.ClearField(bookgenre.FieldDescription, field.TypeString)
	}
	if value, ok := bgu.mutation.ExternalLink(); ok {
		_spec.SetField(bookgenre.FieldExternalLink, field.TypeString, value)
	}
	if bgu.mutation.ExternalLinkCleared() {
		_spec.ClearField(bookgenre.FieldExternalLink, field.TypeString)
	}
	if bgu.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookgenre.BooksTable,
			Columns: bookgenre.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bgu.mutation.RemovedBooksIDs(); len(nodes) > 0 && !bgu.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookgenre.BooksTable,
			Columns: bookgenre.BooksPrimaryKey,
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
	if nodes := bgu.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookgenre.BooksTable,
			Columns: bookgenre.BooksPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, bgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookgenre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bgu.mutation.done = true
	return n, nil
}

// BookGenreUpdateOne is the builder for updating a single BookGenre entity.
type BookGenreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookGenreMutation
}

// SetCreatedBy sets the "created_by" field.
func (bguo *BookGenreUpdateOne) SetCreatedBy(s string) *BookGenreUpdateOne {
	bguo.mutation.SetCreatedBy(s)
	return bguo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (bguo *BookGenreUpdateOne) SetNillableCreatedBy(s *string) *BookGenreUpdateOne {
	if s != nil {
		bguo.SetCreatedBy(*s)
	}
	return bguo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (bguo *BookGenreUpdateOne) ClearCreatedBy() *BookGenreUpdateOne {
	bguo.mutation.ClearCreatedBy()
	return bguo
}

// SetUpdatedAt sets the "updated_at" field.
func (bguo *BookGenreUpdateOne) SetUpdatedAt(t time.Time) *BookGenreUpdateOne {
	bguo.mutation.SetUpdatedAt(t)
	return bguo
}

// SetUpdatedBy sets the "updated_by" field.
func (bguo *BookGenreUpdateOne) SetUpdatedBy(s string) *BookGenreUpdateOne {
	bguo.mutation.SetUpdatedBy(s)
	return bguo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (bguo *BookGenreUpdateOne) SetNillableUpdatedBy(s *string) *BookGenreUpdateOne {
	if s != nil {
		bguo.SetUpdatedBy(*s)
	}
	return bguo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (bguo *BookGenreUpdateOne) ClearUpdatedBy() *BookGenreUpdateOne {
	bguo.mutation.ClearUpdatedBy()
	return bguo
}

// SetDisplayName sets the "display_name" field.
func (bguo *BookGenreUpdateOne) SetDisplayName(s string) *BookGenreUpdateOne {
	bguo.mutation.SetDisplayName(s)
	return bguo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (bguo *BookGenreUpdateOne) SetNillableDisplayName(s *string) *BookGenreUpdateOne {
	if s != nil {
		bguo.SetDisplayName(*s)
	}
	return bguo
}

// ClearDisplayName clears the value of the "display_name" field.
func (bguo *BookGenreUpdateOne) ClearDisplayName() *BookGenreUpdateOne {
	bguo.mutation.ClearDisplayName()
	return bguo
}

// SetAbbreviation sets the "abbreviation" field.
func (bguo *BookGenreUpdateOne) SetAbbreviation(s string) *BookGenreUpdateOne {
	bguo.mutation.SetAbbreviation(s)
	return bguo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (bguo *BookGenreUpdateOne) SetNillableAbbreviation(s *string) *BookGenreUpdateOne {
	if s != nil {
		bguo.SetAbbreviation(*s)
	}
	return bguo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (bguo *BookGenreUpdateOne) ClearAbbreviation() *BookGenreUpdateOne {
	bguo.mutation.ClearAbbreviation()
	return bguo
}

// SetDescription sets the "description" field.
func (bguo *BookGenreUpdateOne) SetDescription(s string) *BookGenreUpdateOne {
	bguo.mutation.SetDescription(s)
	return bguo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bguo *BookGenreUpdateOne) SetNillableDescription(s *string) *BookGenreUpdateOne {
	if s != nil {
		bguo.SetDescription(*s)
	}
	return bguo
}

// ClearDescription clears the value of the "description" field.
func (bguo *BookGenreUpdateOne) ClearDescription() *BookGenreUpdateOne {
	bguo.mutation.ClearDescription()
	return bguo
}

// SetExternalLink sets the "external_link" field.
func (bguo *BookGenreUpdateOne) SetExternalLink(s string) *BookGenreUpdateOne {
	bguo.mutation.SetExternalLink(s)
	return bguo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (bguo *BookGenreUpdateOne) SetNillableExternalLink(s *string) *BookGenreUpdateOne {
	if s != nil {
		bguo.SetExternalLink(*s)
	}
	return bguo
}

// ClearExternalLink clears the value of the "external_link" field.
func (bguo *BookGenreUpdateOne) ClearExternalLink() *BookGenreUpdateOne {
	bguo.mutation.ClearExternalLink()
	return bguo
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (bguo *BookGenreUpdateOne) AddBookIDs(ids ...int) *BookGenreUpdateOne {
	bguo.mutation.AddBookIDs(ids...)
	return bguo
}

// AddBooks adds the "books" edges to the Book entity.
func (bguo *BookGenreUpdateOne) AddBooks(b ...*Book) *BookGenreUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bguo.AddBookIDs(ids...)
}

// Mutation returns the BookGenreMutation object of the builder.
func (bguo *BookGenreUpdateOne) Mutation() *BookGenreMutation {
	return bguo.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (bguo *BookGenreUpdateOne) ClearBooks() *BookGenreUpdateOne {
	bguo.mutation.ClearBooks()
	return bguo
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (bguo *BookGenreUpdateOne) RemoveBookIDs(ids ...int) *BookGenreUpdateOne {
	bguo.mutation.RemoveBookIDs(ids...)
	return bguo
}

// RemoveBooks removes "books" edges to Book entities.
func (bguo *BookGenreUpdateOne) RemoveBooks(b ...*Book) *BookGenreUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bguo.RemoveBookIDs(ids...)
}

// Where appends a list predicates to the BookGenreUpdate builder.
func (bguo *BookGenreUpdateOne) Where(ps ...predicate.BookGenre) *BookGenreUpdateOne {
	bguo.mutation.Where(ps...)
	return bguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bguo *BookGenreUpdateOne) Select(field string, fields ...string) *BookGenreUpdateOne {
	bguo.fields = append([]string{field}, fields...)
	return bguo
}

// Save executes the query and returns the updated BookGenre entity.
func (bguo *BookGenreUpdateOne) Save(ctx context.Context) (*BookGenre, error) {
	if err := bguo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, bguo.sqlSave, bguo.mutation, bguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bguo *BookGenreUpdateOne) SaveX(ctx context.Context) *BookGenre {
	node, err := bguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bguo *BookGenreUpdateOne) Exec(ctx context.Context) error {
	_, err := bguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bguo *BookGenreUpdateOne) ExecX(ctx context.Context) {
	if err := bguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bguo *BookGenreUpdateOne) defaults() error {
	if _, ok := bguo.mutation.UpdatedAt(); !ok {
		if bookgenre.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized bookgenre.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := bookgenre.UpdateDefaultUpdatedAt()
		bguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (bguo *BookGenreUpdateOne) sqlSave(ctx context.Context) (_node *BookGenre, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookgenre.Table, bookgenre.Columns, sqlgraph.NewFieldSpec(bookgenre.FieldID, field.TypeInt))
	id, ok := bguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BookGenre.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookgenre.FieldID)
		for _, f := range fields {
			if !bookgenre.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookgenre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bguo.mutation.CreatedBy(); ok {
		_spec.SetField(bookgenre.FieldCreatedBy, field.TypeString, value)
	}
	if bguo.mutation.CreatedByCleared() {
		_spec.ClearField(bookgenre.FieldCreatedBy, field.TypeString)
	}
	if value, ok := bguo.mutation.UpdatedAt(); ok {
		_spec.SetField(bookgenre.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bguo.mutation.UpdatedBy(); ok {
		_spec.SetField(bookgenre.FieldUpdatedBy, field.TypeString, value)
	}
	if bguo.mutation.UpdatedByCleared() {
		_spec.ClearField(bookgenre.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := bguo.mutation.DisplayName(); ok {
		_spec.SetField(bookgenre.FieldDisplayName, field.TypeString, value)
	}
	if bguo.mutation.DisplayNameCleared() {
		_spec.ClearField(bookgenre.FieldDisplayName, field.TypeString)
	}
	if value, ok := bguo.mutation.Abbreviation(); ok {
		_spec.SetField(bookgenre.FieldAbbreviation, field.TypeString, value)
	}
	if bguo.mutation.AbbreviationCleared() {
		_spec.ClearField(bookgenre.FieldAbbreviation, field.TypeString)
	}
	if value, ok := bguo.mutation.Description(); ok {
		_spec.SetField(bookgenre.FieldDescription, field.TypeString, value)
	}
	if bguo.mutation.DescriptionCleared() {
		_spec.ClearField(bookgenre.FieldDescription, field.TypeString)
	}
	if value, ok := bguo.mutation.ExternalLink(); ok {
		_spec.SetField(bookgenre.FieldExternalLink, field.TypeString, value)
	}
	if bguo.mutation.ExternalLinkCleared() {
		_spec.ClearField(bookgenre.FieldExternalLink, field.TypeString)
	}
	if bguo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookgenre.BooksTable,
			Columns: bookgenre.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bguo.mutation.RemovedBooksIDs(); len(nodes) > 0 && !bguo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookgenre.BooksTable,
			Columns: bookgenre.BooksPrimaryKey,
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
	if nodes := bguo.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookgenre.BooksTable,
			Columns: bookgenre.BooksPrimaryKey,
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
	_node = &BookGenre{config: bguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookgenre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bguo.mutation.done = true
	return _node, nil
}
