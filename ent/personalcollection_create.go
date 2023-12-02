// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/art"
	"github.com/dkrasnovdev/siberiana-api/ent/artifact"
	"github.com/dkrasnovdev/siberiana-api/ent/book"
	"github.com/dkrasnovdev/siberiana-api/ent/personalcollection"
	"github.com/dkrasnovdev/siberiana-api/ent/petroglyph"
	"github.com/dkrasnovdev/siberiana-api/ent/protectedareapicture"
)

// PersonalCollectionCreate is the builder for creating a PersonalCollection entity.
type PersonalCollectionCreate struct {
	config
	mutation *PersonalCollectionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pcc *PersonalCollectionCreate) SetCreatedAt(t time.Time) *PersonalCollectionCreate {
	pcc.mutation.SetCreatedAt(t)
	return pcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcc *PersonalCollectionCreate) SetNillableCreatedAt(t *time.Time) *PersonalCollectionCreate {
	if t != nil {
		pcc.SetCreatedAt(*t)
	}
	return pcc
}

// SetCreatedBy sets the "created_by" field.
func (pcc *PersonalCollectionCreate) SetCreatedBy(s string) *PersonalCollectionCreate {
	pcc.mutation.SetCreatedBy(s)
	return pcc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pcc *PersonalCollectionCreate) SetNillableCreatedBy(s *string) *PersonalCollectionCreate {
	if s != nil {
		pcc.SetCreatedBy(*s)
	}
	return pcc
}

// SetUpdatedAt sets the "updated_at" field.
func (pcc *PersonalCollectionCreate) SetUpdatedAt(t time.Time) *PersonalCollectionCreate {
	pcc.mutation.SetUpdatedAt(t)
	return pcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pcc *PersonalCollectionCreate) SetNillableUpdatedAt(t *time.Time) *PersonalCollectionCreate {
	if t != nil {
		pcc.SetUpdatedAt(*t)
	}
	return pcc
}

// SetUpdatedBy sets the "updated_by" field.
func (pcc *PersonalCollectionCreate) SetUpdatedBy(s string) *PersonalCollectionCreate {
	pcc.mutation.SetUpdatedBy(s)
	return pcc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pcc *PersonalCollectionCreate) SetNillableUpdatedBy(s *string) *PersonalCollectionCreate {
	if s != nil {
		pcc.SetUpdatedBy(*s)
	}
	return pcc
}

// SetDisplayName sets the "display_name" field.
func (pcc *PersonalCollectionCreate) SetDisplayName(s string) *PersonalCollectionCreate {
	pcc.mutation.SetDisplayName(s)
	return pcc
}

// SetIsPublic sets the "is_public" field.
func (pcc *PersonalCollectionCreate) SetIsPublic(b bool) *PersonalCollectionCreate {
	pcc.mutation.SetIsPublic(b)
	return pcc
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (pcc *PersonalCollectionCreate) SetNillableIsPublic(b *bool) *PersonalCollectionCreate {
	if b != nil {
		pcc.SetIsPublic(*b)
	}
	return pcc
}

// AddArtIDs adds the "art" edge to the Art entity by IDs.
func (pcc *PersonalCollectionCreate) AddArtIDs(ids ...int) *PersonalCollectionCreate {
	pcc.mutation.AddArtIDs(ids...)
	return pcc
}

// AddArt adds the "art" edges to the Art entity.
func (pcc *PersonalCollectionCreate) AddArt(a ...*Art) *PersonalCollectionCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pcc.AddArtIDs(ids...)
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (pcc *PersonalCollectionCreate) AddArtifactIDs(ids ...int) *PersonalCollectionCreate {
	pcc.mutation.AddArtifactIDs(ids...)
	return pcc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (pcc *PersonalCollectionCreate) AddArtifacts(a ...*Artifact) *PersonalCollectionCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pcc.AddArtifactIDs(ids...)
}

// AddPetroglyphIDs adds the "petroglyphs" edge to the Petroglyph entity by IDs.
func (pcc *PersonalCollectionCreate) AddPetroglyphIDs(ids ...int) *PersonalCollectionCreate {
	pcc.mutation.AddPetroglyphIDs(ids...)
	return pcc
}

// AddPetroglyphs adds the "petroglyphs" edges to the Petroglyph entity.
func (pcc *PersonalCollectionCreate) AddPetroglyphs(p ...*Petroglyph) *PersonalCollectionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcc.AddPetroglyphIDs(ids...)
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (pcc *PersonalCollectionCreate) AddBookIDs(ids ...int) *PersonalCollectionCreate {
	pcc.mutation.AddBookIDs(ids...)
	return pcc
}

// AddBooks adds the "books" edges to the Book entity.
func (pcc *PersonalCollectionCreate) AddBooks(b ...*Book) *PersonalCollectionCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pcc.AddBookIDs(ids...)
}

