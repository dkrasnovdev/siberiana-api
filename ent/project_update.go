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
	"github.com/dkrasnovdev/heritage-api/ent/project"
	"github.com/dkrasnovdev/heritage-api/ent/projecttype"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *ProjectUpdate) SetCreatedBy(s string) *ProjectUpdate {
	pu.mutation.SetCreatedBy(s)
	return pu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCreatedBy(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetCreatedBy(*s)
	}
	return pu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (pu *ProjectUpdate) ClearCreatedBy() *ProjectUpdate {
	pu.mutation.ClearCreatedBy()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProjectUpdate) SetUpdatedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *ProjectUpdate) SetUpdatedBy(s string) *ProjectUpdate {
	pu.mutation.SetUpdatedBy(s)
	return pu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableUpdatedBy(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetUpdatedBy(*s)
	}
	return pu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pu *ProjectUpdate) ClearUpdatedBy() *ProjectUpdate {
	pu.mutation.ClearUpdatedBy()
	return pu
}

// SetAbbreviation sets the "abbreviation" field.
func (pu *ProjectUpdate) SetAbbreviation(s string) *ProjectUpdate {
	pu.mutation.SetAbbreviation(s)
	return pu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableAbbreviation(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetAbbreviation(*s)
	}
	return pu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (pu *ProjectUpdate) ClearAbbreviation() *ProjectUpdate {
	pu.mutation.ClearAbbreviation()
	return pu
}

// SetDisplayName sets the "display_name" field.
func (pu *ProjectUpdate) SetDisplayName(s string) *ProjectUpdate {
	pu.mutation.SetDisplayName(s)
	return pu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDisplayName(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetDisplayName(*s)
	}
	return pu
}

// ClearDisplayName clears the value of the "display_name" field.
func (pu *ProjectUpdate) ClearDisplayName() *ProjectUpdate {
	pu.mutation.ClearDisplayName()
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProjectUpdate) SetDescription(s string) *ProjectUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDescription(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *ProjectUpdate) ClearDescription() *ProjectUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetExternalLinks sets the "external_links" field.
func (pu *ProjectUpdate) SetExternalLinks(s []string) *ProjectUpdate {
	pu.mutation.SetExternalLinks(s)
	return pu
}

// AppendExternalLinks appends s to the "external_links" field.
func (pu *ProjectUpdate) AppendExternalLinks(s []string) *ProjectUpdate {
	pu.mutation.AppendExternalLinks(s)
	return pu
}

// ClearExternalLinks clears the value of the "external_links" field.
func (pu *ProjectUpdate) ClearExternalLinks() *ProjectUpdate {
	pu.mutation.ClearExternalLinks()
	return pu
}

// SetBeginData sets the "begin_data" field.
func (pu *ProjectUpdate) SetBeginData(t time.Time) *ProjectUpdate {
	pu.mutation.SetBeginData(t)
	return pu
}

// SetNillableBeginData sets the "begin_data" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableBeginData(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetBeginData(*t)
	}
	return pu
}

// ClearBeginData clears the value of the "begin_data" field.
func (pu *ProjectUpdate) ClearBeginData() *ProjectUpdate {
	pu.mutation.ClearBeginData()
	return pu
}

// SetEndDate sets the "end_date" field.
func (pu *ProjectUpdate) SetEndDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetEndDate(t)
	return pu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableEndDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetEndDate(*t)
	}
	return pu
}

// ClearEndDate clears the value of the "end_date" field.
func (pu *ProjectUpdate) ClearEndDate() *ProjectUpdate {
	pu.mutation.ClearEndDate()
	return pu
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (pu *ProjectUpdate) AddArtifactIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddArtifactIDs(ids...)
	return pu
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (pu *ProjectUpdate) AddArtifacts(a ...*Artifact) *ProjectUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddArtifactIDs(ids...)
}

// AddTeamIDs adds the "team" edge to the Person entity by IDs.
func (pu *ProjectUpdate) AddTeamIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddTeamIDs(ids...)
	return pu
}

// AddTeam adds the "team" edges to the Person entity.
func (pu *ProjectUpdate) AddTeam(p ...*Person) *ProjectUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddTeamIDs(ids...)
}

