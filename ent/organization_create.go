// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/holder"
	"github.com/dkrasnovdev/heritage-api/ent/organization"
	"github.com/dkrasnovdev/heritage-api/ent/organizationtype"
	"github.com/dkrasnovdev/heritage-api/ent/person"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrganizationCreate) SetCreatedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreatedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetCreatedBy sets the "created_by" field.
func (oc *OrganizationCreate) SetCreatedBy(s string) *OrganizationCreate {
	oc.mutation.SetCreatedBy(s)
	return oc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreatedBy(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetCreatedBy(*s)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrganizationCreate) SetUpdatedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetUpdatedBy sets the "updated_by" field.
func (oc *OrganizationCreate) SetUpdatedBy(s string) *OrganizationCreate {
	oc.mutation.SetUpdatedBy(s)
	return oc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdatedBy(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetUpdatedBy(*s)
	}
	return oc
}

// SetAddress sets the "address" field.
func (oc *OrganizationCreate) SetAddress(s string) *OrganizationCreate {
	oc.mutation.SetAddress(s)
	return oc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableAddress(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetAddress(*s)
	}
	return oc
}

// SetPhoneNumbers sets the "phone_numbers" field.
func (oc *OrganizationCreate) SetPhoneNumbers(s []string) *OrganizationCreate {
	oc.mutation.SetPhoneNumbers(s)
	return oc
}

// SetEmails sets the "emails" field.
func (oc *OrganizationCreate) SetEmails(s []string) *OrganizationCreate {
	oc.mutation.SetEmails(s)
	return oc
}

// SetDisplayName sets the "display_name" field.
func (oc *OrganizationCreate) SetDisplayName(s string) *OrganizationCreate {
	oc.mutation.SetDisplayName(s)
	return oc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDisplayName(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetDisplayName(*s)
	}
	return oc
}

// SetAbbreviation sets the "abbreviation" field.
func (oc *OrganizationCreate) SetAbbreviation(s string) *OrganizationCreate {
	oc.mutation.SetAbbreviation(s)
	return oc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableAbbreviation(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetAbbreviation(*s)
	}
	return oc
}

// SetDescription sets the "description" field.
func (oc *OrganizationCreate) SetDescription(s string) *OrganizationCreate {
	oc.mutation.SetDescription(s)
	return oc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDescription(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetDescription(*s)
	}
	return oc
}

// SetExternalLinks sets the "external_links" field.
func (oc *OrganizationCreate) SetExternalLinks(s []string) *OrganizationCreate {
	oc.mutation.SetExternalLinks(s)
	return oc
}

// SetPrimaryImageURL sets the "primary_image_url" field.
func (oc *OrganizationCreate) SetPrimaryImageURL(s string) *OrganizationCreate {
	oc.mutation.SetPrimaryImageURL(s)
	return oc
}

// SetNillablePrimaryImageURL sets the "primary_image_url" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillablePrimaryImageURL(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetPrimaryImageURL(*s)
	}
	return oc
}

// SetAdditionalImagesUrls sets the "additional_images_urls" field.
func (oc *OrganizationCreate) SetAdditionalImagesUrls(s []string) *OrganizationCreate {
	oc.mutation.SetAdditionalImagesUrls(s)
	return oc
}

// SetPreviousNames sets the "previous_names" field.
func (oc *OrganizationCreate) SetPreviousNames(s []string) *OrganizationCreate {
	oc.mutation.SetPreviousNames(s)
	return oc
}

// SetIsInAConsortium sets the "is_in_a_consortium" field.
func (oc *OrganizationCreate) SetIsInAConsortium(b bool) *OrganizationCreate {
	oc.mutation.SetIsInAConsortium(b)
	return oc
}

// SetNillableIsInAConsortium sets the "is_in_a_consortium" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableIsInAConsortium(b *bool) *OrganizationCreate {
	if b != nil {
		oc.SetIsInAConsortium(*b)
	}
	return oc
}

// SetConsortiumDocumentURL sets the "consortium_document_url" field.
func (oc *OrganizationCreate) SetConsortiumDocumentURL(s string) *OrganizationCreate {
	oc.mutation.SetConsortiumDocumentURL(s)
	return oc
}

// SetNillableConsortiumDocumentURL sets the "consortium_document_url" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableConsortiumDocumentURL(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetConsortiumDocumentURL(*s)
	}
	return oc
}

// AddPersonIDs adds the "people" edge to the Person entity by IDs.
func (oc *OrganizationCreate) AddPersonIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddPersonIDs(ids...)
	return oc
}

