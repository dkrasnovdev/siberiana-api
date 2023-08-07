// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/art"
	"github.com/dkrasnovdev/heritage-api/ent/artgenre"
	"github.com/dkrasnovdev/heritage-api/ent/artstyle"
)

// ArtCreate is the builder for creating a Art entity.
type ArtCreate struct {
	config
	mutation *ArtMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *ArtCreate) SetCreatedAt(t time.Time) *ArtCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ArtCreate) SetNillableCreatedAt(t *time.Time) *ArtCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetCreatedBy sets the "created_by" field.
func (ac *ArtCreate) SetCreatedBy(s string) *ArtCreate {
	ac.mutation.SetCreatedBy(s)
	return ac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ac *ArtCreate) SetNillableCreatedBy(s *string) *ArtCreate {
	if s != nil {
		ac.SetCreatedBy(*s)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ArtCreate) SetUpdatedAt(t time.Time) *ArtCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ArtCreate) SetNillableUpdatedAt(t *time.Time) *ArtCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetUpdatedBy sets the "updated_by" field.
func (ac *ArtCreate) SetUpdatedBy(s string) *ArtCreate {
	ac.mutation.SetUpdatedBy(s)
	return ac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ac *ArtCreate) SetNillableUpdatedBy(s *string) *ArtCreate {
	if s != nil {
		ac.SetUpdatedBy(*s)
	}
	return ac
}

// SetDisplayName sets the "display_name" field.
func (ac *ArtCreate) SetDisplayName(s string) *ArtCreate {
	ac.mutation.SetDisplayName(s)
	return ac
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ac *ArtCreate) SetNillableDisplayName(s *string) *ArtCreate {
	if s != nil {
		ac.SetDisplayName(*s)
	}
	return ac
}

// SetAbbreviation sets the "abbreviation" field.
func (ac *ArtCreate) SetAbbreviation(s string) *ArtCreate {
	ac.mutation.SetAbbreviation(s)
	return ac
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ac *ArtCreate) SetNillableAbbreviation(s *string) *ArtCreate {
	if s != nil {
		ac.SetAbbreviation(*s)
	}
	return ac
}

// SetDescription sets the "description" field.
func (ac *ArtCreate) SetDescription(s string) *ArtCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *ArtCreate) SetNillableDescription(s *string) *ArtCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetExternalLink sets the "external_link" field.
func (ac *ArtCreate) SetExternalLink(s string) *ArtCreate {
	ac.mutation.SetExternalLink(s)
	return ac
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (ac *ArtCreate) SetNillableExternalLink(s *string) *ArtCreate {
	if s != nil {
		ac.SetExternalLink(*s)
	}
	return ac
}

// SetPrimaryImageURL sets the "primary_image_url" field.
func (ac *ArtCreate) SetPrimaryImageURL(s string) *ArtCreate {
	ac.mutation.SetPrimaryImageURL(s)
	return ac
}

// SetNillablePrimaryImageURL sets the "primary_image_url" field if the given value is not nil.
func (ac *ArtCreate) SetNillablePrimaryImageURL(s *string) *ArtCreate {
	if s != nil {
		ac.SetPrimaryImageURL(*s)
	}
	return ac
}

// SetAdditionalImagesUrls sets the "additional_images_urls" field.
func (ac *ArtCreate) SetAdditionalImagesUrls(s []string) *ArtCreate {
	ac.mutation.SetAdditionalImagesUrls(s)
	return ac
}

// AddArtGenreIDs adds the "art_genre" edge to the ArtGenre entity by IDs.
func (ac *ArtCreate) AddArtGenreIDs(ids ...int) *ArtCreate {
	ac.mutation.AddArtGenreIDs(ids...)
	return ac
}

// AddArtGenre adds the "art_genre" edges to the ArtGenre entity.
func (ac *ArtCreate) AddArtGenre(a ...*ArtGenre) *ArtCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddArtGenreIDs(ids...)
}

// AddArtStyleIDs adds the "art_style" edge to the ArtStyle entity by IDs.
func (ac *ArtCreate) AddArtStyleIDs(ids ...int) *ArtCreate {
	ac.mutation.AddArtStyleIDs(ids...)
	return ac
}

// AddArtStyle adds the "art_style" edges to the ArtStyle entity.
func (ac *ArtCreate) AddArtStyle(a ...*ArtStyle) *ArtCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddArtStyleIDs(ids...)
}

// Mutation returns the ArtMutation object of the builder.
func (ac *ArtCreate) Mutation() *ArtMutation {
	return ac.mutation
}

// Save creates the Art in the database.
func (ac *ArtCreate) Save(ctx context.Context) (*Art, error) {
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtCreate) SaveX(ctx context.Context) *Art {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArtCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if art.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized art.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := art.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		if art.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized art.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := art.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Art.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Art.updated_at"`)}
	}
	return nil
}

func (ac *ArtCreate) sqlSave(ctx context.Context) (*Art, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArtCreate) createSpec() (*Art, *sqlgraph.CreateSpec) {
	var (
		_node = &Art{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(art.Table, sqlgraph.NewFieldSpec(art.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(art.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.SetField(art.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(art.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.UpdatedBy(); ok {
		_spec.SetField(art.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ac.mutation.DisplayName(); ok {
		_spec.SetField(art.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := ac.mutation.Abbreviation(); ok {
		_spec.SetField(art.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(art.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.ExternalLink(); ok {
		_spec.SetField(art.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := ac.mutation.PrimaryImageURL(); ok {
		_spec.SetField(art.FieldPrimaryImageURL, field.TypeString, value)
		_node.PrimaryImageURL = value
	}
	if value, ok := ac.mutation.AdditionalImagesUrls(); ok {
		_spec.SetField(art.FieldAdditionalImagesUrls, field.TypeJSON, value)
		_node.AdditionalImagesUrls = value
	}
	if nodes := ac.mutation.ArtGenreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtGenreTable,
			Columns: art.ArtGenrePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ArtStyleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   art.ArtStyleTable,
			Columns: art.ArtStylePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artstyle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArtCreateBulk is the builder for creating many Art entities in bulk.
type ArtCreateBulk struct {
	config
	builders []*ArtCreate
}

// Save creates the Art entities in the database.
func (acb *ArtCreateBulk) Save(ctx context.Context) ([]*Art, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Art, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArtCreateBulk) SaveX(ctx context.Context) []*Art {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