// SetProjectTypeID sets the "project_type" edge to the ProjectType entity by ID.
func (pu *ProjectUpdate) SetProjectTypeID(id int) *ProjectUpdate {
	pu.mutation.SetProjectTypeID(id)
	return pu
}

// SetNillableProjectTypeID sets the "project_type" edge to the ProjectType entity by ID if the given value is not nil.
func (pu *ProjectUpdate) SetNillableProjectTypeID(id *int) *ProjectUpdate {
	if id != nil {
		pu = pu.SetProjectTypeID(*id)
	}
	return pu
}

// SetProjectType sets the "project_type" edge to the ProjectType entity.
func (pu *ProjectUpdate) SetProjectType(p *ProjectType) *ProjectUpdate {
	return pu.SetProjectTypeID(p.ID)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (pu *ProjectUpdate) ClearArtifacts() *ProjectUpdate {
	pu.mutation.ClearArtifacts()
	return pu
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (pu *ProjectUpdate) RemoveArtifactIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveArtifactIDs(ids...)
	return pu
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (pu *ProjectUpdate) RemoveArtifacts(a ...*Artifact) *ProjectUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveArtifactIDs(ids...)
}

// ClearTeam clears all "team" edges to the Person entity.
func (pu *ProjectUpdate) ClearTeam() *ProjectUpdate {
	pu.mutation.ClearTeam()
	return pu
}

// RemoveTeamIDs removes the "team" edge to Person entities by IDs.
func (pu *ProjectUpdate) RemoveTeamIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveTeamIDs(ids...)
	return pu
}

// RemoveTeam removes "team" edges to Person entities.
func (pu *ProjectUpdate) RemoveTeam(p ...*Person) *ProjectUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveTeamIDs(ids...)
}

