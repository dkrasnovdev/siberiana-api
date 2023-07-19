// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/technique"
)

// Technique is the model entity for the Technique schema.
type Technique struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TechniqueQuery when eager-loading is set.
	Edges        TechniqueEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TechniqueEdges holds the relations/edges for other nodes in the graph.
type TechniqueEdges struct {
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
func (e TechniqueEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[0] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Technique) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case technique.FieldID:
			values[i] = new(sql.NullInt64)
		case technique.FieldCreatedBy, technique.FieldUpdatedBy, technique.FieldDisplayName, technique.FieldDescription:
			values[i] = new(sql.NullString)
		case technique.FieldCreatedAt, technique.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Technique fields.
func (t *Technique) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case technique.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case technique.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case technique.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				t.CreatedBy = value.String
			}
		case technique.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case technique.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				t.UpdatedBy = value.String
			}
		case technique.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				t.DisplayName = value.String
			}
		case technique.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Technique.
// This includes values selected through modifiers, order, etc.
func (t *Technique) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryArtifacts queries the "artifacts" edge of the Technique entity.
func (t *Technique) QueryArtifacts() *ArtifactQuery {
	return NewTechniqueClient(t.config).QueryArtifacts(t)
}

// Update returns a builder for updating this Technique.
// Note that you need to call Technique.Unwrap() before calling this method if this Technique
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Technique) Update() *TechniqueUpdateOne {
	return NewTechniqueClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Technique entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Technique) Unwrap() *Technique {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Technique is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Technique) String() string {
	var builder strings.Builder
	builder.WriteString("Technique(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(t.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(t.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(t.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteByte(')')
	return builder.String()
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Technique) NamedArtifacts(name string) ([]*Artifact, error) {
	if t.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Technique) appendNamedArtifacts(name string, edges ...*Artifact) {
	if t.Edges.namedArtifacts == nil {
		t.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		t.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		t.Edges.namedArtifacts[name] = append(t.Edges.namedArtifacts[name], edges...)
	}
}

// Techniques is a parsable slice of Technique.
type Techniques []*Technique
