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
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
	"github.com/dkrasnovdev/siberiana-api/ent/protectedarea"
	"github.com/dkrasnovdev/siberiana-api/ent/protectedareacategory"
	"github.com/dkrasnovdev/siberiana-api/ent/protectedareapicture"
)

// ProtectedAreaUpdate is the builder for updating ProtectedArea entities.
type ProtectedAreaUpdate struct {
	config
	hooks    []Hook
	mutation *ProtectedAreaMutation
}

// Where appends a list predicates to the ProtectedAreaUpdate builder.
func (pau *ProtectedAreaUpdate) Where(ps ...predicate.ProtectedArea) *ProtectedAreaUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// SetCreatedBy sets the "created_by" field.
func (pau *ProtectedAreaUpdate) SetCreatedBy(s string) *ProtectedAreaUpdate {
	pau.mutation.SetCreatedBy(s)
	return pau
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableCreatedBy(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetCreatedBy(*s)
	}
	return pau
}

// ClearCreatedBy clears the value of the "created_by" field.
func (pau *ProtectedAreaUpdate) ClearCreatedBy() *ProtectedAreaUpdate {
	pau.mutation.ClearCreatedBy()
	return pau
}

// SetUpdatedAt sets the "updated_at" field.
func (pau *ProtectedAreaUpdate) SetUpdatedAt(t time.Time) *ProtectedAreaUpdate {
	pau.mutation.SetUpdatedAt(t)
	return pau
}

// SetUpdatedBy sets the "updated_by" field.
func (pau *ProtectedAreaUpdate) SetUpdatedBy(s string) *ProtectedAreaUpdate {
	pau.mutation.SetUpdatedBy(s)
	return pau
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableUpdatedBy(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetUpdatedBy(*s)
	}
	return pau
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pau *ProtectedAreaUpdate) ClearUpdatedBy() *ProtectedAreaUpdate {
	pau.mutation.ClearUpdatedBy()
	return pau
}

// SetDisplayName sets the "display_name" field.
func (pau *ProtectedAreaUpdate) SetDisplayName(s string) *ProtectedAreaUpdate {
	pau.mutation.SetDisplayName(s)
	return pau
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableDisplayName(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetDisplayName(*s)
	}
	return pau
}

// ClearDisplayName clears the value of the "display_name" field.
func (pau *ProtectedAreaUpdate) ClearDisplayName() *ProtectedAreaUpdate {
	pau.mutation.ClearDisplayName()
	return pau
}

// SetAbbreviation sets the "abbreviation" field.
func (pau *ProtectedAreaUpdate) SetAbbreviation(s string) *ProtectedAreaUpdate {
	pau.mutation.SetAbbreviation(s)
	return pau
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableAbbreviation(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetAbbreviation(*s)
	}
	return pau
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (pau *ProtectedAreaUpdate) ClearAbbreviation() *ProtectedAreaUpdate {
	pau.mutation.ClearAbbreviation()
	return pau
}

// SetDescription sets the "description" field.
func (pau *ProtectedAreaUpdate) SetDescription(s string) *ProtectedAreaUpdate {
	pau.mutation.SetDescription(s)
	return pau
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableDescription(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetDescription(*s)
	}
	return pau
}

// ClearDescription clears the value of the "description" field.
func (pau *ProtectedAreaUpdate) ClearDescription() *ProtectedAreaUpdate {
	pau.mutation.ClearDescription()
	return pau
}

// SetExternalLink sets the "external_link" field.
func (pau *ProtectedAreaUpdate) SetExternalLink(s string) *ProtectedAreaUpdate {
	pau.mutation.SetExternalLink(s)
	return pau
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableExternalLink(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetExternalLink(*s)
	}
	return pau
}

// ClearExternalLink clears the value of the "external_link" field.
func (pau *ProtectedAreaUpdate) ClearExternalLink() *ProtectedAreaUpdate {
	pau.mutation.ClearExternalLink()
	return pau
}

// SetArea sets the "area" field.
func (pau *ProtectedAreaUpdate) SetArea(s string) *ProtectedAreaUpdate {
	pau.mutation.SetArea(s)
	return pau
}

// SetNillableArea sets the "area" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableArea(s *string) *ProtectedAreaUpdate {
	if s != nil {
		pau.SetArea(*s)
	}
	return pau
}