// ClearProjectType clears the "project_type" edge to the ProjectType entity.
func (pu *ProjectUpdate) ClearProjectType() *ProjectUpdate {
	pu.mutation.ClearProjectType()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProjectUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if project.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized project.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := project.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.SetField(project.FieldCreatedBy, field.TypeString, value)
	}
	if pu.mutation.CreatedByCleared() {
		_spec.ClearField(project.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.SetField(project.FieldUpdatedBy, field.TypeString, value)
	}
	if pu.mutation.UpdatedByCleared() {
		_spec.ClearField(project.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.Abbreviation(); ok {
		_spec.SetField(project.FieldAbbreviation, field.TypeString, value)
	}
	if pu.mutation.AbbreviationCleared() {
		_spec.ClearField(project.FieldAbbreviation, field.TypeString)
	}
	if value, ok := pu.mutation.DisplayName(); ok {
		_spec.SetField(project.FieldDisplayName, field.TypeString, value)
	}
	if pu.mutation.DisplayNameCleared() {
		_spec.ClearField(project.FieldDisplayName, field.TypeString)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(project.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.ExternalLinks(); ok {
		_spec.SetField(project.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, project.FieldExternalLinks, value)
		})
	}
	if pu.mutation.ExternalLinksCleared() {
		_spec.ClearField(project.FieldExternalLinks, field.TypeJSON)
	}
	if value, ok := pu.mutation.BeginData(); ok {
		_spec.SetField(project.FieldBeginData, field.TypeTime, value)
	}
	if pu.mutation.BeginDataCleared() {
		_spec.ClearField(project.FieldBeginData, field.TypeTime)
	}
	if value, ok := pu.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if pu.mutation.EndDateCleared() {
		_spec.ClearField(project.FieldEndDate, field.TypeTime)
	}
	if pu.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.ArtifactsTable,
			Columns: project.ArtifactsPrimaryKey,
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
			Table:   project.ArtifactsTable,
			Columns: project.ArtifactsPrimaryKey,
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
			Table:   project.ArtifactsTable,
			Columns: project.ArtifactsPrimaryKey,
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
	if pu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.TeamTable,
			Columns: project.TeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTeamIDs(); len(nodes) > 0 && !pu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.TeamTable,
			Columns: project.TeamPrimaryKey,
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
	if nodes := pu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.TeamTable,
			Columns: project.TeamPrimaryKey,
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
	if pu.mutation.ProjectTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.ProjectTypeTable,
			Columns: []string{project.ProjectTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProjectTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.ProjectTypeTable,
			Columns: []string{project.ProjectTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetCreatedBy sets the "created_by" field.
func (puo *ProjectUpdateOne) SetCreatedBy(s string) *ProjectUpdateOne {
	puo.mutation.SetCreatedBy(s)
	return puo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCreatedBy(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetCreatedBy(*s)
	}
	return puo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (puo *ProjectUpdateOne) ClearCreatedBy() *ProjectUpdateOne {
	puo.mutation.ClearCreatedBy()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProjectUpdateOne) SetUpdatedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *ProjectUpdateOne) SetUpdatedBy(s string) *ProjectUpdateOne {
	puo.mutation.SetUpdatedBy(s)
	return puo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableUpdatedBy(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetUpdatedBy(*s)
	}
	return puo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (puo *ProjectUpdateOne) ClearUpdatedBy() *ProjectUpdateOne {
	puo.mutation.ClearUpdatedBy()
	return puo
}

// SetAbbreviation sets the "abbreviation" field.
func (puo *ProjectUpdateOne) SetAbbreviation(s string) *ProjectUpdateOne {
	puo.mutation.SetAbbreviation(s)
	return puo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableAbbreviation(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetAbbreviation(*s)
	}
	return puo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (puo *ProjectUpdateOne) ClearAbbreviation() *ProjectUpdateOne {
	puo.mutation.ClearAbbreviation()
	return puo
}

// SetDisplayName sets the "display_name" field.
func (puo *ProjectUpdateOne) SetDisplayName(s string) *ProjectUpdateOne {
	puo.mutation.SetDisplayName(s)
	return puo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDisplayName(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetDisplayName(*s)
	}
	return puo
}

// ClearDisplayName clears the value of the "display_name" field.
func (puo *ProjectUpdateOne) ClearDisplayName() *ProjectUpdateOne {
	puo.mutation.ClearDisplayName()
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProjectUpdateOne) SetDescription(s string) *ProjectUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDescription(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *ProjectUpdateOne) ClearDescription() *ProjectUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetExternalLinks sets the "external_links" field.
func (puo *ProjectUpdateOne) SetExternalLinks(s []string) *ProjectUpdateOne {
	puo.mutation.SetExternalLinks(s)
	return puo
}

// AppendExternalLinks appends s to the "external_links" field.
func (puo *ProjectUpdateOne) AppendExternalLinks(s []string) *ProjectUpdateOne {
	puo.mutation.AppendExternalLinks(s)
	return puo
}

// ClearExternalLinks clears the value of the "external_links" field.
func (puo *ProjectUpdateOne) ClearExternalLinks() *ProjectUpdateOne {
	puo.mutation.ClearExternalLinks()
	return puo
}

// SetBeginData sets the "begin_data" field.
func (puo *ProjectUpdateOne) SetBeginData(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetBeginData(t)
	return puo
}

// SetNillableBeginData sets the "begin_data" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableBeginData(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetBeginData(*t)
	}
	return puo
}

// ClearBeginData clears the value of the "begin_data" field.
func (puo *ProjectUpdateOne) ClearBeginData() *ProjectUpdateOne {
	puo.mutation.ClearBeginData()
	return puo
}

// SetEndDate sets the "end_date" field.
func (puo *ProjectUpdateOne) SetEndDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetEndDate(t)
	return puo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableEndDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetEndDate(*t)
	}
	return puo
}

// ClearEndDate clears the value of the "end_date" field.
func (puo *ProjectUpdateOne) ClearEndDate() *ProjectUpdateOne {
	puo.mutation.ClearEndDate()
	return puo
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (puo *ProjectUpdateOne) AddArtifactIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddArtifactIDs(ids...)
	return puo
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (puo *ProjectUpdateOne) AddArtifacts(a ...*Artifact) *ProjectUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddArtifactIDs(ids...)
}

// AddTeamIDs adds the "team" edge to the Person entity by IDs.
func (puo *ProjectUpdateOne) AddTeamIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddTeamIDs(ids...)
	return puo
}

// AddTeam adds the "team" edges to the Person entity.
func (puo *ProjectUpdateOne) AddTeam(p ...*Person) *ProjectUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddTeamIDs(ids...)
}

// SetProjectTypeID sets the "project_type" edge to the ProjectType entity by ID.
func (puo *ProjectUpdateOne) SetProjectTypeID(id int) *ProjectUpdateOne {
	puo.mutation.SetProjectTypeID(id)
	return puo
}

// SetNillableProjectTypeID sets the "project_type" edge to the ProjectType entity by ID if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableProjectTypeID(id *int) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetProjectTypeID(*id)
	}
	return puo
}

