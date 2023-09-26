// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/artifact"
	"github.com/dkrasnovdev/siberiana-api/ent/book"
	"github.com/dkrasnovdev/siberiana-api/ent/collection"
	"github.com/dkrasnovdev/siberiana-api/ent/holder"
	"github.com/dkrasnovdev/siberiana-api/ent/organization"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/personrole"
	"github.com/dkrasnovdev/siberiana-api/ent/project"
	"github.com/dkrasnovdev/siberiana-api/ent/publication"
)

// PersonCreate is the builder for creating a Person entity.
type PersonCreate struct {
	config
	mutation *PersonMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PersonCreate) SetCreatedAt(t time.Time) *PersonCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PersonCreate) SetNillableCreatedAt(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *PersonCreate) SetCreatedBy(s string) *PersonCreate {
	pc.mutation.SetCreatedBy(s)
	return pc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pc *PersonCreate) SetNillableCreatedBy(s *string) *PersonCreate {
	if s != nil {
		pc.SetCreatedBy(*s)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PersonCreate) SetUpdatedAt(t time.Time) *PersonCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUpdatedAt(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *PersonCreate) SetUpdatedBy(s string) *PersonCreate {
	pc.mutation.SetUpdatedBy(s)
	return pc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUpdatedBy(s *string) *PersonCreate {
	if s != nil {
		pc.SetUpdatedBy(*s)
	}
	return pc
}

// SetAddress sets the "address" field.
func (pc *PersonCreate) SetAddress(s string) *PersonCreate {
	pc.mutation.SetAddress(s)
	return pc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (pc *PersonCreate) SetNillableAddress(s *string) *PersonCreate {
	if s != nil {
		pc.SetAddress(*s)
	}
	return pc
}

// SetPhoneNumbers sets the "phone_numbers" field.
func (pc *PersonCreate) SetPhoneNumbers(s []string) *PersonCreate {
	pc.mutation.SetPhoneNumbers(s)
	return pc
}

// SetEmails sets the "emails" field.
func (pc *PersonCreate) SetEmails(s []string) *PersonCreate {
	pc.mutation.SetEmails(s)
	return pc
}

// SetDisplayName sets the "display_name" field.
func (pc *PersonCreate) SetDisplayName(s string) *PersonCreate {
	pc.mutation.SetDisplayName(s)
	return pc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDisplayName(s *string) *PersonCreate {
	if s != nil {
		pc.SetDisplayName(*s)
	}
	return pc
}

// SetAbbreviation sets the "abbreviation" field.
func (pc *PersonCreate) SetAbbreviation(s string) *PersonCreate {
	pc.mutation.SetAbbreviation(s)
	return pc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (pc *PersonCreate) SetNillableAbbreviation(s *string) *PersonCreate {
	if s != nil {
		pc.SetAbbreviation(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *PersonCreate) SetDescription(s string) *PersonCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDescription(s *string) *PersonCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetExternalLink sets the "external_link" field.
func (pc *PersonCreate) SetExternalLink(s string) *PersonCreate {
	pc.mutation.SetExternalLink(s)
	return pc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (pc *PersonCreate) SetNillableExternalLink(s *string) *PersonCreate {
	if s != nil {
		pc.SetExternalLink(*s)
	}
	return pc
}

// SetPrimaryImageURL sets the "primary_image_url" field.
func (pc *PersonCreate) SetPrimaryImageURL(s string) *PersonCreate {
	pc.mutation.SetPrimaryImageURL(s)
	return pc
}

// SetNillablePrimaryImageURL sets the "primary_image_url" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePrimaryImageURL(s *string) *PersonCreate {
	if s != nil {
		pc.SetPrimaryImageURL(*s)
	}
	return pc
}

// SetAdditionalImagesUrls sets the "additional_images_urls" field.
func (pc *PersonCreate) SetAdditionalImagesUrls(s []string) *PersonCreate {
	pc.mutation.SetAdditionalImagesUrls(s)
	return pc
}

// SetGivenName sets the "given_name" field.
func (pc *PersonCreate) SetGivenName(s string) *PersonCreate {
	pc.mutation.SetGivenName(s)
	return pc
}

// SetNillableGivenName sets the "given_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableGivenName(s *string) *PersonCreate {
	if s != nil {
		pc.SetGivenName(*s)
	}
	return pc
}

// SetFamilyName sets the "family_name" field.
func (pc *PersonCreate) SetFamilyName(s string) *PersonCreate {
	pc.mutation.SetFamilyName(s)
	return pc
}

// SetNillableFamilyName sets the "family_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableFamilyName(s *string) *PersonCreate {
	if s != nil {
		pc.SetFamilyName(*s)
	}
	return pc
}

// SetPatronymicName sets the "patronymic_name" field.
func (pc *PersonCreate) SetPatronymicName(s string) *PersonCreate {
	pc.mutation.SetPatronymicName(s)
	return pc
}

// SetNillablePatronymicName sets the "patronymic_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePatronymicName(s *string) *PersonCreate {
	if s != nil {
		pc.SetPatronymicName(*s)
	}
	return pc
}

// SetBeginData sets the "begin_data" field.
func (pc *PersonCreate) SetBeginData(t time.Time) *PersonCreate {
	pc.mutation.SetBeginData(t)
	return pc
}

// SetNillableBeginData sets the "begin_data" field if the given value is not nil.
func (pc *PersonCreate) SetNillableBeginData(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetBeginData(*t)
	}
	return pc
}

// SetEndDate sets the "end_date" field.
func (pc *PersonCreate) SetEndDate(t time.Time) *PersonCreate {
	pc.mutation.SetEndDate(t)
	return pc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (pc *PersonCreate) SetNillableEndDate(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetEndDate(*t)
	}
	return pc
}

// SetGender sets the "gender" field.
func (pc *PersonCreate) SetGender(pe person.Gender) *PersonCreate {
	pc.mutation.SetGender(pe)
	return pc
}

// AddCollectionIDs adds the "collections" edge to the Collection entity by IDs.
func (pc *PersonCreate) AddCollectionIDs(ids ...int) *PersonCreate {
	pc.mutation.AddCollectionIDs(ids...)
	return pc
}

// AddCollections adds the "collections" edges to the Collection entity.
func (pc *PersonCreate) AddCollections(c ...*Collection) *PersonCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCollectionIDs(ids...)
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (pc *PersonCreate) AddArtifactIDs(ids ...int) *PersonCreate {
	pc.mutation.AddArtifactIDs(ids...)
	return pc
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (pc *PersonCreate) AddArtifacts(a ...*Artifact) *PersonCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddArtifactIDs(ids...)
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (pc *PersonCreate) AddBookIDs(ids ...int) *PersonCreate {
	pc.mutation.AddBookIDs(ids...)
	return pc
}

// AddBooks adds the "books" edges to the Book entity.
func (pc *PersonCreate) AddBooks(b ...*Book) *PersonCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBookIDs(ids...)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (pc *PersonCreate) AddProjectIDs(ids ...int) *PersonCreate {
	pc.mutation.AddProjectIDs(ids...)
	return pc
}

// AddProjects adds the "projects" edges to the Project entity.
func (pc *PersonCreate) AddProjects(p ...*Project) *PersonCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProjectIDs(ids...)
}

// AddPublicationIDs adds the "publications" edge to the Publication entity by IDs.
func (pc *PersonCreate) AddPublicationIDs(ids ...int) *PersonCreate {
	pc.mutation.AddPublicationIDs(ids...)
	return pc
}

// AddPublications adds the "publications" edges to the Publication entity.
func (pc *PersonCreate) AddPublications(p ...*Publication) *PersonCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPublicationIDs(ids...)
}

// AddPersonRoleIDs adds the "person_roles" edge to the PersonRole entity by IDs.
func (pc *PersonCreate) AddPersonRoleIDs(ids ...int) *PersonCreate {
	pc.mutation.AddPersonRoleIDs(ids...)
	return pc
}

// AddPersonRoles adds the "person_roles" edges to the PersonRole entity.
func (pc *PersonCreate) AddPersonRoles(p ...*PersonRole) *PersonCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPersonRoleIDs(ids...)
}

// SetHolderID sets the "holder" edge to the Holder entity by ID.
func (pc *PersonCreate) SetHolderID(id int) *PersonCreate {
	pc.mutation.SetHolderID(id)
	return pc
}

// SetNillableHolderID sets the "holder" edge to the Holder entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableHolderID(id *int) *PersonCreate {
	if id != nil {
		pc = pc.SetHolderID(*id)
	}
	return pc
}

