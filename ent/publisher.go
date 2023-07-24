// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/publisher"
)

// Publisher is the model entity for the Publisher schema.
type Publisher struct {
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
	// ExternalLinks holds the value of the "external_links" field.
	ExternalLinks []string `json:"external_links,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PublisherQuery when eager-loading is set.
	Edges        PublisherEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PublisherEdges holds the relations/edges for other nodes in the graph.
type PublisherEdges struct {
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedBooks map[string][]*Book
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e PublisherEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[0] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Publisher) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case publisher.FieldExternalLinks:
			values[i] = new([]byte)
		case publisher.FieldID:
			values[i] = new(sql.NullInt64)
		case publisher.FieldCreatedBy, publisher.FieldUpdatedBy, publisher.FieldDisplayName, publisher.FieldDescription:
			values[i] = new(sql.NullString)
		case publisher.FieldCreatedAt, publisher.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Publisher fields.
func (pu *Publisher) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case publisher.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pu.ID = int(value.Int64)
		case publisher.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pu.CreatedAt = value.Time
			}
		case publisher.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pu.CreatedBy = value.String
			}
		case publisher.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pu.UpdatedAt = value.Time
			}
		case publisher.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pu.UpdatedBy = value.String
			}
		case publisher.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pu.DisplayName = value.String
			}
		case publisher.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pu.Description = value.String
			}
		case publisher.FieldExternalLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field external_links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pu.ExternalLinks); err != nil {
					return fmt.Errorf("unmarshal field external_links: %w", err)
				}
			}
		default:
			pu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Publisher.
// This includes values selected through modifiers, order, etc.
func (pu *Publisher) Value(name string) (ent.Value, error) {
	return pu.selectValues.Get(name)
}

// QueryBooks queries the "books" edge of the Publisher entity.
func (pu *Publisher) QueryBooks() *BookQuery {
	return NewPublisherClient(pu.config).QueryBooks(pu)
}

// Update returns a builder for updating this Publisher.
// Note that you need to call Publisher.Unwrap() before calling this method if this Publisher
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *Publisher) Update() *PublisherUpdateOne {
	return NewPublisherClient(pu.config).UpdateOne(pu)
}

// Unwrap unwraps the Publisher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pu *Publisher) Unwrap() *Publisher {
	_tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Publisher is not a transactional entity")
	}
	pu.config.driver = _tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *Publisher) String() string {
	var builder strings.Builder
	builder.WriteString("Publisher(")
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
	builder.WriteString("description=")
	builder.WriteString(pu.Description)
	builder.WriteString(", ")
	builder.WriteString("external_links=")
	builder.WriteString(fmt.Sprintf("%v", pu.ExternalLinks))
	builder.WriteByte(')')
	return builder.String()
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pu *Publisher) NamedBooks(name string) ([]*Book, error) {
	if pu.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pu.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pu *Publisher) appendNamedBooks(name string, edges ...*Book) {
	if pu.Edges.namedBooks == nil {
		pu.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		pu.Edges.namedBooks[name] = []*Book{}
	} else {
		pu.Edges.namedBooks[name] = append(pu.Edges.namedBooks[name], edges...)
	}
}

// Publishers is a parsable slice of Publisher.
type Publishers []*Publisher
