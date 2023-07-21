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
	"github.com/dkrasnovdev/heritage-api/ent/organizationtype"
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
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ExternalLinks holds the value of the "external_links" field.
	ExternalLinks []string `json:"external_links,omitempty"`
	// PrimaryImageURL holds the value of the "primary_image_url" field.
	PrimaryImageURL string `json:"primary_image_url,omitempty"`
	// AdditionalImagesUrls holds the value of the "additional_images_urls" field.
	AdditionalImagesUrls []string `json:"additional_images_urls,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges                           OrganizationEdges `json:"edges"`
	holder_organization             *int
	organization_type_organizations *int
	selectValues                    sql.SelectValues
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// People holds the value of the people edge.
	People []*Person `json:"people,omitempty"`
	// Holder holds the value of the holder edge.
	Holder *Holder `json:"holder,omitempty"`
	// OrganizationType holds the value of the organization_type edge.
	OrganizationType *OrganizationType `json:"organization_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedPeople map[string][]*Person
}

// PeopleOrErr returns the People value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) PeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[0] {
		return e.People, nil
	}
	return nil, &NotLoadedError{edge: "people"}
}

// HolderOrErr returns the Holder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationEdges) HolderOrErr() (*Holder, error) {
	if e.loadedTypes[1] {
		if e.Holder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: holder.Label}
		}
		return e.Holder, nil
	}
	return nil, &NotLoadedError{edge: "holder"}
}

// OrganizationTypeOrErr returns the OrganizationType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationEdges) OrganizationTypeOrErr() (*OrganizationType, error) {
	if e.loadedTypes[2] {
		if e.OrganizationType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organizationtype.Label}
		}
		return e.OrganizationType, nil
	}
	return nil, &NotLoadedError{edge: "organization_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldPhoneNumbers, organization.FieldEmails, organization.FieldExternalLinks, organization.FieldAdditionalImagesUrls:
			values[i] = new([]byte)
		case organization.FieldID:
			values[i] = new(sql.NullInt64)
		case organization.FieldCreatedBy, organization.FieldUpdatedBy, organization.FieldAddress, organization.FieldDisplayName, organization.FieldDescription, organization.FieldPrimaryImageURL:
			values[i] = new(sql.NullString)
		case organization.FieldCreatedAt, organization.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case organization.ForeignKeys[0]: // holder_organization
			values[i] = new(sql.NullInt64)
		case organization.ForeignKeys[1]: // organization_type_organizations
			values[i] = new(sql.NullInt64)
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
		case organization.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				o.Description = value.String
			}
		case organization.FieldExternalLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field external_links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.ExternalLinks); err != nil {
					return fmt.Errorf("unmarshal field external_links: %w", err)
				}
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
		case organization.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field holder_organization", value)
			} else if value.Valid {
				o.holder_organization = new(int)
				*o.holder_organization = int(value.Int64)
			}
		case organization.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field organization_type_organizations", value)
			} else if value.Valid {
				o.organization_type_organizations = new(int)
				*o.organization_type_organizations = int(value.Int64)
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

// QueryPeople queries the "people" edge of the Organization entity.
func (o *Organization) QueryPeople() *PersonQuery {
	return NewOrganizationClient(o.config).QueryPeople(o)
}

// QueryHolder queries the "holder" edge of the Organization entity.
func (o *Organization) QueryHolder() *HolderQuery {
	return NewOrganizationClient(o.config).QueryHolder(o)
}

// QueryOrganizationType queries the "organization_type" edge of the Organization entity.
func (o *Organization) QueryOrganizationType() *OrganizationTypeQuery {
	return NewOrganizationClient(o.config).QueryOrganizationType(o)
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
	builder.WriteString("description=")
	builder.WriteString(o.Description)
	builder.WriteString(", ")
	builder.WriteString("external_links=")
	builder.WriteString(fmt.Sprintf("%v", o.ExternalLinks))
	builder.WriteString(", ")
	builder.WriteString("primary_image_url=")
	builder.WriteString(o.PrimaryImageURL)
	builder.WriteString(", ")
	builder.WriteString("additional_images_urls=")
	builder.WriteString(fmt.Sprintf("%v", o.AdditionalImagesUrls))
	builder.WriteByte(')')
	return builder.String()
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
