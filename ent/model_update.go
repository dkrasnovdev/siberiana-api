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
	"github.com/dkrasnovdev/siberiana-api/ent/artifact"
	"github.com/dkrasnovdev/siberiana-api/ent/model"
	"github.com/dkrasnovdev/siberiana-api/ent/petroglyph"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// ModelUpdate is the builder for updating Model entities.
type ModelUpdate struct {
	config
	hooks    []Hook
	mutation *ModelMutation
}

// Where appends a list predicates to the ModelUpdate builder.
func (mu *ModelUpdate) Where(ps ...predicate.Model) *ModelUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedBy sets the "created_by" field.
func (mu *ModelUpdate) SetCreatedBy(s string) *ModelUpdate {
	mu.mutation.SetCreatedBy(s)
	return mu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableCreatedBy(s *string) *ModelUpdate {
	if s != nil {
		mu.SetCreatedBy(*s)
	}
	return mu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (mu *ModelUpdate) ClearCreatedBy() *ModelUpdate {
	mu.mutation.ClearCreatedBy()
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *ModelUpdate) SetUpdatedAt(t time.Time) *ModelUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetUpdatedBy sets the "updated_by" field.
func (mu *ModelUpdate) SetUpdatedBy(s string) *ModelUpdate {
	mu.mutation.SetUpdatedBy(s)
	return mu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableUpdatedBy(s *string) *ModelUpdate {
	if s != nil {
		mu.SetUpdatedBy(*s)
	}
	return mu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (mu *ModelUpdate) ClearUpdatedBy() *ModelUpdate {
	mu.mutation.ClearUpdatedBy()
	return mu
}

// SetDisplayName sets the "display_name" field.
func (mu *ModelUpdate) SetDisplayName(s string) *ModelUpdate {
	mu.mutation.SetDisplayName(s)
	return mu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableDisplayName(s *string) *ModelUpdate {
	if s != nil {
		mu.SetDisplayName(*s)
	}
	return mu
}

// ClearDisplayName clears the value of the "display_name" field.
func (mu *ModelUpdate) ClearDisplayName() *ModelUpdate {
	mu.mutation.ClearDisplayName()
	return mu
}

// SetAbbreviation sets the "abbreviation" field.
func (mu *ModelUpdate) SetAbbreviation(s string) *ModelUpdate {
	mu.mutation.SetAbbreviation(s)
	return mu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableAbbreviation(s *string) *ModelUpdate {
	if s != nil {
		mu.SetAbbreviation(*s)
	}
	return mu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (mu *ModelUpdate) ClearAbbreviation() *ModelUpdate {
	mu.mutation.ClearAbbreviation()
	return mu
}

// SetDescription sets the "description" field.
func (mu *ModelUpdate) SetDescription(s string) *ModelUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableDescription(s *string) *ModelUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// ClearDescription clears the value of the "description" field.
func (mu *ModelUpdate) ClearDescription() *ModelUpdate {
	mu.mutation.ClearDescription()
	return mu
}

// SetExternalLink sets the "external_link" field.
func (mu *ModelUpdate) SetExternalLink(s string) *ModelUpdate {
	mu.mutation.SetExternalLink(s)
	return mu
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableExternalLink(s *string) *ModelUpdate {
	if s != nil {
		mu.SetExternalLink(*s)
	}
	return mu
}

// ClearExternalLink clears the value of the "external_link" field.
func (mu *ModelUpdate) ClearExternalLink() *ModelUpdate {
	mu.mutation.ClearExternalLink()
	return mu
}

// SetStatus sets the "status" field.
func (mu *ModelUpdate) SetStatus(m model.Status) *ModelUpdate {
	mu.mutation.SetStatus(m)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableStatus(m *model.Status) *ModelUpdate {
	if m != nil {
		mu.SetStatus(*m)
	}
	return mu
}

// ClearStatus clears the value of the "status" field.
func (mu *ModelUpdate) ClearStatus() *ModelUpdate {
	mu.mutation.ClearStatus()
	return mu
}

// SetFileURL sets the "file_url" field.
func (mu *ModelUpdate) SetFileURL(s string) *ModelUpdate {
	mu.mutation.SetFileURL(s)
	return mu
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (mu *ModelUpdate) AddArtifactIDs(ids ...int) *ModelUpdate {
	mu.mutation.AddArtifactIDs(ids...)
	return mu
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (mu *ModelUpdate) AddArtifacts(a ...*Artifact) *ModelUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mu.AddArtifactIDs(ids...)
}

// AddPetroglyphIDs adds the "petroglyphs" edge to the Petroglyph entity by IDs.
func (mu *ModelUpdate) AddPetroglyphIDs(ids ...int) *ModelUpdate {
	mu.mutation.AddPetroglyphIDs(ids...)
	return mu
}

// AddPetroglyphs adds the "petroglyphs" edges to the Petroglyph entity.
func (mu *ModelUpdate) AddPetroglyphs(p ...*Petroglyph) *ModelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddPetroglyphIDs(ids...)
}

// Mutation returns the ModelMutation object of the builder.
func (mu *ModelUpdate) Mutation() *ModelMutation {
	return mu.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (mu *ModelUpdate) ClearArtifacts() *ModelUpdate {
	mu.mutation.ClearArtifacts()
	return mu
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (mu *ModelUpdate) RemoveArtifactIDs(ids ...int) *ModelUpdate {
	mu.mutation.RemoveArtifactIDs(ids...)
	return mu
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (mu *ModelUpdate) RemoveArtifacts(a ...*Artifact) *ModelUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mu.RemoveArtifactIDs(ids...)
}

// ClearPetroglyphs clears all "petroglyphs" edges to the Petroglyph entity.
func (mu *ModelUpdate) ClearPetroglyphs() *ModelUpdate {
	mu.mutation.ClearPetroglyphs()
	return mu
}

// RemovePetroglyphIDs removes the "petroglyphs" edge to Petroglyph entities by IDs.
func (mu *ModelUpdate) RemovePetroglyphIDs(ids ...int) *ModelUpdate {
	mu.mutation.RemovePetroglyphIDs(ids...)
	return mu
}

// RemovePetroglyphs removes "petroglyphs" edges to Petroglyph entities.
func (mu *ModelUpdate) RemovePetroglyphs(p ...*Petroglyph) *ModelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemovePetroglyphIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *ModelUpdate) Save(ctx context.Context) (int, error) {
	if err := mu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ModelUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ModelUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ModelUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *ModelUpdate) defaults() error {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		if model.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized model.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := model.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mu *ModelUpdate) check() error {
	if v, ok := mu.mutation.Status(); ok {
		if err := model.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Model.status": %w`, err)}
		}
	}
	return nil
}

func (mu *ModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(model.Table, model.Columns, sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedBy(); ok {
		_spec.SetField(model.FieldCreatedBy, field.TypeString, value)
	}
	if mu.mutation.CreatedByCleared() {
		_spec.ClearField(model.FieldCreatedBy, field.TypeString)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(model.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.UpdatedBy(); ok {
		_spec.SetField(model.FieldUpdatedBy, field.TypeString, value)
	}
	if mu.mutation.UpdatedByCleared() {
		_spec.ClearField(model.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := mu.mutation.DisplayName(); ok {
		_spec.SetField(model.FieldDisplayName, field.TypeString, value)
	}
	if mu.mutation.DisplayNameCleared() {
		_spec.ClearField(model.FieldDisplayName, field.TypeString)
	}
	if value, ok := mu.mutation.Abbreviation(); ok {
		_spec.SetField(model.FieldAbbreviation, field.TypeString, value)
	}
	if mu.mutation.AbbreviationCleared() {
		_spec.ClearField(model.FieldAbbreviation, field.TypeString)
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(model.FieldDescription, field.TypeString, value)
	}
	if mu.mutation.DescriptionCleared() {
		_spec.ClearField(model.FieldDescription, field.TypeString)
	}
	if value, ok := mu.mutation.ExternalLink(); ok {
		_spec.SetField(model.FieldExternalLink, field.TypeString, value)
	}
	if mu.mutation.ExternalLinkCleared() {
		_spec.ClearField(model.FieldExternalLink, field.TypeString)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(model.FieldStatus, field.TypeEnum, value)
	}
	if mu.mutation.StatusCleared() {
		_spec.ClearField(model.FieldStatus, field.TypeEnum)
	}
	if value, ok := mu.mutation.FileURL(); ok {
		_spec.SetField(model.FieldFileURL, field.TypeString, value)
	}
	if mu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ArtifactsTable,
			Columns: []string{model.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !mu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ArtifactsTable,
			Columns: []string{model.ArtifactsColumn},
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
	if nodes := mu.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ArtifactsTable,
			Columns: []string{model.ArtifactsColumn},
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
	if mu.mutation.PetroglyphsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.PetroglyphsTable,
			Columns: []string{model.PetroglyphsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(petroglyph.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedPetroglyphsIDs(); len(nodes) > 0 && !mu.mutation.PetroglyphsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.PetroglyphsTable,
			Columns: []string{model.PetroglyphsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(petroglyph.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.PetroglyphsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.PetroglyphsTable,
			Columns: []string{model.PetroglyphsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(petroglyph.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{model.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// ModelUpdateOne is the builder for updating a single Model entity.
type ModelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ModelMutation
}

// SetCreatedBy sets the "created_by" field.
func (muo *ModelUpdateOne) SetCreatedBy(s string) *ModelUpdateOne {
	muo.mutation.SetCreatedBy(s)
	return muo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableCreatedBy(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetCreatedBy(*s)
	}
	return muo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (muo *ModelUpdateOne) ClearCreatedBy() *ModelUpdateOne {
	muo.mutation.ClearCreatedBy()
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *ModelUpdateOne) SetUpdatedAt(t time.Time) *ModelUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetUpdatedBy sets the "updated_by" field.
func (muo *ModelUpdateOne) SetUpdatedBy(s string) *ModelUpdateOne {
	muo.mutation.SetUpdatedBy(s)
	return muo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableUpdatedBy(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetUpdatedBy(*s)
	}
	return muo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (muo *ModelUpdateOne) ClearUpdatedBy() *ModelUpdateOne {
	muo.mutation.ClearUpdatedBy()
	return muo
}

// SetDisplayName sets the "display_name" field.
func (muo *ModelUpdateOne) SetDisplayName(s string) *ModelUpdateOne {
	muo.mutation.SetDisplayName(s)
	return muo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableDisplayName(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetDisplayName(*s)
	}
	return muo
}

// ClearDisplayName clears the value of the "display_name" field.
func (muo *ModelUpdateOne) ClearDisplayName() *ModelUpdateOne {
	muo.mutation.ClearDisplayName()
	return muo
}

// SetAbbreviation sets the "abbreviation" field.
func (muo *ModelUpdateOne) SetAbbreviation(s string) *ModelUpdateOne {
	muo.mutation.SetAbbreviation(s)
	return muo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableAbbreviation(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetAbbreviation(*s)
	}
	return muo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (muo *ModelUpdateOne) ClearAbbreviation() *ModelUpdateOne {
	muo.mutation.ClearAbbreviation()
	return muo
}

// SetDescription sets the "description" field.
func (muo *ModelUpdateOne) SetDescription(s string) *ModelUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableDescription(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// ClearDescription clears the value of the "description" field.
func (muo *ModelUpdateOne) ClearDescription() *ModelUpdateOne {
	muo.mutation.ClearDescription()
	return muo
}

// SetExternalLink sets the "external_link" field.
func (muo *ModelUpdateOne) SetExternalLink(s string) *ModelUpdateOne {
	muo.mutation.SetExternalLink(s)
	return muo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableExternalLink(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetExternalLink(*s)
	}
	return muo
}

// ClearExternalLink clears the value of the "external_link" field.
func (muo *ModelUpdateOne) ClearExternalLink() *ModelUpdateOne {
	muo.mutation.ClearExternalLink()
	return muo
}

// SetStatus sets the "status" field.
func (muo *ModelUpdateOne) SetStatus(m model.Status) *ModelUpdateOne {
	muo.mutation.SetStatus(m)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableStatus(m *model.Status) *ModelUpdateOne {
	if m != nil {
		muo.SetStatus(*m)
	}
	return muo
}

// ClearStatus clears the value of the "status" field.
func (muo *ModelUpdateOne) ClearStatus() *ModelUpdateOne {
	muo.mutation.ClearStatus()
	return muo
}

// SetFileURL sets the "file_url" field.
func (muo *ModelUpdateOne) SetFileURL(s string) *ModelUpdateOne {
	muo.mutation.SetFileURL(s)
	return muo
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (muo *ModelUpdateOne) AddArtifactIDs(ids ...int) *ModelUpdateOne {
	muo.mutation.AddArtifactIDs(ids...)
	return muo
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (muo *ModelUpdateOne) AddArtifacts(a ...*Artifact) *ModelUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return muo.AddArtifactIDs(ids...)
}

// AddPetroglyphIDs adds the "petroglyphs" edge to the Petroglyph entity by IDs.
func (muo *ModelUpdateOne) AddPetroglyphIDs(ids ...int) *ModelUpdateOne {
	muo.mutation.AddPetroglyphIDs(ids...)
	return muo
}

// AddPetroglyphs adds the "petroglyphs" edges to the Petroglyph entity.
func (muo *ModelUpdateOne) AddPetroglyphs(p ...*Petroglyph) *ModelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddPetroglyphIDs(ids...)
}

// Mutation returns the ModelMutation object of the builder.
func (muo *ModelUpdateOne) Mutation() *ModelMutation {
	return muo.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (muo *ModelUpdateOne) ClearArtifacts() *ModelUpdateOne {
	muo.mutation.ClearArtifacts()
	return muo
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (muo *ModelUpdateOne) RemoveArtifactIDs(ids ...int) *ModelUpdateOne {
	muo.mutation.RemoveArtifactIDs(ids...)
	return muo
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (muo *ModelUpdateOne) RemoveArtifacts(a ...*Artifact) *ModelUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return muo.RemoveArtifactIDs(ids...)
}

// ClearPetroglyphs clears all "petroglyphs" edges to the Petroglyph entity.
func (muo *ModelUpdateOne) ClearPetroglyphs() *ModelUpdateOne {
	muo.mutation.ClearPetroglyphs()
	return muo
}

// RemovePetroglyphIDs removes the "petroglyphs" edge to Petroglyph entities by IDs.
func (muo *ModelUpdateOne) RemovePetroglyphIDs(ids ...int) *ModelUpdateOne {
	muo.mutation.RemovePetroglyphIDs(ids...)
	return muo
}

// RemovePetroglyphs removes "petroglyphs" edges to Petroglyph entities.
func (muo *ModelUpdateOne) RemovePetroglyphs(p ...*Petroglyph) *ModelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemovePetroglyphIDs(ids...)
}

// Where appends a list predicates to the ModelUpdate builder.
func (muo *ModelUpdateOne) Where(ps ...predicate.Model) *ModelUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *ModelUpdateOne) Select(field string, fields ...string) *ModelUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Model entity.
func (muo *ModelUpdateOne) Save(ctx context.Context) (*Model, error) {
	if err := muo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ModelUpdateOne) SaveX(ctx context.Context) *Model {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *ModelUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ModelUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *ModelUpdateOne) defaults() error {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		if model.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized model.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := model.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (muo *ModelUpdateOne) check() error {
	if v, ok := muo.mutation.Status(); ok {
		if err := model.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Model.status": %w`, err)}
		}
	}
	return nil
}

func (muo *ModelUpdateOne) sqlSave(ctx context.Context) (_node *Model, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(model.Table, model.Columns, sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Model.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, model.FieldID)
		for _, f := range fields {
			if !model.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != model.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.CreatedBy(); ok {
		_spec.SetField(model.FieldCreatedBy, field.TypeString, value)
	}
	if muo.mutation.CreatedByCleared() {
		_spec.ClearField(model.FieldCreatedBy, field.TypeString)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(model.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.UpdatedBy(); ok {
		_spec.SetField(model.FieldUpdatedBy, field.TypeString, value)
	}
	if muo.mutation.UpdatedByCleared() {
		_spec.ClearField(model.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := muo.mutation.DisplayName(); ok {
		_spec.SetField(model.FieldDisplayName, field.TypeString, value)
	}
	if muo.mutation.DisplayNameCleared() {
		_spec.ClearField(model.FieldDisplayName, field.TypeString)
	}
	if value, ok := muo.mutation.Abbreviation(); ok {
		_spec.SetField(model.FieldAbbreviation, field.TypeString, value)
	}
	if muo.mutation.AbbreviationCleared() {
		_spec.ClearField(model.FieldAbbreviation, field.TypeString)
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(model.FieldDescription, field.TypeString, value)
	}
	if muo.mutation.DescriptionCleared() {
		_spec.ClearField(model.FieldDescription, field.TypeString)
	}
	if value, ok := muo.mutation.ExternalLink(); ok {
		_spec.SetField(model.FieldExternalLink, field.TypeString, value)
	}
	if muo.mutation.ExternalLinkCleared() {
		_spec.ClearField(model.FieldExternalLink, field.TypeString)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(model.FieldStatus, field.TypeEnum, value)
	}
	if muo.mutation.StatusCleared() {
		_spec.ClearField(model.FieldStatus, field.TypeEnum)
	}
	if value, ok := muo.mutation.FileURL(); ok {
		_spec.SetField(model.FieldFileURL, field.TypeString, value)
	}
	if muo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ArtifactsTable,
			Columns: []string{model.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !muo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ArtifactsTable,
			Columns: []string{model.ArtifactsColumn},
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
	if nodes := muo.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ArtifactsTable,
			Columns: []string{model.ArtifactsColumn},
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
	if muo.mutation.PetroglyphsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.PetroglyphsTable,
			Columns: []string{model.PetroglyphsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(petroglyph.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedPetroglyphsIDs(); len(nodes) > 0 && !muo.mutation.PetroglyphsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.PetroglyphsTable,
			Columns: []string{model.PetroglyphsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(petroglyph.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.PetroglyphsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.PetroglyphsTable,
			Columns: []string{model.PetroglyphsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(petroglyph.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Model{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{model.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
