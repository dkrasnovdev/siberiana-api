// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/holderresponsibility"
)

// HolderResponsibility is the model entity for the HolderResponsibility schema.
type HolderResponsibility struct {
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
	// The values are being populated by the HolderResponsibilityQuery when eager-loading is set.
	Edges        HolderResponsibilityEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HolderResponsibilityEdges holds the relations/edges for other nodes in the graph.
type HolderResponsibilityEdges struct {
	// Holder holds the value of the holder edge.
	Holder []*Holder `json:"holder,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedHolder map[string][]*Holder
}

// HolderOrErr returns the Holder value or an error if the edge
// was not loaded in eager-loading.
func (e HolderResponsibilityEdges) HolderOrErr() ([]*Holder, error) {
	if e.loadedTypes[0] {
		return e.Holder, nil
	}
	return nil, &NotLoadedError{edge: "holder"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HolderResponsibility) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case holderresponsibility.FieldID:
			values[i] = new(sql.NullInt64)
		case holderresponsibility.FieldCreatedBy, holderresponsibility.FieldUpdatedBy, holderresponsibility.FieldDisplayName, holderresponsibility.FieldAbbreviation, holderresponsibility.FieldDescription, holderresponsibility.FieldExternalLink:
			values[i] = new(sql.NullString)
		case holderresponsibility.FieldCreatedAt, holderresponsibility.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HolderResponsibility fields.
func (hr *HolderResponsibility) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case holderresponsibility.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			hr.ID = int(value.Int64)
		case holderresponsibility.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				hr.CreatedAt = value.Time
			}
		case holderresponsibility.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				hr.CreatedBy = value.String
			}
		case holderresponsibility.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				hr.UpdatedAt = value.Time
			}
		case holderresponsibility.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				hr.UpdatedBy = value.String
			}
		case holderresponsibility.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				hr.DisplayName = value.String
			}
		case holderresponsibility.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				hr.Abbreviation = value.String
			}
		case holderresponsibility.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				hr.Description = value.String
			}
		case holderresponsibility.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				hr.ExternalLink = value.String
			}
		default:
			hr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HolderResponsibility.
// This includes values selected through modifiers, order, etc.
func (hr *HolderResponsibility) Value(name string) (ent.Value, error) {
	return hr.selectValues.Get(name)
}

// QueryHolder queries the "holder" edge of the HolderResponsibility entity.
func (hr *HolderResponsibility) QueryHolder() *HolderQuery {
	return NewHolderResponsibilityClient(hr.config).QueryHolder(hr)
}

// Update returns a builder for updating this HolderResponsibility.
// Note that you need to call HolderResponsibility.Unwrap() before calling this method if this HolderResponsibility
// was returned from a transaction, and the transaction was committed or rolled back.
func (hr *HolderResponsibility) Update() *HolderResponsibilityUpdateOne {
	return NewHolderResponsibilityClient(hr.config).UpdateOne(hr)
}

// Unwrap unwraps the HolderResponsibility entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hr *HolderResponsibility) Unwrap() *HolderResponsibility {
	_tx, ok := hr.config.driver.(*txDriver)
	if !ok {
		panic("ent: HolderResponsibility is not a transactional entity")
	}
	hr.config.driver = _tx.drv
	return hr
}

// String implements the fmt.Stringer.
func (hr *HolderResponsibility) String() string {
	var builder strings.Builder
	builder.WriteString("HolderResponsibility(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(hr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(hr.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(hr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(hr.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(hr.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(hr.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(hr.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(hr.ExternalLink)
	builder.WriteByte(')')
	return builder.String()
}

// NamedHolder returns the Holder named value or an error if the edge was not
// loaded in eager-loading with this name.
func (hr *HolderResponsibility) NamedHolder(name string) ([]*Holder, error) {
	if hr.Edges.namedHolder == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := hr.Edges.namedHolder[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (hr *HolderResponsibility) appendNamedHolder(name string, edges ...*Holder) {
	if hr.Edges.namedHolder == nil {
		hr.Edges.namedHolder = make(map[string][]*Holder)
	}
	if len(edges) == 0 {
		hr.Edges.namedHolder[name] = []*Holder{}
	} else {
		hr.Edges.namedHolder[name] = append(hr.Edges.namedHolder[name], edges...)
	}
}

// HolderResponsibilities is a parsable slice of HolderResponsibility.
type HolderResponsibilities []*HolderResponsibility
