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
	"github.com/dkrasnovdev/siberiana-api/ent/collection"
	"github.com/dkrasnovdev/siberiana-api/ent/culture"
	"github.com/dkrasnovdev/siberiana-api/ent/holder"
	"github.com/dkrasnovdev/siberiana-api/ent/license"
	"github.com/dkrasnovdev/siberiana-api/ent/location"
	"github.com/dkrasnovdev/siberiana-api/ent/medium"
	"github.com/dkrasnovdev/siberiana-api/ent/model"
	"github.com/dkrasnovdev/siberiana-api/ent/monument"
	"github.com/dkrasnovdev/siberiana-api/ent/period"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/project"
	"github.com/dkrasnovdev/siberiana-api/ent/publication"
	"github.com/dkrasnovdev/siberiana-api/ent/set"
	"github.com/dkrasnovdev/siberiana-api/ent/technique"
)

// ArtifactCreate is the builder for creating a Artifact entity.
type ArtifactCreate struct {
	config
	mutation *ArtifactMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *ArtifactCreate) SetCreatedAt(t time.Time) *ArtifactCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableCreatedAt(t *time.Time) *ArtifactCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetCreatedBy sets the "created_by" field.
func (ac *ArtifactCreate) SetCreatedBy(s string) *ArtifactCreate {
	ac.mutation.SetCreatedBy(s)
	return ac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableCreatedBy(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetCreatedBy(*s)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ArtifactCreate) SetUpdatedAt(t time.Time) *ArtifactCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableUpdatedAt(t *time.Time) *ArtifactCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetUpdatedBy sets the "updated_by" field.
func (ac *ArtifactCreate) SetUpdatedBy(s string) *ArtifactCreate {
	ac.mutation.SetUpdatedBy(s)
	return ac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableUpdatedBy(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetUpdatedBy(*s)
	}
	return ac
}

// SetDisplayName sets the "display_name" field.
func (ac *ArtifactCreate) SetDisplayName(s string) *ArtifactCreate {
	ac.mutation.SetDisplayName(s)
	return ac
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableDisplayName(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetDisplayName(*s)
	}
	return ac
}

// SetAbbreviation sets the "abbreviation" field.
func (ac *ArtifactCreate) SetAbbreviation(s string) *ArtifactCreate {
	ac.mutation.SetAbbreviation(s)
	return ac
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableAbbreviation(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetAbbreviation(*s)
	}
	return ac
}

// SetDescription sets the "description" field.
func (ac *ArtifactCreate) SetDescription(s string) *ArtifactCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableDescription(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetExternalLink sets the "external_link" field.
func (ac *ArtifactCreate) SetExternalLink(s string) *ArtifactCreate {
	ac.mutation.SetExternalLink(s)
	return ac
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableExternalLink(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetExternalLink(*s)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *ArtifactCreate) SetStatus(a artifact.Status) *ArtifactCreate {
	ac.mutation.SetStatus(a)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableStatus(a *artifact.Status) *ArtifactCreate {
	if a != nil {
		ac.SetStatus(*a)
	}
	return ac
}

// SetPrimaryImageURL sets the "primary_image_url" field.
func (ac *ArtifactCreate) SetPrimaryImageURL(s string) *ArtifactCreate {
	ac.mutation.SetPrimaryImageURL(s)
	return ac
}

// SetNillablePrimaryImageURL sets the "primary_image_url" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillablePrimaryImageURL(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetPrimaryImageURL(*s)
	}
	return ac
}

// SetAdditionalImagesUrls sets the "additional_images_urls" field.
func (ac *ArtifactCreate) SetAdditionalImagesUrls(s []string) *ArtifactCreate {
	ac.mutation.SetAdditionalImagesUrls(s)
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *ArtifactCreate) SetDeletedAt(t time.Time) *ArtifactCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableDeletedAt(t *time.Time) *ArtifactCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetDeletedBy sets the "deleted_by" field.
func (ac *ArtifactCreate) SetDeletedBy(s string) *ArtifactCreate {
	ac.mutation.SetDeletedBy(s)
	return ac
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableDeletedBy(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetDeletedBy(*s)
	}
	return ac
}

// SetDating sets the "dating" field.
func (ac *ArtifactCreate) SetDating(s string) *ArtifactCreate {
	ac.mutation.SetDating(s)
	return ac
}

// SetNillableDating sets the "dating" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableDating(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetDating(*s)
	}
	return ac
}

