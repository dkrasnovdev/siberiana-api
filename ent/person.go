// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/holder"
	"github.com/dkrasnovdev/heritage-api/ent/organization"
	"github.com/dkrasnovdev/heritage-api/ent/person"
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
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ExternalLinks holds the value of the "external_links" field.
	ExternalLinks []string `json:"external_links,omitempty"`
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonQuery when eager-loading is set.
	Edges               PersonEdges `json:"edges"`
	holder_person       *int
	organization_people *int
	selectValues        sql.SelectValues
}

// PersonEdges holds the relations/edges for other nodes in the graph.
type PersonEdges struct {
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// Publications holds the value of the publications edge.
	Publications []*Publication `json:"publications,omitempty"`
	// PersonRoles holds the value of the person_roles edge.
	PersonRoles []*PersonRole `json:"person_roles,omitempty"`
	// Holder holds the value of the holder edge.
	Holder *Holder `json:"holder,omitempty"`
	// Affiliation holds the value of the affiliation edge.
	Affiliation *Organization `json:"affiliation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedArtifacts    map[string][]*Artifact
	namedProjects     map[string][]*Project
	namedPublications map[string][]*Publication
	namedPersonRoles  map[string][]*PersonRole
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[0] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[1] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// PublicationsOrErr returns the Publications value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) PublicationsOrErr() ([]*Publication, error) {
	if e.loadedTypes[2] {
		return e.Publications, nil
	}
	return nil, &NotLoadedError{edge: "publications"}
}

// PersonRolesOrErr returns the PersonRoles value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) PersonRolesOrErr() ([]*PersonRole, error) {
	if e.loadedTypes[3] {
		return e.PersonRoles, nil
	}
	return nil, &NotLoadedError{edge: "person_roles"}
}

// HolderOrErr returns the Holder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) HolderOrErr() (*Holder, error) {
	if e.loadedTypes[4] {
		if e.Holder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: holder.Label}
		}
		return e.Holder, nil
	}
	return nil, &NotLoadedError{edge: "holder"}
}

// AffiliationOrErr returns the Affiliation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) AffiliationOrErr() (*Organization, error) {
	if e.loadedTypes[5] {
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
		case person.FieldPhoneNumbers, person.FieldEmails, person.FieldExternalLinks, person.FieldAdditionalImagesUrls:
			values[i] = new([]byte)
		case person.FieldID:
			values[i] = new(sql.NullInt64)
		case person.FieldCreatedBy, person.FieldUpdatedBy, person.FieldAddress, person.FieldDisplayName, person.FieldDescription, person.FieldPrimaryImageURL, person.FieldGivenName, person.FieldFamilyName, person.FieldPatronymicName, person.FieldGender:
			values[i] = new(sql.NullString)
		case person.FieldCreatedAt, person.FieldUpdatedAt, person.FieldBeginData, person.FieldEndDate:
			values[i] = new(sql.NullTime)
		case person.ForeignKeys[0]: // holder_person
			values[i] = new(sql.NullInt64)
		case person.ForeignKeys[1]: // organization_people
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
		case person.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = value.String
			}
		case person.FieldExternalLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field external_links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.ExternalLinks); err != nil {
					return fmt.Errorf("unmarshal field external_links: %w", err)
				}
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
		case person.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field holder_person", value)
			} else if value.Valid {
				pe.holder_person = new(int)
				*pe.holder_person = int(value.Int64)
			}
		case person.ForeignKeys[1]:
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

// QueryArtifacts queries the "artifacts" edge of the Person entity.
func (pe *Person) QueryArtifacts() *ArtifactQuery {
	return NewPersonClient(pe.config).QueryArtifacts(pe)
}

// QueryProjects queries the "projects" edge of the Person entity.
func (pe *Person) QueryProjects() *ProjectQuery {
	return NewPersonClient(pe.config).QueryProjects(pe)
}

// QueryPublications queries the "publications" edge of the Person entity.
func (pe *Person) QueryPublications() *PublicationQuery {
	return NewPersonClient(pe.config).QueryPublications(pe)
}

// QueryPersonRoles queries the "person_roles" edge of the Person entity.
func (pe *Person) QueryPersonRoles() *PersonRoleQuery {
	return NewPersonClient(pe.config).QueryPersonRoles(pe)
}

// QueryHolder queries the "holder" edge of the Person entity.
func (pe *Person) QueryHolder() *HolderQuery {
	return NewPersonClient(pe.config).QueryHolder(pe)
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
	builder.WriteString("description=")
	builder.WriteString(pe.Description)
	builder.WriteString(", ")
	builder.WriteString("external_links=")
	builder.WriteString(fmt.Sprintf("%v", pe.ExternalLinks))
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
	builder.WriteByte(')')
	return builder.String()
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

// NamedPersonRoles returns the PersonRoles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedPersonRoles(name string) ([]*PersonRole, error) {
	if pe.Edges.namedPersonRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedPersonRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedPersonRoles(name string, edges ...*PersonRole) {
	if pe.Edges.namedPersonRoles == nil {
		pe.Edges.namedPersonRoles = make(map[string][]*PersonRole)
	}
	if len(edges) == 0 {
		pe.Edges.namedPersonRoles[name] = []*PersonRole{}
	} else {
		pe.Edges.namedPersonRoles[name] = append(pe.Edges.namedPersonRoles[name], edges...)
	}
}

// Persons is a parsable slice of Person.
type Persons []*Person
