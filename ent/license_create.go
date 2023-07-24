// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/artifact"
	"github.com/dkrasnovdev/heritage-api/ent/book"
	"github.com/dkrasnovdev/heritage-api/ent/license"
	"github.com/dkrasnovdev/heritage-api/ent/protectedareapicture"
)

// LicenseCreate is the builder for creating a License entity.
type LicenseCreate struct {
	config
	mutation *LicenseMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lc *LicenseCreate) SetCreatedAt(t time.Time) *LicenseCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableCreatedAt(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetCreatedBy sets the "created_by" field.
func (lc *LicenseCreate) SetCreatedBy(s string) *LicenseCreate {
	lc.mutation.SetCreatedBy(s)
	return lc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableCreatedBy(s *string) *LicenseCreate {
	if s != nil {
		lc.SetCreatedBy(*s)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LicenseCreate) SetUpdatedAt(t time.Time) *LicenseCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableUpdatedAt(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetUpdatedBy sets the "updated_by" field.
func (lc *LicenseCreate) SetUpdatedBy(s string) *LicenseCreate {
	lc.mutation.SetUpdatedBy(s)
	return lc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableUpdatedBy(s *string) *LicenseCreate {
	if s != nil {
		lc.SetUpdatedBy(*s)
	}
	return lc
}

// SetDisplayName sets the "display_name" field.
func (lc *LicenseCreate) SetDisplayName(s string) *LicenseCreate {
	lc.mutation.SetDisplayName(s)
	return lc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableDisplayName(s *string) *LicenseCreate {
	if s != nil {
		lc.SetDisplayName(*s)
	}
	return lc
}

// SetAbbreviation sets the "abbreviation" field.
func (lc *LicenseCreate) SetAbbreviation(s string) *LicenseCreate {
	lc.mutation.SetAbbreviation(s)
	return lc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableAbbreviation(s *string) *LicenseCreate {
	if s != nil {
		lc.SetAbbreviation(*s)
	}
	return lc
}

// SetDescription sets the "description" field.
func (lc *LicenseCreate) SetDescription(s string) *LicenseCreate {
	lc.mutation.SetDescription(s)
	return lc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableDescription(s *string) *LicenseCreate {
	if s != nil {
		lc.SetDescription(*s)
	}
	return lc
}

// SetExternalLinks sets the "external_links" field.
func (lc *LicenseCreate) SetExternalLinks(s []string) *LicenseCreate {
	lc.mutation.SetExternalLinks(s)
	return lc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (lc *LicenseCreate) AddArtifactIDs(ids ...int) *LicenseCreate {
	lc.mutation.AddArtifactIDs(ids...)
	return lc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (lc *LicenseCreate) AddArtifacts(a ...*Artifact) *LicenseCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lc.AddArtifactIDs(ids...)
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (lc *LicenseCreate) AddBookIDs(ids ...int) *LicenseCreate {
	lc.mutation.AddBookIDs(ids...)
	return lc
}

// AddBooks adds the "books" edges to the Book entity.
func (lc *LicenseCreate) AddBooks(b ...*Book) *LicenseCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return lc.AddBookIDs(ids...)
}

// AddProtectedAreaPictureIDs adds the "protected_area_pictures" edge to the ProtectedAreaPicture entity by IDs.
func (lc *LicenseCreate) AddProtectedAreaPictureIDs(ids ...int) *LicenseCreate {
	lc.mutation.AddProtectedAreaPictureIDs(ids...)
	return lc
}

// AddProtectedAreaPictures adds the "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (lc *LicenseCreate) AddProtectedAreaPictures(p ...*ProtectedAreaPicture) *LicenseCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lc.AddProtectedAreaPictureIDs(ids...)
}

// Mutation returns the LicenseMutation object of the builder.
func (lc *LicenseCreate) Mutation() *LicenseMutation {
	return lc.mutation
}

// Save creates the License in the database.
func (lc *LicenseCreate) Save(ctx context.Context) (*License, error) {
	if err := lc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LicenseCreate) SaveX(ctx context.Context) *License {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LicenseCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LicenseCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LicenseCreate) defaults() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		if license.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized license.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := license.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		if license.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized license.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := license.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lc *LicenseCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "License.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "License.updated_at"`)}
	}
	return nil
}

func (lc *LicenseCreate) sqlSave(ctx context.Context) (*License, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LicenseCreate) createSpec() (*License, *sqlgraph.CreateSpec) {
	var (
		_node = &License{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(license.Table, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(license.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.CreatedBy(); ok {
		_spec.SetField(license.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(license.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.UpdatedBy(); ok {
		_spec.SetField(license.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := lc.mutation.DisplayName(); ok {
		_spec.SetField(license.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := lc.mutation.Abbreviation(); ok {
		_spec.SetField(license.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := lc.mutation.Description(); ok {
		_spec.SetField(license.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := lc.mutation.ExternalLinks(); ok {
		_spec.SetField(license.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
	if nodes := lc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ArtifactsTable,
			Columns: []string{license.ArtifactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.BooksTable,
			Columns: []string{license.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ProtectedAreaPicturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   license.ProtectedAreaPicturesTable,
			Columns: []string{license.ProtectedAreaPicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(protectedareapicture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LicenseCreateBulk is the builder for creating many License entities in bulk.
type LicenseCreateBulk struct {
	config
	builders []*LicenseCreate
}

// Save creates the License entities in the database.
func (lcb *LicenseCreateBulk) Save(ctx context.Context) ([]*License, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*License, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LicenseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LicenseCreateBulk) SaveX(ctx context.Context) []*License {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LicenseCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LicenseCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
