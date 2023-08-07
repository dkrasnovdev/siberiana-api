// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/publication"
)

// Publication is the model entity for the Publication schema.
type Publication struct {
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
	// Abbreviation holds the value of the "abbreviation" field.
	Abbreviation string `json:"abbreviation,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ExternalLink holds the value of the "external_link" field.
	ExternalLink string `json:"external_link,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PublicationQuery when eager-loading is set.
	Edges        PublicationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PublicationEdges holds the relations/edges for other nodes in the graph.
type PublicationEdges struct {
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Authors holds the value of the authors edge.
	Authors []*Person `json:"authors,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedArtifacts map[string][]*Artifact
	namedAuthors   map[string][]*Person
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e PublicationEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[0] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// AuthorsOrErr returns the Authors value or an error if the edge
// was not loaded in eager-loading.
func (e PublicationEdges) AuthorsOrErr() ([]*Person, error) {
	if e.loadedTypes[1] {
		return e.Authors, nil
	}
	return nil, &NotLoadedError{edge: "authors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Publication) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case publication.FieldID:
			values[i] = new(sql.NullInt64)
		case publication.FieldCreatedBy, publication.FieldUpdatedBy, publication.FieldDisplayName, publication.FieldAbbreviation, publication.FieldDescription, publication.FieldExternalLink:
			values[i] = new(sql.NullString)
		case publication.FieldCreatedAt, publication.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Publication fields.
func (pu *Publication) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case publication.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pu.ID = int(value.Int64)
		case publication.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pu.CreatedAt = value.Time
			}
		case publication.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pu.CreatedBy = value.String
			}
		case publication.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pu.UpdatedAt = value.Time
			}
		case publication.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pu.UpdatedBy = value.String
			}
		case publication.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pu.DisplayName = value.String
			}
		case publication.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				pu.Abbreviation = value.String
			}
		case publication.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pu.Description = value.String
			}
		case publication.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				pu.ExternalLink = value.String
			}
		default:
			pu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Publication.
// This includes values selected through modifiers, order, etc.
func (pu *Publication) Value(name string) (ent.Value, error) {
	return pu.selectValues.Get(name)
}

// QueryArtifacts queries the "artifacts" edge of the Publication entity.
func (pu *Publication) QueryArtifacts() *ArtifactQuery {
	return NewPublicationClient(pu.config).QueryArtifacts(pu)
}

// QueryAuthors queries the "authors" edge of the Publication entity.
func (pu *Publication) QueryAuthors() *PersonQuery {
	return NewPublicationClient(pu.config).QueryAuthors(pu)
}

// Update returns a builder for updating this Publication.
// Note that you need to call Publication.Unwrap() before calling this method if this Publication
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *Publication) Update() *PublicationUpdateOne {
	return NewPublicationClient(pu.config).UpdateOne(pu)
}

// Unwrap unwraps the Publication entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pu *Publication) Unwrap() *Publication {
	_tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Publication is not a transactional entity")
	}
	pu.config.driver = _tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *Publication) String() string {
	var builder strings.Builder
	builder.WriteString("Publication(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pu.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pu.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pu.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(pu.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(pu.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pu.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(pu.ExternalLink)
	builder.WriteByte(')')
	return builder.String()
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pu *Publication) NamedArtifacts(name string) ([]*Artifact, error) {
	if pu.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pu.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pu *Publication) appendNamedArtifacts(name string, edges ...*Artifact) {
	if pu.Edges.namedArtifacts == nil {
		pu.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		pu.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		pu.Edges.namedArtifacts[name] = append(pu.Edges.namedArtifacts[name], edges...)
	}
}

// NamedAuthors returns the Authors named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pu *Publication) NamedAuthors(name string) ([]*Person, error) {
	if pu.Edges.namedAuthors == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pu.Edges.namedAuthors[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pu *Publication) appendNamedAuthors(name string, edges ...*Person) {
	if pu.Edges.namedAuthors == nil {
		pu.Edges.namedAuthors = make(map[string][]*Person)
	}
	if len(edges) == 0 {
		pu.Edges.namedAuthors[name] = []*Person{}
	} else {
		pu.Edges.namedAuthors[name] = append(pu.Edges.namedAuthors[name], edges...)
	}
}

// Publications is a parsable slice of Publication.
type Publications []*Publication