// SetHolder sets the "holder" edge to the Holder entity.
func (pc *PersonCreate) SetHolder(h *Holder) *PersonCreate {
	return pc.SetHolderID(h.ID)
}

// SetAffiliationID sets the "affiliation" edge to the Organization entity by ID.
func (pc *PersonCreate) SetAffiliationID(id int) *PersonCreate {
	pc.mutation.SetAffiliationID(id)
	return pc
}

// SetNillableAffiliationID sets the "affiliation" edge to the Organization entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableAffiliationID(id *int) *PersonCreate {
	if id != nil {
		pc = pc.SetAffiliationID(*id)
	}
	return pc
}

// SetAffiliation sets the "affiliation" edge to the Organization entity.
func (pc *PersonCreate) SetAffiliation(o *Organization) *PersonCreate {
	return pc.SetAffiliationID(o.ID)
}

// Mutation returns the PersonMutation object of the builder.
func (pc *PersonCreate) Mutation() *PersonMutation {
	return pc.mutation
}

// Save creates the Person in the database.
func (pc *PersonCreate) Save(ctx context.Context) (*Person, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonCreate) SaveX(ctx context.Context) *Person {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PersonCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PersonCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PersonCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if person.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized person.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := person.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if person.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized person.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := person.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Person.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Person.updated_at"`)}
	}
	if _, ok := pc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Person.gender"`)}
	}
	if v, ok := pc.mutation.Gender(); ok {
		if err := person.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Person.gender": %w`, err)}
		}
	}
	return nil
}