// AddPeople adds the "people" edges to the Person entity.
func (oc *OrganizationCreate) AddPeople(p ...*Person) *OrganizationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oc.AddPersonIDs(ids...)
}

// SetHolderID sets the "holder" edge to the Holder entity by ID.
func (oc *OrganizationCreate) SetHolderID(id int) *OrganizationCreate {
	oc.mutation.SetHolderID(id)
	return oc
}

// SetNillableHolderID sets the "holder" edge to the Holder entity by ID if the given value is not nil.
func (oc *OrganizationCreate) SetNillableHolderID(id *int) *OrganizationCreate {
	if id != nil {
		oc = oc.SetHolderID(*id)
	}
	return oc
}

// SetHolder sets the "holder" edge to the Holder entity.
func (oc *OrganizationCreate) SetHolder(h *Holder) *OrganizationCreate {
	return oc.SetHolderID(h.ID)
}

// SetOrganizationTypeID sets the "organization_type" edge to the OrganizationType entity by ID.
func (oc *OrganizationCreate) SetOrganizationTypeID(id int) *OrganizationCreate {
	oc.mutation.SetOrganizationTypeID(id)
	return oc
}

// SetNillableOrganizationTypeID sets the "organization_type" edge to the OrganizationType entity by ID if the given value is not nil.
func (oc *OrganizationCreate) SetNillableOrganizationTypeID(id *int) *OrganizationCreate {
	if id != nil {
		oc = oc.SetOrganizationTypeID(*id)
	}
	return oc
}

// SetOrganizationType sets the "organization_type" edge to the OrganizationType entity.
func (oc *OrganizationCreate) SetOrganizationType(o *OrganizationType) *OrganizationCreate {
	return oc.SetOrganizationTypeID(o.ID)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	if err := oc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrganizationCreate) defaults() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		if organization.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized organization.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := organization.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		if organization.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized organization.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := organization.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Organization.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Organization.updated_at"`)}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(organization.Table, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(organization.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.CreatedBy(); ok {
		_spec.SetField(organization.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.UpdatedBy(); ok {
		_spec.SetField(organization.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := oc.mutation.Address(); ok {
		_spec.SetField(organization.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := oc.mutation.PhoneNumbers(); ok {
		_spec.SetField(organization.FieldPhoneNumbers, field.TypeJSON, value)
		_node.PhoneNumbers = value
	}
	if value, ok := oc.mutation.Emails(); ok {
		_spec.SetField(organization.FieldEmails, field.TypeJSON, value)
		_node.Emails = value
	}
	if value, ok := oc.mutation.DisplayName(); ok {
		_spec.SetField(organization.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := oc.mutation.Abbreviation(); ok {
		_spec.SetField(organization.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := oc.mutation.Description(); ok {
		_spec.SetField(organization.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := oc.mutation.ExternalLinks(); ok {
		_spec.SetField(organization.FieldExternalLinks, field.TypeJSON, value)
		_node.ExternalLinks = value
	}
	if value, ok := oc.mutation.PrimaryImageURL(); ok {
		_spec.SetField(organization.FieldPrimaryImageURL, field.TypeString, value)
		_node.PrimaryImageURL = value
	}
	if value, ok := oc.mutation.AdditionalImagesUrls(); ok {
		_spec.SetField(organization.FieldAdditionalImagesUrls, field.TypeJSON, value)
		_node.AdditionalImagesUrls = value
	}
	if value, ok := oc.mutation.PreviousNames(); ok {
		_spec.SetField(organization.FieldPreviousNames, field.TypeJSON, value)
		_node.PreviousNames = value
	}
	if value, ok := oc.mutation.IsInAConsortium(); ok {
		_spec.SetField(organization.FieldIsInAConsortium, field.TypeBool, value)
		_node.IsInAConsortium = value
	}
	if value, ok := oc.mutation.ConsortiumDocumentURL(); ok {
		_spec.SetField(organization.FieldConsortiumDocumentURL, field.TypeString, value)
		_node.ConsortiumDocumentURL = value
	}
	if nodes := oc.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.PeopleTable,
			Columns: []string{organization.PeopleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.HolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   organization.HolderTable,
			Columns: []string{organization.HolderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.holder_organization = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrganizationTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.OrganizationTypeTable,
			Columns: []string{organization.OrganizationTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_type_organizations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	builders []*OrganizationCreate
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
