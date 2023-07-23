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
	"github.com/dkrasnovdev/heritage-api/ent/country"
	"github.com/dkrasnovdev/heritage-api/ent/district"
	"github.com/dkrasnovdev/heritage-api/ent/location"
	"github.com/dkrasnovdev/heritage-api/ent/region"
	"github.com/dkrasnovdev/heritage-api/ent/settlement"
)

// LocationCreate is the builder for creating a Location entity.
type LocationCreate struct {
	config
	mutation *LocationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lc *LocationCreate) SetCreatedAt(t time.Time) *LocationCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LocationCreate) SetNillableCreatedAt(t *time.Time) *LocationCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetCreatedBy sets the "created_by" field.
func (lc *LocationCreate) SetCreatedBy(s string) *LocationCreate {
	lc.mutation.SetCreatedBy(s)
	return lc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (lc *LocationCreate) SetNillableCreatedBy(s *string) *LocationCreate {
	if s != nil {
		lc.SetCreatedBy(*s)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LocationCreate) SetUpdatedAt(t time.Time) *LocationCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LocationCreate) SetNillableUpdatedAt(t *time.Time) *LocationCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetUpdatedBy sets the "updated_by" field.
func (lc *LocationCreate) SetUpdatedBy(s string) *LocationCreate {
	lc.mutation.SetUpdatedBy(s)
	return lc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (lc *LocationCreate) SetNillableUpdatedBy(s *string) *LocationCreate {
	if s != nil {
		lc.SetUpdatedBy(*s)
	}
	return lc
}

// SetDisplayName sets the "display_name" field.
func (lc *LocationCreate) SetDisplayName(s string) *LocationCreate {
	lc.mutation.SetDisplayName(s)
	return lc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (lc *LocationCreate) SetNillableDisplayName(s *string) *LocationCreate {
	if s != nil {
		lc.SetDisplayName(*s)
	}
	return lc
}

// SetDescription sets the "description" field.
func (lc *LocationCreate) SetDescription(s string) *LocationCreate {
	lc.mutation.SetDescription(s)
	return lc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lc *LocationCreate) SetNillableDescription(s *string) *LocationCreate {
	if s != nil {
		lc.SetDescription(*s)
	}
	return lc
}

// SetExternalLinks sets the "external_links" field.
func (lc *LocationCreate) SetExternalLinks(s []string) *LocationCreate {
	lc.mutation.SetExternalLinks(s)
	return lc
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (lc *LocationCreate) AddArtifactIDs(ids ...int) *LocationCreate {
	lc.mutation.AddArtifactIDs(ids...)
	return lc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (lc *LocationCreate) AddArtifacts(a ...*Artifact) *LocationCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lc.AddArtifactIDs(ids...)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (lc *LocationCreate) SetCountryID(id int) *LocationCreate {
	lc.mutation.SetCountryID(id)
	return lc
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (lc *LocationCreate) SetNillableCountryID(id *int) *LocationCreate {
	if id != nil {
		lc = lc.SetCountryID(*id)
	}
	return lc
}

// SetCountry sets the "country" edge to the Country entity.
func (lc *LocationCreate) SetCountry(c *Country) *LocationCreate {
	return lc.SetCountryID(c.ID)
}

// SetDistrictID sets the "district" edge to the District entity by ID.
func (lc *LocationCreate) SetDistrictID(id int) *LocationCreate {
	lc.mutation.SetDistrictID(id)
	return lc
}

// SetNillableDistrictID sets the "district" edge to the District entity by ID if the given value is not nil.
func (lc *LocationCreate) SetNillableDistrictID(id *int) *LocationCreate {
	if id != nil {
		lc = lc.SetDistrictID(*id)
	}
	return lc
}

// SetDistrict sets the "district" edge to the District entity.
func (lc *LocationCreate) SetDistrict(d *District) *LocationCreate {
	return lc.SetDistrictID(d.ID)
}

// SetSettlementID sets the "settlement" edge to the Settlement entity by ID.
func (lc *LocationCreate) SetSettlementID(id int) *LocationCreate {
	lc.mutation.SetSettlementID(id)
	return lc
}

// SetNillableSettlementID sets the "settlement" edge to the Settlement entity by ID if the given value is not nil.
func (lc *LocationCreate) SetNillableSettlementID(id *int) *LocationCreate {
	if id != nil {
		lc = lc.SetSettlementID(*id)
	}
	return lc
}

// SetSettlement sets the "settlement" edge to the Settlement entity.
func (lc *LocationCreate) SetSettlement(s *Settlement) *LocationCreate {
	return lc.SetSettlementID(s.ID)
}

// SetRegionID sets the "region" edge to the Region entity by ID.
func (lc *LocationCreate) SetRegionID(id int) *LocationCreate {
	lc.mutation.SetRegionID(id)
	return lc
}

// SetNillableRegionID sets the "region" edge to the Region entity by ID if the given value is not nil.
func (lc *LocationCreate) SetNillableRegionID(id *int) *LocationCreate {
	if id != nil {
		lc = lc.SetRegionID(*id)
	}
	return lc
}

// SetRegion sets the "region" edge to the Region entity.
func (lc *LocationCreate) SetRegion(r *Region) *LocationCreate {
	return lc.SetRegionID(r.ID)
}

// Mutation returns the LocationMutation object of the builder.
func (lc *LocationCreate) Mutation() *LocationMutation {
	return lc.mutation
}

// Save creates the Location in the database.
func (lc *LocationCreate) Save(ctx context.Context) (*Location, error) {
	if err := lc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LocationCreate) SaveX(ctx context.Context) *Location {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LocationCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LocationCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LocationCreate) defaults() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		if location.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized location.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := location.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		if location.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized location.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := location.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lc *LocationCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Location.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Location.updated_at"`)}
	}
	return nil
}

func (lc *LocationCreate) sqlSave(ctx context.Context) (*Location, error) {
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

func (lc *LocationCreate) createSpec() (*Location, *sqlgraph.CreateSpec) {
	var (
		_node = &Location{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(location.Table, sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(location.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.CreatedBy(); ok {
		_spec.SetField(location.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(location.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.UpdatedBy(); ok {
		_spec.SetField(location.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := lc.mutation.DisplayName(); ok {
		_spec.SetField(location.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := lc.mutation.Description(); ok {
		_spec.SetField(location.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := lc.mutation.ExternalLinks(); ok {
		_spec.SetField(location.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
	if nodes := lc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ArtifactsTable,
			Columns: []string{location.ArtifactsColumn},
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
	if nodes := lc.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   location.CountryTable,
			Columns: []string{location.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.DistrictIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   location.DistrictTable,
			Columns: []string{location.DistrictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(district.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.SettlementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   location.SettlementTable,
			Columns: []string{location.SettlementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(settlement.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.RegionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   location.RegionTable,
			Columns: []string{location.RegionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LocationCreateBulk is the builder for creating many Location entities in bulk.
type LocationCreateBulk struct {
	config
	builders []*LocationCreate
}

// Save creates the Location entities in the database.
func (lcb *LocationCreateBulk) Save(ctx context.Context) ([]*Location, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Location, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LocationMutation)
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
func (lcb *LocationCreateBulk) SaveX(ctx context.Context) []*Location {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LocationCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LocationCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