// SetDimensions sets the "dimensions" field.
func (ac *ArtifactCreate) SetDimensions(s string) *ArtifactCreate {
	ac.mutation.SetDimensions(s)
	return ac
}

// SetNillableDimensions sets the "dimensions" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableDimensions(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetDimensions(*s)
	}
	return ac
}

// SetChemicalComposition sets the "chemical_composition" field.
func (ac *ArtifactCreate) SetChemicalComposition(s string) *ArtifactCreate {
	ac.mutation.SetChemicalComposition(s)
	return ac
}

// SetNillableChemicalComposition sets the "chemical_composition" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableChemicalComposition(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetChemicalComposition(*s)
	}
	return ac
}

// SetNumber sets the "number" field.
func (ac *ArtifactCreate) SetNumber(s string) *ArtifactCreate {
	ac.mutation.SetNumber(s)
	return ac
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableNumber(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetNumber(*s)
	}
	return ac
}

// SetTypology sets the "typology" field.
func (ac *ArtifactCreate) SetTypology(s string) *ArtifactCreate {
	ac.mutation.SetTypology(s)
	return ac
}

// SetNillableTypology sets the "typology" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableTypology(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetTypology(*s)
	}
	return ac
}

// SetWeight sets the "weight" field.
func (ac *ArtifactCreate) SetWeight(s string) *ArtifactCreate {
	ac.mutation.SetWeight(s)
	return ac
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableWeight(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetWeight(*s)
	}
	return ac
}

// SetAdmissionDate sets the "admission_date" field.
func (ac *ArtifactCreate) SetAdmissionDate(t time.Time) *ArtifactCreate {
	ac.mutation.SetAdmissionDate(t)
	return ac
}

// SetNillableAdmissionDate sets the "admission_date" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableAdmissionDate(t *time.Time) *ArtifactCreate {
	if t != nil {
		ac.SetAdmissionDate(*t)
	}
	return ac
}

// AddAuthorIDs adds the "authors" edge to the Person entity by IDs.
func (ac *ArtifactCreate) AddAuthorIDs(ids ...int) *ArtifactCreate {
	ac.mutation.AddAuthorIDs(ids...)
	return ac
}

// AddAuthors adds the "authors" edges to the Person entity.
func (ac *ArtifactCreate) AddAuthors(p ...*Person) *ArtifactCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddAuthorIDs(ids...)
}

// AddMediumIDs adds the "mediums" edge to the Medium entity by IDs.
func (ac *ArtifactCreate) AddMediumIDs(ids ...int) *ArtifactCreate {
	ac.mutation.AddMediumIDs(ids...)
	return ac
}

// AddMediums adds the "mediums" edges to the Medium entity.
func (ac *ArtifactCreate) AddMediums(m ...*Medium) *ArtifactCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ac.AddMediumIDs(ids...)
}

// AddTechniqueIDs adds the "techniques" edge to the Technique entity by IDs.
func (ac *ArtifactCreate) AddTechniqueIDs(ids ...int) *ArtifactCreate {
	ac.mutation.AddTechniqueIDs(ids...)
	return ac
}

// AddTechniques adds the "techniques" edges to the Technique entity.
func (ac *ArtifactCreate) AddTechniques(t ...*Technique) *ArtifactCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTechniqueIDs(ids...)
}

