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
	"github.com/dkrasnovdev/heritage-api/ent/artifact"
	"github.com/dkrasnovdev/heritage-api/ent/person"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
	"github.com/dkrasnovdev/heritage-api/ent/publication"
)

// PublicationUpdate is the builder for updating Publication entities.
type PublicationUpdate struct {
	config
	hooks    []Hook
	mutation *PublicationMutation
}

// Where appends a list predicates to the PublicationUpdate builder.
func (pu *PublicationUpdate) Where(ps ...predicate.Publication) *PublicationUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *PublicationUpdate) SetCreatedBy(s string) *PublicationUpdate {
	pu.mutation.SetCreatedBy(s)
	return pu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pu *PublicationUpdate) SetNillableCreatedBy(s *string) *PublicationUpdate {
	if s != nil {
		pu.SetCreatedBy(*s)
	}
	return pu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (pu *PublicationUpdate) ClearCreatedBy() *PublicationUpdate {
	pu.mutation.ClearCreatedBy()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PublicationUpdate) SetUpdatedAt(t time.Time) *PublicationUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *PublicationUpdate) SetUpdatedBy(s string) *PublicationUpdate {
	pu.mutation.SetUpdatedBy(s)
	return pu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pu *PublicationUpdate) SetNillableUpdatedBy(s *string) *PublicationUpdate {
	if s != nil {
		pu.SetUpdatedBy(*s)
	}
	return pu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pu *PublicationUpdate) ClearUpdatedBy() *PublicationUpdate {
	pu.mutation.ClearUpdatedBy()
	return pu
}

// SetAbbreviation sets the "abbreviation" field.
func (pu *PublicationUpdate) SetAbbreviation(s string) *PublicationUpdate {
	pu.mutation.SetAbbreviation(s)
	return pu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pu *PublicationUpdate) SetNillableAbbreviation(s *string) *PublicationUpdate {
	if s != nil {
		pu.SetAbbreviation(*s)
	}
	return pu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (pu *PublicationUpdate) ClearAbbreviation() *PublicationUpdate {
	pu.mutation.ClearAbbreviation()
	return pu
}

// SetDisplayName sets the "display_name" field.
func (pu *PublicationUpdate) SetDisplayName(s string) *PublicationUpdate {
	pu.mutation.SetDisplayName(s)
	return pu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pu *PublicationUpdate) SetNillableDisplayName(s *string) *PublicationUpdate {
	if s != nil {
		pu.SetDisplayName(*s)
	}
	return pu
}

// ClearDisplayName clears the value of the "display_name" field.
func (pu *PublicationUpdate) ClearDisplayName() *PublicationUpdate {
	pu.mutation.ClearDisplayName()
	return pu
}

// SetDescription sets the "description" field.
func (pu *PublicationUpdate) SetDescription(s string) *PublicationUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PublicationUpdate) SetNillableDescription(s *string) *PublicationUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PublicationUpdate) ClearDescription() *PublicationUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetExternalLinks sets the "external_links" field.
func (pu *PublicationUpdate) SetExternalLinks(s []string) *PublicationUpdate {
	pu.mutation.SetExternalLinks(s)
	return pu
}

// AppendExternalLinks appends s to the "external_links" field.
func (pu *PublicationUpdate) AppendExternalLinks(s []string) *PublicationUpdate {
	pu.mutation.AppendExternalLinks(s)
	return pu
}

// ClearExternalLinks clears the value of the "external_links" field.
func (pu *PublicationUpdate) ClearExternalLinks() *PublicationUpdate {
	pu.mutation.ClearExternalLinks()
	return pu
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (pu *PublicationUpdate) AddArtifactIDs(ids ...int) *PublicationUpdate {
	pu.mutation.AddArtifactIDs(ids...)
	return pu
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (pu *PublicationUpdate) AddArtifacts(a ...*Artifact) *PublicationUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddArtifactIDs(ids...)
}

// AddAuthorIDs adds the "authors" edge to the Person entity by IDs.
func (pu *PublicationUpdate) AddAuthorIDs(ids ...int) *PublicationUpdate {
	pu.mutation.AddAuthorIDs(ids...)
	return pu
}

// AddAuthors adds the "authors" edges to the Person entity.
func (pu *PublicationUpdate) AddAuthors(p ...*Person) *PublicationUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddAuthorIDs(ids...)
}

