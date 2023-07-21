// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/monument"
)

// Monument is the model entity for the Monument schema.
type Monument struct {
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
	// The values are being populated by the MonumentQuery when eager-loading is set.
	Edges        MonumentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MonumentEdges holds the relations/edges for other nodes in the graph.
type MonumentEdges struct {
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedArtifacts map[string][]*Artifact
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e MonumentEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[0] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Monument) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case monument.FieldID:
			values[i] = new(sql.NullInt64)
		case monument.FieldCreatedBy, monument.FieldUpdatedBy, monument.FieldDisplayName, monument.FieldDescription, monument.FieldExternalLink:
			values[i] = new(sql.NullString)
		case monument.FieldCreatedAt, monument.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Monument fields.
func (m *Monument) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case monument.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case monument.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case monument.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				m.CreatedBy = value.String
			}
		case monument.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case monument.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				m.UpdatedBy = value.String
			}
		case monument.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				m.DisplayName = value.String
			}
		case monument.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				m.Description = value.String
			}
		case monument.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				m.ExternalLink = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Monument.
// This includes values selected through modifiers, order, etc.
func (m *Monument) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryArtifacts queries the "artifacts" edge of the Monument entity.
func (m *Monument) QueryArtifacts() *ArtifactQuery {
	return NewMonumentClient(m.config).QueryArtifacts(m)
}

// Update returns a builder for updating this Monument.
// Note that you need to call Monument.Unwrap() before calling this method if this Monument
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Monument) Update() *MonumentUpdateOne {
	return NewMonumentClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Monument entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Monument) Unwrap() *Monument {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Monument is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Monument) String() string {
	var builder strings.Builder
	builder.WriteString("Monument(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(m.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(m.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(m.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(m.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(m.ExternalLink)
	builder.WriteByte(')')
	return builder.String()
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Monument) NamedArtifacts(name string) ([]*Artifact, error) {
	if m.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Monument) appendNamedArtifacts(name string, edges ...*Artifact) {
	if m.Edges.namedArtifacts == nil {
		m.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		m.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		m.Edges.namedArtifacts[name] = append(m.Edges.namedArtifacts[name], edges...)
	}
}

// Monuments is a parsable slice of Monument.
type Monuments []*Monument
