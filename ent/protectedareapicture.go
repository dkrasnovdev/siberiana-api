// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/protectedareapicture"
)

// ProtectedAreaPicture is the model entity for the ProtectedAreaPicture schema.
type ProtectedAreaPicture struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProtectedAreaPicture) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case protectedareapicture.FieldID:
			values[i] = new(sql.NullInt64)
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
	builder.WriteString(fmt.Sprintf("id=%v", pap.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ProtectedAreaPictures is a parsable slice of ProtectedAreaPicture.
type ProtectedAreaPictures []*ProtectedAreaPicture