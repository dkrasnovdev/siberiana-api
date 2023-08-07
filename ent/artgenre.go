// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/artgenre"
)

// ArtGenre is the model entity for the ArtGenre schema.
type ArtGenre struct {
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
	// The values are being populated by the ArtGenreQuery when eager-loading is set.
	Edges        ArtGenreEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ArtGenreEdges holds the relations/edges for other nodes in the graph.
type ArtGenreEdges struct {
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
func (e ArtGenreEdges) ArtOrErr() ([]*Art, error) {
	if e.loadedTypes[0] {
		return e.Art, nil
	}
	return nil, &NotLoadedError{edge: "art"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ArtGenre) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case artgenre.FieldID:
			values[i] = new(sql.NullInt64)
		case artgenre.FieldCreatedBy, artgenre.FieldUpdatedBy, artgenre.FieldDisplayName, artgenre.FieldAbbreviation, artgenre.FieldDescription, artgenre.FieldExternalLink:
			values[i] = new(sql.NullString)
		case artgenre.FieldCreatedAt, artgenre.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ArtGenre fields.
func (ag *ArtGenre) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artgenre.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ag.ID = int(value.Int64)
		case artgenre.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ag.CreatedAt = value.Time
			}
		case artgenre.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ag.CreatedBy = value.String
			}
		case artgenre.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ag.UpdatedAt = value.Time
			}
		case artgenre.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ag.UpdatedBy = value.String
			}
		case artgenre.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				ag.DisplayName = value.String
			}
		case artgenre.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				ag.Abbreviation = value.String
			}
		case artgenre.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ag.Description = value.String
			}
		case artgenre.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				ag.ExternalLink = value.String
			}
		default:
			ag.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ArtGenre.
// This includes values selected through modifiers, order, etc.
func (ag *ArtGenre) Value(name string) (ent.Value, error) {
	return ag.selectValues.Get(name)
}

// QueryArt queries the "art" edge of the ArtGenre entity.
func (ag *ArtGenre) QueryArt() *ArtQuery {
	return NewArtGenreClient(ag.config).QueryArt(ag)
}

// Update returns a builder for updating this ArtGenre.
// Note that you need to call ArtGenre.Unwrap() before calling this method if this ArtGenre
// was returned from a transaction, and the transaction was committed or rolled back.
func (ag *ArtGenre) Update() *ArtGenreUpdateOne {
	return NewArtGenreClient(ag.config).UpdateOne(ag)
}

// Unwrap unwraps the ArtGenre entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ag *ArtGenre) Unwrap() *ArtGenre {
	_tx, ok := ag.config.driver.(*txDriver)
	if !ok {
		panic("ent: ArtGenre is not a transactional entity")
	}
	ag.config.driver = _tx.drv
	return ag
}

// String implements the fmt.Stringer.
func (ag *ArtGenre) String() string {
	var builder strings.Builder
	builder.WriteString("ArtGenre(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ag.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ag.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(ag.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ag.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(ag.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(ag.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(ag.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ag.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(ag.ExternalLink)
	builder.WriteByte(')')
	return builder.String()
}

// NamedArt returns the Art named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ag *ArtGenre) NamedArt(name string) ([]*Art, error) {
	if ag.Edges.namedArt == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ag.Edges.namedArt[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ag *ArtGenre) appendNamedArt(name string, edges ...*Art) {
	if ag.Edges.namedArt == nil {
		ag.Edges.namedArt = make(map[string][]*Art)
	}
	if len(edges) == 0 {
		ag.Edges.namedArt[name] = []*Art{}
	} else {
		ag.Edges.namedArt[name] = append(ag.Edges.namedArt[name], edges...)
	}
}

// ArtGenres is a parsable slice of ArtGenre.
type ArtGenres []*ArtGenre
