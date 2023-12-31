// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/organization"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
)

// Person is the model entity for the Person schema.
type Person struct {
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
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// PhoneNumbers holds the value of the "phone_numbers" field.
	PhoneNumbers []string `json:"phone_numbers,omitempty"`
	// Emails holds the value of the "emails" field.
	Emails []string `json:"emails,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Abbreviation holds the value of the "abbreviation" field.
	Abbreviation string `json:"abbreviation,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ExternalLink holds the value of the "external_link" field.
	ExternalLink string `json:"external_link,omitempty"`
	// PrimaryImageURL holds the value of the "primary_image_url" field.
	PrimaryImageURL string `json:"primary_image_url,omitempty"`
	// AdditionalImagesUrls holds the value of the "additional_images_urls" field.
	AdditionalImagesUrls []string `json:"additional_images_urls,omitempty"`
	// GivenName holds the value of the "given_name" field.
	GivenName string `json:"given_name,omitempty"`
	// FamilyName holds the value of the "family_name" field.
	FamilyName string `json:"family_name,omitempty"`
	// PatronymicName holds the value of the "patronymic_name" field.
	PatronymicName string `json:"patronymic_name,omitempty"`
	// BeginData holds the value of the "begin_data" field.
	BeginData time.Time `json:"begin_data,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender person.Gender `json:"gender,omitempty"`
	// Occupation holds the value of the "occupation" field.
	Occupation string `json:"occupation,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonQuery when eager-loading is set.
	Edges               PersonEdges `json:"edges"`
	organization_people *int
	selectValues        sql.SelectValues
}

// PersonEdges holds the relations/edges for other nodes in the graph.
type PersonEdges struct {
	// Collections holds the value of the collections edge.
	Collections []*Collection `json:"collections,omitempty"`
	// Art holds the value of the art edge.
	Art []*Art `json:"art,omitempty"`
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Herbaria holds the value of the herbaria edge.
	Herbaria []*Herbarium `json:"herbaria,omitempty"`
	// ProtectedAreaPictures holds the value of the protected_area_pictures edge.
	ProtectedAreaPictures []*ProtectedAreaPicture `json:"protected_area_pictures,omitempty"`
	// DonatedArtifacts holds the value of the donated_artifacts edge.
	DonatedArtifacts []*Artifact `json:"donated_artifacts,omitempty"`
	// PetroglyphsAccountingDocumentation holds the value of the petroglyphs_accounting_documentation edge.
	PetroglyphsAccountingDocumentation []*Petroglyph `json:"petroglyphs_accounting_documentation,omitempty"`
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// Visits holds the value of the visits edge.
	Visits []*Visit `json:"visits,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// Publications holds the value of the publications edge.
	Publications []*Publication `json:"publications,omitempty"`
	// Affiliation holds the value of the affiliation edge.
	Affiliation *Organization `json:"affiliation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [12]bool
	// totalCount holds the count of the edges above.
	totalCount [12]map[string]int

	namedCollections                        map[string][]*Collection
	namedArt                                map[string][]*Art
	namedArtifacts                          map[string][]*Artifact
	namedHerbaria                           map[string][]*Herbarium
	namedProtectedAreaPictures              map[string][]*ProtectedAreaPicture
	namedDonatedArtifacts                   map[string][]*Artifact
	namedPetroglyphsAccountingDocumentation map[string][]*Petroglyph
	namedBooks                              map[string][]*Book
	namedVisits                             map[string][]*Visit
	namedProjects                           map[string][]*Project
	namedPublications                       map[string][]*Publication
}

// CollectionsOrErr returns the Collections value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) CollectionsOrErr() ([]*Collection, error) {
	if e.loadedTypes[0] {
		return e.Collections, nil
	}
	return nil, &NotLoadedError{edge: "collections"}
}

