// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/district"
)

// District is the model entity for the District schema.
type District struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*District) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case district.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the District fields.
func (d *District) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case district.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the District.
// This includes values selected through modifiers, order, etc.
func (d *District) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// Update returns a builder for updating this District.
// Note that you need to call District.Unwrap() before calling this method if this District
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *District) Update() *DistrictUpdateOne {
	return NewDistrictClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the District entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *District) Unwrap() *District {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: District is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *District) String() string {
	var builder strings.Builder
	builder.WriteString("District(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Districts is a parsable slice of District.
type Districts []*District