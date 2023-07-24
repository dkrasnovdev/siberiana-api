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
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
	"github.com/dkrasnovdev/heritage-api/ent/project"
	"github.com/dkrasnovdev/heritage-api/ent/projecttype"
)

// ProjectTypeUpdate is the builder for updating ProjectType entities.
type ProjectTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectTypeMutation
}

// Where appends a list predicates to the ProjectTypeUpdate builder.
func (ptu *ProjectTypeUpdate) Where(ps ...predicate.ProjectType) *ProjectTypeUpdate {
	ptu.mutation.Where(ps...)
	return ptu
}

// SetCreatedBy sets the "created_by" field.
func (ptu *ProjectTypeUpdate) SetCreatedBy(s string) *ProjectTypeUpdate {
	ptu.mutation.SetCreatedBy(s)
	return ptu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ptu *ProjectTypeUpdate) SetNillableCreatedBy(s *string) *ProjectTypeUpdate {
	if s != nil {
		ptu.SetCreatedBy(*s)
	}
	return ptu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ptu *ProjectTypeUpdate) ClearCreatedBy() *ProjectTypeUpdate {
	ptu.mutation.ClearCreatedBy()
	return ptu
}

// SetUpdatedAt sets the "updated_at" field.
func (ptu *ProjectTypeUpdate) SetUpdatedAt(t time.Time) *ProjectTypeUpdate {
	ptu.mutation.SetUpdatedAt(t)
	return ptu
}

// SetUpdatedBy sets the "updated_by" field.
func (ptu *ProjectTypeUpdate) SetUpdatedBy(s string) *ProjectTypeUpdate {
	ptu.mutation.SetUpdatedBy(s)
	return ptu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ptu *ProjectTypeUpdate) SetNillableUpdatedBy(s *string) *ProjectTypeUpdate {
	if s != nil {
		ptu.SetUpdatedBy(*s)
	}
	return ptu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ptu *ProjectTypeUpdate) ClearUpdatedBy() *ProjectTypeUpdate {
	ptu.mutation.ClearUpdatedBy()
	return ptu
}

