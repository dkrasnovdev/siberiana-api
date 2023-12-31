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
	"github.com/dkrasnovdev/siberiana-api/ent/art"
	"github.com/dkrasnovdev/siberiana-api/ent/artgenre"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// ArtGenreUpdate is the builder for updating ArtGenre entities.
type ArtGenreUpdate struct {
	config
	hooks    []Hook
	mutation *ArtGenreMutation
}

// Where appends a list predicates to the ArtGenreUpdate builder.
func (agu *ArtGenreUpdate) Where(ps ...predicate.ArtGenre) *ArtGenreUpdate {
	agu.mutation.Where(ps...)
	return agu
}

// SetCreatedBy sets the "created_by" field.
func (agu *ArtGenreUpdate) SetCreatedBy(s string) *ArtGenreUpdate {
	agu.mutation.SetCreatedBy(s)
	return agu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (agu *ArtGenreUpdate) SetNillableCreatedBy(s *string) *ArtGenreUpdate {
	if s != nil {
		agu.SetCreatedBy(*s)
	}
	return agu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (agu *ArtGenreUpdate) ClearCreatedBy() *ArtGenreUpdate {
	agu.mutation.ClearCreatedBy()
	return agu
}

// SetUpdatedAt sets the "updated_at" field.
func (agu *ArtGenreUpdate) SetUpdatedAt(t time.Time) *ArtGenreUpdate {
	agu.mutation.SetUpdatedAt(t)
	return agu
}

// SetUpdatedBy sets the "updated_by" field.
func (agu *ArtGenreUpdate) SetUpdatedBy(s string) *ArtGenreUpdate {
	agu.mutation.SetUpdatedBy(s)
	return agu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (agu *ArtGenreUpdate) SetNillableUpdatedBy(s *string) *ArtGenreUpdate {
	if s != nil {
		agu.SetUpdatedBy(*s)
	}
	return agu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (agu *ArtGenreUpdate) ClearUpdatedBy() *ArtGenreUpdate {
	agu.mutation.ClearUpdatedBy()
	return agu
}

// SetDisplayName sets the "display_name" field.
func (agu *ArtGenreUpdate) SetDisplayName(s string) *ArtGenreUpdate {
	agu.mutation.SetDisplayName(s)
	return agu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (agu *ArtGenreUpdate) SetNillableDisplayName(s *string) *ArtGenreUpdate {
	if s != nil {
		agu.SetDisplayName(*s)
	}
	return agu
}

// ClearDisplayName clears the value of the "display_name" field.
func (agu *ArtGenreUpdate) ClearDisplayName() *ArtGenreUpdate {
	agu.mutation.ClearDisplayName()
	return agu
}

// SetAbbreviation sets the "abbreviation" field.
func (agu *ArtGenreUpdate) SetAbbreviation(s string) *ArtGenreUpdate {
	agu.mutation.SetAbbreviation(s)
	return agu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (agu *ArtGenreUpdate) SetNillableAbbreviation(s *string) *ArtGenreUpdate {
	if s != nil {
		agu.SetAbbreviation(*s)
	}
	return agu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (agu *ArtGenreUpdate) ClearAbbreviation() *ArtGenreUpdate {
	agu.mutation.ClearAbbreviation()
	return agu
}

// SetDescription sets the "description" field.
func (agu *ArtGenreUpdate) SetDescription(s string) *ArtGenreUpdate {
	agu.mutation.SetDescription(s)
	return agu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (agu *ArtGenreUpdate) SetNillableDescription(s *string) *ArtGenreUpdate {
	if s != nil {
		agu.SetDescription(*s)
	}
	return agu
}

// ClearDescription clears the value of the "description" field.
func (agu *ArtGenreUpdate) ClearDescription() *ArtGenreUpdate {
	agu.mutation.ClearDescription()
	return agu
}

// SetExternalLink sets the "external_link" field.
func (agu *ArtGenreUpdate) SetExternalLink(s string) *ArtGenreUpdate {
	agu.mutation.SetExternalLink(s)
	return agu
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (agu *ArtGenreUpdate) SetNillableExternalLink(s *string) *ArtGenreUpdate {
	if s != nil {
		agu.SetExternalLink(*s)
	}
	return agu
}

// ClearExternalLink clears the value of the "external_link" field.
func (agu *ArtGenreUpdate) ClearExternalLink() *ArtGenreUpdate {
	agu.mutation.ClearExternalLink()
	return agu
}

// AddArtIDs adds the "art" edge to the Art entity by IDs.
func (agu *ArtGenreUpdate) AddArtIDs(ids ...int) *ArtGenreUpdate {
	agu.mutation.AddArtIDs(ids...)
	return agu
}

// AddArt adds the "art" edges to the Art entity.
func (agu *ArtGenreUpdate) AddArt(a ...*Art) *ArtGenreUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return agu.AddArtIDs(ids...)
}

// Mutation returns the ArtGenreMutation object of the builder.
func (agu *ArtGenreUpdate) Mutation() *ArtGenreMutation {
	return agu.mutation
}

// ClearArt clears all "art" edges to the Art entity.
func (agu *ArtGenreUpdate) ClearArt() *ArtGenreUpdate {
	agu.mutation.ClearArt()
	return agu
}

// RemoveArtIDs removes the "art" edge to Art entities by IDs.
func (agu *ArtGenreUpdate) RemoveArtIDs(ids ...int) *ArtGenreUpdate {
	agu.mutation.RemoveArtIDs(ids...)
	return agu
}

// RemoveArt removes "art" edges to Art entities.
func (agu *ArtGenreUpdate) RemoveArt(a ...*Art) *ArtGenreUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return agu.RemoveArtIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (agu *ArtGenreUpdate) Save(ctx context.Context) (int, error) {
	if err := agu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, agu.sqlSave, agu.mutation, agu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (agu *ArtGenreUpdate) SaveX(ctx context.Context) int {
	affected, err := agu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (agu *ArtGenreUpdate) Exec(ctx context.Context) error {
	_, err := agu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agu *ArtGenreUpdate) ExecX(ctx context.Context) {
	if err := agu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agu *ArtGenreUpdate) defaults() error {
	if _, ok := agu.mutation.UpdatedAt(); !ok {
		if artgenre.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized artgenre.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := artgenre.UpdateDefaultUpdatedAt()
		agu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (agu *ArtGenreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(artgenre.Table, artgenre.Columns, sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt))
	if ps := agu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := agu.mutation.CreatedBy(); ok {
		_spec.SetField(artgenre.FieldCreatedBy, field.TypeString, value)
	}
	if agu.mutation.CreatedByCleared() {
		_spec.ClearField(artgenre.FieldCreatedBy, field.TypeString)
	}
	if value, ok := agu.mutation.UpdatedAt(); ok {
		_spec.SetField(artgenre.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := agu.mutation.UpdatedBy(); ok {
		_spec.SetField(artgenre.FieldUpdatedBy, field.TypeString, value)
	}
	if agu.mutation.UpdatedByCleared() {
		_spec.ClearField(artgenre.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := agu.mutation.DisplayName(); ok {
		_spec.SetField(artgenre.FieldDisplayName, field.TypeString, value)
	}
	if agu.mutation.DisplayNameCleared() {
		_spec.ClearField(artgenre.FieldDisplayName, field.TypeString)
	}
	if value, ok := agu.mutation.Abbreviation(); ok {
		_spec.SetField(artgenre.FieldAbbreviation, field.TypeString, value)
	}
	if agu.mutation.AbbreviationCleared() {
		_spec.ClearField(artgenre.FieldAbbreviation, field.TypeString)
	}
	if value, ok := agu.mutation.Description(); ok {
		_spec.SetField(artgenre.FieldDescription, field.TypeString, value)
	}
	if agu.mutation.DescriptionCleared() {
		_spec.ClearField(artgenre.FieldDescription, field.TypeString)
	}
	if value, ok := agu.mutation.ExternalLink(); ok {
		_spec.SetField(artgenre.FieldExternalLink, field.TypeString, value)
	}
	if agu.mutation.ExternalLinkCleared() {
		_spec.ClearField(artgenre.FieldExternalLink, field.TypeString)
	}
	if agu.mutation.ArtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artgenre.ArtTable,
			Columns: artgenre.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := agu.mutation.RemovedArtIDs(); len(nodes) > 0 && !agu.mutation.ArtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artgenre.ArtTable,
			Columns: artgenre.ArtPrimaryKey,
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
	if nodes := agu.mutation.ArtIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artgenre.ArtTable,
			Columns: artgenre.ArtPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, agu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artgenre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	agu.mutation.done = true
	return n, nil
}

// ArtGenreUpdateOne is the builder for updating a single ArtGenre entity.
type ArtGenreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtGenreMutation
}

// SetCreatedBy sets the "created_by" field.
func (aguo *ArtGenreUpdateOne) SetCreatedBy(s string) *ArtGenreUpdateOne {
	aguo.mutation.SetCreatedBy(s)
	return aguo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (aguo *ArtGenreUpdateOne) SetNillableCreatedBy(s *string) *ArtGenreUpdateOne {
	if s != nil {
		aguo.SetCreatedBy(*s)
	}
	return aguo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (aguo *ArtGenreUpdateOne) ClearCreatedBy() *ArtGenreUpdateOne {
	aguo.mutation.ClearCreatedBy()
	return aguo
}

// SetUpdatedAt sets the "updated_at" field.
func (aguo *ArtGenreUpdateOne) SetUpdatedAt(t time.Time) *ArtGenreUpdateOne {
	aguo.mutation.SetUpdatedAt(t)
	return aguo
}

// SetUpdatedBy sets the "updated_by" field.
func (aguo *ArtGenreUpdateOne) SetUpdatedBy(s string) *ArtGenreUpdateOne {
	aguo.mutation.SetUpdatedBy(s)
	return aguo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aguo *ArtGenreUpdateOne) SetNillableUpdatedBy(s *string) *ArtGenreUpdateOne {
	if s != nil {
		aguo.SetUpdatedBy(*s)
	}
	return aguo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aguo *ArtGenreUpdateOne) ClearUpdatedBy() *ArtGenreUpdateOne {
	aguo.mutation.ClearUpdatedBy()
	return aguo
}

// SetDisplayName sets the "display_name" field.
func (aguo *ArtGenreUpdateOne) SetDisplayName(s string) *ArtGenreUpdateOne {
	aguo.mutation.SetDisplayName(s)
	return aguo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (aguo *ArtGenreUpdateOne) SetNillableDisplayName(s *string) *ArtGenreUpdateOne {
	if s != nil {
		aguo.SetDisplayName(*s)
	}
	return aguo
}

// ClearDisplayName clears the value of the "display_name" field.
func (aguo *ArtGenreUpdateOne) ClearDisplayName() *ArtGenreUpdateOne {
	aguo.mutation.ClearDisplayName()
	return aguo
}

// SetAbbreviation sets the "abbreviation" field.
func (aguo *ArtGenreUpdateOne) SetAbbreviation(s string) *ArtGenreUpdateOne {
	aguo.mutation.SetAbbreviation(s)
	return aguo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (aguo *ArtGenreUpdateOne) SetNillableAbbreviation(s *string) *ArtGenreUpdateOne {
	if s != nil {
		aguo.SetAbbreviation(*s)
	}
	return aguo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (aguo *ArtGenreUpdateOne) ClearAbbreviation() *ArtGenreUpdateOne {
	aguo.mutation.ClearAbbreviation()
	return aguo
}

// SetDescription sets the "description" field.
func (aguo *ArtGenreUpdateOne) SetDescription(s string) *ArtGenreUpdateOne {
	aguo.mutation.SetDescription(s)
	return aguo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (aguo *ArtGenreUpdateOne) SetNillableDescription(s *string) *ArtGenreUpdateOne {
	if s != nil {
		aguo.SetDescription(*s)
	}
	return aguo
}

// ClearDescription clears the value of the "description" field.
func (aguo *ArtGenreUpdateOne) ClearDescription() *ArtGenreUpdateOne {
	aguo.mutation.ClearDescription()
	return aguo
}

// SetExternalLink sets the "external_link" field.
func (aguo *ArtGenreUpdateOne) SetExternalLink(s string) *ArtGenreUpdateOne {
	aguo.mutation.SetExternalLink(s)
	return aguo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (aguo *ArtGenreUpdateOne) SetNillableExternalLink(s *string) *ArtGenreUpdateOne {
	if s != nil {
		aguo.SetExternalLink(*s)
	}
	return aguo
}

// ClearExternalLink clears the value of the "external_link" field.
func (aguo *ArtGenreUpdateOne) ClearExternalLink() *ArtGenreUpdateOne {
	aguo.mutation.ClearExternalLink()
	return aguo
}

// AddArtIDs adds the "art" edge to the Art entity by IDs.
func (aguo *ArtGenreUpdateOne) AddArtIDs(ids ...int) *ArtGenreUpdateOne {
	aguo.mutation.AddArtIDs(ids...)
	return aguo
}

// AddArt adds the "art" edges to the Art entity.
func (aguo *ArtGenreUpdateOne) AddArt(a ...*Art) *ArtGenreUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aguo.AddArtIDs(ids...)
}

// Mutation returns the ArtGenreMutation object of the builder.
func (aguo *ArtGenreUpdateOne) Mutation() *ArtGenreMutation {
	return aguo.mutation
}

// ClearArt clears all "art" edges to the Art entity.
func (aguo *ArtGenreUpdateOne) ClearArt() *ArtGenreUpdateOne {
	aguo.mutation.ClearArt()
	return aguo
}

// RemoveArtIDs removes the "art" edge to Art entities by IDs.
func (aguo *ArtGenreUpdateOne) RemoveArtIDs(ids ...int) *ArtGenreUpdateOne {
	aguo.mutation.RemoveArtIDs(ids...)
	return aguo
}

// RemoveArt removes "art" edges to Art entities.
func (aguo *ArtGenreUpdateOne) RemoveArt(a ...*Art) *ArtGenreUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aguo.RemoveArtIDs(ids...)
}

// Where appends a list predicates to the ArtGenreUpdate builder.
func (aguo *ArtGenreUpdateOne) Where(ps ...predicate.ArtGenre) *ArtGenreUpdateOne {
	aguo.mutation.Where(ps...)
	return aguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aguo *ArtGenreUpdateOne) Select(field string, fields ...string) *ArtGenreUpdateOne {
	aguo.fields = append([]string{field}, fields...)
	return aguo
}

// Save executes the query and returns the updated ArtGenre entity.
func (aguo *ArtGenreUpdateOne) Save(ctx context.Context) (*ArtGenre, error) {
	if err := aguo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, aguo.sqlSave, aguo.mutation, aguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aguo *ArtGenreUpdateOne) SaveX(ctx context.Context) *ArtGenre {
	node, err := aguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aguo *ArtGenreUpdateOne) Exec(ctx context.Context) error {
	_, err := aguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aguo *ArtGenreUpdateOne) ExecX(ctx context.Context) {
	if err := aguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aguo *ArtGenreUpdateOne) defaults() error {
	if _, ok := aguo.mutation.UpdatedAt(); !ok {
		if artgenre.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized artgenre.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := artgenre.UpdateDefaultUpdatedAt()
		aguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (aguo *ArtGenreUpdateOne) sqlSave(ctx context.Context) (_node *ArtGenre, err error) {
	_spec := sqlgraph.NewUpdateSpec(artgenre.Table, artgenre.Columns, sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt))
	id, ok := aguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArtGenre.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artgenre.FieldID)
		for _, f := range fields {
			if !artgenre.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != artgenre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aguo.mutation.CreatedBy(); ok {
		_spec.SetField(artgenre.FieldCreatedBy, field.TypeString, value)
	}
	if aguo.mutation.CreatedByCleared() {
		_spec.ClearField(artgenre.FieldCreatedBy, field.TypeString)
	}
	if value, ok := aguo.mutation.UpdatedAt(); ok {
		_spec.SetField(artgenre.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := aguo.mutation.UpdatedBy(); ok {
		_spec.SetField(artgenre.FieldUpdatedBy, field.TypeString, value)
	}
	if aguo.mutation.UpdatedByCleared() {
		_spec.ClearField(artgenre.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := aguo.mutation.DisplayName(); ok {
		_spec.SetField(artgenre.FieldDisplayName, field.TypeString, value)
	}
	if aguo.mutation.DisplayNameCleared() {
		_spec.ClearField(artgenre.FieldDisplayName, field.TypeString)
	}
	if value, ok := aguo.mutation.Abbreviation(); ok {
		_spec.SetField(artgenre.FieldAbbreviation, field.TypeString, value)
	}
	if aguo.mutation.AbbreviationCleared() {
		_spec.ClearField(artgenre.FieldAbbreviation, field.TypeString)
	}
	if value, ok := aguo.mutation.Description(); ok {
		_spec.SetField(artgenre.FieldDescription, field.TypeString, value)
	}
	if aguo.mutation.DescriptionCleared() {
		_spec.ClearField(artgenre.FieldDescription, field.TypeString)
	}
	if value, ok := aguo.mutation.ExternalLink(); ok {
		_spec.SetField(artgenre.FieldExternalLink, field.TypeString, value)
	}
	if aguo.mutation.ExternalLinkCleared() {
		_spec.ClearField(artgenre.FieldExternalLink, field.TypeString)
	}
	if aguo.mutation.ArtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artgenre.ArtTable,
			Columns: artgenre.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aguo.mutation.RemovedArtIDs(); len(nodes) > 0 && !aguo.mutation.ArtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artgenre.ArtTable,
			Columns: artgenre.ArtPrimaryKey,
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
	if nodes := aguo.mutation.ArtIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artgenre.ArtTable,
			Columns: artgenre.ArtPrimaryKey,
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
	_node = &ArtGenre{config: aguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artgenre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aguo.mutation.done = true
	return _node, nil
}
