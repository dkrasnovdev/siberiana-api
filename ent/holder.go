// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/holder"
	"github.com/dkrasnovdev/heritage-api/ent/organization"
	"github.com/dkrasnovdev/heritage-api/ent/person"
)

// Holder is the model entity for the Holder schema.
type Holder struct {
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
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ExternalLink holds the value of the "external_link" field.
	ExternalLink string `json:"external_link,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HolderQuery when eager-loading is set.
	Edges        HolderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HolderEdges holds the relations/edges for other nodes in the graph.
type HolderEdges struct {
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Person holds the value of the person edge.
	Person *Person `json:"person,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedArtifacts map[string][]*Artifact
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e HolderEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[0] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// PersonOrErr returns the Person value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HolderEdges) PersonOrErr() (*Person, error) {
	if e.loadedTypes[1] {
		if e.Person == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: person.Label}
		}
		return e.Person, nil
	}
	return nil, &NotLoadedError{edge: "person"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HolderEdges) OrganizationOrErr() (*Organization, error) {
	if e.loadedTypes[2] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Holder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case holder.FieldID:
			values[i] = new(sql.NullInt64)
		case holder.FieldCreatedBy, holder.FieldUpdatedBy, holder.FieldDisplayName, holder.FieldDescription, holder.FieldExternalLink:
			values[i] = new(sql.NullString)
		case holder.FieldCreatedAt, holder.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Holder fields.
func (h *Holder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case holder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case holder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				h.CreatedAt = value.Time
			}
		case holder.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				h.CreatedBy = value.String
			}
		case holder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				h.UpdatedAt = value.Time
			}
		case holder.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				h.UpdatedBy = value.String
			}
		case holder.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				h.DisplayName = value.String
			}
		case holder.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				h.Description = value.String
			}
		case holder.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				h.ExternalLink = value.String
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Holder.
// This includes values selected through modifiers, order, etc.
func (h *Holder) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryArtifacts queries the "artifacts" edge of the Holder entity.
func (h *Holder) QueryArtifacts() *ArtifactQuery {
	return NewHolderClient(h.config).QueryArtifacts(h)
}

// QueryPerson queries the "person" edge of the Holder entity.
func (h *Holder) QueryPerson() *PersonQuery {
	return NewHolderClient(h.config).QueryPerson(h)
}

// QueryOrganization queries the "organization" edge of the Holder entity.
func (h *Holder) QueryOrganization() *OrganizationQuery {
	return NewHolderClient(h.config).QueryOrganization(h)
}

// Update returns a builder for updating this Holder.
// Note that you need to call Holder.Unwrap() before calling this method if this Holder
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Holder) Update() *HolderUpdateOne {
	return NewHolderClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Holder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Holder) Unwrap() *Holder {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Holder is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Holder) String() string {
	var builder strings.Builder
	builder.WriteString("Holder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("created_at=")
	builder.WriteString(h.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(h.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(h.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(h.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(h.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(h.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(h.ExternalLink)
	builder.WriteByte(')')
	return builder.String()
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (h *Holder) NamedArtifacts(name string) ([]*Artifact, error) {
	if h.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := h.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (h *Holder) appendNamedArtifacts(name string, edges ...*Artifact) {
	if h.Edges.namedArtifacts == nil {
		h.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		h.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		h.Edges.namedArtifacts[name] = append(h.Edges.namedArtifacts[name], edges...)
	}
}

// Holders is a parsable slice of Holder.
type Holders []*Holder
