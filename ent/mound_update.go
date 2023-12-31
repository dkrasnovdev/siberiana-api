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
	"github.com/dkrasnovdev/siberiana-api/ent/mound"
	"github.com/dkrasnovdev/siberiana-api/ent/petroglyph"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
	"github.com/dkrasnovdev/siberiana-api/ent/visit"
)

// MoundUpdate is the builder for updating Mound entities.
type MoundUpdate struct {
	config
	hooks    []Hook
	mutation *MoundMutation
}

// Where appends a list predicates to the MoundUpdate builder.
func (mu *MoundUpdate) Where(ps ...predicate.Mound) *MoundUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedBy sets the "created_by" field.
func (mu *MoundUpdate) SetCreatedBy(s string) *MoundUpdate {
	mu.mutation.SetCreatedBy(s)
	return mu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mu *MoundUpdate) SetNillableCreatedBy(s *string) *MoundUpdate {
	if s != nil {
		mu.SetCreatedBy(*s)
	}
	return mu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (mu *MoundUpdate) ClearCreatedBy() *MoundUpdate {
	mu.mutation.ClearCreatedBy()
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MoundUpdate) SetUpdatedAt(t time.Time) *MoundUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetUpdatedBy sets the "updated_by" field.
func (mu *MoundUpdate) SetUpdatedBy(s string) *MoundUpdate {
	mu.mutation.SetUpdatedBy(s)
	return mu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mu *MoundUpdate) SetNillableUpdatedBy(s *string) *MoundUpdate {
	if s != nil {
		mu.SetUpdatedBy(*s)
	}
	return mu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (mu *MoundUpdate) ClearUpdatedBy() *MoundUpdate {
	mu.mutation.ClearUpdatedBy()
	return mu
}

// SetDisplayName sets the "display_name" field.
func (mu *MoundUpdate) SetDisplayName(s string) *MoundUpdate {
	mu.mutation.SetDisplayName(s)
	return mu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (mu *MoundUpdate) SetNillableDisplayName(s *string) *MoundUpdate {
	if s != nil {
		mu.SetDisplayName(*s)
	}
	return mu
}

// ClearDisplayName clears the value of the "display_name" field.
func (mu *MoundUpdate) ClearDisplayName() *MoundUpdate {
	mu.mutation.ClearDisplayName()
	return mu
}

// SetAbbreviation sets the "abbreviation" field.
func (mu *MoundUpdate) SetAbbreviation(s string) *MoundUpdate {
	mu.mutation.SetAbbreviation(s)
	return mu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (mu *MoundUpdate) SetNillableAbbreviation(s *string) *MoundUpdate {
	if s != nil {
		mu.SetAbbreviation(*s)
	}
	return mu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (mu *MoundUpdate) ClearAbbreviation() *MoundUpdate {
	mu.mutation.ClearAbbreviation()
	return mu
}

// SetDescription sets the "description" field.
func (mu *MoundUpdate) SetDescription(s string) *MoundUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MoundUpdate) SetNillableDescription(s *string) *MoundUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// ClearDescription clears the value of the "description" field.
func (mu *MoundUpdate) ClearDescription() *MoundUpdate {
	mu.mutation.ClearDescription()
	return mu
}

// SetExternalLink sets the "external_link" field.
func (mu *MoundUpdate) SetExternalLink(s string) *MoundUpdate {
	mu.mutation.SetExternalLink(s)
	return mu
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (mu *MoundUpdate) SetNillableExternalLink(s *string) *MoundUpdate {
	if s != nil {
		mu.SetExternalLink(*s)
	}
	return mu
}

// ClearExternalLink clears the value of the "external_link" field.
func (mu *MoundUpdate) ClearExternalLink() *MoundUpdate {
	mu.mutation.ClearExternalLink()
	return mu
}

// SetNumber sets the "number" field.
func (mu *MoundUpdate) SetNumber(s string) *MoundUpdate {
	mu.mutation.SetNumber(s)
	return mu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (mu *MoundUpdate) SetNillableNumber(s *string) *MoundUpdate {
	if s != nil {
		mu.SetNumber(*s)
	}
	return mu
}

// ClearNumber clears the value of the "number" field.
func (mu *MoundUpdate) ClearNumber() *MoundUpdate {
	mu.mutation.ClearNumber()
	return mu
}

// AddPetroglyphIDs adds the "petroglyphs" edge to the Petroglyph entity by IDs.
func (mu *MoundUpdate) AddPetroglyphIDs(ids ...int) *MoundUpdate {
	mu.mutation.AddPetroglyphIDs(ids...)
	return mu
}

// AddPetroglyphs adds the "petroglyphs" edges to the Petroglyph entity.
func (mu *MoundUpdate) AddPetroglyphs(p ...*Petroglyph) *MoundUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddPetroglyphIDs(ids...)
}

// AddVisitIDs adds the "visits" edge to the Visit entity by IDs.
func (mu *MoundUpdate) AddVisitIDs(ids ...int) *MoundUpdate {
	mu.mutation.AddVisitIDs(ids...)
	return mu
}

// AddVisits adds the "visits" edges to the Visit entity.
func (mu *MoundUpdate) AddVisits(v ...*Visit) *MoundUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return mu.AddVisitIDs(ids...)
}

// Mutation returns the MoundMutation object of the builder.
func (mu *MoundUpdate) Mutation() *MoundMutation {
	return mu.mutation
}

// ClearPetroglyphs clears all "petroglyphs" edges to the Petroglyph entity.
func (mu *MoundUpdate) ClearPetroglyphs() *MoundUpdate {
	mu.mutation.ClearPetroglyphs()
	return mu
}

// RemovePetroglyphIDs removes the "petroglyphs" edge to Petroglyph entities by IDs.
func (mu *MoundUpdate) RemovePetroglyphIDs(ids ...int) *MoundUpdate {
	mu.mutation.RemovePetroglyphIDs(ids...)
	return mu
}

// RemovePetroglyphs removes "petroglyphs" edges to Petroglyph entities.
func (mu *MoundUpdate) RemovePetroglyphs(p ...*Petroglyph) *MoundUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemovePetroglyphIDs(ids...)
}

// ClearVisits clears all "visits" edges to the Visit entity.
func (mu *MoundUpdate) ClearVisits() *MoundUpdate {
	mu.mutation.ClearVisits()
	return mu
}

// RemoveVisitIDs removes the "visits" edge to Visit entities by IDs.
func (mu *MoundUpdate) RemoveVisitIDs(ids ...int) *MoundUpdate {
	mu.mutation.RemoveVisitIDs(ids...)
	return mu
}

// RemoveVisits removes "visits" edges to Visit entities.
func (mu *MoundUpdate) RemoveVisits(v ...*Visit) *MoundUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return mu.RemoveVisitIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MoundUpdate) Save(ctx context.Context) (int, error) {
	if err := mu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MoundUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MoundUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MoundUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MoundUpdate) defaults() error {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		if mound.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized mound.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := mound.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (mu *MoundUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mound.Table, mound.Columns, sqlgraph.NewFieldSpec(mound.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedBy(); ok {
		_spec.SetField(mound.FieldCreatedBy, field.TypeString, value)
	}
	if mu.mutation.CreatedByCleared() {
		_spec.ClearField(mound.FieldCreatedBy, field.TypeString)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(mound.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.UpdatedBy(); ok {
		_spec.SetField(mound.FieldUpdatedBy, field.TypeString, value)
	}
	if mu.mutation.UpdatedByCleared() {
		_spec.ClearField(mound.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := mu.mutation.DisplayName(); ok {
		_spec.SetField(mound.FieldDisplayName, field.TypeString, value)
	}
	if mu.mutation.DisplayNameCleared() {
		_spec.ClearField(mound.FieldDisplayName, field.TypeString)
	}
	if value, ok := mu.mutation.Abbreviation(); ok {
		_spec.SetField(mound.FieldAbbreviation, field.TypeString, value)
	}
	if mu.mutation.AbbreviationCleared() {
		_spec.ClearField(mound.FieldAbbreviation, field.TypeString)
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(mound.FieldDescription, field.TypeString, value)
	}
	if mu.mutation.DescriptionCleared() {
		_spec.ClearField(mound.FieldDescription, field.TypeString)
	}
	if value, ok := mu.mutation.ExternalLink(); ok {
		_spec.SetField(mound.FieldExternalLink, field.TypeString, value)
	}
	if mu.mutation.ExternalLinkCleared() {
		_spec.ClearField(mound.FieldExternalLink, field.TypeString)
	}
	if value, ok := mu.mutation.Number(); ok {
		_spec.SetField(mound.FieldNumber, field.TypeString, value)
	}
	if mu.mutation.NumberCleared() {
		_spec.ClearField(mound.FieldNumber, field.TypeString)
	}
	if mu.mutation.PetroglyphsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mound.PetroglyphsTable,
			Columns: []string{mound.PetroglyphsColumn},
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
			Table:   mound.PetroglyphsTable,
			Columns: []string{mound.PetroglyphsColumn},
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
			Table:   mound.PetroglyphsTable,
			Columns: []string{mound.PetroglyphsColumn},
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
	if mu.mutation.VisitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   mound.VisitsTable,
			Columns: mound.VisitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedVisitsIDs(); len(nodes) > 0 && !mu.mutation.VisitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   mound.VisitsTable,
			Columns: mound.VisitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.VisitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   mound.VisitsTable,
			Columns: mound.VisitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mound.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MoundUpdateOne is the builder for updating a single Mound entity.
type MoundUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MoundMutation
}

// SetCreatedBy sets the "created_by" field.
func (muo *MoundUpdateOne) SetCreatedBy(s string) *MoundUpdateOne {
	muo.mutation.SetCreatedBy(s)
	return muo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (muo *MoundUpdateOne) SetNillableCreatedBy(s *string) *MoundUpdateOne {
	if s != nil {
		muo.SetCreatedBy(*s)
	}
	return muo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (muo *MoundUpdateOne) ClearCreatedBy() *MoundUpdateOne {
	muo.mutation.ClearCreatedBy()
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MoundUpdateOne) SetUpdatedAt(t time.Time) *MoundUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetUpdatedBy sets the "updated_by" field.
func (muo *MoundUpdateOne) SetUpdatedBy(s string) *MoundUpdateOne {
	muo.mutation.SetUpdatedBy(s)
	return muo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (muo *MoundUpdateOne) SetNillableUpdatedBy(s *string) *MoundUpdateOne {
	if s != nil {
		muo.SetUpdatedBy(*s)
	}
	return muo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (muo *MoundUpdateOne) ClearUpdatedBy() *MoundUpdateOne {
	muo.mutation.ClearUpdatedBy()
	return muo
}

// SetDisplayName sets the "display_name" field.
func (muo *MoundUpdateOne) SetDisplayName(s string) *MoundUpdateOne {
	muo.mutation.SetDisplayName(s)
	return muo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (muo *MoundUpdateOne) SetNillableDisplayName(s *string) *MoundUpdateOne {
	if s != nil {
		muo.SetDisplayName(*s)
	}
	return muo
}

// ClearDisplayName clears the value of the "display_name" field.
func (muo *MoundUpdateOne) ClearDisplayName() *MoundUpdateOne {
	muo.mutation.ClearDisplayName()
	return muo
}

// SetAbbreviation sets the "abbreviation" field.
func (muo *MoundUpdateOne) SetAbbreviation(s string) *MoundUpdateOne {
	muo.mutation.SetAbbreviation(s)
	return muo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (muo *MoundUpdateOne) SetNillableAbbreviation(s *string) *MoundUpdateOne {
	if s != nil {
		muo.SetAbbreviation(*s)
	}
	return muo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (muo *MoundUpdateOne) ClearAbbreviation() *MoundUpdateOne {
	muo.mutation.ClearAbbreviation()
	return muo
}

// SetDescription sets the "description" field.
func (muo *MoundUpdateOne) SetDescription(s string) *MoundUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MoundUpdateOne) SetNillableDescription(s *string) *MoundUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// ClearDescription clears the value of the "description" field.
func (muo *MoundUpdateOne) ClearDescription() *MoundUpdateOne {
	muo.mutation.ClearDescription()
	return muo
}

// SetExternalLink sets the "external_link" field.
func (muo *MoundUpdateOne) SetExternalLink(s string) *MoundUpdateOne {
	muo.mutation.SetExternalLink(s)
	return muo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (muo *MoundUpdateOne) SetNillableExternalLink(s *string) *MoundUpdateOne {
	if s != nil {
		muo.SetExternalLink(*s)
	}
	return muo
}

// ClearExternalLink clears the value of the "external_link" field.
func (muo *MoundUpdateOne) ClearExternalLink() *MoundUpdateOne {
	muo.mutation.ClearExternalLink()
	return muo
}

// SetNumber sets the "number" field.
func (muo *MoundUpdateOne) SetNumber(s string) *MoundUpdateOne {
	muo.mutation.SetNumber(s)
	return muo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (muo *MoundUpdateOne) SetNillableNumber(s *string) *MoundUpdateOne {
	if s != nil {
		muo.SetNumber(*s)
	}
	return muo
}

// ClearNumber clears the value of the "number" field.
func (muo *MoundUpdateOne) ClearNumber() *MoundUpdateOne {
	muo.mutation.ClearNumber()
	return muo
}

// AddPetroglyphIDs adds the "petroglyphs" edge to the Petroglyph entity by IDs.
func (muo *MoundUpdateOne) AddPetroglyphIDs(ids ...int) *MoundUpdateOne {
	muo.mutation.AddPetroglyphIDs(ids...)
	return muo
}

// AddPetroglyphs adds the "petroglyphs" edges to the Petroglyph entity.
func (muo *MoundUpdateOne) AddPetroglyphs(p ...*Petroglyph) *MoundUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddPetroglyphIDs(ids...)
}

// AddVisitIDs adds the "visits" edge to the Visit entity by IDs.
func (muo *MoundUpdateOne) AddVisitIDs(ids ...int) *MoundUpdateOne {
	muo.mutation.AddVisitIDs(ids...)
	return muo
}

// AddVisits adds the "visits" edges to the Visit entity.
func (muo *MoundUpdateOne) AddVisits(v ...*Visit) *MoundUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return muo.AddVisitIDs(ids...)
}

// Mutation returns the MoundMutation object of the builder.
func (muo *MoundUpdateOne) Mutation() *MoundMutation {
	return muo.mutation
}

// ClearPetroglyphs clears all "petroglyphs" edges to the Petroglyph entity.
func (muo *MoundUpdateOne) ClearPetroglyphs() *MoundUpdateOne {
	muo.mutation.ClearPetroglyphs()
	return muo
}

// RemovePetroglyphIDs removes the "petroglyphs" edge to Petroglyph entities by IDs.
func (muo *MoundUpdateOne) RemovePetroglyphIDs(ids ...int) *MoundUpdateOne {
	muo.mutation.RemovePetroglyphIDs(ids...)
	return muo
}

// RemovePetroglyphs removes "petroglyphs" edges to Petroglyph entities.
func (muo *MoundUpdateOne) RemovePetroglyphs(p ...*Petroglyph) *MoundUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemovePetroglyphIDs(ids...)
}

// ClearVisits clears all "visits" edges to the Visit entity.
func (muo *MoundUpdateOne) ClearVisits() *MoundUpdateOne {
	muo.mutation.ClearVisits()
	return muo
}

// RemoveVisitIDs removes the "visits" edge to Visit entities by IDs.
func (muo *MoundUpdateOne) RemoveVisitIDs(ids ...int) *MoundUpdateOne {
	muo.mutation.RemoveVisitIDs(ids...)
	return muo
}

// RemoveVisits removes "visits" edges to Visit entities.
func (muo *MoundUpdateOne) RemoveVisits(v ...*Visit) *MoundUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return muo.RemoveVisitIDs(ids...)
}

// Where appends a list predicates to the MoundUpdate builder.
func (muo *MoundUpdateOne) Where(ps ...predicate.Mound) *MoundUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MoundUpdateOne) Select(field string, fields ...string) *MoundUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mound entity.
func (muo *MoundUpdateOne) Save(ctx context.Context) (*Mound, error) {
	if err := muo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MoundUpdateOne) SaveX(ctx context.Context) *Mound {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MoundUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MoundUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MoundUpdateOne) defaults() error {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		if mound.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized mound.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := mound.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (muo *MoundUpdateOne) sqlSave(ctx context.Context) (_node *Mound, err error) {
	_spec := sqlgraph.NewUpdateSpec(mound.Table, mound.Columns, sqlgraph.NewFieldSpec(mound.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mound.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mound.FieldID)
		for _, f := range fields {
			if !mound.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mound.FieldID {
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
		_spec.SetField(mound.FieldCreatedBy, field.TypeString, value)
	}
	if muo.mutation.CreatedByCleared() {
		_spec.ClearField(mound.FieldCreatedBy, field.TypeString)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(mound.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.UpdatedBy(); ok {
		_spec.SetField(mound.FieldUpdatedBy, field.TypeString, value)
	}
	if muo.mutation.UpdatedByCleared() {
		_spec.ClearField(mound.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := muo.mutation.DisplayName(); ok {
		_spec.SetField(mound.FieldDisplayName, field.TypeString, value)
	}
	if muo.mutation.DisplayNameCleared() {
		_spec.ClearField(mound.FieldDisplayName, field.TypeString)
	}
	if value, ok := muo.mutation.Abbreviation(); ok {
		_spec.SetField(mound.FieldAbbreviation, field.TypeString, value)
	}
	if muo.mutation.AbbreviationCleared() {
		_spec.ClearField(mound.FieldAbbreviation, field.TypeString)
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(mound.FieldDescription, field.TypeString, value)
	}
	if muo.mutation.DescriptionCleared() {
		_spec.ClearField(mound.FieldDescription, field.TypeString)
	}
	if value, ok := muo.mutation.ExternalLink(); ok {
		_spec.SetField(mound.FieldExternalLink, field.TypeString, value)
	}
	if muo.mutation.ExternalLinkCleared() {
		_spec.ClearField(mound.FieldExternalLink, field.TypeString)
	}
	if value, ok := muo.mutation.Number(); ok {
		_spec.SetField(mound.FieldNumber, field.TypeString, value)
	}
	if muo.mutation.NumberCleared() {
		_spec.ClearField(mound.FieldNumber, field.TypeString)
	}
	if muo.mutation.PetroglyphsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mound.PetroglyphsTable,
			Columns: []string{mound.PetroglyphsColumn},
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
			Table:   mound.PetroglyphsTable,
			Columns: []string{mound.PetroglyphsColumn},
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
			Table:   mound.PetroglyphsTable,
			Columns: []string{mound.PetroglyphsColumn},
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
	if muo.mutation.VisitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   mound.VisitsTable,
			Columns: mound.VisitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedVisitsIDs(); len(nodes) > 0 && !muo.mutation.VisitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   mound.VisitsTable,
			Columns: mound.VisitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.VisitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   mound.VisitsTable,
			Columns: mound.VisitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Mound{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mound.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