// AddProtectedAreaPictureIDs adds the "protected_area_pictures" edge to the ProtectedAreaPicture entity by IDs.
func (pcc *PersonalCollectionCreate) AddProtectedAreaPictureIDs(ids ...int) *PersonalCollectionCreate {
	pcc.mutation.AddProtectedAreaPictureIDs(ids...)
	return pcc
}

// AddProtectedAreaPictures adds the "protected_area_pictures" edges to the ProtectedAreaPicture entity.
func (pcc *PersonalCollectionCreate) AddProtectedAreaPictures(p ...*ProtectedAreaPicture) *PersonalCollectionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcc.AddProtectedAreaPictureIDs(ids...)
}

// Mutation returns the PersonalCollectionMutation object of the builder.
func (pcc *PersonalCollectionCreate) Mutation() *PersonalCollectionMutation {
	return pcc.mutation
}

// Save creates the PersonalCollection in the database.
func (pcc *PersonalCollectionCreate) Save(ctx context.Context) (*PersonalCollection, error) {
	if err := pcc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pcc.sqlSave, pcc.mutation, pcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *PersonalCollectionCreate) SaveX(ctx context.Context) *PersonalCollection {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *PersonalCollectionCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *PersonalCollectionCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcc *PersonalCollectionCreate) defaults() error {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		if personalcollection.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized personalcollection.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := personalcollection.DefaultCreatedAt()
		pcc.mutation.SetCreatedAt(v)
	}
	if _, ok := pcc.mutation.UpdatedAt(); !ok {
		if personalcollection.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized personalcollection.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := personalcollection.DefaultUpdatedAt()
		pcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pcc.mutation.IsPublic(); !ok {
		v := personalcollection.DefaultIsPublic
		pcc.mutation.SetIsPublic(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pcc *PersonalCollectionCreate) check() error {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PersonalCollection.created_at"`)}
	}
	if _, ok := pcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PersonalCollection.updated_at"`)}
	}
	if _, ok := pcc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "PersonalCollection.display_name"`)}
	}
	if v, ok := pcc.mutation.DisplayName(); ok {
		if err := personalcollection.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "PersonalCollection.display_name": %w`, err)}
		}
	}
	if _, ok := pcc.mutation.IsPublic(); !ok {
		return &ValidationError{Name: "is_public", err: errors.New(`ent: missing required field "PersonalCollection.is_public"`)}
	}
	return nil
}

func (pcc *PersonalCollectionCreate) sqlSave(ctx context.Context) (*PersonalCollection, error) {
	if err := pcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pcc.mutation.id = &_node.ID
	pcc.mutation.done = true
	return _node, nil
}

func (pcc *PersonalCollectionCreate) createSpec() (*PersonalCollection, *sqlgraph.CreateSpec) {
	var (
		_node = &PersonalCollection{config: pcc.config}
		_spec = sqlgraph.NewCreateSpec(personalcollection.Table, sqlgraph.NewFieldSpec(personalcollection.FieldID, field.TypeInt))
	)
	if value, ok := pcc.mutation.CreatedAt(); ok {
		_spec.SetField(personalcollection.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pcc.mutation.CreatedBy(); ok {
		_spec.SetField(personalcollection.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pcc.mutation.UpdatedAt(); ok {
		_spec.SetField(personalcollection.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pcc.mutation.UpdatedBy(); ok {
		_spec.SetField(personalcollection.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pcc.mutation.DisplayName(); ok {
		_spec.SetField(personalcollection.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := pcc.mutation.IsPublic(); ok {
		_spec.SetField(personalcollection.FieldIsPublic, field.TypeBool, value)
		_node.IsPublic = value
	}
	if nodes := pcc.mutation.ArtIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   personalcollection.ArtTable,
			Columns: personalcollection.ArtPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   personalcollection.ArtifactsTable,
			Columns: personalcollection.ArtifactsPrimaryKey,
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
	if nodes := pcc.mutation.PetroglyphsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   personalcollection.PetroglyphsTable,
			Columns: personalcollection.PetroglyphsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(petroglyph.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcc.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   personalcollection.BooksTable,
			Columns: personalcollection.BooksPrimaryKey,
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
	if nodes := pcc.mutation.ProtectedAreaPicturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   personalcollection.ProtectedAreaPicturesTable,
			Columns: personalcollection.ProtectedAreaPicturesPrimaryKey,
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

// PersonalCollectionCreateBulk is the builder for creating many PersonalCollection entities in bulk.
type PersonalCollectionCreateBulk struct {
	config
	err      error
	builders []*PersonalCollectionCreate
}

// Save creates the PersonalCollection entities in the database.
func (pccb *PersonalCollectionCreateBulk) Save(ctx context.Context) ([]*PersonalCollection, error) {
	if pccb.err != nil {
		return nil, pccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*PersonalCollection, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonalCollectionMutation)
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
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *PersonalCollectionCreateBulk) SaveX(ctx context.Context) []*PersonalCollection {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *PersonalCollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *PersonalCollectionCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
