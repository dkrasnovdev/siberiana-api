// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/artifact"
	"github.com/dkrasnovdev/heritage-api/ent/collection"
	"github.com/dkrasnovdev/heritage-api/ent/culture"
	"github.com/dkrasnovdev/heritage-api/ent/license"
	"github.com/dkrasnovdev/heritage-api/ent/location"
	"github.com/dkrasnovdev/heritage-api/ent/model"
	"github.com/dkrasnovdev/heritage-api/ent/monument"
	"github.com/dkrasnovdev/heritage-api/ent/period"
	"github.com/dkrasnovdev/heritage-api/ent/set"
)

// Artifact is the model entity for the Artifact schema.
type Artifact struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Abbreviation holds the value of the "abbreviation" field.
	Abbreviation string `json:"abbreviation,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ExternalLinks holds the value of the "external_links" field.
	ExternalLinks []string `json:"external_links,omitempty"`
	// PrimaryImageURL holds the value of the "primary_image_url" field.
	PrimaryImageURL string `json:"primary_image_url,omitempty"`
	// AdditionalImagesUrls holds the value of the "additional_images_urls" field.
	AdditionalImagesUrls []string `json:"additional_images_urls,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// Dating holds the value of the "dating" field.
	Dating string `json:"dating,omitempty"`
	// Dimensions holds the value of the "dimensions" field.
	Dimensions string `json:"dimensions,omitempty"`
	// ChemicalComposition holds the value of the "chemical_composition" field.
	ChemicalComposition string `json:"chemical_composition,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Typology holds the value of the "typology" field.
	Typology string `json:"typology,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight string `json:"weight,omitempty"`
	// AdmissionDate holds the value of the "admission_date" field.
	AdmissionDate time.Time `json:"admission_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtifactQuery when eager-loading is set.
	Edges                ArtifactEdges `json:"edges"`
	collection_artifacts *int
	culture_artifacts    *int
	license_artifacts    *int
	location_artifacts   *int
	model_artifacts      *int
	monument_artifacts   *int
	period_artifacts     *int
	set_artifacts        *int
	selectValues         sql.SelectValues
}

// ArtifactEdges holds the relations/edges for other nodes in the graph.
type ArtifactEdges struct {
	// Authors holds the value of the authors edge.
	Authors []*Person `json:"authors,omitempty"`
	// Mediums holds the value of the mediums edge.
	Mediums []*Medium `json:"mediums,omitempty"`
	// Techniques holds the value of the techniques edge.
	Techniques []*Technique `json:"techniques,omitempty"`
	// Period holds the value of the period edge.
	Period *Period `json:"period,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// Publications holds the value of the publications edge.
	Publications []*Publication `json:"publications,omitempty"`
	// Holders holds the value of the holders edge.
	Holders []*Holder `json:"holders,omitempty"`
	// CulturalAffiliation holds the value of the cultural_affiliation edge.
	CulturalAffiliation *Culture `json:"cultural_affiliation,omitempty"`
	// Monument holds the value of the monument edge.
	Monument *Monument `json:"monument,omitempty"`
	// Model holds the value of the model edge.
	Model *Model `json:"model,omitempty"`
	// Set holds the value of the set edge.
	Set *Set `json:"set,omitempty"`
	// Location holds the value of the location edge.
	Location *Location `json:"location,omitempty"`
	// Collection holds the value of the collection edge.
	Collection *Collection `json:"collection,omitempty"`
	// License holds the value of the license edge.
	License *License `json:"license,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [14]bool
	// totalCount holds the count of the edges above.
	totalCount [14]map[string]int

	namedAuthors      map[string][]*Person
	namedMediums      map[string][]*Medium
	namedTechniques   map[string][]*Technique
	namedProjects     map[string][]*Project
	namedPublications map[string][]*Publication
	namedHolders      map[string][]*Holder
}

// AuthorsOrErr returns the Authors value or an error if the edge
// was not loaded in eager-loading.
func (e ArtifactEdges) AuthorsOrErr() ([]*Person, error) {
	if e.loadedTypes[0] {
		return e.Authors, nil
	}
	return nil, &NotLoadedError{edge: "authors"}
}

// MediumsOrErr returns the Mediums value or an error if the edge
// was not loaded in eager-loading.
func (e ArtifactEdges) MediumsOrErr() ([]*Medium, error) {
	if e.loadedTypes[1] {
		return e.Mediums, nil
	}
	return nil, &NotLoadedError{edge: "mediums"}
}

// TechniquesOrErr returns the Techniques value or an error if the edge
// was not loaded in eager-loading.
func (e ArtifactEdges) TechniquesOrErr() ([]*Technique, error) {
	if e.loadedTypes[2] {
		return e.Techniques, nil
	}
	return nil, &NotLoadedError{edge: "techniques"}
}

// PeriodOrErr returns the Period value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) PeriodOrErr() (*Period, error) {
	if e.loadedTypes[3] {
		if e.Period == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: period.Label}
		}
		return e.Period, nil
	}
	return nil, &NotLoadedError{edge: "period"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e ArtifactEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[4] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// PublicationsOrErr returns the Publications value or an error if the edge
// was not loaded in eager-loading.
func (e ArtifactEdges) PublicationsOrErr() ([]*Publication, error) {
	if e.loadedTypes[5] {
		return e.Publications, nil
	}
	return nil, &NotLoadedError{edge: "publications"}
}

// HoldersOrErr returns the Holders value or an error if the edge
// was not loaded in eager-loading.
func (e ArtifactEdges) HoldersOrErr() ([]*Holder, error) {
	if e.loadedTypes[6] {
		return e.Holders, nil
	}
	return nil, &NotLoadedError{edge: "holders"}
}

// CulturalAffiliationOrErr returns the CulturalAffiliation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) CulturalAffiliationOrErr() (*Culture, error) {
	if e.loadedTypes[7] {
		if e.CulturalAffiliation == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: culture.Label}
		}
		return e.CulturalAffiliation, nil
	}
	return nil, &NotLoadedError{edge: "cultural_affiliation"}
}

// MonumentOrErr returns the Monument value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) MonumentOrErr() (*Monument, error) {
	if e.loadedTypes[8] {
		if e.Monument == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: monument.Label}
		}
		return e.Monument, nil
	}
	return nil, &NotLoadedError{edge: "monument"}
}

// ModelOrErr returns the Model value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) ModelOrErr() (*Model, error) {
	if e.loadedTypes[9] {
		if e.Model == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: model.Label}
		}
		return e.Model, nil
	}
	return nil, &NotLoadedError{edge: "model"}
}

// SetOrErr returns the Set value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) SetOrErr() (*Set, error) {
	if e.loadedTypes[10] {
		if e.Set == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: set.Label}
		}
		return e.Set, nil
	}
	return nil, &NotLoadedError{edge: "set"}
}

// LocationOrErr returns the Location value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) LocationOrErr() (*Location, error) {
	if e.loadedTypes[11] {
		if e.Location == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: location.Label}
		}
		return e.Location, nil
	}
	return nil, &NotLoadedError{edge: "location"}
}

// CollectionOrErr returns the Collection value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) CollectionOrErr() (*Collection, error) {
	if e.loadedTypes[12] {
		if e.Collection == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: collection.Label}
		}
		return e.Collection, nil
	}
	return nil, &NotLoadedError{edge: "collection"}
}

// LicenseOrErr returns the License value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) LicenseOrErr() (*License, error) {
	if e.loadedTypes[13] {
		if e.License == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: license.Label}
		}
		return e.License, nil
	}
	return nil, &NotLoadedError{edge: "license"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Artifact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case artifact.FieldExternalLinks, artifact.FieldAdditionalImagesUrls:
			values[i] = new([]byte)
		case artifact.FieldID:
			values[i] = new(sql.NullInt64)
		case artifact.FieldCreatedBy, artifact.FieldUpdatedBy, artifact.FieldAbbreviation, artifact.FieldDisplayName, artifact.FieldDescription, artifact.FieldPrimaryImageURL, artifact.FieldDeletedBy, artifact.FieldDating, artifact.FieldDimensions, artifact.FieldChemicalComposition, artifact.FieldNumber, artifact.FieldTypology, artifact.FieldWeight:
			values[i] = new(sql.NullString)
		case artifact.FieldCreatedAt, artifact.FieldUpdatedAt, artifact.FieldDeletedAt, artifact.FieldAdmissionDate:
			values[i] = new(sql.NullTime)
		case artifact.ForeignKeys[0]: // collection_artifacts
			values[i] = new(sql.NullInt64)
		case artifact.ForeignKeys[1]: // culture_artifacts
			values[i] = new(sql.NullInt64)
		case artifact.ForeignKeys[2]: // license_artifacts
			values[i] = new(sql.NullInt64)
		case artifact.ForeignKeys[3]: // location_artifacts
			values[i] = new(sql.NullInt64)
		case artifact.ForeignKeys[4]: // model_artifacts
			values[i] = new(sql.NullInt64)
		case artifact.ForeignKeys[5]: // monument_artifacts
			values[i] = new(sql.NullInt64)
		case artifact.ForeignKeys[6]: // period_artifacts
			values[i] = new(sql.NullInt64)
		case artifact.ForeignKeys[7]: // set_artifacts
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Artifact fields.
func (a *Artifact) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artifact.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case artifact.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case artifact.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				a.CreatedBy = value.String
			}
		case artifact.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case artifact.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				a.UpdatedBy = value.String
			}
		case artifact.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				a.Abbreviation = value.String
			}
		case artifact.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				a.DisplayName = value.String
			}
		case artifact.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case artifact.FieldExternalLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field external_links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.ExternalLinks); err != nil {
					return fmt.Errorf("unmarshal field external_links: %w", err)
				}
			}
		case artifact.FieldPrimaryImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_image_url", values[i])
			} else if value.Valid {
				a.PrimaryImageURL = value.String
			}
		case artifact.FieldAdditionalImagesUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_images_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.AdditionalImagesUrls); err != nil {
					return fmt.Errorf("unmarshal field additional_images_urls: %w", err)
				}
			}
		case artifact.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case artifact.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				a.DeletedBy = value.String
			}
		case artifact.FieldDating:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dating", values[i])
			} else if value.Valid {
				a.Dating = value.String
			}
		case artifact.FieldDimensions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dimensions", values[i])
			} else if value.Valid {
				a.Dimensions = value.String
			}
		case artifact.FieldChemicalComposition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chemical_composition", values[i])
			} else if value.Valid {
				a.ChemicalComposition = value.String
			}
		case artifact.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				a.Number = value.String
			}
		case artifact.FieldTypology:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field typology", values[i])
			} else if value.Valid {
				a.Typology = value.String
			}
		case artifact.FieldWeight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				a.Weight = value.String
			}
		case artifact.FieldAdmissionDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field admission_date", values[i])
			} else if value.Valid {
				a.AdmissionDate = value.Time
			}
		case artifact.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field collection_artifacts", value)
			} else if value.Valid {
				a.collection_artifacts = new(int)
				*a.collection_artifacts = int(value.Int64)
			}
		case artifact.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field culture_artifacts", value)
			} else if value.Valid {
				a.culture_artifacts = new(int)
				*a.culture_artifacts = int(value.Int64)
			}
		case artifact.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field license_artifacts", value)
			} else if value.Valid {
				a.license_artifacts = new(int)
				*a.license_artifacts = int(value.Int64)
			}
		case artifact.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field location_artifacts", value)
			} else if value.Valid {
				a.location_artifacts = new(int)
				*a.location_artifacts = int(value.Int64)
			}
		case artifact.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field model_artifacts", value)
			} else if value.Valid {
				a.model_artifacts = new(int)
				*a.model_artifacts = int(value.Int64)
			}
		case artifact.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field monument_artifacts", value)
			} else if value.Valid {
				a.monument_artifacts = new(int)
				*a.monument_artifacts = int(value.Int64)
			}
		case artifact.ForeignKeys[6]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field period_artifacts", value)
			} else if value.Valid {
				a.period_artifacts = new(int)
				*a.period_artifacts = int(value.Int64)
			}
		case artifact.ForeignKeys[7]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field set_artifacts", value)
			} else if value.Valid {
				a.set_artifacts = new(int)
				*a.set_artifacts = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Artifact.
// This includes values selected through modifiers, order, etc.
func (a *Artifact) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryAuthors queries the "authors" edge of the Artifact entity.
func (a *Artifact) QueryAuthors() *PersonQuery {
	return NewArtifactClient(a.config).QueryAuthors(a)
}

// QueryMediums queries the "mediums" edge of the Artifact entity.
func (a *Artifact) QueryMediums() *MediumQuery {
	return NewArtifactClient(a.config).QueryMediums(a)
}

// QueryTechniques queries the "techniques" edge of the Artifact entity.
func (a *Artifact) QueryTechniques() *TechniqueQuery {
	return NewArtifactClient(a.config).QueryTechniques(a)
}

// QueryPeriod queries the "period" edge of the Artifact entity.
func (a *Artifact) QueryPeriod() *PeriodQuery {
	return NewArtifactClient(a.config).QueryPeriod(a)
}

// QueryProjects queries the "projects" edge of the Artifact entity.
func (a *Artifact) QueryProjects() *ProjectQuery {
	return NewArtifactClient(a.config).QueryProjects(a)
}

// QueryPublications queries the "publications" edge of the Artifact entity.
func (a *Artifact) QueryPublications() *PublicationQuery {
	return NewArtifactClient(a.config).QueryPublications(a)
}

// QueryHolders queries the "holders" edge of the Artifact entity.
func (a *Artifact) QueryHolders() *HolderQuery {
	return NewArtifactClient(a.config).QueryHolders(a)
}

// QueryCulturalAffiliation queries the "cultural_affiliation" edge of the Artifact entity.
func (a *Artifact) QueryCulturalAffiliation() *CultureQuery {
	return NewArtifactClient(a.config).QueryCulturalAffiliation(a)
}

// QueryMonument queries the "monument" edge of the Artifact entity.
func (a *Artifact) QueryMonument() *MonumentQuery {
	return NewArtifactClient(a.config).QueryMonument(a)
}

// QueryModel queries the "model" edge of the Artifact entity.
func (a *Artifact) QueryModel() *ModelQuery {
	return NewArtifactClient(a.config).QueryModel(a)
}

// QuerySet queries the "set" edge of the Artifact entity.
func (a *Artifact) QuerySet() *SetQuery {
	return NewArtifactClient(a.config).QuerySet(a)
}

// QueryLocation queries the "location" edge of the Artifact entity.
func (a *Artifact) QueryLocation() *LocationQuery {
	return NewArtifactClient(a.config).QueryLocation(a)
}

// QueryCollection queries the "collection" edge of the Artifact entity.
func (a *Artifact) QueryCollection() *CollectionQuery {
	return NewArtifactClient(a.config).QueryCollection(a)
}

// QueryLicense queries the "license" edge of the Artifact entity.
func (a *Artifact) QueryLicense() *LicenseQuery {
	return NewArtifactClient(a.config).QueryLicense(a)
}

// Update returns a builder for updating this Artifact.
// Note that you need to call Artifact.Unwrap() before calling this method if this Artifact
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Artifact) Update() *ArtifactUpdateOne {
	return NewArtifactClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Artifact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Artifact) Unwrap() *Artifact {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Artifact is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Artifact) String() string {
	var builder strings.Builder
	builder.WriteString("Artifact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(a.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(a.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(a.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(a.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("external_links=")
	builder.WriteString(fmt.Sprintf("%v", a.ExternalLinks))
	builder.WriteString(", ")
	builder.WriteString("primary_image_url=")
	builder.WriteString(a.PrimaryImageURL)
	builder.WriteString(", ")
	builder.WriteString("additional_images_urls=")
	builder.WriteString(fmt.Sprintf("%v", a.AdditionalImagesUrls))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(a.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("dating=")
	builder.WriteString(a.Dating)
	builder.WriteString(", ")
	builder.WriteString("dimensions=")
	builder.WriteString(a.Dimensions)
	builder.WriteString(", ")
	builder.WriteString("chemical_composition=")
	builder.WriteString(a.ChemicalComposition)
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(a.Number)
	builder.WriteString(", ")
	builder.WriteString("typology=")
	builder.WriteString(a.Typology)
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(a.Weight)
	builder.WriteString(", ")
	builder.WriteString("admission_date=")
	builder.WriteString(a.AdmissionDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedAuthors returns the Authors named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Artifact) NamedAuthors(name string) ([]*Person, error) {
	if a.Edges.namedAuthors == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedAuthors[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Artifact) appendNamedAuthors(name string, edges ...*Person) {
	if a.Edges.namedAuthors == nil {
		a.Edges.namedAuthors = make(map[string][]*Person)
	}
	if len(edges) == 0 {
		a.Edges.namedAuthors[name] = []*Person{}
	} else {
		a.Edges.namedAuthors[name] = append(a.Edges.namedAuthors[name], edges...)
	}
}

// NamedMediums returns the Mediums named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Artifact) NamedMediums(name string) ([]*Medium, error) {
	if a.Edges.namedMediums == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedMediums[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Artifact) appendNamedMediums(name string, edges ...*Medium) {
	if a.Edges.namedMediums == nil {
		a.Edges.namedMediums = make(map[string][]*Medium)
	}
	if len(edges) == 0 {
		a.Edges.namedMediums[name] = []*Medium{}
	} else {
		a.Edges.namedMediums[name] = append(a.Edges.namedMediums[name], edges...)
	}
}

// NamedTechniques returns the Techniques named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Artifact) NamedTechniques(name string) ([]*Technique, error) {
	if a.Edges.namedTechniques == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedTechniques[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Artifact) appendNamedTechniques(name string, edges ...*Technique) {
	if a.Edges.namedTechniques == nil {
		a.Edges.namedTechniques = make(map[string][]*Technique)
	}
	if len(edges) == 0 {
		a.Edges.namedTechniques[name] = []*Technique{}
	} else {
		a.Edges.namedTechniques[name] = append(a.Edges.namedTechniques[name], edges...)
	}
}

// NamedProjects returns the Projects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Artifact) NamedProjects(name string) ([]*Project, error) {
	if a.Edges.namedProjects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedProjects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Artifact) appendNamedProjects(name string, edges ...*Project) {
	if a.Edges.namedProjects == nil {
		a.Edges.namedProjects = make(map[string][]*Project)
	}
	if len(edges) == 0 {
		a.Edges.namedProjects[name] = []*Project{}
	} else {
		a.Edges.namedProjects[name] = append(a.Edges.namedProjects[name], edges...)
	}
}

// NamedPublications returns the Publications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Artifact) NamedPublications(name string) ([]*Publication, error) {
	if a.Edges.namedPublications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedPublications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Artifact) appendNamedPublications(name string, edges ...*Publication) {
	if a.Edges.namedPublications == nil {
		a.Edges.namedPublications = make(map[string][]*Publication)
	}
	if len(edges) == 0 {
		a.Edges.namedPublications[name] = []*Publication{}
	} else {
		a.Edges.namedPublications[name] = append(a.Edges.namedPublications[name], edges...)
	}
}

// NamedHolders returns the Holders named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Artifact) NamedHolders(name string) ([]*Holder, error) {
	if a.Edges.namedHolders == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedHolders[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Artifact) appendNamedHolders(name string, edges ...*Holder) {
	if a.Edges.namedHolders == nil {
		a.Edges.namedHolders = make(map[string][]*Holder)
	}
	if len(edges) == 0 {
		a.Edges.namedHolders[name] = []*Holder{}
	} else {
		a.Edges.namedHolders[name] = append(a.Edges.namedHolders[name], edges...)
	}
}

// Artifacts is a parsable slice of Artifact.
type Artifacts []*Artifact
