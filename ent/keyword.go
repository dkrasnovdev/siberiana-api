// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/keyword"
)

// Keyword is the model entity for the Keyword schema.
type Keyword struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Keyword) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case keyword.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Keyword fields.
func (k *Keyword) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keyword.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			k.ID = int(value.Int64)
		default:
			k.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Keyword.
// This includes values selected through modifiers, order, etc.
func (k *Keyword) Value(name string) (ent.Value, error) {
	return k.selectValues.Get(name)
}

// Update returns a builder for updating this Keyword.
// Note that you need to call Keyword.Unwrap() before calling this method if this Keyword
// was returned from a transaction, and the transaction was committed or rolled back.
func (k *Keyword) Update() *KeywordUpdateOne {
	return NewKeywordClient(k.config).UpdateOne(k)
}

// Unwrap unwraps the Keyword entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (k *Keyword) Unwrap() *Keyword {
	_tx, ok := k.config.driver.(*txDriver)
	if !ok {
		panic("ent: Keyword is not a transactional entity")
	}
	k.config.driver = _tx.drv
	return k
}

// String implements the fmt.Stringer.
func (k *Keyword) String() string {
	var builder strings.Builder
	builder.WriteString("Keyword(")
	builder.WriteString(fmt.Sprintf("id=%v", k.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Keywords is a parsable slice of Keyword.
type Keywords []*Keyword
