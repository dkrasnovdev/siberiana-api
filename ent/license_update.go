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
	"github.com/dkrasnovdev/heritage-api/ent/artifact"
	"github.com/dkrasnovdev/heritage-api/ent/book"
	"github.com/dkrasnovdev/heritage-api/ent/license"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
	"github.com/dkrasnovdev/heritage-api/ent/protectedareapicture"
)

// LicenseUpdate is the builder for updating License entities.
type LicenseUpdate struct {
	config
	hooks    []Hook
	mutation *LicenseMutation
}

// Where appends a list predicates to the LicenseUpdate builder.
func (lu *LicenseUpdate) Where(ps ...predicate.License) *LicenseUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetCreatedBy sets the "created_by" field.
func (lu *LicenseUpdate) SetCreatedBy(s string) *LicenseUpdate {
	lu.mutation.SetCreatedBy(s)
	return lu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableCreatedBy(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetCreatedBy(*s)
	}
	return lu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (lu *LicenseUpdate) ClearCreatedBy() *LicenseUpdate {
	lu.mutation.ClearCreatedBy()
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LicenseUpdate) SetUpdatedAt(t time.Time) *LicenseUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetUpdatedBy sets the "updated_by" field.
func (lu *LicenseUpdate) SetUpdatedBy(s string) *LicenseUpdate {
	lu.mutation.SetUpdatedBy(s)
	return lu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableUpdatedBy(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetUpdatedBy(*s)
	}
	return lu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (lu *LicenseUpdate) ClearUpdatedBy() *LicenseUpdate {
	lu.mutation.ClearUpdatedBy()
	return lu
}

// SetDisplayName sets the "display_name" field.
func (lu *LicenseUpdate) SetDisplayName(s string) *LicenseUpdate {
	lu.mutation.SetDisplayName(s)
	return lu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableDisplayName(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetDisplayName(*s)
	}
	return lu
}

// ClearDisplayName clears the value of the "display_name" field.
func (lu *LicenseUpdate) ClearDisplayName() *LicenseUpdate {
	lu.mutation.ClearDisplayName()
	return lu
}

// SetAbbreviation sets the "abbreviation" field.
func (lu *LicenseUpdate) SetAbbreviation(s string) *LicenseUpdate {
	lu.mutation.SetAbbreviation(s)
	return lu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableAbbreviation(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetAbbreviation(*s)
	}
	return lu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (lu *LicenseUpdate) ClearAbbreviation() *LicenseUpdate {
	lu.mutation.ClearAbbreviation()
	return lu
}

// SetDescription sets the "description" field.
func (lu *LicenseUpdate) SetDescription(s string) *LicenseUpdate {
	lu.mutation.SetDescription(s)
	return lu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableDescription(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetDescription(*s)
	}
	return lu
}

// ClearDescription clears the value of the "description" field.
func (lu *LicenseUpdate) ClearDescription() *LicenseUpdate {
	lu.mutation.ClearDescription()
	return lu
}

// SetExternalLink sets the "external_link" field.
func (lu *LicenseUpdate) SetExternalLink(s string) *LicenseUpdate {
	lu.mutation.SetExternalLink(s)
	return lu
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableExternalLink(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetExternalLink(*s)
	}
	return lu
}

// ClearExternalLink clears the value of the "external_link" field.
func (lu *LicenseUpdate) ClearExternalLink() *LicenseUpdate {
	lu.mutation.ClearExternalLink()
	return lu
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (lu *LicenseUpdate) AddArtifactIDs(ids ...int) *LicenseUpdate {
	lu.mutation.AddArtifactIDs(ids...)
	return lu
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (lu *LicenseUpdate) AddArtifacts(a ...*Artifact) *LicenseUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.AddArtifactIDs(ids...)
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (lu *LicenseUpdate) AddBookIDs(ids ...int) *LicenseUpdate {
	lu.mutation.AddBookIDs(ids...)
	return lu
}

// AddBooks adds the "books" edges to the Book entity.
func (lu *LicenseUpdate) AddBooks(b ...*Book) *LicenseUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return lu.AddBookIDs(ids...)
}

// AddProtectedAreaPictureIDs adds the "protected_area_pictures" edge to the ProtectedAreaPicture entity by IDs.
func (lu *LicenseUpdate) AddProtectedAreaPictureIDs(ids ...int) *LicenseUpdate {
	lu.mutation.AddProtectedAreaPictureIDs(ids...)
	return lu
}

// AddProtectedAreaPictures adds the "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (lu *LicenseUpdate) AddProtectedAreaPictures(p ...*ProtectedAreaPicture) *LicenseUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lu.AddProtectedAreaPictureIDs(ids...)
}

// Mutation returns the LicenseMutation object of the builder.
func (lu *LicenseUpdate) Mutation() *LicenseMutation {
	return lu.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (lu *LicenseUpdate) ClearArtifacts() *LicenseUpdate {
	lu.mutation.ClearArtifacts()
	return lu
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (lu *LicenseUpdate) RemoveArtifactIDs(ids ...int) *LicenseUpdate {
	lu.mutation.RemoveArtifactIDs(ids...)
	return lu
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (lu *LicenseUpdate) RemoveArtifacts(a ...*Artifact) *LicenseUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.RemoveArtifactIDs(ids...)
}

// ClearBooks clears all "books" edges to the Book entity.
func (lu *LicenseUpdate) ClearBooks() *LicenseUpdate {
	lu.mutation.ClearBooks()
	return lu
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (lu *LicenseUpdate) RemoveBookIDs(ids ...int) *LicenseUpdate {
	lu.mutation.RemoveBookIDs(ids...)
	return lu
}

// RemoveBooks removes "books" edges to Book entities.
func (lu *LicenseUpdate) RemoveBooks(b ...*Book) *LicenseUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return lu.RemoveBookIDs(ids...)
}

// ClearProtectedAreaPictures clears all "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (lu *LicenseUpdate) ClearProtectedAreaPictures() *LicenseUpdate {
	lu.mutation.ClearProtectedAreaPictures()
	return lu
}

// RemoveProtectedAreaPictureIDs removes the "protected_area_pictures" edge to ProtectedAreaPicture entities by IDs.
func (lu *LicenseUpdate) RemoveProtectedAreaPictureIDs(ids ...int) *LicenseUpdate {
	lu.mutation.RemoveProtectedAreaPictureIDs(ids...)
	return lu
}

// RemoveProtectedAreaPictures removes "protected_area_pictures" edges to ProtectedAreaPicture entities.
func (lu *LicenseUpdate) RemoveProtectedAreaPictures(p ...*ProtectedAreaPicture) *LicenseUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lu.RemoveProtectedAreaPictureIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LicenseUpdate) Save(ctx context.Context) (int, error) {
	if err := lu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LicenseUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LicenseUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LicenseUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LicenseUpdate) defaults() error {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		if license.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized license.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := license.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (lu *LicenseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.CreatedBy(); ok {
		_spec.SetField(license.FieldCreatedBy, field.TypeString, value)
	}
	if lu.mutation.CreatedByCleared() {
		_spec.ClearField(license.FieldCreatedBy, field.TypeString)
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(license.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.UpdatedBy(); ok {
		_spec.SetField(license.FieldUpdatedBy, field.TypeString, value)
	}
	if lu.mutation.UpdatedByCleared() {
		_spec.ClearField(license.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := lu.mutation.DisplayName(); ok {
		_spec.SetField(license.FieldDisplayName, field.TypeString, value)
	}
	if lu.mutation.DisplayNameCleared() {
		_spec.ClearField(license.FieldDisplayName, field.TypeString)
	}
	if value, ok := lu.mutation.Abbreviation(); ok {
		_spec.SetField(license.FieldAbbreviation, field.TypeString, value)
	}
	if lu.mutation.AbbreviationCleared() {
		_spec.ClearField(license.FieldAbbreviation, field.TypeString)
	}
	if value, ok := lu.mutation.Description(); ok {
		_spec.SetField(license.FieldDescription, field.TypeString, value)
	}
	if lu.mutation.DescriptionCleared() {
		_spec.ClearField(license.FieldDescription, field.TypeString)
	}
	if value, ok := lu.mutation.ExternalLink(); ok {
		_spec.SetField(license.FieldExternalLink, field.TypeString, value)
	}
	if lu.mutation.ExternalLinkCleared() {
		_spec.ClearField(license.FieldExternalLink, field.TypeString)
	}
	if lu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ArtifactsTable,
			Columns: []string{license.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !lu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ArtifactsTable,
			Columns: []string{license.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ArtifactsTable,
			Columns: []string{license.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.BooksTable,
			Columns: []string{license.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedBooksIDs(); len(nodes) > 0 && !lu.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.BooksTable,
			Columns: []string{license.BooksColumn},
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
	if nodes := lu.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.BooksTable,
			Columns: []string{license.BooksColumn},
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
	if lu.mutation.ProtectedAreaPicturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ProtectedAreaPicturesTable,
			Columns: []string{license.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedProtectedAreaPicturesIDs(); len(nodes) > 0 && !lu.mutation.ProtectedAreaPicturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ProtectedAreaPicturesTable,
			Columns: []string{license.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ProtectedAreaPicturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ProtectedAreaPicturesTable,
			Columns: []string{license.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{license.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LicenseUpdateOne is the builder for updating a single License entity.
type LicenseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LicenseMutation
}

// SetCreatedBy sets the "created_by" field.
func (luo *LicenseUpdateOne) SetCreatedBy(s string) *LicenseUpdateOne {
	luo.mutation.SetCreatedBy(s)
	return luo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableCreatedBy(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetCreatedBy(*s)
	}
	return luo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (luo *LicenseUpdateOne) ClearCreatedBy() *LicenseUpdateOne {
	luo.mutation.ClearCreatedBy()
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LicenseUpdateOne) SetUpdatedAt(t time.Time) *LicenseUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetUpdatedBy sets the "updated_by" field.
func (luo *LicenseUpdateOne) SetUpdatedBy(s string) *LicenseUpdateOne {
	luo.mutation.SetUpdatedBy(s)
	return luo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableUpdatedBy(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetUpdatedBy(*s)
	}
	return luo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (luo *LicenseUpdateOne) ClearUpdatedBy() *LicenseUpdateOne {
	luo.mutation.ClearUpdatedBy()
	return luo
}

// SetDisplayName sets the "display_name" field.
func (luo *LicenseUpdateOne) SetDisplayName(s string) *LicenseUpdateOne {
	luo.mutation.SetDisplayName(s)
	return luo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableDisplayName(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetDisplayName(*s)
	}
	return luo
}

// ClearDisplayName clears the value of the "display_name" field.
func (luo *LicenseUpdateOne) ClearDisplayName() *LicenseUpdateOne {
	luo.mutation.ClearDisplayName()
	return luo
}

// SetAbbreviation sets the "abbreviation" field.
func (luo *LicenseUpdateOne) SetAbbreviation(s string) *LicenseUpdateOne {
	luo.mutation.SetAbbreviation(s)
	return luo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableAbbreviation(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetAbbreviation(*s)
	}
	return luo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (luo *LicenseUpdateOne) ClearAbbreviation() *LicenseUpdateOne {
	luo.mutation.ClearAbbreviation()
	return luo
}

// SetDescription sets the "description" field.
func (luo *LicenseUpdateOne) SetDescription(s string) *LicenseUpdateOne {
	luo.mutation.SetDescription(s)
	return luo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableDescription(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetDescription(*s)
	}
	return luo
}

// ClearDescription clears the value of the "description" field.
func (luo *LicenseUpdateOne) ClearDescription() *LicenseUpdateOne {
	luo.mutation.ClearDescription()
	return luo
}

// SetExternalLink sets the "external_link" field.
func (luo *LicenseUpdateOne) SetExternalLink(s string) *LicenseUpdateOne {
	luo.mutation.SetExternalLink(s)
	return luo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableExternalLink(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetExternalLink(*s)
	}
	return luo
}

// ClearExternalLink clears the value of the "external_link" field.
func (luo *LicenseUpdateOne) ClearExternalLink() *LicenseUpdateOne {
	luo.mutation.ClearExternalLink()
	return luo
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (luo *LicenseUpdateOne) AddArtifactIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.AddArtifactIDs(ids...)
	return luo
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (luo *LicenseUpdateOne) AddArtifacts(a ...*Artifact) *LicenseUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.AddArtifactIDs(ids...)
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (luo *LicenseUpdateOne) AddBookIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.AddBookIDs(ids...)
	return luo
}

// AddBooks adds the "books" edges to the Book entity.
func (luo *LicenseUpdateOne) AddBooks(b ...*Book) *LicenseUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return luo.AddBookIDs(ids...)
}

// AddProtectedAreaPictureIDs adds the "protected_area_pictures" edge to the ProtectedAreaPicture entity by IDs.
func (luo *LicenseUpdateOne) AddProtectedAreaPictureIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.AddProtectedAreaPictureIDs(ids...)
	return luo
}

// AddProtectedAreaPictures adds the "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (luo *LicenseUpdateOne) AddProtectedAreaPictures(p ...*ProtectedAreaPicture) *LicenseUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return luo.AddProtectedAreaPictureIDs(ids...)
}

// Mutation returns the LicenseMutation object of the builder.
func (luo *LicenseUpdateOne) Mutation() *LicenseMutation {
	return luo.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (luo *LicenseUpdateOne) ClearArtifacts() *LicenseUpdateOne {
	luo.mutation.ClearArtifacts()
	return luo
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (luo *LicenseUpdateOne) RemoveArtifactIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.RemoveArtifactIDs(ids...)
	return luo
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (luo *LicenseUpdateOne) RemoveArtifacts(a ...*Artifact) *LicenseUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.RemoveArtifactIDs(ids...)
}

// ClearBooks clears all "books" edges to the Book entity.
func (luo *LicenseUpdateOne) ClearBooks() *LicenseUpdateOne {
	luo.mutation.ClearBooks()
	return luo
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (luo *LicenseUpdateOne) RemoveBookIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.RemoveBookIDs(ids...)
	return luo
}

// RemoveBooks removes "books" edges to Book entities.
func (luo *LicenseUpdateOne) RemoveBooks(b ...*Book) *LicenseUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return luo.RemoveBookIDs(ids...)
}

// ClearProtectedAreaPictures clears all "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (luo *LicenseUpdateOne) ClearProtectedAreaPictures() *LicenseUpdateOne {
	luo.mutation.ClearProtectedAreaPictures()
	return luo
}

// RemoveProtectedAreaPictureIDs removes the "protected_area_pictures" edge to ProtectedAreaPicture entities by IDs.
func (luo *LicenseUpdateOne) RemoveProtectedAreaPictureIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.RemoveProtectedAreaPictureIDs(ids...)
	return luo
}

// RemoveProtectedAreaPictures removes "protected_area_pictures" edges to ProtectedAreaPicture entities.
func (luo *LicenseUpdateOne) RemoveProtectedAreaPictures(p ...*ProtectedAreaPicture) *LicenseUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return luo.RemoveProtectedAreaPictureIDs(ids...)
}

// Where appends a list predicates to the LicenseUpdate builder.
func (luo *LicenseUpdateOne) Where(ps ...predicate.License) *LicenseUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LicenseUpdateOne) Select(field string, fields ...string) *LicenseUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated License entity.
func (luo *LicenseUpdateOne) Save(ctx context.Context) (*License, error) {
	if err := luo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LicenseUpdateOne) SaveX(ctx context.Context) *License {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LicenseUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LicenseUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LicenseUpdateOne) defaults() error {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		if license.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized license.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := license.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (luo *LicenseUpdateOne) sqlSave(ctx context.Context) (_node *License, err error) {
	_spec := sqlgraph.NewUpdateSpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "License.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, license.FieldID)
		for _, f := range fields {
			if !license.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != license.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.CreatedBy(); ok {
		_spec.SetField(license.FieldCreatedBy, field.TypeString, value)
	}
	if luo.mutation.CreatedByCleared() {
		_spec.ClearField(license.FieldCreatedBy, field.TypeString)
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(license.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.UpdatedBy(); ok {
		_spec.SetField(license.FieldUpdatedBy, field.TypeString, value)
	}
	if luo.mutation.UpdatedByCleared() {
		_spec.ClearField(license.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := luo.mutation.DisplayName(); ok {
		_spec.SetField(license.FieldDisplayName, field.TypeString, value)
	}
	if luo.mutation.DisplayNameCleared() {
		_spec.ClearField(license.FieldDisplayName, field.TypeString)
	}
	if value, ok := luo.mutation.Abbreviation(); ok {
		_spec.SetField(license.FieldAbbreviation, field.TypeString, value)
	}
	if luo.mutation.AbbreviationCleared() {
		_spec.ClearField(license.FieldAbbreviation, field.TypeString)
	}
	if value, ok := luo.mutation.Description(); ok {
		_spec.SetField(license.FieldDescription, field.TypeString, value)
	}
	if luo.mutation.DescriptionCleared() {
		_spec.ClearField(license.FieldDescription, field.TypeString)
	}
	if value, ok := luo.mutation.ExternalLink(); ok {
		_spec.SetField(license.FieldExternalLink, field.TypeString, value)
	}
	if luo.mutation.ExternalLinkCleared() {
		_spec.ClearField(license.FieldExternalLink, field.TypeString)
	}
	if luo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ArtifactsTable,
			Columns: []string{license.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !luo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ArtifactsTable,
			Columns: []string{license.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ArtifactsTable,
			Columns: []string{license.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.BooksTable,
			Columns: []string{license.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedBooksIDs(); len(nodes) > 0 && !luo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.BooksTable,
			Columns: []string{license.BooksColumn},
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
	if nodes := luo.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.BooksTable,
			Columns: []string{license.BooksColumn},
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
	if luo.mutation.ProtectedAreaPicturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ProtectedAreaPicturesTable,
			Columns: []string{license.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedProtectedAreaPicturesIDs(); len(nodes) > 0 && !luo.mutation.ProtectedAreaPicturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ProtectedAreaPicturesTable,
			Columns: []string{license.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ProtectedAreaPicturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ProtectedAreaPicturesTable,
			Columns: []string{license.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &License{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{license.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