// SetPeriodID sets the "period" edge to the Period entity by ID.
func (ac *ArtifactCreate) SetPeriodID(id int) *ArtifactCreate {
	ac.mutation.SetPeriodID(id)
	return ac
}

// SetNillablePeriodID sets the "period" edge to the Period entity by ID if the given value is not nil.
func (ac *ArtifactCreate) SetNillablePeriodID(id *int) *ArtifactCreate {
	if id != nil {
		ac = ac.SetPeriodID(*id)
	}
	return ac
}

// SetPeriod sets the "period" edge to the Period entity.
func (ac *ArtifactCreate) SetPeriod(p *Period) *ArtifactCreate {
	return ac.SetPeriodID(p.ID)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (ac *ArtifactCreate) AddProjectIDs(ids ...int) *ArtifactCreate {
	ac.mutation.AddProjectIDs(ids...)
	return ac
}

// AddProjects adds the "projects" edges to the Project entity.
func (ac *ArtifactCreate) AddProjects(p ...*Project) *ArtifactCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddProjectIDs(ids...)
}

// AddPublicationIDs adds the "publications" edge to the Publication entity by IDs.
func (ac *ArtifactCreate) AddPublicationIDs(ids ...int) *ArtifactCreate {
	ac.mutation.AddPublicationIDs(ids...)
	return ac
}

// AddPublications adds the "publications" edges to the Publication entity.
func (ac *ArtifactCreate) AddPublications(p ...*Publication) *ArtifactCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPublicationIDs(ids...)
}

// AddHolderIDs adds the "holders" edge to the Holder entity by IDs.
func (ac *ArtifactCreate) AddHolderIDs(ids ...int) *ArtifactCreate {
	ac.mutation.AddHolderIDs(ids...)
	return ac
}

// AddHolders adds the "holders" edges to the Holder entity.
func (ac *ArtifactCreate) AddHolders(h ...*Holder) *ArtifactCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return ac.AddHolderIDs(ids...)
}

// SetCulturalAffiliationID sets the "cultural_affiliation" edge to the Culture entity by ID.
func (ac *ArtifactCreate) SetCulturalAffiliationID(id int) *ArtifactCreate {
	ac.mutation.SetCulturalAffiliationID(id)
	return ac
}

// SetNillableCulturalAffiliationID sets the "cultural_affiliation" edge to the Culture entity by ID if the given value is not nil.
func (ac *ArtifactCreate) SetNillableCulturalAffiliationID(id *int) *ArtifactCreate {
	if id != nil {
		ac = ac.SetCulturalAffiliationID(*id)
	}
	return ac
}

// SetCulturalAffiliation sets the "cultural_affiliation" edge to the Culture entity.
func (ac *ArtifactCreate) SetCulturalAffiliation(c *Culture) *ArtifactCreate {
	return ac.SetCulturalAffiliationID(c.ID)
}

// SetMonumentID sets the "monument" edge to the Monument entity by ID.
func (ac *ArtifactCreate) SetMonumentID(id int) *ArtifactCreate {
	ac.mutation.SetMonumentID(id)
	return ac
}

// SetNillableMonumentID sets the "monument" edge to the Monument entity by ID if the given value is not nil.
func (ac *ArtifactCreate) SetNillableMonumentID(id *int) *ArtifactCreate {
	if id != nil {
		ac = ac.SetMonumentID(*id)
	}
	return ac
}

// SetMonument sets the "monument" edge to the Monument entity.
func (ac *ArtifactCreate) SetMonument(m *Monument) *ArtifactCreate {
	return ac.SetMonumentID(m.ID)
}

// SetModelID sets the "model" edge to the Model entity by ID.
func (ac *ArtifactCreate) SetModelID(id int) *ArtifactCreate {
	ac.mutation.SetModelID(id)
	return ac
}

// SetNillableModelID sets the "model" edge to the Model entity by ID if the given value is not nil.
func (ac *ArtifactCreate) SetNillableModelID(id *int) *ArtifactCreate {
	if id != nil {
		ac = ac.SetModelID(*id)
	}
	return ac
}

