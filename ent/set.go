// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/set"
)

// Set is the model entity for the Set schema.
type Set struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Set) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case set.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Set fields.
func (s *Set) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case set.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Set.
// This includes values selected through modifiers, order, etc.
func (s *Set) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Set.
// Note that you need to call Set.Unwrap() before calling this method if this Set
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Set) Update() *SetUpdateOne {
	return NewSetClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Set entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Set) Unwrap() *Set {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Set is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Set) String() string {
	var builder strings.Builder
	builder.WriteString("Set(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Sets is a parsable slice of Set.
type Sets []*Set