// Mutation returns the PublicationMutation object of the builder.
func (pu *PublicationUpdate) Mutation() *PublicationMutation {
	return pu.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (pu *PublicationUpdate) ClearArtifacts() *PublicationUpdate {
	pu.mutation.ClearArtifacts()
	return pu
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (pu *PublicationUpdate) RemoveArtifactIDs(ids ...int) *PublicationUpdate {
	pu.mutation.RemoveArtifactIDs(ids...)
	return pu
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (pu *PublicationUpdate) RemoveArtifacts(a ...*Artifact) *PublicationUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveArtifactIDs(ids...)
}

// ClearAuthors clears all "authors" edges to the Person entity.
func (pu *PublicationUpdate) ClearAuthors() *PublicationUpdate {
	pu.mutation.ClearAuthors()
	return pu
}

// RemoveAuthorIDs removes the "authors" edge to Person entities by IDs.
func (pu *PublicationUpdate) RemoveAuthorIDs(ids ...int) *PublicationUpdate {
	pu.mutation.RemoveAuthorIDs(ids...)
	return pu
}

// RemoveAuthors removes "authors" edges to Person entities.
func (pu *PublicationUpdate) RemoveAuthors(p ...*Person) *PublicationUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveAuthorIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PublicationUpdate) Save(ctx context.Context) (int, error) {
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PublicationUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PublicationUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PublicationUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PublicationUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if publication.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized publication.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := publication.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pu *PublicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(publication.Table, publication.Columns, sqlgraph.NewFieldSpec(publication.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.SetField(publication.FieldCreatedBy, field.TypeString, value)
	}
	if pu.mutation.CreatedByCleared() {
		_spec.ClearField(publication.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(publication.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.SetField(publication.FieldUpdatedBy, field.TypeString, value)
	}
	if pu.mutation.UpdatedByCleared() {
		_spec.ClearField(publication.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.Abbreviation(); ok {
		_spec.SetField(publication.FieldAbbreviation, field.TypeString, value)
	}
	if pu.mutation.AbbreviationCleared() {
		_spec.ClearField(publication.FieldAbbreviation, field.TypeString)
	}
	if value, ok := pu.mutation.DisplayName(); ok {
		_spec.SetField(publication.FieldDisplayName, field.TypeString, value)
	}
	if pu.mutation.DisplayNameCleared() {
		_spec.ClearField(publication.FieldDisplayName, field.TypeString)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(publication.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(publication.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.ExternalLinks(); ok {
		_spec.SetField(publication.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, publication.FieldExternalLinks, value)
		})
	}
	if pu.mutation.ExternalLinksCleared() {
		_spec.ClearField(publication.FieldExternalLinks, field.TypeJSON)
	}
	if pu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   publication.ArtifactsTable,
			Columns: publication.ArtifactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !pu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   publication.ArtifactsTable,
			Columns: publication.ArtifactsPrimaryKey,
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
	if nodes := pu.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   publication.ArtifactsTable,
			Columns: publication.ArtifactsPrimaryKey,
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
	if pu.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publication.AuthorsTable,
			Columns: publication.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedAuthorsIDs(); len(nodes) > 0 && !pu.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publication.AuthorsTable,
			Columns: publication.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publication.AuthorsTable,
			Columns: publication.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{publication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PublicationUpdateOne is the builder for updating a single Publication entity.
type PublicationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PublicationMutation
}

// SetCreatedBy sets the "created_by" field.
func (puo *PublicationUpdateOne) SetCreatedBy(s string) *PublicationUpdateOne {
	puo.mutation.SetCreatedBy(s)
	return puo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (puo *PublicationUpdateOne) SetNillableCreatedBy(s *string) *PublicationUpdateOne {
	if s != nil {
		puo.SetCreatedBy(*s)
	}
	return puo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (puo *PublicationUpdateOne) ClearCreatedBy() *PublicationUpdateOne {
	puo.mutation.ClearCreatedBy()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PublicationUpdateOne) SetUpdatedAt(t time.Time) *PublicationUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *PublicationUpdateOne) SetUpdatedBy(s string) *PublicationUpdateOne {
	puo.mutation.SetUpdatedBy(s)
	return puo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (puo *PublicationUpdateOne) SetNillableUpdatedBy(s *string) *PublicationUpdateOne {
	if s != nil {
		puo.SetUpdatedBy(*s)
	}
	return puo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (puo *PublicationUpdateOne) ClearUpdatedBy() *PublicationUpdateOne {
	puo.mutation.ClearUpdatedBy()
	return puo
}

// SetAbbreviation sets the "abbreviation" field.
func (puo *PublicationUpdateOne) SetAbbreviation(s string) *PublicationUpdateOne {
	puo.mutation.SetAbbreviation(s)
	return puo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (puo *PublicationUpdateOne) SetNillableAbbreviation(s *string) *PublicationUpdateOne {
	if s != nil {
		puo.SetAbbreviation(*s)
	}
	return puo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (puo *PublicationUpdateOne) ClearAbbreviation() *PublicationUpdateOne {
	puo.mutation.ClearAbbreviation()
	return puo
}

// SetDisplayName sets the "display_name" field.
func (puo *PublicationUpdateOne) SetDisplayName(s string) *PublicationUpdateOne {
	puo.mutation.SetDisplayName(s)
	return puo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (puo *PublicationUpdateOne) SetNillableDisplayName(s *string) *PublicationUpdateOne {
	if s != nil {
		puo.SetDisplayName(*s)
	}
	return puo
}

// ClearDisplayName clears the value of the "display_name" field.
func (puo *PublicationUpdateOne) ClearDisplayName() *PublicationUpdateOne {
	puo.mutation.ClearDisplayName()
	return puo
}

// SetDescription sets the "description" field.
func (puo *PublicationUpdateOne) SetDescription(s string) *PublicationUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PublicationUpdateOne) SetNillableDescription(s *string) *PublicationUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PublicationUpdateOne) ClearDescription() *PublicationUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetExternalLinks sets the "external_links" field.
func (puo *PublicationUpdateOne) SetExternalLinks(s []string) *PublicationUpdateOne {
	puo.mutation.SetExternalLinks(s)
	return puo
}

// AppendExternalLinks appends s to the "external_links" field.
func (puo *PublicationUpdateOne) AppendExternalLinks(s []string) *PublicationUpdateOne {
	puo.mutation.AppendExternalLinks(s)
	return puo
}

// ClearExternalLinks clears the value of the "external_links" field.
func (puo *PublicationUpdateOne) ClearExternalLinks() *PublicationUpdateOne {
	puo.mutation.ClearExternalLinks()
	return puo
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (puo *PublicationUpdateOne) AddArtifactIDs(ids ...int) *PublicationUpdateOne {
	puo.mutation.AddArtifactIDs(ids...)
	return puo
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (puo *PublicationUpdateOne) AddArtifacts(a ...*Artifact) *PublicationUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddArtifactIDs(ids...)
}

// AddAuthorIDs adds the "authors" edge to the Person entity by IDs.
func (puo *PublicationUpdateOne) AddAuthorIDs(ids ...int) *PublicationUpdateOne {
	puo.mutation.AddAuthorIDs(ids...)
	return puo
}

// AddAuthors adds the "authors" edges to the Person entity.
func (puo *PublicationUpdateOne) AddAuthors(p ...*Person) *PublicationUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddAuthorIDs(ids...)
}

// Mutation returns the PublicationMutation object of the builder.
func (puo *PublicationUpdateOne) Mutation() *PublicationMutation {
	return puo.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (puo *PublicationUpdateOne) ClearArtifacts() *PublicationUpdateOne {
	puo.mutation.ClearArtifacts()
	return puo
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (puo *PublicationUpdateOne) RemoveArtifactIDs(ids ...int) *PublicationUpdateOne {
	puo.mutation.RemoveArtifactIDs(ids...)
	return puo
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (puo *PublicationUpdateOne) RemoveArtifacts(a ...*Artifact) *PublicationUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveArtifactIDs(ids...)
}

// ClearAuthors clears all "authors" edges to the Person entity.
func (puo *PublicationUpdateOne) ClearAuthors() *PublicationUpdateOne {
	puo.mutation.ClearAuthors()
	return puo
}

// RemoveAuthorIDs removes the "authors" edge to Person entities by IDs.
func (puo *PublicationUpdateOne) RemoveAuthorIDs(ids ...int) *PublicationUpdateOne {
	puo.mutation.RemoveAuthorIDs(ids...)
	return puo
}

// RemoveAuthors removes "authors" edges to Person entities.
func (puo *PublicationUpdateOne) RemoveAuthors(p ...*Person) *PublicationUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveAuthorIDs(ids...)
}

// Where appends a list predicates to the PublicationUpdate builder.
func (puo *PublicationUpdateOne) Where(ps ...predicate.Publication) *PublicationUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PublicationUpdateOne) Select(field string, fields ...string) *PublicationUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Publication entity.
func (puo *PublicationUpdateOne) Save(ctx context.Context) (*Publication, error) {
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PublicationUpdateOne) SaveX(ctx context.Context) *Publication {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PublicationUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PublicationUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PublicationUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if publication.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized publication.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := publication.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (puo *PublicationUpdateOne) sqlSave(ctx context.Context) (_node *Publication, err error) {
	_spec := sqlgraph.NewUpdateSpec(publication.Table, publication.Columns, sqlgraph.NewFieldSpec(publication.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Publication.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, publication.FieldID)
		for _, f := range fields {
			if !publication.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != publication.FieldID {
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
		_spec.SetField(publication.FieldCreatedBy, field.TypeString, value)
	}
	if puo.mutation.CreatedByCleared() {
		_spec.ClearField(publication.FieldCreatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(publication.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.SetField(publication.FieldUpdatedBy, field.TypeString, value)
	}
	if puo.mutation.UpdatedByCleared() {
		_spec.ClearField(publication.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.Abbreviation(); ok {
		_spec.SetField(publication.FieldAbbreviation, field.TypeString, value)
	}
	if puo.mutation.AbbreviationCleared() {
		_spec.ClearField(publication.FieldAbbreviation, field.TypeString)
	}
	if value, ok := puo.mutation.DisplayName(); ok {
		_spec.SetField(publication.FieldDisplayName, field.TypeString, value)
	}
	if puo.mutation.DisplayNameCleared() {
		_spec.ClearField(publication.FieldDisplayName, field.TypeString)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(publication.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(publication.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.ExternalLinks(); ok {
		_spec.SetField(publication.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, publication.FieldExternalLinks, value)
		})
	}
	if puo.mutation.ExternalLinksCleared() {
		_spec.ClearField(publication.FieldExternalLinks, field.TypeJSON)
	}
	if puo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   publication.ArtifactsTable,
			Columns: publication.ArtifactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !puo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   publication.ArtifactsTable,
			Columns: publication.ArtifactsPrimaryKey,
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
	if nodes := puo.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   publication.ArtifactsTable,
			Columns: publication.ArtifactsPrimaryKey,
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
	if puo.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publication.AuthorsTable,
			Columns: publication.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedAuthorsIDs(); len(nodes) > 0 && !puo.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publication.AuthorsTable,
			Columns: publication.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publication.AuthorsTable,
			Columns: publication.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Publication{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{publication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
