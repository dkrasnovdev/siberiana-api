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
	"github.com/dkrasnovdev/siberiana-api/ent/monument"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
	"github.com/dkrasnovdev/siberiana-api/ent/set"
)

// SetUpdate is the builder for updating Set entities.
type SetUpdate struct {
	config
	hooks    []Hook
	mutation *SetMutation
}

// Where appends a list predicates to the SetUpdate builder.
func (su *SetUpdate) Where(ps ...predicate.Set) *SetUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCreatedBy sets the "created_by" field.
func (su *SetUpdate) SetCreatedBy(s string) *SetUpdate {
	su.mutation.SetCreatedBy(s)
	return su
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (su *SetUpdate) SetNillableCreatedBy(s *string) *SetUpdate {
	if s != nil {
		su.SetCreatedBy(*s)
	}
	return su
}

// ClearCreatedBy clears the value of the "created_by" field.
func (su *SetUpdate) ClearCreatedBy() *SetUpdate {
	su.mutation.ClearCreatedBy()
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SetUpdate) SetUpdatedAt(t time.Time) *SetUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetUpdatedBy sets the "updated_by" field.
func (su *SetUpdate) SetUpdatedBy(s string) *SetUpdate {
	su.mutation.SetUpdatedBy(s)
	return su
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (su *SetUpdate) SetNillableUpdatedBy(s *string) *SetUpdate {
	if s != nil {
		su.SetUpdatedBy(*s)
	}
	return su
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (su *SetUpdate) ClearUpdatedBy() *SetUpdate {
	su.mutation.ClearUpdatedBy()
	return su
}

// SetDisplayName sets the "display_name" field.
func (su *SetUpdate) SetDisplayName(s string) *SetUpdate {
	su.mutation.SetDisplayName(s)
	return su
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (su *SetUpdate) SetNillableDisplayName(s *string) *SetUpdate {
	if s != nil {
		su.SetDisplayName(*s)
	}
	return su
}

// ClearDisplayName clears the value of the "display_name" field.
func (su *SetUpdate) ClearDisplayName() *SetUpdate {
	su.mutation.ClearDisplayName()
	return su
}

// SetAbbreviation sets the "abbreviation" field.
func (su *SetUpdate) SetAbbreviation(s string) *SetUpdate {
	su.mutation.SetAbbreviation(s)
	return su
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (su *SetUpdate) SetNillableAbbreviation(s *string) *SetUpdate {
	if s != nil {
		su.SetAbbreviation(*s)
	}
	return su
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (su *SetUpdate) ClearAbbreviation() *SetUpdate {
	su.mutation.ClearAbbreviation()
	return su
}

// SetDescription sets the "description" field.
func (su *SetUpdate) SetDescription(s string) *SetUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *SetUpdate) SetNillableDescription(s *string) *SetUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of the "description" field.
func (su *SetUpdate) ClearDescription() *SetUpdate {
	su.mutation.ClearDescription()
	return su
}

// SetExternalLink sets the "external_link" field.
func (su *SetUpdate) SetExternalLink(s string) *SetUpdate {
	su.mutation.SetExternalLink(s)
	return su
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (su *SetUpdate) SetNillableExternalLink(s *string) *SetUpdate {
	if s != nil {
		su.SetExternalLink(*s)
	}
	return su
}

// ClearExternalLink clears the value of the "external_link" field.
func (su *SetUpdate) ClearExternalLink() *SetUpdate {
	su.mutation.ClearExternalLink()
	return su
}

// SetSlug sets the "slug" field.
func (su *SetUpdate) SetSlug(s string) *SetUpdate {
	su.mutation.SetSlug(s)
	return su
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (su *SetUpdate) SetNillableSlug(s *string) *SetUpdate {
	if s != nil {
		su.SetSlug(*s)
	}
	return su
}

// ClearSlug clears the value of the "slug" field.
func (su *SetUpdate) ClearSlug() *SetUpdate {
	su.mutation.ClearSlug()
	return su
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (su *SetUpdate) AddArtifactIDs(ids ...int) *SetUpdate {
	su.mutation.AddArtifactIDs(ids...)
	return su
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (su *SetUpdate) AddArtifacts(a ...*Artifact) *SetUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.AddArtifactIDs(ids...)
}

// AddMonumentIDs adds the "monuments" edge to the Monument entity by IDs.
func (su *SetUpdate) AddMonumentIDs(ids ...int) *SetUpdate {
	su.mutation.AddMonumentIDs(ids...)
	return su
}

// AddMonuments adds the "monuments" edges to the Monument entity.
func (su *SetUpdate) AddMonuments(m ...*Monument) *SetUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return su.AddMonumentIDs(ids...)
}

// Mutation returns the SetMutation object of the builder.
func (su *SetUpdate) Mutation() *SetMutation {
	return su.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (su *SetUpdate) ClearArtifacts() *SetUpdate {
	su.mutation.ClearArtifacts()
	return su
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (su *SetUpdate) RemoveArtifactIDs(ids ...int) *SetUpdate {
	su.mutation.RemoveArtifactIDs(ids...)
	return su
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (su *SetUpdate) RemoveArtifacts(a ...*Artifact) *SetUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.RemoveArtifactIDs(ids...)
}

// ClearMonuments clears all "monuments" edges to the Monument entity.
func (su *SetUpdate) ClearMonuments() *SetUpdate {
	su.mutation.ClearMonuments()
	return su
}

// RemoveMonumentIDs removes the "monuments" edge to Monument entities by IDs.
func (su *SetUpdate) RemoveMonumentIDs(ids ...int) *SetUpdate {
	su.mutation.RemoveMonumentIDs(ids...)
	return su
}

// RemoveMonuments removes "monuments" edges to Monument entities.
func (su *SetUpdate) RemoveMonuments(m ...*Monument) *SetUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return su.RemoveMonumentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SetUpdate) Save(ctx context.Context) (int, error) {
	if err := su.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SetUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SetUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SetUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SetUpdate) defaults() error {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		if set.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized set.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := set.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (su *SetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(set.Table, set.Columns, sqlgraph.NewFieldSpec(set.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.CreatedBy(); ok {
		_spec.SetField(set.FieldCreatedBy, field.TypeString, value)
	}
	if su.mutation.CreatedByCleared() {
		_spec.ClearField(set.FieldCreatedBy, field.TypeString)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(set.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedBy(); ok {
		_spec.SetField(set.FieldUpdatedBy, field.TypeString, value)
	}
	if su.mutation.UpdatedByCleared() {
		_spec.ClearField(set.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := su.mutation.DisplayName(); ok {
		_spec.SetField(set.FieldDisplayName, field.TypeString, value)
	}
	if su.mutation.DisplayNameCleared() {
		_spec.ClearField(set.FieldDisplayName, field.TypeString)
	}
	if value, ok := su.mutation.Abbreviation(); ok {
		_spec.SetField(set.FieldAbbreviation, field.TypeString, value)
	}
	if su.mutation.AbbreviationCleared() {
		_spec.ClearField(set.FieldAbbreviation, field.TypeString)
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.SetField(set.FieldDescription, field.TypeString, value)
	}
	if su.mutation.DescriptionCleared() {
		_spec.ClearField(set.FieldDescription, field.TypeString)
	}
	if value, ok := su.mutation.ExternalLink(); ok {
		_spec.SetField(set.FieldExternalLink, field.TypeString, value)
	}
	if su.mutation.ExternalLinkCleared() {
		_spec.ClearField(set.FieldExternalLink, field.TypeString)
	}
	if value, ok := su.mutation.Slug(); ok {
		_spec.SetField(set.FieldSlug, field.TypeString, value)
	}
	if su.mutation.SlugCleared() {
		_spec.ClearField(set.FieldSlug, field.TypeString)
	}
	if su.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   set.ArtifactsTable,
			Columns: []string{set.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !su.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   set.ArtifactsTable,
			Columns: []string{set.ArtifactsColumn},
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
	if nodes := su.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   set.ArtifactsTable,
			Columns: []string{set.ArtifactsColumn},
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
	if su.mutation.MonumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   set.MonumentsTable,
			Columns: set.MonumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedMonumentsIDs(); len(nodes) > 0 && !su.mutation.MonumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   set.MonumentsTable,
			Columns: set.MonumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.MonumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   set.MonumentsTable,
			Columns: set.MonumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{set.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SetUpdateOne is the builder for updating a single Set entity.
type SetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SetMutation
}

// SetCreatedBy sets the "created_by" field.
func (suo *SetUpdateOne) SetCreatedBy(s string) *SetUpdateOne {
	suo.mutation.SetCreatedBy(s)
	return suo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suo *SetUpdateOne) SetNillableCreatedBy(s *string) *SetUpdateOne {
	if s != nil {
		suo.SetCreatedBy(*s)
	}
	return suo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (suo *SetUpdateOne) ClearCreatedBy() *SetUpdateOne {
	suo.mutation.ClearCreatedBy()
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SetUpdateOne) SetUpdatedAt(t time.Time) *SetUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetUpdatedBy sets the "updated_by" field.
func (suo *SetUpdateOne) SetUpdatedBy(s string) *SetUpdateOne {
	suo.mutation.SetUpdatedBy(s)
	return suo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (suo *SetUpdateOne) SetNillableUpdatedBy(s *string) *SetUpdateOne {
	if s != nil {
		suo.SetUpdatedBy(*s)
	}
	return suo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (suo *SetUpdateOne) ClearUpdatedBy() *SetUpdateOne {
	suo.mutation.ClearUpdatedBy()
	return suo
}

// SetDisplayName sets the "display_name" field.
func (suo *SetUpdateOne) SetDisplayName(s string) *SetUpdateOne {
	suo.mutation.SetDisplayName(s)
	return suo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (suo *SetUpdateOne) SetNillableDisplayName(s *string) *SetUpdateOne {
	if s != nil {
		suo.SetDisplayName(*s)
	}
	return suo
}

// ClearDisplayName clears the value of the "display_name" field.
func (suo *SetUpdateOne) ClearDisplayName() *SetUpdateOne {
	suo.mutation.ClearDisplayName()
	return suo
}

// SetAbbreviation sets the "abbreviation" field.
func (suo *SetUpdateOne) SetAbbreviation(s string) *SetUpdateOne {
	suo.mutation.SetAbbreviation(s)
	return suo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (suo *SetUpdateOne) SetNillableAbbreviation(s *string) *SetUpdateOne {
	if s != nil {
		suo.SetAbbreviation(*s)
	}
	return suo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (suo *SetUpdateOne) ClearAbbreviation() *SetUpdateOne {
	suo.mutation.ClearAbbreviation()
	return suo
}

// SetDescription sets the "description" field.
func (suo *SetUpdateOne) SetDescription(s string) *SetUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *SetUpdateOne) SetNillableDescription(s *string) *SetUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of the "description" field.
func (suo *SetUpdateOne) ClearDescription() *SetUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// SetExternalLink sets the "external_link" field.
func (suo *SetUpdateOne) SetExternalLink(s string) *SetUpdateOne {
	suo.mutation.SetExternalLink(s)
	return suo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (suo *SetUpdateOne) SetNillableExternalLink(s *string) *SetUpdateOne {
	if s != nil {
		suo.SetExternalLink(*s)
	}
	return suo
}

// ClearExternalLink clears the value of the "external_link" field.
func (suo *SetUpdateOne) ClearExternalLink() *SetUpdateOne {
	suo.mutation.ClearExternalLink()
	return suo
}

// SetSlug sets the "slug" field.
func (suo *SetUpdateOne) SetSlug(s string) *SetUpdateOne {
	suo.mutation.SetSlug(s)
	return suo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (suo *SetUpdateOne) SetNillableSlug(s *string) *SetUpdateOne {
	if s != nil {
		suo.SetSlug(*s)
	}
	return suo
}

// ClearSlug clears the value of the "slug" field.
func (suo *SetUpdateOne) ClearSlug() *SetUpdateOne {
	suo.mutation.ClearSlug()
	return suo
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (suo *SetUpdateOne) AddArtifactIDs(ids ...int) *SetUpdateOne {
	suo.mutation.AddArtifactIDs(ids...)
	return suo
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (suo *SetUpdateOne) AddArtifacts(a ...*Artifact) *SetUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.AddArtifactIDs(ids...)
}

// AddMonumentIDs adds the "monuments" edge to the Monument entity by IDs.
func (suo *SetUpdateOne) AddMonumentIDs(ids ...int) *SetUpdateOne {
	suo.mutation.AddMonumentIDs(ids...)
	return suo
}

// AddMonuments adds the "monuments" edges to the Monument entity.
func (suo *SetUpdateOne) AddMonuments(m ...*Monument) *SetUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return suo.AddMonumentIDs(ids...)
}

// Mutation returns the SetMutation object of the builder.
func (suo *SetUpdateOne) Mutation() *SetMutation {
	return suo.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (suo *SetUpdateOne) ClearArtifacts() *SetUpdateOne {
	suo.mutation.ClearArtifacts()
	return suo
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (suo *SetUpdateOne) RemoveArtifactIDs(ids ...int) *SetUpdateOne {
	suo.mutation.RemoveArtifactIDs(ids...)
	return suo
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (suo *SetUpdateOne) RemoveArtifacts(a ...*Artifact) *SetUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.RemoveArtifactIDs(ids...)
}

// ClearMonuments clears all "monuments" edges to the Monument entity.
func (suo *SetUpdateOne) ClearMonuments() *SetUpdateOne {
	suo.mutation.ClearMonuments()
	return suo
}

// RemoveMonumentIDs removes the "monuments" edge to Monument entities by IDs.
func (suo *SetUpdateOne) RemoveMonumentIDs(ids ...int) *SetUpdateOne {
	suo.mutation.RemoveMonumentIDs(ids...)
	return suo
}

// RemoveMonuments removes "monuments" edges to Monument entities.
func (suo *SetUpdateOne) RemoveMonuments(m ...*Monument) *SetUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return suo.RemoveMonumentIDs(ids...)
}

// Where appends a list predicates to the SetUpdate builder.
func (suo *SetUpdateOne) Where(ps ...predicate.Set) *SetUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SetUpdateOne) Select(field string, fields ...string) *SetUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Set entity.
func (suo *SetUpdateOne) Save(ctx context.Context) (*Set, error) {
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SetUpdateOne) SaveX(ctx context.Context) *Set {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SetUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SetUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SetUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		if set.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized set.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := set.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (suo *SetUpdateOne) sqlSave(ctx context.Context) (_node *Set, err error) {
	_spec := sqlgraph.NewUpdateSpec(set.Table, set.Columns, sqlgraph.NewFieldSpec(set.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Set.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, set.FieldID)
		for _, f := range fields {
			if !set.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != set.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.CreatedBy(); ok {
		_spec.SetField(set.FieldCreatedBy, field.TypeString, value)
	}
	if suo.mutation.CreatedByCleared() {
		_spec.ClearField(set.FieldCreatedBy, field.TypeString)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(set.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedBy(); ok {
		_spec.SetField(set.FieldUpdatedBy, field.TypeString, value)
	}
	if suo.mutation.UpdatedByCleared() {
		_spec.ClearField(set.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := suo.mutation.DisplayName(); ok {
		_spec.SetField(set.FieldDisplayName, field.TypeString, value)
	}
	if suo.mutation.DisplayNameCleared() {
		_spec.ClearField(set.FieldDisplayName, field.TypeString)
	}
	if value, ok := suo.mutation.Abbreviation(); ok {
		_spec.SetField(set.FieldAbbreviation, field.TypeString, value)
	}
	if suo.mutation.AbbreviationCleared() {
		_spec.ClearField(set.FieldAbbreviation, field.TypeString)
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.SetField(set.FieldDescription, field.TypeString, value)
	}
	if suo.mutation.DescriptionCleared() {
		_spec.ClearField(set.FieldDescription, field.TypeString)
	}
	if value, ok := suo.mutation.ExternalLink(); ok {
		_spec.SetField(set.FieldExternalLink, field.TypeString, value)
	}
	if suo.mutation.ExternalLinkCleared() {
		_spec.ClearField(set.FieldExternalLink, field.TypeString)
	}
	if value, ok := suo.mutation.Slug(); ok {
		_spec.SetField(set.FieldSlug, field.TypeString, value)
	}
	if suo.mutation.SlugCleared() {
		_spec.ClearField(set.FieldSlug, field.TypeString)
	}
	if suo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   set.ArtifactsTable,
			Columns: []string{set.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !suo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   set.ArtifactsTable,
			Columns: []string{set.ArtifactsColumn},
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
	if nodes := suo.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   set.ArtifactsTable,
			Columns: []string{set.ArtifactsColumn},
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
	if suo.mutation.MonumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   set.MonumentsTable,
			Columns: set.MonumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedMonumentsIDs(); len(nodes) > 0 && !suo.mutation.MonumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   set.MonumentsTable,
			Columns: set.MonumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.MonumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   set.MonumentsTable,
			Columns: set.MonumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Set{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{set.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