func (pc *PersonCreate) sqlSave(ctx context.Context) (*Person, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PersonCreate) createSpec() (*Person, *sqlgraph.CreateSpec) {
	var (
		_node = &Person{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(person.Table, sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(person.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(person.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(person.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(person.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.Address(); ok {
		_spec.SetField(person.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := pc.mutation.PhoneNumbers(); ok {
		_spec.SetField(person.FieldPhoneNumbers, field.TypeJSON, value)
		_node.PhoneNumbers = value
	}
	if value, ok := pc.mutation.Emails(); ok {
		_spec.SetField(person.FieldEmails, field.TypeJSON, value)
		_node.Emails = value
	}
	if value, ok := pc.mutation.DisplayName(); ok {
		_spec.SetField(person.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := pc.mutation.Abbreviation(); ok {
		_spec.SetField(person.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(person.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.ExternalLink(); ok {
		_spec.SetField(person.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := pc.mutation.PrimaryImageURL(); ok {
		_spec.SetField(person.FieldPrimaryImageURL, field.TypeString, value)
		_node.PrimaryImageURL = value
	}
	if value, ok := pc.mutation.AdditionalImagesUrls(); ok {
		_spec.SetField(person.FieldAdditionalImagesUrls, field.TypeJSON, value)
		_node.AdditionalImagesUrls = value
	}
	if value, ok := pc.mutation.GivenName(); ok {
		_spec.SetField(person.FieldGivenName, field.TypeString, value)
		_node.GivenName = value
	}
	if value, ok := pc.mutation.FamilyName(); ok {
		_spec.SetField(person.FieldFamilyName, field.TypeString, value)
		_node.FamilyName = value
	}
	if value, ok := pc.mutation.PatronymicName(); ok {
		_spec.SetField(person.FieldPatronymicName, field.TypeString, value)
		_node.PatronymicName = value
	}
	if value, ok := pc.mutation.BeginData(); ok {
		_spec.SetField(person.FieldBeginData, field.TypeTime, value)
		_node.BeginData = value
	}
	if value, ok := pc.mutation.EndDate(); ok {
		_spec.SetField(person.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.SetField(person.FieldGender, field.TypeEnum, value)
		_node.Gender = value
	}
	if nodes := pc.mutation.CollectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.CollectionsTable,
			Columns: person.CollectionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.ArtifactsTable,
			Columns: person.ArtifactsPrimaryKey,
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
	if nodes := pc.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.BooksTable,
			Columns: person.BooksPrimaryKey,
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
	if nodes := pc.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.ProjectsTable,
			Columns: person.ProjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PublicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.PublicationsTable,
			Columns: person.PublicationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(publication.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PersonRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.PersonRolesTable,
			Columns: person.PersonRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.HolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   person.HolderTable,
			Columns: []string{person.HolderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.holder_person = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.AffiliationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.AffiliationTable,
			Columns: []string{person.AffiliationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_people = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PersonCreateBulk is the builder for creating many Person entities in bulk.
type PersonCreateBulk struct {
	config
	err      error
	builders []*PersonCreate
}

// Save creates the Person entities in the database.
func (pcb *PersonCreateBulk) Save(ctx context.Context) ([]*Person, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Person, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PersonCreateBulk) SaveX(ctx context.Context) []*Person {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PersonCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PersonCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
