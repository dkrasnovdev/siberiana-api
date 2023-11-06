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
)

// Organization is the model entity for the Organization schema.
type Organization struct {
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
	// PreviousNames holds the value of the "previous_names" field.
	PreviousNames []string `json:"previous_names,omitempty"`
	// IsInAConsortium holds the value of the "is_in_a_consortium" field.
	IsInAConsortium bool `json:"is_in_a_consortium,omitempty"`
	// ConsortiumDocumentURL holds the value of the "consortium_document_url" field.
	ConsortiumDocumentURL string `json:"consortium_document_url,omitempty"`
	// Type holds the value of the "type" field.
	Type organization.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges        OrganizationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// People holds the value of the people edge.
	People []*Person `json:"people,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedArtifacts map[string][]*Artifact
	namedBooks     map[string][]*Book
	namedPeople    map[string][]*Person
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[0] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[1] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// PeopleOrErr returns the People value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) PeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[2] {
		return e.People, nil
	}
	return nil, &NotLoadedError{edge: "people"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldPhoneNumbers, organization.FieldEmails, organization.FieldAdditionalImagesUrls, organization.FieldPreviousNames:
			values[i] = new([]byte)
		case organization.FieldIsInAConsortium:
			values[i] = new(sql.NullBool)
		case organization.FieldID:
			values[i] = new(sql.NullInt64)
		case organization.FieldCreatedBy, organization.FieldUpdatedBy, organization.FieldAddress, organization.FieldDisplayName, organization.FieldAbbreviation, organization.FieldDescription, organization.FieldExternalLink, organization.FieldPrimaryImageURL, organization.FieldConsortiumDocumentURL, organization.FieldType:
			values[i] = new(sql.NullString)
		case organization.FieldCreatedAt, organization.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case organization.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case organization.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				o.CreatedBy = value.String
			}
		case organization.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case organization.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				o.UpdatedBy = value.String
			}
		case organization.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				o.Address = value.String
			}
		case organization.FieldPhoneNumbers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field phone_numbers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.PhoneNumbers); err != nil {
					return fmt.Errorf("unmarshal field phone_numbers: %w", err)
				}
			}
		case organization.FieldEmails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field emails", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.Emails); err != nil {
					return fmt.Errorf("unmarshal field emails: %w", err)
				}
			}
		case organization.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				o.DisplayName = value.String
			}
		case organization.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				o.Abbreviation = value.String
			}
		case organization.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				o.Description = value.String
			}
		case organization.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				o.ExternalLink = value.String
			}
		case organization.FieldPrimaryImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_image_url", values[i])
			} else if value.Valid {
				o.PrimaryImageURL = value.String
			}
		case organization.FieldAdditionalImagesUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_images_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.AdditionalImagesUrls); err != nil {
					return fmt.Errorf("unmarshal field additional_images_urls: %w", err)
				}
			}
		case organization.FieldPreviousNames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field previous_names", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.PreviousNames); err != nil {
					return fmt.Errorf("unmarshal field previous_names: %w", err)
				}
			}
		case organization.FieldIsInAConsortium:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_in_a_consortium", values[i])
			} else if value.Valid {
				o.IsInAConsortium = value.Bool
			}
		case organization.FieldConsortiumDocumentURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field consortium_document_url", values[i])
			} else if value.Valid {
				o.ConsortiumDocumentURL = value.String
			}
		case organization.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				o.Type = organization.Type(value.String)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Organization.
// This includes values selected through modifiers, order, etc.
func (o *Organization) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryArtifacts queries the "artifacts" edge of the Organization entity.
func (o *Organization) QueryArtifacts() *ArtifactQuery {
	return NewOrganizationClient(o.config).QueryArtifacts(o)
}

// QueryBooks queries the "books" edge of the Organization entity.
func (o *Organization) QueryBooks() *BookQuery {
	return NewOrganizationClient(o.config).QueryBooks(o)
}

// QueryPeople queries the "people" edge of the Organization entity.
func (o *Organization) QueryPeople() *PersonQuery {
	return NewOrganizationClient(o.config).QueryPeople(o)
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return NewOrganizationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(o.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(o.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(o.Address)
	builder.WriteString(", ")
	builder.WriteString("phone_numbers=")
	builder.WriteString(fmt.Sprintf("%v", o.PhoneNumbers))
	builder.WriteString(", ")
	builder.WriteString("emails=")
	builder.WriteString(fmt.Sprintf("%v", o.Emails))
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(o.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(o.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(o.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(o.ExternalLink)
	builder.WriteString(", ")
	builder.WriteString("primary_image_url=")
	builder.WriteString(o.PrimaryImageURL)
	builder.WriteString(", ")
	builder.WriteString("additional_images_urls=")
	builder.WriteString(fmt.Sprintf("%v", o.AdditionalImagesUrls))
	builder.WriteString(", ")
	builder.WriteString("previous_names=")
	builder.WriteString(fmt.Sprintf("%v", o.PreviousNames))
	builder.WriteString(", ")
	builder.WriteString("is_in_a_consortium=")
	builder.WriteString(fmt.Sprintf("%v", o.IsInAConsortium))
	builder.WriteString(", ")
	builder.WriteString("consortium_document_url=")
	builder.WriteString(o.ConsortiumDocumentURL)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", o.Type))
	builder.WriteByte(')')
	return builder.String()
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedArtifacts(name string) ([]*Artifact, error) {
	if o.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedArtifacts(name string, edges ...*Artifact) {
	if o.Edges.namedArtifacts == nil {
		o.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		o.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		o.Edges.namedArtifacts[name] = append(o.Edges.namedArtifacts[name], edges...)
	}
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedBooks(name string) ([]*Book, error) {
	if o.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedBooks(name string, edges ...*Book) {
	if o.Edges.namedBooks == nil {
		o.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		o.Edges.namedBooks[name] = []*Book{}
	} else {
		o.Edges.namedBooks[name] = append(o.Edges.namedBooks[name], edges...)
	}
}

// NamedPeople returns the People named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedPeople(name string) ([]*Person, error) {
	if o.Edges.namedPeople == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedPeople[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedPeople(name string, edges ...*Person) {
	if o.Edges.namedPeople == nil {
		o.Edges.namedPeople = make(map[string][]*Person)
	}
	if len(edges) == 0 {
		o.Edges.namedPeople[name] = []*Person{}
	} else {
		o.Edges.namedPeople[name] = append(o.Edges.namedPeople[name], edges...)
	}
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization
