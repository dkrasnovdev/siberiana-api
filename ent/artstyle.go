// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/artstyle"
)

// ArtStyle is the model entity for the ArtStyle schema.
type ArtStyle struct {
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
	// ExternalLinks holds the value of the "external_links" field.
	ExternalLinks []string `json:"external_links,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtStyleQuery when eager-loading is set.
	Edges        ArtStyleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ArtStyleEdges holds the relations/edges for other nodes in the graph.
type ArtStyleEdges struct {
	// Art holds the value of the art edge.
	Art []*Art `json:"art,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedArt map[string][]*Art
}

// ArtOrErr returns the Art value or an error if the edge
// was not loaded in eager-loading.
func (e ArtStyleEdges) ArtOrErr() ([]*Art, error) {
	if e.loadedTypes[0] {
		return e.Art, nil
	}
	return nil, &NotLoadedError{edge: "art"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ArtStyle) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case artstyle.FieldExternalLinks:
			values[i] = new([]byte)
		case artstyle.FieldID:
			values[i] = new(sql.NullInt64)
		case artstyle.FieldCreatedBy, artstyle.FieldUpdatedBy, artstyle.FieldDisplayName, artstyle.FieldAbbreviation, artstyle.FieldDescription:
			values[i] = new(sql.NullString)
		case artstyle.FieldCreatedAt, artstyle.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ArtStyle fields.
func (as *ArtStyle) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artstyle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = int(value.Int64)
		case artstyle.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				as.CreatedAt = value.Time
			}
		case artstyle.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				as.CreatedBy = value.String
			}
		case artstyle.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				as.UpdatedAt = value.Time
			}
		case artstyle.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				as.UpdatedBy = value.String
			}
		case artstyle.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				as.DisplayName = value.String
			}
		case artstyle.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				as.Abbreviation = value.String
			}
		case artstyle.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				as.Description = value.String
			}
		case artstyle.FieldExternalLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field external_links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &as.ExternalLinks); err != nil {
					return fmt.Errorf("unmarshal field external_links: %w", err)
				}
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ArtStyle.
// This includes values selected through modifiers, order, etc.
func (as *ArtStyle) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// QueryArt queries the "art" edge of the ArtStyle entity.
func (as *ArtStyle) QueryArt() *ArtQuery {
	return NewArtStyleClient(as.config).QueryArt(as)
}

// Update returns a builder for updating this ArtStyle.
// Note that you need to call ArtStyle.Unwrap() before calling this method if this ArtStyle
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *ArtStyle) Update() *ArtStyleUpdateOne {
	return NewArtStyleClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the ArtStyle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *ArtStyle) Unwrap() *ArtStyle {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: ArtStyle is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *ArtStyle) String() string {
	var builder strings.Builder
	builder.WriteString("ArtStyle(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("created_at=")
	builder.WriteString(as.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(as.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(as.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(as.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(as.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(as.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(as.Description)
	builder.WriteString(", ")
	builder.WriteString("external_links=")
	builder.WriteString(fmt.Sprintf("%v", as.ExternalLinks))
	builder.WriteByte(')')
	return builder.String()
}

// NamedArt returns the Art named value or an error if the edge was not
// loaded in eager-loading with this name.
func (as *ArtStyle) NamedArt(name string) ([]*Art, error) {
	if as.Edges.namedArt == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := as.Edges.namedArt[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (as *ArtStyle) appendNamedArt(name string, edges ...*Art) {
	if as.Edges.namedArt == nil {
		as.Edges.namedArt = make(map[string][]*Art)
	}
	if len(edges) == 0 {
		as.Edges.namedArt[name] = []*Art{}
	} else {
		as.Edges.namedArt[name] = append(as.Edges.namedArt[name], edges...)
	}
}

// ArtStyles is a parsable slice of ArtStyle.
type ArtStyles []*ArtStyle
