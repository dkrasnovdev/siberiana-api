// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/protectedareapicture"
)

// ProtectedAreaPicture is the model entity for the ProtectedAreaPicture schema.
type ProtectedAreaPicture struct {
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
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProtectedAreaPicture) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case protectedareapicture.FieldExternalLinks:
			values[i] = new([]byte)
		case protectedareapicture.FieldID:
			values[i] = new(sql.NullInt64)
		case protectedareapicture.FieldCreatedBy, protectedareapicture.FieldUpdatedBy, protectedareapicture.FieldDisplayName, protectedareapicture.FieldAbbreviation, protectedareapicture.FieldDescription:
			values[i] = new(sql.NullString)
		case protectedareapicture.FieldCreatedAt, protectedareapicture.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProtectedAreaPicture fields.
func (pap *ProtectedAreaPicture) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case protectedareapicture.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pap.ID = int(value.Int64)
		case protectedareapicture.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pap.CreatedAt = value.Time
			}
		case protectedareapicture.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pap.CreatedBy = value.String
			}
		case protectedareapicture.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pap.UpdatedAt = value.Time
			}
		case protectedareapicture.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pap.UpdatedBy = value.String
			}
		case protectedareapicture.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pap.DisplayName = value.String
			}
		case protectedareapicture.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				pap.Abbreviation = value.String
			}
		case protectedareapicture.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pap.Description = value.String
			}
		case protectedareapicture.FieldExternalLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field external_links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pap.ExternalLinks); err != nil {
					return fmt.Errorf("unmarshal field external_links: %w", err)
				}
			}
		default:
			pap.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProtectedAreaPicture.
// This includes values selected through modifiers, order, etc.
func (pap *ProtectedAreaPicture) Value(name string) (ent.Value, error) {
	return pap.selectValues.Get(name)
}

// Update returns a builder for updating this ProtectedAreaPicture.
// Note that you need to call ProtectedAreaPicture.Unwrap() before calling this method if this ProtectedAreaPicture
// was returned from a transaction, and the transaction was committed or rolled back.
func (pap *ProtectedAreaPicture) Update() *ProtectedAreaPictureUpdateOne {
	return NewProtectedAreaPictureClient(pap.config).UpdateOne(pap)
}

// Unwrap unwraps the ProtectedAreaPicture entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pap *ProtectedAreaPicture) Unwrap() *ProtectedAreaPicture {
	_tx, ok := pap.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProtectedAreaPicture is not a transactional entity")
	}
	pap.config.driver = _tx.drv
	return pap
}

// String implements the fmt.Stringer.
func (pap *ProtectedAreaPicture) String() string {
	var builder strings.Builder
	builder.WriteString("ProtectedAreaPicture(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pap.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pap.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pap.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pap.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pap.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(pap.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(pap.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pap.Description)
	builder.WriteString(", ")
	builder.WriteString("external_links=")
	builder.WriteString(fmt.Sprintf("%v", pap.ExternalLinks))
	builder.WriteByte(')')
	return builder.String()
}

// ProtectedAreaPictures is a parsable slice of ProtectedAreaPicture.
type ProtectedAreaPictures []*ProtectedAreaPicture