// ClearArea clears the value of the "area" field.
func (pau *ProtectedAreaUpdate) ClearArea() *ProtectedAreaUpdate {
	pau.mutation.ClearArea()
	return pau
}

// SetEstablishmentDate sets the "establishment_date" field.
func (pau *ProtectedAreaUpdate) SetEstablishmentDate(t time.Time) *ProtectedAreaUpdate {
	pau.mutation.SetEstablishmentDate(t)
	return pau
}

// SetNillableEstablishmentDate sets the "establishment_date" field if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableEstablishmentDate(t *time.Time) *ProtectedAreaUpdate {
	if t != nil {
		pau.SetEstablishmentDate(*t)
	}
	return pau
}

// ClearEstablishmentDate clears the value of the "establishment_date" field.
func (pau *ProtectedAreaUpdate) ClearEstablishmentDate() *ProtectedAreaUpdate {
	pau.mutation.ClearEstablishmentDate()
	return pau
}

// AddProtectedAreaPictureIDs adds the "protected_area_pictures" edge to the ProtectedAreaPicture entity by IDs.
func (pau *ProtectedAreaUpdate) AddProtectedAreaPictureIDs(ids ...int) *ProtectedAreaUpdate {
	pau.mutation.AddProtectedAreaPictureIDs(ids...)
	return pau
}

// AddProtectedAreaPictures adds the "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (pau *ProtectedAreaUpdate) AddProtectedAreaPictures(p ...*ProtectedAreaPicture) *ProtectedAreaUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pau.AddProtectedAreaPictureIDs(ids...)
}

// SetProtectedAreaCategoryID sets the "protected_area_category" edge to the ProtectedAreaCategory entity by ID.
func (pau *ProtectedAreaUpdate) SetProtectedAreaCategoryID(id int) *ProtectedAreaUpdate {
	pau.mutation.SetProtectedAreaCategoryID(id)
	return pau
}

// SetNillableProtectedAreaCategoryID sets the "protected_area_category" edge to the ProtectedAreaCategory entity by ID if the given value is not nil.
func (pau *ProtectedAreaUpdate) SetNillableProtectedAreaCategoryID(id *int) *ProtectedAreaUpdate {
	if id != nil {
		pau = pau.SetProtectedAreaCategoryID(*id)
	}
	return pau
}

// SetProtectedAreaCategory sets the "protected_area_category" edge to the ProtectedAreaCategory entity.
func (pau *ProtectedAreaUpdate) SetProtectedAreaCategory(p *ProtectedAreaCategory) *ProtectedAreaUpdate {
	return pau.SetProtectedAreaCategoryID(p.ID)
}

// Mutation returns the ProtectedAreaMutation object of the builder.
func (pau *ProtectedAreaUpdate) Mutation() *ProtectedAreaMutation {
	return pau.mutation
}

// ClearProtectedAreaPictures clears all "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (pau *ProtectedAreaUpdate) ClearProtectedAreaPictures() *ProtectedAreaUpdate {
	pau.mutation.ClearProtectedAreaPictures()
	return pau
}

// RemoveProtectedAreaPictureIDs removes the "protected_area_pictures" edge to ProtectedAreaPicture entities by IDs.
func (pau *ProtectedAreaUpdate) RemoveProtectedAreaPictureIDs(ids ...int) *ProtectedAreaUpdate {
	pau.mutation.RemoveProtectedAreaPictureIDs(ids...)
	return pau
}

// RemoveProtectedAreaPictures removes "protected_area_pictures" edges to ProtectedAreaPicture entities.
func (pau *ProtectedAreaUpdate) RemoveProtectedAreaPictures(p ...*ProtectedAreaPicture) *ProtectedAreaUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pau.RemoveProtectedAreaPictureIDs(ids...)
}