// SetAbbreviation sets the "abbreviation" field.
func (ptu *ProjectTypeUpdate) SetAbbreviation(s string) *ProjectTypeUpdate {
	ptu.mutation.SetAbbreviation(s)
	return ptu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ptu *ProjectTypeUpdate) SetNillableAbbreviation(s *string) *ProjectTypeUpdate {
	if s != nil {
		ptu.SetAbbreviation(*s)
	}
	return ptu
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (ptu *ProjectTypeUpdate) ClearAbbreviation() *ProjectTypeUpdate {
	ptu.mutation.ClearAbbreviation()
	return ptu
}

// SetDisplayName sets the "display_name" field.
func (ptu *ProjectTypeUpdate) SetDisplayName(s string) *ProjectTypeUpdate {
	ptu.mutation.SetDisplayName(s)
	return ptu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ptu *ProjectTypeUpdate) SetNillableDisplayName(s *string) *ProjectTypeUpdate {
	if s != nil {
		ptu.SetDisplayName(*s)
	}
	return ptu
}

// ClearDisplayName clears the value of the "display_name" field.
func (ptu *ProjectTypeUpdate) ClearDisplayName() *ProjectTypeUpdate {
	ptu.mutation.ClearDisplayName()
	return ptu
}

// SetDescription sets the "description" field.
func (ptu *ProjectTypeUpdate) SetDescription(s string) *ProjectTypeUpdate {
	ptu.mutation.SetDescription(s)
	return ptu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ptu *ProjectTypeUpdate) SetNillableDescription(s *string) *ProjectTypeUpdate {
	if s != nil {
		ptu.SetDescription(*s)
	}
	return ptu
}

// ClearDescription clears the value of the "description" field.
func (ptu *ProjectTypeUpdate) ClearDescription() *ProjectTypeUpdate {
	ptu.mutation.ClearDescription()
	return ptu
}

// SetExternalLinks sets the "external_links" field.
func (ptu *ProjectTypeUpdate) SetExternalLinks(s []string) *ProjectTypeUpdate {
	ptu.mutation.SetExternalLinks(s)
	return ptu
}

// AppendExternalLinks appends s to the "external_links" field.
func (ptu *ProjectTypeUpdate) AppendExternalLinks(s []string) *ProjectTypeUpdate {
	ptu.mutation.AppendExternalLinks(s)
	return ptu
}

// ClearExternalLinks clears the value of the "external_links" field.
func (ptu *ProjectTypeUpdate) ClearExternalLinks() *ProjectTypeUpdate {
	ptu.mutation.ClearExternalLinks()
	return ptu
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (ptu *ProjectTypeUpdate) AddProjectIDs(ids ...int) *ProjectTypeUpdate {
	ptu.mutation.AddProjectIDs(ids...)
	return ptu
}

// AddProjects adds the "projects" edges to the Project entity.
func (ptu *ProjectTypeUpdate) AddProjects(p ...*Project) *ProjectTypeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptu.AddProjectIDs(ids...)
}

// Mutation returns the ProjectTypeMutation object of the builder.
func (ptu *ProjectTypeUpdate) Mutation() *ProjectTypeMutation {
	return ptu.mutation
}

// ClearProjects clears all "projects" edges to the Project entity.
func (ptu *ProjectTypeUpdate) ClearProjects() *ProjectTypeUpdate {
	ptu.mutation.ClearProjects()
	return ptu
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (ptu *ProjectTypeUpdate) RemoveProjectIDs(ids ...int) *ProjectTypeUpdate {
	ptu.mutation.RemoveProjectIDs(ids...)
	return ptu
}

// RemoveProjects removes "projects" edges to Project entities.
func (ptu *ProjectTypeUpdate) RemoveProjects(p ...*Project) *ProjectTypeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptu.RemoveProjectIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *ProjectTypeUpdate) Save(ctx context.Context) (int, error) {
	if err := ptu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ptu.sqlSave, ptu.mutation, ptu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *ProjectTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *ProjectTypeUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *ProjectTypeUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptu *ProjectTypeUpdate) defaults() error {
	if _, ok := ptu.mutation.UpdatedAt(); !ok {
		if projecttype.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized projecttype.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := projecttype.UpdateDefaultUpdatedAt()
		ptu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ptu *ProjectTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(projecttype.Table, projecttype.Columns, sqlgraph.NewFieldSpec(projecttype.FieldID, field.TypeInt))
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.CreatedBy(); ok {
		_spec.SetField(projecttype.FieldCreatedBy, field.TypeString, value)
	}
	if ptu.mutation.CreatedByCleared() {
		_spec.ClearField(projecttype.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ptu.mutation.UpdatedAt(); ok {
		_spec.SetField(projecttype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ptu.mutation.UpdatedBy(); ok {
		_spec.SetField(projecttype.FieldUpdatedBy, field.TypeString, value)
	}
	if ptu.mutation.UpdatedByCleared() {
		_spec.ClearField(projecttype.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ptu.mutation.Abbreviation(); ok {
		_spec.SetField(projecttype.FieldAbbreviation, field.TypeString, value)
	}
	if ptu.mutation.AbbreviationCleared() {
		_spec.ClearField(projecttype.FieldAbbreviation, field.TypeString)
	}
	if value, ok := ptu.mutation.DisplayName(); ok {
		_spec.SetField(projecttype.FieldDisplayName, field.TypeString, value)
	}
	if ptu.mutation.DisplayNameCleared() {
		_spec.ClearField(projecttype.FieldDisplayName, field.TypeString)
	}
	if value, ok := ptu.mutation.Description(); ok {
		_spec.SetField(projecttype.FieldDescription, field.TypeString, value)
	}
	if ptu.mutation.DescriptionCleared() {
		_spec.ClearField(projecttype.FieldDescription, field.TypeString)
	}
	if value, ok := ptu.mutation.ExternalLinks(); ok {
		_spec.SetField(projecttype.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := ptu.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, projecttype.FieldExternalLinks, value)
		})
	}
	if ptu.mutation.ExternalLinksCleared() {
		_spec.ClearField(projecttype.FieldExternalLinks, field.TypeJSON)
	}
	if ptu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttype.ProjectsTable,
			Columns: []string{projecttype.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !ptu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttype.ProjectsTable,
			Columns: []string{projecttype.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttype.ProjectsTable,
			Columns: []string{projecttype.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ptu.mutation.done = true
	return n, nil
}

// ProjectTypeUpdateOne is the builder for updating a single ProjectType entity.
type ProjectTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectTypeMutation
}

// SetCreatedBy sets the "created_by" field.
func (ptuo *ProjectTypeUpdateOne) SetCreatedBy(s string) *ProjectTypeUpdateOne {
	ptuo.mutation.SetCreatedBy(s)
	return ptuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ptuo *ProjectTypeUpdateOne) SetNillableCreatedBy(s *string) *ProjectTypeUpdateOne {
	if s != nil {
		ptuo.SetCreatedBy(*s)
	}
	return ptuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ptuo *ProjectTypeUpdateOne) ClearCreatedBy() *ProjectTypeUpdateOne {
	ptuo.mutation.ClearCreatedBy()
	return ptuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ptuo *ProjectTypeUpdateOne) SetUpdatedAt(t time.Time) *ProjectTypeUpdateOne {
	ptuo.mutation.SetUpdatedAt(t)
	return ptuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ptuo *ProjectTypeUpdateOne) SetUpdatedBy(s string) *ProjectTypeUpdateOne {
	ptuo.mutation.SetUpdatedBy(s)
	return ptuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ptuo *ProjectTypeUpdateOne) SetNillableUpdatedBy(s *string) *ProjectTypeUpdateOne {
	if s != nil {
		ptuo.SetUpdatedBy(*s)
	}
	return ptuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ptuo *ProjectTypeUpdateOne) ClearUpdatedBy() *ProjectTypeUpdateOne {
	ptuo.mutation.ClearUpdatedBy()
	return ptuo
}

// SetAbbreviation sets the "abbreviation" field.
func (ptuo *ProjectTypeUpdateOne) SetAbbreviation(s string) *ProjectTypeUpdateOne {
	ptuo.mutation.SetAbbreviation(s)
	return ptuo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ptuo *ProjectTypeUpdateOne) SetNillableAbbreviation(s *string) *ProjectTypeUpdateOne {
	if s != nil {
		ptuo.SetAbbreviation(*s)
	}
	return ptuo
}

// ClearAbbreviation clears the value of the "abbreviation" field.
func (ptuo *ProjectTypeUpdateOne) ClearAbbreviation() *ProjectTypeUpdateOne {
	ptuo.mutation.ClearAbbreviation()
	return ptuo
}

// SetDisplayName sets the "display_name" field.
func (ptuo *ProjectTypeUpdateOne) SetDisplayName(s string) *ProjectTypeUpdateOne {
	ptuo.mutation.SetDisplayName(s)
	return ptuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ptuo *ProjectTypeUpdateOne) SetNillableDisplayName(s *string) *ProjectTypeUpdateOne {
	if s != nil {
		ptuo.SetDisplayName(*s)
	}
	return ptuo
}

// ClearDisplayName clears the value of the "display_name" field.
func (ptuo *ProjectTypeUpdateOne) ClearDisplayName() *ProjectTypeUpdateOne {
	ptuo.mutation.ClearDisplayName()
	return ptuo
}

// SetDescription sets the "description" field.
func (ptuo *ProjectTypeUpdateOne) SetDescription(s string) *ProjectTypeUpdateOne {
	ptuo.mutation.SetDescription(s)
	return ptuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ptuo *ProjectTypeUpdateOne) SetNillableDescription(s *string) *ProjectTypeUpdateOne {
	if s != nil {
		ptuo.SetDescription(*s)
	}
	return ptuo
}

// ClearDescription clears the value of the "description" field.
func (ptuo *ProjectTypeUpdateOne) ClearDescription() *ProjectTypeUpdateOne {
	ptuo.mutation.ClearDescription()
	return ptuo
}

// SetExternalLinks sets the "external_links" field.
func (ptuo *ProjectTypeUpdateOne) SetExternalLinks(s []string) *ProjectTypeUpdateOne {
	ptuo.mutation.SetExternalLinks(s)
	return ptuo
}

// AppendExternalLinks appends s to the "external_links" field.
func (ptuo *ProjectTypeUpdateOne) AppendExternalLinks(s []string) *ProjectTypeUpdateOne {
	ptuo.mutation.AppendExternalLinks(s)
	return ptuo
}

// ClearExternalLinks clears the value of the "external_links" field.
func (ptuo *ProjectTypeUpdateOne) ClearExternalLinks() *ProjectTypeUpdateOne {
	ptuo.mutation.ClearExternalLinks()
	return ptuo
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (ptuo *ProjectTypeUpdateOne) AddProjectIDs(ids ...int) *ProjectTypeUpdateOne {
	ptuo.mutation.AddProjectIDs(ids...)
	return ptuo
}

// AddProjects adds the "projects" edges to the Project entity.
func (ptuo *ProjectTypeUpdateOne) AddProjects(p ...*Project) *ProjectTypeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptuo.AddProjectIDs(ids...)
}

// Mutation returns the ProjectTypeMutation object of the builder.
func (ptuo *ProjectTypeUpdateOne) Mutation() *ProjectTypeMutation {
	return ptuo.mutation
}

// ClearProjects clears all "projects" edges to the Project entity.
func (ptuo *ProjectTypeUpdateOne) ClearProjects() *ProjectTypeUpdateOne {
	ptuo.mutation.ClearProjects()
	return ptuo
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (ptuo *ProjectTypeUpdateOne) RemoveProjectIDs(ids ...int) *ProjectTypeUpdateOne {
	ptuo.mutation.RemoveProjectIDs(ids...)
	return ptuo
}

// RemoveProjects removes "projects" edges to Project entities.
func (ptuo *ProjectTypeUpdateOne) RemoveProjects(p ...*Project) *ProjectTypeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptuo.RemoveProjectIDs(ids...)
}

// Where appends a list predicates to the ProjectTypeUpdate builder.
func (ptuo *ProjectTypeUpdateOne) Where(ps ...predicate.ProjectType) *ProjectTypeUpdateOne {
	ptuo.mutation.Where(ps...)
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *ProjectTypeUpdateOne) Select(field string, fields ...string) *ProjectTypeUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated ProjectType entity.
func (ptuo *ProjectTypeUpdateOne) Save(ctx context.Context) (*ProjectType, error) {
	if err := ptuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ptuo.sqlSave, ptuo.mutation, ptuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *ProjectTypeUpdateOne) SaveX(ctx context.Context) *ProjectType {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *ProjectTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *ProjectTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptuo *ProjectTypeUpdateOne) defaults() error {
	if _, ok := ptuo.mutation.UpdatedAt(); !ok {
		if projecttype.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized projecttype.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := projecttype.UpdateDefaultUpdatedAt()
		ptuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ptuo *ProjectTypeUpdateOne) sqlSave(ctx context.Context) (_node *ProjectType, err error) {
	_spec := sqlgraph.NewUpdateSpec(projecttype.Table, projecttype.Columns, sqlgraph.NewFieldSpec(projecttype.FieldID, field.TypeInt))
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProjectType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projecttype.FieldID)
		for _, f := range fields {
			if !projecttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projecttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.CreatedBy(); ok {
		_spec.SetField(projecttype.FieldCreatedBy, field.TypeString, value)
	}
	if ptuo.mutation.CreatedByCleared() {
		_spec.ClearField(projecttype.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ptuo.mutation.UpdatedAt(); ok {
		_spec.SetField(projecttype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ptuo.mutation.UpdatedBy(); ok {
		_spec.SetField(projecttype.FieldUpdatedBy, field.TypeString, value)
	}
	if ptuo.mutation.UpdatedByCleared() {
		_spec.ClearField(projecttype.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ptuo.mutation.Abbreviation(); ok {
		_spec.SetField(projecttype.FieldAbbreviation, field.TypeString, value)
	}
	if ptuo.mutation.AbbreviationCleared() {
		_spec.ClearField(projecttype.FieldAbbreviation, field.TypeString)
	}
	if value, ok := ptuo.mutation.DisplayName(); ok {
		_spec.SetField(projecttype.FieldDisplayName, field.TypeString, value)
	}
	if ptuo.mutation.DisplayNameCleared() {
		_spec.ClearField(projecttype.FieldDisplayName, field.TypeString)
	}
	if value, ok := ptuo.mutation.Description(); ok {
		_spec.SetField(projecttype.FieldDescription, field.TypeString, value)
	}
	if ptuo.mutation.DescriptionCleared() {
		_spec.ClearField(projecttype.FieldDescription, field.TypeString)
	}
	if value, ok := ptuo.mutation.ExternalLinks(); ok {
		_spec.SetField(projecttype.FieldExternalLinks, field.TypeJSON, value)
	}
	if value, ok := ptuo.mutation.AppendedExternalLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, projecttype.FieldExternalLinks, value)
		})
	}
	if ptuo.mutation.ExternalLinksCleared() {
		_spec.ClearField(projecttype.FieldExternalLinks, field.TypeJSON)
	}
	if ptuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttype.ProjectsTable,
			Columns: []string{projecttype.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !ptuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttype.ProjectsTable,
			Columns: []string{projecttype.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttype.ProjectsTable,
			Columns: []string{projecttype.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectType{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ptuo.mutation.done = true
	return _node, nil
}