// ArtOrErr returns the Art value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) ArtOrErr() ([]*Art, error) {
	if e.loadedTypes[1] {
		return e.Art, nil
	}
	return nil, &NotLoadedError{edge: "art"}
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[2] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// HerbariaOrErr returns the Herbaria value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) HerbariaOrErr() ([]*Herbarium, error) {
	if e.loadedTypes[3] {
		return e.Herbaria, nil
	}
	return nil, &NotLoadedError{edge: "herbaria"}
}

// ProtectedAreaPicturesOrErr returns the ProtectedAreaPictures value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) ProtectedAreaPicturesOrErr() ([]*ProtectedAreaPicture, error) {
	if e.loadedTypes[4] {
		return e.ProtectedAreaPictures, nil
	}
	return nil, &NotLoadedError{edge: "protected_area_pictures"}
}

// DonatedArtifactsOrErr returns the DonatedArtifacts value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) DonatedArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[5] {
		return e.DonatedArtifacts, nil
	}
	return nil, &NotLoadedError{edge: "donated_artifacts"}
}

// PetroglyphsAccountingDocumentationOrErr returns the PetroglyphsAccountingDocumentation value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) PetroglyphsAccountingDocumentationOrErr() ([]*Petroglyph, error) {
	if e.loadedTypes[6] {
		return e.PetroglyphsAccountingDocumentation, nil
	}
	return nil, &NotLoadedError{edge: "petroglyphs_accounting_documentation"}
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[7] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// VisitsOrErr returns the Visits value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) VisitsOrErr() ([]*Visit, error) {
	if e.loadedTypes[8] {
		return e.Visits, nil
	}
	return nil, &NotLoadedError{edge: "visits"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[9] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// PublicationsOrErr returns the Publications value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) PublicationsOrErr() ([]*Publication, error) {
	if e.loadedTypes[10] {
		return e.Publications, nil
	}
	return nil, &NotLoadedError{edge: "publications"}
}

