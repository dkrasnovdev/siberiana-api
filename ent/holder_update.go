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
	"github.com/dkrasnovdev/siberiana-api/ent/book"
	"github.com/dkrasnovdev/siberiana-api/ent/holder"
	"github.com/dkrasnovdev/siberiana-api/ent/holderresponsibility"
	"github.com/dkrasnovdev/siberiana-api/ent/organization"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// HolderUpdate is the builder for updating Holder entities.
type HolderUpdate struct {
	config
	hooks    []Hook
	mutation *HolderMutation
}

// Where appends a list predicates to the HolderUpdate builder.
func (hu *HolderUpdate) Where(ps ...predicate.Holder) *HolderUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetCreatedBy sets the "created_by" field.
func (hu *HolderUpdate) SetCreatedBy(s string) *HolderUpdate {
	hu.mutation.SetCreatedBy(s)
	return hu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (hu *HolderUpdate) SetNillableCreatedBy(s *string) *HolderUpdate {
	if s != nil {
		hu.SetCreatedBy(*s)
	}
	return hu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (hu *HolderUpdate) ClearCreatedBy() *HolderUpdate {
	hu.mutation.ClearCreatedBy()
	return hu
}

// SetUpdatedAt sets the "updated_at" field.
func (hu *HolderUpdate) SetUpdatedAt(t time.Time) *HolderUpdate {
	hu.mutation.SetUpdatedAt(t)
	return hu
}

// SetUpdatedBy sets the "updated_by" field.
func (hu *HolderUpdate) SetUpdatedBy(s string) *HolderUpdate {
	hu.mutation.SetUpdatedBy(s)
	return hu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (hu *HolderUpdate) SetNillableUpdatedBy(s *string) *HolderUpdate {
	if s != nil {
		hu.SetUpdatedBy(*s)
	}
	return hu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (hu *HolderUpdate) ClearUpdatedBy() *HolderUpdate {
	hu.mutation.ClearUpdatedBy()
	return hu
}

// SetBeginData sets the "begin_data" field.
func (hu *HolderUpdate) SetBeginData(t time.Time) *HolderUpdate {
	hu.mutation.SetBeginData(t)
	return hu
}

// SetEndDate sets the "end_date" field.
func (hu *HolderUpdate) SetEndDate(t time.Time) *HolderUpdate {
	hu.mutation.SetEndDate(t)
	return hu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (hu *HolderUpdate) SetNillableEndDate(t *time.Time) *HolderUpdate {
	if t != nil {
		hu.SetEndDate(*t)
	}
	return hu
}

// ClearEndDate clears the value of the "end_date" field.
func (hu *HolderUpdate) ClearEndDate() *HolderUpdate {
	hu.mutation.ClearEndDate()
	return hu
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (hu *HolderUpdate) AddArtifactIDs(ids ...int) *HolderUpdate {
	hu.mutation.AddArtifactIDs(ids...)
	return hu
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (hu *HolderUpdate) AddArtifacts(a ...*Artifact) *HolderUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return hu.AddArtifactIDs(ids...)
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (hu *HolderUpdate) AddBookIDs(ids ...int) *HolderUpdate {
	hu.mutation.AddBookIDs(ids...)
	return hu
}

// AddBooks adds the "books" edges to the Book entity.
func (hu *HolderUpdate) AddBooks(b ...*Book) *HolderUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return hu.AddBookIDs(ids...)
}

// AddHolderResponsibilityIDs adds the "holder_responsibilities" edge to the HolderResponsibility entity by IDs.
func (hu *HolderUpdate) AddHolderResponsibilityIDs(ids ...int) *HolderUpdate {
	hu.mutation.AddHolderResponsibilityIDs(ids...)
	return hu
}

// AddHolderResponsibilities adds the "holder_responsibilities" edges to the HolderResponsibility entity.
func (hu *HolderUpdate) AddHolderResponsibilities(h ...*HolderResponsibility) *HolderUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hu.AddHolderResponsibilityIDs(ids...)
}

// SetPersonID sets the "person" edge to the Person entity by ID.
func (hu *HolderUpdate) SetPersonID(id int) *HolderUpdate {
	hu.mutation.SetPersonID(id)
	return hu
}

// SetNillablePersonID sets the "person" edge to the Person entity by ID if the given value is not nil.
func (hu *HolderUpdate) SetNillablePersonID(id *int) *HolderUpdate {
	if id != nil {
		hu = hu.SetPersonID(*id)
	}
	return hu
}

// SetPerson sets the "person" edge to the Person entity.
func (hu *HolderUpdate) SetPerson(p *Person) *HolderUpdate {
	return hu.SetPersonID(p.ID)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (hu *HolderUpdate) SetOrganizationID(id int) *HolderUpdate {
	hu.mutation.SetOrganizationID(id)
	return hu
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (hu *HolderUpdate) SetNillableOrganizationID(id *int) *HolderUpdate {
	if id != nil {
		hu = hu.SetOrganizationID(*id)
	}
	return hu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (hu *HolderUpdate) SetOrganization(o *Organization) *HolderUpdate {
	return hu.SetOrganizationID(o.ID)
}

// Mutation returns the HolderMutation object of the builder.
func (hu *HolderUpdate) Mutation() *HolderMutation {
	return hu.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (hu *HolderUpdate) ClearArtifacts() *HolderUpdate {
	hu.mutation.ClearArtifacts()
	return hu
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (hu *HolderUpdate) RemoveArtifactIDs(ids ...int) *HolderUpdate {
	hu.mutation.RemoveArtifactIDs(ids...)
	return hu
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (hu *HolderUpdate) RemoveArtifacts(a ...*Artifact) *HolderUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return hu.RemoveArtifactIDs(ids...)
}

// ClearBooks clears all "books" edges to the Book entity.
func (hu *HolderUpdate) ClearBooks() *HolderUpdate {
	hu.mutation.ClearBooks()
	return hu
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (hu *HolderUpdate) RemoveBookIDs(ids ...int) *HolderUpdate {
	hu.mutation.RemoveBookIDs(ids...)
	return hu
}

// RemoveBooks removes "books" edges to Book entities.
func (hu *HolderUpdate) RemoveBooks(b ...*Book) *HolderUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return hu.RemoveBookIDs(ids...)
}

// ClearHolderResponsibilities clears all "holder_responsibilities" edges to the HolderResponsibility entity.
func (hu *HolderUpdate) ClearHolderResponsibilities() *HolderUpdate {
	hu.mutation.ClearHolderResponsibilities()
	return hu
}

// RemoveHolderResponsibilityIDs removes the "holder_responsibilities" edge to HolderResponsibility entities by IDs.
func (hu *HolderUpdate) RemoveHolderResponsibilityIDs(ids ...int) *HolderUpdate {
	hu.mutation.RemoveHolderResponsibilityIDs(ids...)
	return hu
}

// RemoveHolderResponsibilities removes "holder_responsibilities" edges to HolderResponsibility entities.
func (hu *HolderUpdate) RemoveHolderResponsibilities(h ...*HolderResponsibility) *HolderUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hu.RemoveHolderResponsibilityIDs(ids...)
}

// ClearPerson clears the "person" edge to the Person entity.
func (hu *HolderUpdate) ClearPerson() *HolderUpdate {
	hu.mutation.ClearPerson()
	return hu
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (hu *HolderUpdate) ClearOrganization() *HolderUpdate {
	hu.mutation.ClearOrganization()
	return hu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HolderUpdate) Save(ctx context.Context) (int, error) {
	if err := hu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HolderUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HolderUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HolderUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hu *HolderUpdate) defaults() error {
	if _, ok := hu.mutation.UpdatedAt(); !ok {
		if holder.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized holder.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := holder.UpdateDefaultUpdatedAt()
		hu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (hu *HolderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(holder.Table, holder.Columns, sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.CreatedBy(); ok {
		_spec.SetField(holder.FieldCreatedBy, field.TypeString, value)
	}
	if hu.mutation.CreatedByCleared() {
		_spec.ClearField(holder.FieldCreatedBy, field.TypeString)
	}
	if value, ok := hu.mutation.UpdatedAt(); ok {
		_spec.SetField(holder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hu.mutation.UpdatedBy(); ok {
		_spec.SetField(holder.FieldUpdatedBy, field.TypeString, value)
	}
	if hu.mutation.UpdatedByCleared() {
		_spec.ClearField(holder.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := hu.mutation.BeginData(); ok {
		_spec.SetField(holder.FieldBeginData, field.TypeTime, value)
	}
	if value, ok := hu.mutation.EndDate(); ok {
		_spec.SetField(holder.FieldEndDate, field.TypeTime, value)
	}
	if hu.mutation.EndDateCleared() {
		_spec.ClearField(holder.FieldEndDate, field.TypeTime)
	}
	if hu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.ArtifactsTable,
			Columns: holder.ArtifactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !hu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.ArtifactsTable,
			Columns: holder.ArtifactsPrimaryKey,
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
	if nodes := hu.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.ArtifactsTable,
			Columns: holder.ArtifactsPrimaryKey,
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
	if hu.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.BooksTable,
			Columns: holder.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedBooksIDs(); len(nodes) > 0 && !hu.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.BooksTable,
			Columns: holder.BooksPrimaryKey,
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
	if nodes := hu.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.BooksTable,
			Columns: holder.BooksPrimaryKey,
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
	if hu.mutation.HolderResponsibilitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.HolderResponsibilitiesTable,
			Columns: holder.HolderResponsibilitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedHolderResponsibilitiesIDs(); len(nodes) > 0 && !hu.mutation.HolderResponsibilitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.HolderResponsibilitiesTable,
			Columns: holder.HolderResponsibilitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.HolderResponsibilitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.HolderResponsibilitiesTable,
			Columns: holder.HolderResponsibilitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.PersonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   holder.PersonTable,
			Columns: []string{holder.PersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   holder.PersonTable,
			Columns: []string{holder.PersonColumn},
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
	if hu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   holder.OrganizationTable,
			Columns: []string{holder.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   holder.OrganizationTable,
			Columns: []string{holder.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{holder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HolderUpdateOne is the builder for updating a single Holder entity.
type HolderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HolderMutation
}

// SetCreatedBy sets the "created_by" field.
func (huo *HolderUpdateOne) SetCreatedBy(s string) *HolderUpdateOne {
	huo.mutation.SetCreatedBy(s)
	return huo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (huo *HolderUpdateOne) SetNillableCreatedBy(s *string) *HolderUpdateOne {
	if s != nil {
		huo.SetCreatedBy(*s)
	}
	return huo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (huo *HolderUpdateOne) ClearCreatedBy() *HolderUpdateOne {
	huo.mutation.ClearCreatedBy()
	return huo
}

// SetUpdatedAt sets the "updated_at" field.
func (huo *HolderUpdateOne) SetUpdatedAt(t time.Time) *HolderUpdateOne {
	huo.mutation.SetUpdatedAt(t)
	return huo
}

// SetUpdatedBy sets the "updated_by" field.
func (huo *HolderUpdateOne) SetUpdatedBy(s string) *HolderUpdateOne {
	huo.mutation.SetUpdatedBy(s)
	return huo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (huo *HolderUpdateOne) SetNillableUpdatedBy(s *string) *HolderUpdateOne {
	if s != nil {
		huo.SetUpdatedBy(*s)
	}
	return huo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (huo *HolderUpdateOne) ClearUpdatedBy() *HolderUpdateOne {
	huo.mutation.ClearUpdatedBy()
	return huo
}

// SetBeginData sets the "begin_data" field.
func (huo *HolderUpdateOne) SetBeginData(t time.Time) *HolderUpdateOne {
	huo.mutation.SetBeginData(t)
	return huo
}

// SetEndDate sets the "end_date" field.
func (huo *HolderUpdateOne) SetEndDate(t time.Time) *HolderUpdateOne {
	huo.mutation.SetEndDate(t)
	return huo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (huo *HolderUpdateOne) SetNillableEndDate(t *time.Time) *HolderUpdateOne {
	if t != nil {
		huo.SetEndDate(*t)
	}
	return huo
}

// ClearEndDate clears the value of the "end_date" field.
func (huo *HolderUpdateOne) ClearEndDate() *HolderUpdateOne {
	huo.mutation.ClearEndDate()
	return huo
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (huo *HolderUpdateOne) AddArtifactIDs(ids ...int) *HolderUpdateOne {
	huo.mutation.AddArtifactIDs(ids...)
	return huo
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (huo *HolderUpdateOne) AddArtifacts(a ...*Artifact) *HolderUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return huo.AddArtifactIDs(ids...)
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (huo *HolderUpdateOne) AddBookIDs(ids ...int) *HolderUpdateOne {
	huo.mutation.AddBookIDs(ids...)
	return huo
}

// AddBooks adds the "books" edges to the Book entity.
func (huo *HolderUpdateOne) AddBooks(b ...*Book) *HolderUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return huo.AddBookIDs(ids...)
}

// AddHolderResponsibilityIDs adds the "holder_responsibilities" edge to the HolderResponsibility entity by IDs.
func (huo *HolderUpdateOne) AddHolderResponsibilityIDs(ids ...int) *HolderUpdateOne {
	huo.mutation.AddHolderResponsibilityIDs(ids...)
	return huo
}

// AddHolderResponsibilities adds the "holder_responsibilities" edges to the HolderResponsibility entity.
func (huo *HolderUpdateOne) AddHolderResponsibilities(h ...*HolderResponsibility) *HolderUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return huo.AddHolderResponsibilityIDs(ids...)
}

// SetPersonID sets the "person" edge to the Person entity by ID.
func (huo *HolderUpdateOne) SetPersonID(id int) *HolderUpdateOne {
	huo.mutation.SetPersonID(id)
	return huo
}

// SetNillablePersonID sets the "person" edge to the Person entity by ID if the given value is not nil.
func (huo *HolderUpdateOne) SetNillablePersonID(id *int) *HolderUpdateOne {
	if id != nil {
		huo = huo.SetPersonID(*id)
	}
	return huo
}

// SetPerson sets the "person" edge to the Person entity.
func (huo *HolderUpdateOne) SetPerson(p *Person) *HolderUpdateOne {
	return huo.SetPersonID(p.ID)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (huo *HolderUpdateOne) SetOrganizationID(id int) *HolderUpdateOne {
	huo.mutation.SetOrganizationID(id)
	return huo
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (huo *HolderUpdateOne) SetNillableOrganizationID(id *int) *HolderUpdateOne {
	if id != nil {
		huo = huo.SetOrganizationID(*id)
	}
	return huo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (huo *HolderUpdateOne) SetOrganization(o *Organization) *HolderUpdateOne {
	return huo.SetOrganizationID(o.ID)
}

// Mutation returns the HolderMutation object of the builder.
func (huo *HolderUpdateOne) Mutation() *HolderMutation {
	return huo.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (huo *HolderUpdateOne) ClearArtifacts() *HolderUpdateOne {
	huo.mutation.ClearArtifacts()
	return huo
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (huo *HolderUpdateOne) RemoveArtifactIDs(ids ...int) *HolderUpdateOne {
	huo.mutation.RemoveArtifactIDs(ids...)
	return huo
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (huo *HolderUpdateOne) RemoveArtifacts(a ...*Artifact) *HolderUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return huo.RemoveArtifactIDs(ids...)
}

// ClearBooks clears all "books" edges to the Book entity.
func (huo *HolderUpdateOne) ClearBooks() *HolderUpdateOne {
	huo.mutation.ClearBooks()
	return huo
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (huo *HolderUpdateOne) RemoveBookIDs(ids ...int) *HolderUpdateOne {
	huo.mutation.RemoveBookIDs(ids...)
	return huo
}

// RemoveBooks removes "books" edges to Book entities.
func (huo *HolderUpdateOne) RemoveBooks(b ...*Book) *HolderUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return huo.RemoveBookIDs(ids...)
}

// ClearHolderResponsibilities clears all "holder_responsibilities" edges to the HolderResponsibility entity.
func (huo *HolderUpdateOne) ClearHolderResponsibilities() *HolderUpdateOne {
	huo.mutation.ClearHolderResponsibilities()
	return huo
}

// RemoveHolderResponsibilityIDs removes the "holder_responsibilities" edge to HolderResponsibility entities by IDs.
func (huo *HolderUpdateOne) RemoveHolderResponsibilityIDs(ids ...int) *HolderUpdateOne {
	huo.mutation.RemoveHolderResponsibilityIDs(ids...)
	return huo
}

// RemoveHolderResponsibilities removes "holder_responsibilities" edges to HolderResponsibility entities.
func (huo *HolderUpdateOne) RemoveHolderResponsibilities(h ...*HolderResponsibility) *HolderUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return huo.RemoveHolderResponsibilityIDs(ids...)
}

// ClearPerson clears the "person" edge to the Person entity.
func (huo *HolderUpdateOne) ClearPerson() *HolderUpdateOne {
	huo.mutation.ClearPerson()
	return huo
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (huo *HolderUpdateOne) ClearOrganization() *HolderUpdateOne {
	huo.mutation.ClearOrganization()
	return huo
}

// Where appends a list predicates to the HolderUpdate builder.
func (huo *HolderUpdateOne) Where(ps ...predicate.Holder) *HolderUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HolderUpdateOne) Select(field string, fields ...string) *HolderUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Holder entity.
func (huo *HolderUpdateOne) Save(ctx context.Context) (*Holder, error) {
	if err := huo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HolderUpdateOne) SaveX(ctx context.Context) *Holder {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HolderUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HolderUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (huo *HolderUpdateOne) defaults() error {
	if _, ok := huo.mutation.UpdatedAt(); !ok {
		if holder.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized holder.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := holder.UpdateDefaultUpdatedAt()
		huo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (huo *HolderUpdateOne) sqlSave(ctx context.Context) (_node *Holder, err error) {
	_spec := sqlgraph.NewUpdateSpec(holder.Table, holder.Columns, sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Holder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, holder.FieldID)
		for _, f := range fields {
			if !holder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != holder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.CreatedBy(); ok {
		_spec.SetField(holder.FieldCreatedBy, field.TypeString, value)
	}
	if huo.mutation.CreatedByCleared() {
		_spec.ClearField(holder.FieldCreatedBy, field.TypeString)
	}
	if value, ok := huo.mutation.UpdatedAt(); ok {
		_spec.SetField(holder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := huo.mutation.UpdatedBy(); ok {
		_spec.SetField(holder.FieldUpdatedBy, field.TypeString, value)
	}
	if huo.mutation.UpdatedByCleared() {
		_spec.ClearField(holder.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := huo.mutation.BeginData(); ok {
		_spec.SetField(holder.FieldBeginData, field.TypeTime, value)
	}
	if value, ok := huo.mutation.EndDate(); ok {
		_spec.SetField(holder.FieldEndDate, field.TypeTime, value)
	}
	if huo.mutation.EndDateCleared() {
		_spec.ClearField(holder.FieldEndDate, field.TypeTime)
	}
	if huo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.ArtifactsTable,
			Columns: holder.ArtifactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedArtifactsIDs(); len(nodes) > 0 && !huo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.ArtifactsTable,
			Columns: holder.ArtifactsPrimaryKey,
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
	if nodes := huo.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.ArtifactsTable,
			Columns: holder.ArtifactsPrimaryKey,
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
	if huo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.BooksTable,
			Columns: holder.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedBooksIDs(); len(nodes) > 0 && !huo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.BooksTable,
			Columns: holder.BooksPrimaryKey,
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
	if nodes := huo.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.BooksTable,
			Columns: holder.BooksPrimaryKey,
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
	if huo.mutation.HolderResponsibilitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.HolderResponsibilitiesTable,
			Columns: holder.HolderResponsibilitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedHolderResponsibilitiesIDs(); len(nodes) > 0 && !huo.mutation.HolderResponsibilitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.HolderResponsibilitiesTable,
			Columns: holder.HolderResponsibilitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.HolderResponsibilitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   holder.HolderResponsibilitiesTable,
			Columns: holder.HolderResponsibilitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.PersonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   holder.PersonTable,
			Columns: []string{holder.PersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   holder.PersonTable,
			Columns: []string{holder.PersonColumn},
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
	if huo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   holder.OrganizationTable,
			Columns: []string{holder.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   holder.OrganizationTable,
			Columns: []string{holder.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Holder{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{holder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
