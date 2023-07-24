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
	"github.com/dkrasnovdev/heritage-api/ent/artgenre"
	"github.com/dkrasnovdev/heritage-api/ent/artstyle"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// ArtUpdate is the builder for updating Art entities.
type ArtUpdate struct {
	config
	hooks    []Hook
	mutation *ArtMutation
}

// Where appends a list predicates to the ArtUpdate builder.
func (au *ArtUpdate) Where(ps ...predicate.Art) *ArtUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedBy sets the "created_by" field.
func (au *ArtUpdate) SetCreatedBy(s string) *ArtUpdate {
	au.mutation.SetCreatedBy(s)
	return au
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (au *ArtUpdate) SetNillableCreatedBy(s *string) *ArtUpdate {
	if s != nil {
		au.SetCreatedBy(*s)
	}
	return au
}

// ClearCreatedBy clears the value of the "created_by" field.
func (au *ArtUpdate) ClearCreatedBy() *ArtUpdate {
	au.mutation.ClearCreatedBy()
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ArtUpdate) SetUpdatedAt(t time.Time) *ArtUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetUpdatedBy sets the "updated_by" field.
func (au *ArtUpdate) SetUpdatedBy(s string) *ArtUpdate {
	au.mutation.SetUpdatedBy(s)
	return au
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (au *ArtUpdate) SetNillableUpdatedBy(s *string) *ArtUpdate {
	if s != nil {
		au.SetUpdatedBy(*s)
	}
	return au
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (au *ArtUpdate) ClearUpdatedBy() *ArtUpdate {
	au.mutation.ClearUpdatedBy()
	return au
}

// SetDisplayName sets the "display_name" field.
func (au *ArtUpdate) SetDisplayName(s string) *ArtUpdate {
	au.mutation.SetDisplayName(s)
	return au
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (au *ArtUpdate) SetNillableDisplayName(s *string) *ArtUpdate {
	if s != nil {
		au.SetDisplayName(*s)
	}
	return au
}

// ClearDisplayName clears the value of the "display_name" field.
func (au *ArtUpdate) ClearDisplayName() *ArtUpdate {
	au.mutation.ClearDisplayName()
	return au
}

// SetAbbreviation sets the "abbreviation" field.
func (au *ArtUpdate) SetAbbreviation(s string) *ArtUpdate {
	au.mutation.SetAbbreviation(s)
	return au
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (au *ArtUpdate) SetNillableAbbreviation(s *string) *ArtUpdate {
	if s != nil {
		au.SetAbbreviation(*s)
	}
	return au
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (au *ArtUpdate) ClearAbbreviation() *ArtUpdate {
	au.mutation.ClearAbbreviation()
	return au
}

// SetDescription sets the "description" field.
func (au *ArtUpdate) SetDescription(s string) *ArtUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (au *ArtUpdate) SetNillableDescription(s *string) *ArtUpdate {
	if s != nil {
		au.SetDescription(*s)
	}
	return au
}

// ClearDescription clears the value of the "description" field.
func (au *ArtUpdate) ClearDescription() *ArtUpdate {
	au.mutation.ClearDescription()
	return au
}

// SetExternalLinks sets the "external_links" field.
func (au *ArtUpdate) SetExternalLinks(s []string) *ArtUpdate {
	au.mutation.SetExternalLinks(s)
	return au
}

// AppendExternalLinks appends s to the "external_links" field.
func (au *ArtUpdate) AppendExternalLinks(s []string) *ArtUpdate {
	au.mutation.AppendExternalLinks(s)
	return au
}

// ClearExternalLinks clears the value of the "external_links" field.
func (au *ArtUpdate) ClearExternalLinks() *ArtUpdate {
	au.mutation.ClearExternalLinks()
	return au
}

// SetPrimaryImageURL sets the "primary_image_url" field.
func (au *ArtUpdate) SetPrimaryImageURL(s string) *ArtUpdate {
	au.mutation.SetPrimaryImageURL(s)
	return au
}

// SetNillablePrimaryImageURL sets the "primary_image_url" field if the given value is not nil.
func (au *ArtUpdate) SetNillablePrimaryImageURL(s *string) *ArtUpdate {
	if s != nil {
		au.SetPrimaryImageURL(*s)
	}
	return au
}

// ClearPrimaryImageURL clears the value of the "primary_image_url" field.
func (au *ArtUpdate) ClearPrimaryImageURL() *ArtUpdate {
	au.mutation.ClearPrimaryImageURL()
	return au
}

// SetAdditionalImagesUrls sets the "additional_images_urls" field.
func (au *ArtUpdate) SetAdditionalImagesUrls(s []string) *ArtUpdate {
	au.mutation.SetAdditionalImagesUrls(s)
	return au
}

// AppendAdditionalImagesUrls appends s to the "additional_images_urls" field.
func (au *ArtUpdate) AppendAdditionalImagesUrls(s []string) *ArtUpdate {
	au.mutation.AppendAdditionalImagesUrls(s)
	return au
}

// ClearAdditionalImagesUrls clears the value of the "additional_images_urls" field.
func (au *ArtUpdate) ClearAdditionalImagesUrls() *ArtUpdate {
	au.mutation.ClearAdditionalImagesUrls()
	return au
}

// AddArtGenreIDs adds the "art_genre" edge to the ArtGenre entity by IDs.
func (au *ArtUpdate) AddArtGenreIDs(ids ...int) *ArtUpdate {
	au.mutation.AddArtGenreIDs(ids...)
	return au
}

// AddArtGenre adds the "art_genre" edges to the ArtGenre entity.
func (au *ArtUpdate) AddArtGenre(a ...*ArtGenre) *ArtUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddArtGenreIDs(ids...)
}

// AddArtStyleIDs adds the "art_style" edge to the ArtStyle entity by IDs.
func (au *ArtUpdate) AddArtStyleIDs(ids ...int) *ArtUpdate {
	au.mutation.AddArtStyleIDs(ids...)
	return au
}

// AddArtStyle adds the "art_style" edges to the ArtStyle entity.
func (au *ArtUpdate) AddArtStyle(a ...*ArtStyle) *ArtUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddArtStyleIDs(ids...)
}

// Mutation returns the ArtMutation object of the builder.
func (au *ArtUpdate) Mutation() *ArtMutation {
	return au.mutation
}

// ClearArtGenre clears all "art_genre" edges to the ArtGenre entity.
func (au *ArtUpdate) ClearArtGenre() *ArtUpdate {
	au.mutation.ClearArtGenre()
	return au
}

// RemoveArtGenreIDs removes the "art_genre" edge to ArtGenre entities by IDs.
func (au *ArtUpdate) RemoveArtGenreIDs(ids ...int) *ArtUpdate {
	au.mutation.RemoveArtGenreIDs(ids...)
	return au
}

// RemoveArtGenre removes "art_genre" edges to ArtGenre entities.
func (au *ArtUpdate) RemoveArtGenre(a ...*ArtGenre) *ArtUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveArtGenreIDs(ids...)
}

// ClearArtStyle clears all "art_style" edges to the ArtStyle entity.
func (au *ArtUpdate) ClearArtStyle() *ArtUpdate {
	au.mutation.ClearArtStyle()
	return au
}

// RemoveArtStyleIDs removes the "art_style" edge to ArtStyle entities by IDs.
func (au *ArtUpdate) RemoveArtStyleIDs(ids ...int) *ArtUpdate {
	au.mutation.RemoveArtStyleIDs(ids...)
	return au
}

// RemoveArtStyle removes "art_style" edges to ArtStyle entities.
func (au *ArtUpdate) RemoveArtStyle(a ...*ArtStyle) *ArtUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveArtStyleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArtUpdate) Save(ctx context.Context) (int, error) {
	if err := au.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArtUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArtUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArtUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ArtUpdate) defaults() error {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		if art.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized art.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := art.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (au *ArtUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(art.Table, art.Columns, sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedBy(); ok {
		_spec.SetField(art.FieldCreatedBy, field.TypeString, value)
	}
	if au.mutation.CreatedByCleared() {
		_spec.ClearField(art.FieldCreatedBy, field.TypeString)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(art.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedBy(); ok {
		_spec.SetField(art.FieldUpdatedBy, field.TypeString, value)
	}
	if au.mutation.UpdatedByCleared() {
		_spec.ClearField(art.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := au.mutation.DisplayName(); ok {
		_spec.SetField(art.FieldDisplayName, field.TypeString, value)
	}
	if au.mutation.DisplayNameCleared() {
		_spec.ClearField(art.FieldDisplayName, field.TypeString)
	}
	if value, ok := au.mutation.Abbreviation(); ok {
		_spec.SetField(art.FieldAbbreviation, field.TypeString, value)
	}
	if au.mutation.AbbreviationCleared() {
		_spec.ClearField(art.FieldAbbreviation, field.TypeString)
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.SetField(art.FieldDescription, field.TypeString, value)
	}
	if au.mutation.DescriptionCleared() {
		_spec.ClearField(art.FieldDescription, field.TypeString)
	}
	if value, ok := au.mutation.ExternalLinks(); ok {
		_spec.SetField(art.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := au.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, art.FieldExternalLinks, value)
		})
	}
	if au.mutation.ExternalLinksCleared() {
		_spec.ClearField(art.FieldExternalLinks, field.TypeJSON)
	}
	if value, ok := au.mutation.PrimaryImageURL(); ok {
		_spec.SetField(art.FieldPrimaryImageURL, field.TypeString, value)
	}
	if au.mutation.PrimaryImageURLCleared() {
		_spec.ClearField(art.FieldPrimaryImageURL, field.TypeString)
	}
	if value, ok := au.mutation.AdditionalImagesUrls(); ok {
		_spec.SetField(art.FieldAdditionalImagesUrls, field.TypeJSON, value)
	}
	if value, ok := au.mutation.AppendedAdditionalImagesUrls(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, art.FieldAdditionalImagesUrls, value)
		})
	}
	if au.mutation.AdditionalImagesUrlsCleared() {
		_spec.ClearField(art.FieldAdditionalImagesUrls, field.TypeJSON)
	}
	if au.mutation.ArtGenreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtGenreTable,
			Columns: art.ArtGenrePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedArtGenreIDs(); len(nodes) > 0 && !au.mutation.ArtGenreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtGenreTable,
			Columns: art.ArtGenrePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ArtGenreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtGenreTable,
			Columns: art.ArtGenrePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.ArtStyleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtStyleTable,
			Columns: art.ArtStylePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedArtStyleIDs(); len(nodes) > 0 && !au.mutation.ArtStyleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtStyleTable,
			Columns: art.ArtStylePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ArtStyleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtStyleTable,
			Columns: art.ArtStylePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{art.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArtUpdateOne is the builder for updating a single Art entity.
type ArtUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtMutation
}

// SetCreatedBy sets the "created_by" field.
func (auo *ArtUpdateOne) SetCreatedBy(s string) *ArtUpdateOne {
	auo.mutation.SetCreatedBy(s)
	return auo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (auo *ArtUpdateOne) SetNillableCreatedBy(s *string) *ArtUpdateOne {
	if s != nil {
		auo.SetCreatedBy(*s)
	}
	return auo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (auo *ArtUpdateOne) ClearCreatedBy() *ArtUpdateOne {
	auo.mutation.ClearCreatedBy()
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ArtUpdateOne) SetUpdatedAt(t time.Time) *ArtUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetUpdatedBy sets the "updated_by" field.
func (auo *ArtUpdateOne) SetUpdatedBy(s string) *ArtUpdateOne {
	auo.mutation.SetUpdatedBy(s)
	return auo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (auo *ArtUpdateOne) SetNillableUpdatedBy(s *string) *ArtUpdateOne {
	if s != nil {
		auo.SetUpdatedBy(*s)
	}
	return auo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (auo *ArtUpdateOne) ClearUpdatedBy() *ArtUpdateOne {
	auo.mutation.ClearUpdatedBy()
	return auo
}

// SetDisplayName sets the "display_name" field.
func (auo *ArtUpdateOne) SetDisplayName(s string) *ArtUpdateOne {
	auo.mutation.SetDisplayName(s)
	return auo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (auo *ArtUpdateOne) SetNillableDisplayName(s *string) *ArtUpdateOne {
	if s != nil {
		auo.SetDisplayName(*s)
	}
	return auo
}

// ClearDisplayName clears the value of the "display_name" field.
func (auo *ArtUpdateOne) ClearDisplayName() *ArtUpdateOne {
	auo.mutation.ClearDisplayName()
	return auo
}

// SetAbbreviation sets the "abbreviation" field.
func (auo *ArtUpdateOne) SetAbbreviation(s string) *ArtUpdateOne {
	auo.mutation.SetAbbreviation(s)
	return auo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (auo *ArtUpdateOne) SetNillableAbbreviation(s *string) *ArtUpdateOne {
	if s != nil {
		auo.SetAbbreviation(*s)
	}
	return auo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (auo *ArtUpdateOne) ClearAbbreviation() *ArtUpdateOne {
	auo.mutation.ClearAbbreviation()
	return auo
}

// SetDescription sets the "description" field.
func (auo *ArtUpdateOne) SetDescription(s string) *ArtUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (auo *ArtUpdateOne) SetNillableDescription(s *string) *ArtUpdateOne {
	if s != nil {
		auo.SetDescription(*s)
	}
	return auo
}

// ClearDescription clears the value of the "description" field.
func (auo *ArtUpdateOne) ClearDescription() *ArtUpdateOne {
	auo.mutation.ClearDescription()
	return auo
}

// SetExternalLinks sets the "external_links" field.
func (auo *ArtUpdateOne) SetExternalLinks(s []string) *ArtUpdateOne {
	auo.mutation.SetExternalLinks(s)
	return auo
}

// AppendExternalLinks appends s to the "external_links" field.
func (auo *ArtUpdateOne) AppendExternalLinks(s []string) *ArtUpdateOne {
	auo.mutation.AppendExternalLinks(s)
	return auo
}

// ClearExternalLinks clears the value of the "external_links" field.
func (auo *ArtUpdateOne) ClearExternalLinks() *ArtUpdateOne {
	auo.mutation.ClearExternalLinks()
	return auo
}

// SetPrimaryImageURL sets the "primary_image_url" field.
func (auo *ArtUpdateOne) SetPrimaryImageURL(s string) *ArtUpdateOne {
	auo.mutation.SetPrimaryImageURL(s)
	return auo
}

// SetNillablePrimaryImageURL sets the "primary_image_url" field if the given value is not nil.
func (auo *ArtUpdateOne) SetNillablePrimaryImageURL(s *string) *ArtUpdateOne {
	if s != nil {
		auo.SetPrimaryImageURL(*s)
	}
	return auo
}

// ClearPrimaryImageURL clears the value of the "primary_image_url" field.
func (auo *ArtUpdateOne) ClearPrimaryImageURL() *ArtUpdateOne {
	auo.mutation.ClearPrimaryImageURL()
	return auo
}

// SetAdditionalImagesUrls sets the "additional_images_urls" field.
func (auo *ArtUpdateOne) SetAdditionalImagesUrls(s []string) *ArtUpdateOne {
	auo.mutation.SetAdditionalImagesUrls(s)
	return auo
}

// AppendAdditionalImagesUrls appends s to the "additional_images_urls" field.
func (auo *ArtUpdateOne) AppendAdditionalImagesUrls(s []string) *ArtUpdateOne {
	auo.mutation.AppendAdditionalImagesUrls(s)
	return auo
}

// ClearAdditionalImagesUrls clears the value of the "additional_images_urls" field.
func (auo *ArtUpdateOne) ClearAdditionalImagesUrls() *ArtUpdateOne {
	auo.mutation.ClearAdditionalImagesUrls()
	return auo
}

// AddArtGenreIDs adds the "art_genre" edge to the ArtGenre entity by IDs.
func (auo *ArtUpdateOne) AddArtGenreIDs(ids ...int) *ArtUpdateOne {
	auo.mutation.AddArtGenreIDs(ids...)
	return auo
}

// AddArtGenre adds the "art_genre" edges to the ArtGenre entity.
func (auo *ArtUpdateOne) AddArtGenre(a ...*ArtGenre) *ArtUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddArtGenreIDs(ids...)
}

// AddArtStyleIDs adds the "art_style" edge to the ArtStyle entity by IDs.
func (auo *ArtUpdateOne) AddArtStyleIDs(ids ...int) *ArtUpdateOne {
	auo.mutation.AddArtStyleIDs(ids...)
	return auo
}

// AddArtStyle adds the "art_style" edges to the ArtStyle entity.
func (auo *ArtUpdateOne) AddArtStyle(a ...*ArtStyle) *ArtUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddArtStyleIDs(ids...)
}

// Mutation returns the ArtMutation object of the builder.
func (auo *ArtUpdateOne) Mutation() *ArtMutation {
	return auo.mutation
}

// ClearArtGenre clears all "art_genre" edges to the ArtGenre entity.
func (auo *ArtUpdateOne) ClearArtGenre() *ArtUpdateOne {
	auo.mutation.ClearArtGenre()
	return auo
}

// RemoveArtGenreIDs removes the "art_genre" edge to ArtGenre entities by IDs.
func (auo *ArtUpdateOne) RemoveArtGenreIDs(ids ...int) *ArtUpdateOne {
	auo.mutation.RemoveArtGenreIDs(ids...)
	return auo
}

// RemoveArtGenre removes "art_genre" edges to ArtGenre entities.
func (auo *ArtUpdateOne) RemoveArtGenre(a ...*ArtGenre) *ArtUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveArtGenreIDs(ids...)
}

// ClearArtStyle clears all "art_style" edges to the ArtStyle entity.
func (auo *ArtUpdateOne) ClearArtStyle() *ArtUpdateOne {
	auo.mutation.ClearArtStyle()
	return auo
}

// RemoveArtStyleIDs removes the "art_style" edge to ArtStyle entities by IDs.
func (auo *ArtUpdateOne) RemoveArtStyleIDs(ids ...int) *ArtUpdateOne {
	auo.mutation.RemoveArtStyleIDs(ids...)
	return auo
}

// RemoveArtStyle removes "art_style" edges to ArtStyle entities.
func (auo *ArtUpdateOne) RemoveArtStyle(a ...*ArtStyle) *ArtUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveArtStyleIDs(ids...)
}

// Where appends a list predicates to the ArtUpdate builder.
func (auo *ArtUpdateOne) Where(ps ...predicate.Art) *ArtUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArtUpdateOne) Select(field string, fields ...string) *ArtUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Art entity.
func (auo *ArtUpdateOne) Save(ctx context.Context) (*Art, error) {
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArtUpdateOne) SaveX(ctx context.Context) *Art {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArtUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArtUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ArtUpdateOne) defaults() error {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		if art.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized art.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := art.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (auo *ArtUpdateOne) sqlSave(ctx context.Context) (_node *Art, err error) {
	_spec := sqlgraph.NewUpdateSpec(art.Table, art.Columns, sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Art.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, art.FieldID)
		for _, f := range fields {
			if !art.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != art.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedBy(); ok {
		_spec.SetField(art.FieldCreatedBy, field.TypeString, value)
	}
	if auo.mutation.CreatedByCleared() {
		_spec.ClearField(art.FieldCreatedBy, field.TypeString)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(art.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedBy(); ok {
		_spec.SetField(art.FieldUpdatedBy, field.TypeString, value)
	}
	if auo.mutation.UpdatedByCleared() {
		_spec.ClearField(art.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := auo.mutation.DisplayName(); ok {
		_spec.SetField(art.FieldDisplayName, field.TypeString, value)
	}
	if auo.mutation.DisplayNameCleared() {
		_spec.ClearField(art.FieldDisplayName, field.TypeString)
	}
	if value, ok := auo.mutation.Abbreviation(); ok {
		_spec.SetField(art.FieldAbbreviation, field.TypeString, value)
	}
	if auo.mutation.AbbreviationCleared() {
		_spec.ClearField(art.FieldAbbreviation, field.TypeString)
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.SetField(art.FieldDescription, field.TypeString, value)
	}
	if auo.mutation.DescriptionCleared() {
		_spec.ClearField(art.FieldDescription, field.TypeString)
	}
	if value, ok := auo.mutation.ExternalLinks(); ok {
		_spec.SetField(art.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, art.FieldExternalLinks, value)
		})
	}
	if auo.mutation.ExternalLinksCleared() {
		_spec.ClearField(art.FieldExternalLinks, field.TypeJSON)
	}
	if value, ok := auo.mutation.PrimaryImageURL(); ok {
		_spec.SetField(art.FieldPrimaryImageURL, field.TypeString, value)
	}
	if auo.mutation.PrimaryImageURLCleared() {
		_spec.ClearField(art.FieldPrimaryImageURL, field.TypeString)
	}
	if value, ok := auo.mutation.AdditionalImagesUrls(); ok {
		_spec.SetField(art.FieldAdditionalImagesUrls, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.AppendedAdditionalImagesUrls(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, art.FieldAdditionalImagesUrls, value)
		})
	}
	if auo.mutation.AdditionalImagesUrlsCleared() {
		_spec.ClearField(art.FieldAdditionalImagesUrls, field.TypeJSON)
	}
	if auo.mutation.ArtGenreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtGenreTable,
			Columns: art.ArtGenrePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedArtGenreIDs(); len(nodes) > 0 && !auo.mutation.ArtGenreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtGenreTable,
			Columns: art.ArtGenrePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ArtGenreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtGenreTable,
			Columns: art.ArtGenrePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.ArtStyleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtStyleTable,
			Columns: art.ArtStylePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedArtStyleIDs(); len(nodes) > 0 && !auo.mutation.ArtStyleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtStyleTable,
			Columns: art.ArtStylePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ArtStyleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtStyleTable,
			Columns: art.ArtStylePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Art{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{art.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