// SetProjectType sets the "project_type" edge to the ProjectType entity.
func (puo *ProjectUpdateOne) SetProjectType(p *ProjectType) *ProjectUpdateOne {
	return puo.SetProjectTypeID(p.ID)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearArtifacts clears all "artifacts" edges to the Artifact entity.
func (puo *ProjectUpdateOne) ClearArtifacts() *ProjectUpdateOne {
	puo.mutation.ClearArtifacts()
	return puo
}

// RemoveArtifactIDs removes the "artifacts" edge to Artifact entities by IDs.
func (puo *ProjectUpdateOne) RemoveArtifactIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveArtifactIDs(ids...)
	return puo
}

// RemoveArtifacts removes "artifacts" edges to Artifact entities.
func (puo *ProjectUpdateOne) RemoveArtifacts(a ...*Artifact) *ProjectUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveArtifactIDs(ids...)
}

// ClearTeam clears all "team" edges to the Person entity.
func (puo *ProjectUpdateOne) ClearTeam() *ProjectUpdateOne {
	puo.mutation.ClearTeam()
	return puo
}

// RemoveTeamIDs removes the "team" edge to Person entities by IDs.
func (puo *ProjectUpdateOne) RemoveTeamIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveTeamIDs(ids...)
	return puo
}

// RemoveTeam removes "team" edges to Person entities.
func (puo *ProjectUpdateOne) RemoveTeam(p ...*Person) *ProjectUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveTeamIDs(ids...)
}

// ClearProjectType clears the "project_type" edge to the ProjectType entity.
func (puo *ProjectUpdateOne) ClearProjectType() *ProjectUpdateOne {
	puo.mutation.ClearProjectType()
	return puo
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProjectUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if project.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized project.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := project.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
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
		_spec.SetField(project.FieldCreatedBy, field.TypeString, value)
	}
	if puo.mutation.CreatedByCleared() {
		_spec.ClearField(project.FieldCreatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.SetField(project.FieldUpdatedBy, field.TypeString, value)
	}
	if puo.mutation.UpdatedByCleared() {
		_spec.ClearField(project.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.Abbreviation(); ok {
		_spec.SetField(project.FieldAbbreviation, field.TypeString, value)
	}
	if puo.mutation.AbbreviationCleared() {
		_spec.ClearField(project.FieldAbbreviation, field.TypeString)
	}
	if value, ok := puo.mutation.DisplayName(); ok {
		_spec.SetField(project.FieldDisplayName, field.TypeString, value)
	}
	if puo.mutation.DisplayNameCleared() {
		_spec.ClearField(project.FieldDisplayName, field.TypeString)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(project.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.ExternalLinks(); ok {
		_spec.SetField(project.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, project.FieldExternalLinks, value)
		})
	}
	if puo.mutation.ExternalLinksCleared() {
		_spec.ClearField(project.FieldExternalLinks, field.TypeJSON)
	}
	if value, ok := puo.mutation.BeginData(); ok {
		_spec.SetField(project.FieldBeginData, field.TypeTime, value)
	}
	if puo.mutation.BeginDataCleared() {
		_spec.ClearField(project.FieldBeginData, field.TypeTime)
	}
	if value, ok := puo.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if puo.mutation.EndDateCleared() {
		_spec.ClearField(project.FieldEndDate, field.TypeTime)
	}
	if puo.mutation.ArtifactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.ArtifactsTable,
			Columns: project.ArtifactsPrimaryKey,
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
			Table:   project.ArtifactsTable,
			Columns: project.ArtifactsPrimaryKey,
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
			Table:   project.ArtifactsTable,
			Columns: project.ArtifactsPrimaryKey,
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
	if puo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.TeamTable,
			Columns: project.TeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTeamIDs(); len(nodes) > 0 && !puo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.TeamTable,
			Columns: project.TeamPrimaryKey,
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
	if nodes := puo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.TeamTable,
			Columns: project.TeamPrimaryKey,
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
	if puo.mutation.ProjectTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.ProjectTypeTable,
			Columns: []string{project.ProjectTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProjectTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.ProjectTypeTable,
			Columns: []string{project.ProjectTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