// AffiliationOrErr returns the Affiliation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) AffiliationOrErr() (*Organization, error) {
	if e.loadedTypes[11] {
		if e.Affiliation == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Affiliation, nil
	}
	return nil, &NotLoadedError{edge: "affiliation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Person) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case person.FieldPhoneNumbers, person.FieldEmails, person.FieldAdditionalImagesUrls:
			values[i] = new([]byte)
		case person.FieldID:
			values[i] = new(sql.NullInt64)
		case person.FieldCreatedBy, person.FieldUpdatedBy, person.FieldAddress, person.FieldDisplayName, person.FieldAbbreviation, person.FieldDescription, person.FieldExternalLink, person.FieldPrimaryImageURL, person.FieldGivenName, person.FieldFamilyName, person.FieldPatronymicName, person.FieldGender, person.FieldOccupation:
			values[i] = new(sql.NullString)
		case person.FieldCreatedAt, person.FieldUpdatedAt, person.FieldBeginData, person.FieldEndDate:
			values[i] = new(sql.NullTime)
		case person.ForeignKeys[0]: // organization_people
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Person fields.
func (pe *Person) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case person.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case person.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Time
			}
		case person.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pe.CreatedBy = value.String
			}
		case person.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Time
			}
		case person.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pe.UpdatedBy = value.String
			}
		case person.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				pe.Address = value.String
			}
		case person.FieldPhoneNumbers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field phone_numbers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.PhoneNumbers); err != nil {
					return fmt.Errorf("unmarshal field phone_numbers: %w", err)
				}
			}
		case person.FieldEmails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field emails", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.Emails); err != nil {
					return fmt.Errorf("unmarshal field emails: %w", err)
				}
			}
		case person.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pe.DisplayName = value.String
			}
		case person.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				pe.Abbreviation = value.String
			}
		case person.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = value.String
			}
		case person.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				pe.ExternalLink = value.String
			}
		case person.FieldPrimaryImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_image_url", values[i])
			} else if value.Valid {
				pe.PrimaryImageURL = value.String
			}
		case person.FieldAdditionalImagesUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_images_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.AdditionalImagesUrls); err != nil {
					return fmt.Errorf("unmarshal field additional_images_urls: %w", err)
				}
			}
		case person.FieldGivenName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field given_name", values[i])
			} else if value.Valid {
				pe.GivenName = value.String
			}
		case person.FieldFamilyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family_name", values[i])
			} else if value.Valid {
				pe.FamilyName = value.String
			}
		case person.FieldPatronymicName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field patronymic_name", values[i])
			} else if value.Valid {
				pe.PatronymicName = value.String
			}
		case person.FieldBeginData:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field begin_data", values[i])
			} else if value.Valid {
				pe.BeginData = value.Time
			}
		case person.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				pe.EndDate = value.Time
			}
		case person.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pe.Gender = person.Gender(value.String)
			}
		case person.FieldOccupation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field occupation", values[i])
			} else if value.Valid {
				pe.Occupation = value.String
			}
		case person.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field organization_people", value)
			} else if value.Valid {
				pe.organization_people = new(int)
				*pe.organization_people = int(value.Int64)
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Person.
// This includes values selected through modifiers, order, etc.
func (pe *Person) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryCollections queries the "collections" edge of the Person entity.
func (pe *Person) QueryCollections() *CollectionQuery {
	return NewPersonClient(pe.config).QueryCollections(pe)
}

// QueryArt queries the "art" edge of the Person entity.
func (pe *Person) QueryArt() *ArtQuery {
	return NewPersonClient(pe.config).QueryArt(pe)
}

// QueryArtifacts queries the "artifacts" edge of the Person entity.
func (pe *Person) QueryArtifacts() *ArtifactQuery {
	return NewPersonClient(pe.config).QueryArtifacts(pe)
}

// QueryHerbaria queries the "herbaria" edge of the Person entity.
func (pe *Person) QueryHerbaria() *HerbariumQuery {
	return NewPersonClient(pe.config).QueryHerbaria(pe)
}

// QueryProtectedAreaPictures queries the "protected_area_pictures" edge of the Person entity.
func (pe *Person) QueryProtectedAreaPictures() *ProtectedAreaPictureQuery {
	return NewPersonClient(pe.config).QueryProtectedAreaPictures(pe)
}

// QueryDonatedArtifacts queries the "donated_artifacts" edge of the Person entity.
func (pe *Person) QueryDonatedArtifacts() *ArtifactQuery {
	return NewPersonClient(pe.config).QueryDonatedArtifacts(pe)
}

// QueryPetroglyphsAccountingDocumentation queries the "petroglyphs_accounting_documentation" edge of the Person entity.
func (pe *Person) QueryPetroglyphsAccountingDocumentation() *PetroglyphQuery {
	return NewPersonClient(pe.config).QueryPetroglyphsAccountingDocumentation(pe)
}

// QueryBooks queries the "books" edge of the Person entity.
func (pe *Person) QueryBooks() *BookQuery {
	return NewPersonClient(pe.config).QueryBooks(pe)
}

// QueryVisits queries the "visits" edge of the Person entity.
func (pe *Person) QueryVisits() *VisitQuery {
	return NewPersonClient(pe.config).QueryVisits(pe)
}

// QueryProjects queries the "projects" edge of the Person entity.
func (pe *Person) QueryProjects() *ProjectQuery {
	return NewPersonClient(pe.config).QueryProjects(pe)
}

// QueryPublications queries the "publications" edge of the Person entity.
func (pe *Person) QueryPublications() *PublicationQuery {
	return NewPersonClient(pe.config).QueryPublications(pe)
}

// QueryAffiliation queries the "affiliation" edge of the Person entity.
func (pe *Person) QueryAffiliation() *OrganizationQuery {
	return NewPersonClient(pe.config).QueryAffiliation(pe)
}

// Update returns a builder for updating this Person.
// Note that you need to call Person.Unwrap() before calling this method if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return NewPersonClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Person entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Person is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pe.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pe.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(pe.Address)
	builder.WriteString(", ")
	builder.WriteString("phone_numbers=")
	builder.WriteString(fmt.Sprintf("%v", pe.PhoneNumbers))
	builder.WriteString(", ")
	builder.WriteString("emails=")
	builder.WriteString(fmt.Sprintf("%v", pe.Emails))
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(pe.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(pe.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pe.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(pe.ExternalLink)
	builder.WriteString(", ")
	builder.WriteString("primary_image_url=")
	builder.WriteString(pe.PrimaryImageURL)
	builder.WriteString(", ")
	builder.WriteString("additional_images_urls=")
	builder.WriteString(fmt.Sprintf("%v", pe.AdditionalImagesUrls))
	builder.WriteString(", ")
	builder.WriteString("given_name=")
	builder.WriteString(pe.GivenName)
	builder.WriteString(", ")
	builder.WriteString("family_name=")
	builder.WriteString(pe.FamilyName)
	builder.WriteString(", ")
	builder.WriteString("patronymic_name=")
	builder.WriteString(pe.PatronymicName)
	builder.WriteString(", ")
	builder.WriteString("begin_data=")
	builder.WriteString(pe.BeginData.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_date=")
	builder.WriteString(pe.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", pe.Gender))
	builder.WriteString(", ")
	builder.WriteString("occupation=")
	builder.WriteString(pe.Occupation)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCollections returns the Collections named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedCollections(name string) ([]*Collection, error) {
	if pe.Edges.namedCollections == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedCollections[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedCollections(name string, edges ...*Collection) {
	if pe.Edges.namedCollections == nil {
		pe.Edges.namedCollections = make(map[string][]*Collection)
	}
	if len(edges) == 0 {
		pe.Edges.namedCollections[name] = []*Collection{}
	} else {
		pe.Edges.namedCollections[name] = append(pe.Edges.namedCollections[name], edges...)
	}
}

// NamedArt returns the Art named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedArt(name string) ([]*Art, error) {
	if pe.Edges.namedArt == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedArt[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedArt(name string, edges ...*Art) {
	if pe.Edges.namedArt == nil {
		pe.Edges.namedArt = make(map[string][]*Art)
	}
	if len(edges) == 0 {
		pe.Edges.namedArt[name] = []*Art{}
	} else {
		pe.Edges.namedArt[name] = append(pe.Edges.namedArt[name], edges...)
	}
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedArtifacts(name string) ([]*Artifact, error) {
	if pe.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedArtifacts(name string, edges ...*Artifact) {
	if pe.Edges.namedArtifacts == nil {
		pe.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		pe.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		pe.Edges.namedArtifacts[name] = append(pe.Edges.namedArtifacts[name], edges...)
	}
}

// NamedHerbaria returns the Herbaria named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedHerbaria(name string) ([]*Herbarium, error) {
	if pe.Edges.namedHerbaria == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedHerbaria[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedHerbaria(name string, edges ...*Herbarium) {
	if pe.Edges.namedHerbaria == nil {
		pe.Edges.namedHerbaria = make(map[string][]*Herbarium)
	}
	if len(edges) == 0 {
		pe.Edges.namedHerbaria[name] = []*Herbarium{}
	} else {
		pe.Edges.namedHerbaria[name] = append(pe.Edges.namedHerbaria[name], edges...)
	}
}

// NamedProtectedAreaPictures returns the ProtectedAreaPictures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedProtectedAreaPictures(name string) ([]*ProtectedAreaPicture, error) {
	if pe.Edges.namedProtectedAreaPictures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedProtectedAreaPictures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedProtectedAreaPictures(name string, edges ...*ProtectedAreaPicture) {
	if pe.Edges.namedProtectedAreaPictures == nil {
		pe.Edges.namedProtectedAreaPictures = make(map[string][]*ProtectedAreaPicture)
	}
	if len(edges) == 0 {
		pe.Edges.namedProtectedAreaPictures[name] = []*ProtectedAreaPicture{}
	} else {
		pe.Edges.namedProtectedAreaPictures[name] = append(pe.Edges.namedProtectedAreaPictures[name], edges...)
	}
}

// NamedDonatedArtifacts returns the DonatedArtifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedDonatedArtifacts(name string) ([]*Artifact, error) {
	if pe.Edges.namedDonatedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedDonatedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedDonatedArtifacts(name string, edges ...*Artifact) {
	if pe.Edges.namedDonatedArtifacts == nil {
		pe.Edges.namedDonatedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		pe.Edges.namedDonatedArtifacts[name] = []*Artifact{}
	} else {
		pe.Edges.namedDonatedArtifacts[name] = append(pe.Edges.namedDonatedArtifacts[name], edges...)
	}
}

// NamedPetroglyphsAccountingDocumentation returns the PetroglyphsAccountingDocumentation named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedPetroglyphsAccountingDocumentation(name string) ([]*Petroglyph, error) {
	if pe.Edges.namedPetroglyphsAccountingDocumentation == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedPetroglyphsAccountingDocumentation[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedPetroglyphsAccountingDocumentation(name string, edges ...*Petroglyph) {
	if pe.Edges.namedPetroglyphsAccountingDocumentation == nil {
		pe.Edges.namedPetroglyphsAccountingDocumentation = make(map[string][]*Petroglyph)
	}
	if len(edges) == 0 {
		pe.Edges.namedPetroglyphsAccountingDocumentation[name] = []*Petroglyph{}
	} else {
		pe.Edges.namedPetroglyphsAccountingDocumentation[name] = append(pe.Edges.namedPetroglyphsAccountingDocumentation[name], edges...)
	}
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedBooks(name string) ([]*Book, error) {
	if pe.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedBooks(name string, edges ...*Book) {
	if pe.Edges.namedBooks == nil {
		pe.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		pe.Edges.namedBooks[name] = []*Book{}
	} else {
		pe.Edges.namedBooks[name] = append(pe.Edges.namedBooks[name], edges...)
	}
}

// NamedVisits returns the Visits named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedVisits(name string) ([]*Visit, error) {
	if pe.Edges.namedVisits == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedVisits[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedVisits(name string, edges ...*Visit) {
	if pe.Edges.namedVisits == nil {
		pe.Edges.namedVisits = make(map[string][]*Visit)
	}
	if len(edges) == 0 {
		pe.Edges.namedVisits[name] = []*Visit{}
	} else {
		pe.Edges.namedVisits[name] = append(pe.Edges.namedVisits[name], edges...)
	}
}

// NamedProjects returns the Projects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedProjects(name string) ([]*Project, error) {
	if pe.Edges.namedProjects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedProjects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedProjects(name string, edges ...*Project) {
	if pe.Edges.namedProjects == nil {
		pe.Edges.namedProjects = make(map[string][]*Project)
	}
	if len(edges) == 0 {
		pe.Edges.namedProjects[name] = []*Project{}
	} else {
		pe.Edges.namedProjects[name] = append(pe.Edges.namedProjects[name], edges...)
	}
}

// NamedPublications returns the Publications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedPublications(name string) ([]*Publication, error) {
	if pe.Edges.namedPublications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedPublications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedPublications(name string, edges ...*Publication) {
	if pe.Edges.namedPublications == nil {
		pe.Edges.namedPublications = make(map[string][]*Publication)
	}
	if len(edges) == 0 {
		pe.Edges.namedPublications[name] = []*Publication{}
	} else {
		pe.Edges.namedPublications[name] = append(pe.Edges.namedPublications[name], edges...)
	}
}

// Persons is a parsable slice of Person.
type Persons []*Person