// SetModel sets the "model" edge to the Model entity.
func (ac *ArtifactCreate) SetModel(m *Model) *ArtifactCreate {
	return ac.SetModelID(m.ID)
}

// SetSetID sets the "set" edge to the Set entity by ID.
func (ac *ArtifactCreate) SetSetID(id int) *ArtifactCreate {
	ac.mutation.SetSetID(id)
	return ac
}

// SetNillableSetID sets the "set" edge to the Set entity by ID if the given value is not nil.
func (ac *ArtifactCreate) SetNillableSetID(id *int) *ArtifactCreate {
	if id != nil {
		ac = ac.SetSetID(*id)
	}
	return ac
}

// SetSet sets the "set" edge to the Set entity.
func (ac *ArtifactCreate) SetSet(s *Set) *ArtifactCreate {
	return ac.SetSetID(s.ID)
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (ac *ArtifactCreate) SetLocationID(id int) *ArtifactCreate {
	ac.mutation.SetLocationID(id)
	return ac
}

// SetNillableLocationID sets the "location" edge to the Location entity by ID if the given value is not nil.
func (ac *ArtifactCreate) SetNillableLocationID(id *int) *ArtifactCreate {
	if id != nil {
		ac = ac.SetLocationID(*id)
	}
	return ac
}

// SetLocation sets the "location" edge to the Location entity.
func (ac *ArtifactCreate) SetLocation(l *Location) *ArtifactCreate {
	return ac.SetLocationID(l.ID)
}

// SetCollectionID sets the "collection" edge to the Collection entity by ID.
func (ac *ArtifactCreate) SetCollectionID(id int) *ArtifactCreate {
	ac.mutation.SetCollectionID(id)
	return ac
}

// SetCollection sets the "collection" edge to the Collection entity.
func (ac *ArtifactCreate) SetCollection(c *Collection) *ArtifactCreate {
	return ac.SetCollectionID(c.ID)
}

// SetLicenseID sets the "license" edge to the License entity by ID.
func (ac *ArtifactCreate) SetLicenseID(id int) *ArtifactCreate {
	ac.mutation.SetLicenseID(id)
	return ac
}

// SetNillableLicenseID sets the "license" edge to the License entity by ID if the given value is not nil.
func (ac *ArtifactCreate) SetNillableLicenseID(id *int) *ArtifactCreate {
	if id != nil {
		ac = ac.SetLicenseID(*id)
	}
	return ac
}

// SetLicense sets the "license" edge to the License entity.
func (ac *ArtifactCreate) SetLicense(l *License) *ArtifactCreate {
	return ac.SetLicenseID(l.ID)
}

// Mutation returns the ArtifactMutation object of the builder.
func (ac *ArtifactCreate) Mutation() *ArtifactMutation {
	return ac.mutation
}

// Save creates the Artifact in the database.
func (ac *ArtifactCreate) Save(ctx context.Context) (*Artifact, error) {
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtifactCreate) SaveX(ctx context.Context) *Artifact {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtifactCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtifactCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArtifactCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if artifact.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized artifact.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := artifact.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		if artifact.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized artifact.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := artifact.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.Status(); !ok {
		v := artifact.DefaultStatus
		ac.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtifactCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Artifact.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Artifact.updated_at"`)}
	}
	if v, ok := ac.mutation.Status(); ok {
		if err := artifact.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Artifact.status": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CollectionID(); !ok {
		return &ValidationError{Name: "collection", err: errors.New(`ent: missing required edge "Artifact.collection"`)}
	}
	return nil
}

func (ac *ArtifactCreate) sqlSave(ctx context.Context) (*Artifact, error) {
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

func (ac *ArtifactCreate) createSpec() (*Artifact, *sqlgraph.CreateSpec) {
	var (
		_node = &Artifact{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(artifact.Table, sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(artifact.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.SetField(artifact.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(artifact.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.UpdatedBy(); ok {
		_spec.SetField(artifact.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ac.mutation.DisplayName(); ok {
		_spec.SetField(artifact.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := ac.mutation.Abbreviation(); ok {
		_spec.SetField(artifact.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(artifact.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.ExternalLink(); ok {
		_spec.SetField(artifact.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(artifact.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := ac.mutation.PrimaryImageURL(); ok {
		_spec.SetField(artifact.FieldPrimaryImageURL, field.TypeString, value)
		_node.PrimaryImageURL = value
	}
	if value, ok := ac.mutation.AdditionalImagesUrls(); ok {
		_spec.SetField(artifact.FieldAdditionalImagesUrls, field.TypeJSON, value)
		_node.AdditionalImagesUrls = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(artifact.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.DeletedBy(); ok {
		_spec.SetField(artifact.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ac.mutation.Dating(); ok {
		_spec.SetField(artifact.FieldDating, field.TypeString, value)
		_node.Dating = value
	}
	if value, ok := ac.mutation.Dimensions(); ok {
		_spec.SetField(artifact.FieldDimensions, field.TypeString, value)
		_node.Dimensions = value
	}
	if value, ok := ac.mutation.ChemicalComposition(); ok {
		_spec.SetField(artifact.FieldChemicalComposition, field.TypeString, value)
		_node.ChemicalComposition = value
	}
	if value, ok := ac.mutation.Number(); ok {
		_spec.SetField(artifact.FieldNumber, field.TypeString, value)
		_node.Number = value
	}
	if value, ok := ac.mutation.Typology(); ok {
		_spec.SetField(artifact.FieldTypology, field.TypeString, value)
		_node.Typology = value
	}
	if value, ok := ac.mutation.Weight(); ok {
		_spec.SetField(artifact.FieldWeight, field.TypeString, value)
		_node.Weight = value
	}
	if value, ok := ac.mutation.AdmissionDate(); ok {
		_spec.SetField(artifact.FieldAdmissionDate, field.TypeTime, value)
		_node.AdmissionDate = value
	}
	if nodes := ac.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artifact.AuthorsTable,
			Columns: artifact.AuthorsPrimaryKey,
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
	if nodes := ac.mutation.MediumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artifact.MediumsTable,
			Columns: artifact.MediumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medium.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TechniquesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artifact.TechniquesTable,
			Columns: artifact.TechniquesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(technique.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.PeriodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.PeriodTable,
			Columns: []string{artifact.PeriodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(period.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.period_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artifact.ProjectsTable,
			Columns: artifact.ProjectsPrimaryKey,
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
	if nodes := ac.mutation.PublicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artifact.PublicationsTable,
			Columns: artifact.PublicationsPrimaryKey,
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
	if nodes := ac.mutation.HoldersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artifact.HoldersTable,
			Columns: artifact.HoldersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CulturalAffiliationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.CulturalAffiliationTable,
			Columns: []string{artifact.CulturalAffiliationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(culture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.culture_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.MonumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.MonumentTable,
			Columns: []string{artifact.MonumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.monument_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.ModelTable,
			Columns: []string{artifact.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.model_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.SetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.SetTable,
			Columns: []string{artifact.SetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(set.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.set_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.LocationTable,
			Columns: []string{artifact.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.location_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.CollectionTable,
			Columns: []string{artifact.CollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.collection_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.LicenseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.LicenseTable,
			Columns: []string{artifact.LicenseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.license_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArtifactCreateBulk is the builder for creating many Artifact entities in bulk.
type ArtifactCreateBulk struct {
	config
	err      error
	builders []*ArtifactCreate
}

// Save creates the Artifact entities in the database.
func (acb *ArtifactCreateBulk) Save(ctx context.Context) ([]*Artifact, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Artifact, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtifactMutation)
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
func (acb *ArtifactCreateBulk) SaveX(ctx context.Context) []*Artifact {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtifactCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtifactCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
