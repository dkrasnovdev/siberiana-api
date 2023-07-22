// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/library"
)

// Library is the model entity for the Library schema.
type Library struct {
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
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Library) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case library.FieldExternalLinks:
			values[i] = new([]byte)
		case library.FieldID:
			values[i] = new(sql.NullInt64)
		case library.FieldCreatedBy, library.FieldUpdatedBy, library.FieldDisplayName, library.FieldDescription:
			values[i] = new(sql.NullString)
		case library.FieldCreatedAt, library.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Library fields.
func (l *Library) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case library.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case library.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case library.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				l.CreatedBy = value.String
			}
		case library.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case library.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				l.UpdatedBy = value.String
			}
		case library.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				l.DisplayName = value.String
			}
		case library.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				l.Description = value.String
			}
		case library.FieldExternalLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field external_links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &l.ExternalLinks); err != nil {
					return fmt.Errorf("unmarshal field external_links: %w", err)
				}
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Library.
// This includes values selected through modifiers, order, etc.
func (l *Library) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// Update returns a builder for updating this Library.
// Note that you need to call Library.Unwrap() before calling this method if this Library
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Library) Update() *LibraryUpdateOne {
	return NewLibraryClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Library entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Library) Unwrap() *Library {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Library is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Library) String() string {
	var builder strings.Builder
	builder.WriteString("Library(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(l.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(l.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(l.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(l.Description)
	builder.WriteString(", ")
	builder.WriteString("external_links=")
	builder.WriteString(fmt.Sprintf("%v", l.ExternalLinks))
	builder.WriteByte(')')
	return builder.String()
}

// Libraries is a parsable slice of Library.
type Libraries []*Library