// ClearProtectedAreaCategory clears the "protected_area_category" edge to the ProtectedAreaCategory entity.
func (pau *ProtectedAreaUpdate) ClearProtectedAreaCategory() *ProtectedAreaUpdate {
	pau.mutation.ClearProtectedAreaCategory()
	return pau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *ProtectedAreaUpdate) Save(ctx context.Context) (int, error) {
	if err := pau.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *ProtectedAreaUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *ProtectedAreaUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *ProtectedAreaUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pau *ProtectedAreaUpdate) defaults() error {
	if _, ok := pau.mutation.UpdatedAt(); !ok {
		if protectedarea.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedarea.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := protectedarea.UpdateDefaultUpdatedAt()
		pau.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pau *ProtectedAreaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(protectedarea.Table, protectedarea.Columns, sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.CreatedBy(); ok {
		_spec.SetField(protectedarea.FieldCreatedBy, field.TypeString, value)
	}
	if pau.mutation.CreatedByCleared() {
		_spec.ClearField(protectedarea.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pau.mutation.UpdatedAt(); ok {
		_spec.SetField(protectedarea.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pau.mutation.UpdatedBy(); ok {
		_spec.SetField(protectedarea.FieldUpdatedBy, field.TypeString, value)
	}
	if pau.mutation.UpdatedByCleared() {
		_spec.ClearField(protectedarea.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pau.mutation.DisplayName(); ok {
		_spec.SetField(protectedarea.FieldDisplayName, field.TypeString, value)
	}
	if pau.mutation.DisplayNameCleared() {
		_spec.ClearField(protectedarea.FieldDisplayName, field.TypeString)
	}
	if value, ok := pau.mutation.Abbreviation(); ok {
		_spec.SetField(protectedarea.FieldAbbreviation, field.TypeString, value)
	}
	if pau.mutation.AbbreviationCleared() {
		_spec.ClearField(protectedarea.FieldAbbreviation, field.TypeString)
	}
	if value, ok := pau.mutation.Description(); ok {
		_spec.SetField(protectedarea.FieldDescription, field.TypeString, value)
	}
	if pau.mutation.DescriptionCleared() {
		_spec.ClearField(protectedarea.FieldDescription, field.TypeString)
	}
	if value, ok := pau.mutation.ExternalLink(); ok {
		_spec.SetField(protectedarea.FieldExternalLink, field.TypeString, value)
	}
	if pau.mutation.ExternalLinkCleared() {
		_spec.ClearField(protectedarea.FieldExternalLink, field.TypeString)
	}
	if value, ok := pau.mutation.Area(); ok {
		_spec.SetField(protectedarea.FieldArea, field.TypeString, value)
	}
	if pau.mutation.AreaCleared() {
		_spec.ClearField(protectedarea.FieldArea, field.TypeString)
	}
	if value, ok := pau.mutation.EstablishmentDate(); ok {
		_spec.SetField(protectedarea.FieldEstablishmentDate, field.TypeTime, value)
	}
	if pau.mutation.EstablishmentDateCleared() {
		_spec.ClearField(protectedarea.FieldEstablishmentDate, field.TypeTime)
	}
	if pau.mutation.ProtectedAreaPicturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   protectedarea.ProtectedAreaPicturesTable,
			Columns: []string{protectedarea.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.RemovedProtectedAreaPicturesIDs(); len(nodes) > 0 && !pau.mutation.ProtectedAreaPicturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   protectedarea.ProtectedAreaPicturesTable,
			Columns: []string{protectedarea.ProtectedAreaPicturesColumn},
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
	if nodes := pau.mutation.ProtectedAreaPicturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   protectedarea.ProtectedAreaPicturesTable,
			Columns: []string{protectedarea.ProtectedAreaPicturesColumn},
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
	if pau.mutation.ProtectedAreaCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   protectedarea.ProtectedAreaCategoryTable,
			Columns: []string{protectedarea.ProtectedAreaCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareacategory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.ProtectedAreaCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   protectedarea.ProtectedAreaCategoryTable,
			Columns: []string{protectedarea.ProtectedAreaCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareacategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{protectedarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// ProtectedAreaUpdateOne is the builder for updating a single ProtectedArea entity.
type ProtectedAreaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProtectedAreaMutation
}

// SetCreatedBy sets the "created_by" field.
func (pauo *ProtectedAreaUpdateOne) SetCreatedBy(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetCreatedBy(s)
	return pauo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableCreatedBy(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetCreatedBy(*s)
	}
	return pauo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (pauo *ProtectedAreaUpdateOne) ClearCreatedBy() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearCreatedBy()
	return pauo
}

// SetUpdatedAt sets the "updated_at" field.
func (pauo *ProtectedAreaUpdateOne) SetUpdatedAt(t time.Time) *ProtectedAreaUpdateOne {
	pauo.mutation.SetUpdatedAt(t)
	return pauo
}

// SetUpdatedBy sets the "updated_by" field.
func (pauo *ProtectedAreaUpdateOne) SetUpdatedBy(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetUpdatedBy(s)
	return pauo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableUpdatedBy(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetUpdatedBy(*s)
	}
	return pauo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pauo *ProtectedAreaUpdateOne) ClearUpdatedBy() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearUpdatedBy()
	return pauo
}

// SetDisplayName sets the "display_name" field.
func (pauo *ProtectedAreaUpdateOne) SetDisplayName(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetDisplayName(s)
	return pauo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableDisplayName(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetDisplayName(*s)
	}
	return pauo
}

// ClearDisplayName clears the value of the "display_name" field.
func (pauo *ProtectedAreaUpdateOne) ClearDisplayName() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearDisplayName()
	return pauo
}

// SetAbbreviation sets the "abbreviation" field.
func (pauo *ProtectedAreaUpdateOne) SetAbbreviation(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetAbbreviation(s)
	return pauo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableAbbreviation(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetAbbreviation(*s)
	}
	return pauo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (pauo *ProtectedAreaUpdateOne) ClearAbbreviation() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearAbbreviation()
	return pauo
}

// SetDescription sets the "description" field.
func (pauo *ProtectedAreaUpdateOne) SetDescription(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetDescription(s)
	return pauo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableDescription(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetDescription(*s)
	}
	return pauo
}

// ClearDescription clears the value of the "description" field.
func (pauo *ProtectedAreaUpdateOne) ClearDescription() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearDescription()
	return pauo
}

// SetExternalLink sets the "external_link" field.
func (pauo *ProtectedAreaUpdateOne) SetExternalLink(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetExternalLink(s)
	return pauo
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableExternalLink(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetExternalLink(*s)
	}
	return pauo
}

// ClearExternalLink clears the value of the "external_link" field.
func (pauo *ProtectedAreaUpdateOne) ClearExternalLink() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearExternalLink()
	return pauo
}

// SetArea sets the "area" field.
func (pauo *ProtectedAreaUpdateOne) SetArea(s string) *ProtectedAreaUpdateOne {
	pauo.mutation.SetArea(s)
	return pauo
}

// SetNillableArea sets the "area" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableArea(s *string) *ProtectedAreaUpdateOne {
	if s != nil {
		pauo.SetArea(*s)
	}
	return pauo
}

// ClearArea clears the value of the "area" field.
func (pauo *ProtectedAreaUpdateOne) ClearArea() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearArea()
	return pauo
}

// SetEstablishmentDate sets the "establishment_date" field.
func (pauo *ProtectedAreaUpdateOne) SetEstablishmentDate(t time.Time) *ProtectedAreaUpdateOne {
	pauo.mutation.SetEstablishmentDate(t)
	return pauo
}

// SetNillableEstablishmentDate sets the "establishment_date" field if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableEstablishmentDate(t *time.Time) *ProtectedAreaUpdateOne {
	if t != nil {
		pauo.SetEstablishmentDate(*t)
	}
	return pauo
}

// ClearEstablishmentDate clears the value of the "establishment_date" field.
func (pauo *ProtectedAreaUpdateOne) ClearEstablishmentDate() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearEstablishmentDate()
	return pauo
}

// AddProtectedAreaPictureIDs adds the "protected_area_pictures" edge to the ProtectedAreaPicture entity by IDs.
func (pauo *ProtectedAreaUpdateOne) AddProtectedAreaPictureIDs(ids ...int) *ProtectedAreaUpdateOne {
	pauo.mutation.AddProtectedAreaPictureIDs(ids...)
	return pauo
}

// AddProtectedAreaPictures adds the "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (pauo *ProtectedAreaUpdateOne) AddProtectedAreaPictures(p ...*ProtectedAreaPicture) *ProtectedAreaUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pauo.AddProtectedAreaPictureIDs(ids...)
}

// SetProtectedAreaCategoryID sets the "protected_area_category" edge to the ProtectedAreaCategory entity by ID.
func (pauo *ProtectedAreaUpdateOne) SetProtectedAreaCategoryID(id int) *ProtectedAreaUpdateOne {
	pauo.mutation.SetProtectedAreaCategoryID(id)
	return pauo
}

// SetNillableProtectedAreaCategoryID sets the "protected_area_category" edge to the ProtectedAreaCategory entity by ID if the given value is not nil.
func (pauo *ProtectedAreaUpdateOne) SetNillableProtectedAreaCategoryID(id *int) *ProtectedAreaUpdateOne {
	if id != nil {
		pauo = pauo.SetProtectedAreaCategoryID(*id)
	}
	return pauo
}

// SetProtectedAreaCategory sets the "protected_area_category" edge to the ProtectedAreaCategory entity.
func (pauo *ProtectedAreaUpdateOne) SetProtectedAreaCategory(p *ProtectedAreaCategory) *ProtectedAreaUpdateOne {
	return pauo.SetProtectedAreaCategoryID(p.ID)
}

// Mutation returns the ProtectedAreaMutation object of the builder.
func (pauo *ProtectedAreaUpdateOne) Mutation() *ProtectedAreaMutation {
	return pauo.mutation
}

// ClearProtectedAreaPictures clears all "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (pauo *ProtectedAreaUpdateOne) ClearProtectedAreaPictures() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearProtectedAreaPictures()
	return pauo
}

// RemoveProtectedAreaPictureIDs removes the "protected_area_pictures" edge to ProtectedAreaPicture entities by IDs.
func (pauo *ProtectedAreaUpdateOne) RemoveProtectedAreaPictureIDs(ids ...int) *ProtectedAreaUpdateOne {
	pauo.mutation.RemoveProtectedAreaPictureIDs(ids...)
	return pauo
}

// RemoveProtectedAreaPictures removes "protected_area_pictures" edges to ProtectedAreaPicture entities.
func (pauo *ProtectedAreaUpdateOne) RemoveProtectedAreaPictures(p ...*ProtectedAreaPicture) *ProtectedAreaUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pauo.RemoveProtectedAreaPictureIDs(ids...)
}

// ClearProtectedAreaCategory clears the "protected_area_category" edge to the ProtectedAreaCategory entity.
func (pauo *ProtectedAreaUpdateOne) ClearProtectedAreaCategory() *ProtectedAreaUpdateOne {
	pauo.mutation.ClearProtectedAreaCategory()
	return pauo
}

// Where appends a list predicates to the ProtectedAreaUpdate builder.
func (pauo *ProtectedAreaUpdateOne) Where(ps ...predicate.ProtectedArea) *ProtectedAreaUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *ProtectedAreaUpdateOne) Select(field string, fields ...string) *ProtectedAreaUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated ProtectedArea entity.
func (pauo *ProtectedAreaUpdateOne) Save(ctx context.Context) (*ProtectedArea, error) {
	if err := pauo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *ProtectedAreaUpdateOne) SaveX(ctx context.Context) *ProtectedArea {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *ProtectedAreaUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *ProtectedAreaUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pauo *ProtectedAreaUpdateOne) defaults() error {
	if _, ok := pauo.mutation.UpdatedAt(); !ok {
		if protectedarea.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized protectedarea.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := protectedarea.UpdateDefaultUpdatedAt()
		pauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pauo *ProtectedAreaUpdateOne) sqlSave(ctx context.Context) (_node *ProtectedArea, err error) {
	_spec := sqlgraph.NewUpdateSpec(protectedarea.Table, protectedarea.Columns, sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProtectedArea.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, protectedarea.FieldID)
		for _, f := range fields {
			if !protectedarea.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != protectedarea.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.CreatedBy(); ok {
		_spec.SetField(protectedarea.FieldCreatedBy, field.TypeString, value)
	}
	if pauo.mutation.CreatedByCleared() {
		_spec.ClearField(protectedarea.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pauo.mutation.UpdatedAt(); ok {
		_spec.SetField(protectedarea.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pauo.mutation.UpdatedBy(); ok {
		_spec.SetField(protectedarea.FieldUpdatedBy, field.TypeString, value)
	}
	if pauo.mutation.UpdatedByCleared() {
		_spec.ClearField(protectedarea.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pauo.mutation.DisplayName(); ok {
		_spec.SetField(protectedarea.FieldDisplayName, field.TypeString, value)
	}
	if pauo.mutation.DisplayNameCleared() {
		_spec.ClearField(protectedarea.FieldDisplayName, field.TypeString)
	}
	if value, ok := pauo.mutation.Abbreviation(); ok {
		_spec.SetField(protectedarea.FieldAbbreviation, field.TypeString, value)
	}
	if pauo.mutation.AbbreviationCleared() {
		_spec.ClearField(protectedarea.FieldAbbreviation, field.TypeString)
	}
	if value, ok := pauo.mutation.Description(); ok {
		_spec.SetField(protectedarea.FieldDescription, field.TypeString, value)
	}
	if pauo.mutation.DescriptionCleared() {
		_spec.ClearField(protectedarea.FieldDescription, field.TypeString)
	}
	if value, ok := pauo.mutation.ExternalLink(); ok {
		_spec.SetField(protectedarea.FieldExternalLink, field.TypeString, value)
	}
	if pauo.mutation.ExternalLinkCleared() {
		_spec.ClearField(protectedarea.FieldExternalLink, field.TypeString)
	}
	if value, ok := pauo.mutation.Area(); ok {
		_spec.SetField(protectedarea.FieldArea, field.TypeString, value)
	}
	if pauo.mutation.AreaCleared() {
		_spec.ClearField(protectedarea.FieldArea, field.TypeString)
	}
	if value, ok := pauo.mutation.EstablishmentDate(); ok {
		_spec.SetField(protectedarea.FieldEstablishmentDate, field.TypeTime, value)
	}
	if pauo.mutation.EstablishmentDateCleared() {
		_spec.ClearField(protectedarea.FieldEstablishmentDate, field.TypeTime)
	}
	if pauo.mutation.ProtectedAreaPicturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   protectedarea.ProtectedAreaPicturesTable,
			Columns: []string{protectedarea.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.RemovedProtectedAreaPicturesIDs(); len(nodes) > 0 && !pauo.mutation.ProtectedAreaPicturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   protectedarea.ProtectedAreaPicturesTable,
			Columns: []string{protectedarea.ProtectedAreaPicturesColumn},
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
	if nodes := pauo.mutation.ProtectedAreaPicturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   protectedarea.ProtectedAreaPicturesTable,
			Columns: []string{protectedarea.ProtectedAreaPicturesColumn},
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
	if pauo.mutation.ProtectedAreaCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   protectedarea.ProtectedAreaCategoryTable,
			Columns: []string{protectedarea.ProtectedAreaCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareacategory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.ProtectedAreaCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   protectedarea.ProtectedAreaCategoryTable,
			Columns: []string{protectedarea.ProtectedAreaCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareacategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProtectedArea{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{protectedarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}
