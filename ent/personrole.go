// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/personrole"
)

// PersonRole is the model entity for the PersonRole schema.
type PersonRole struct {
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
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonRoleQuery when eager-loading is set.
	Edges        PersonRoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PersonRoleEdges holds the relations/edges for other nodes in the graph.
type PersonRoleEdges struct {
	// Person holds the value of the person edge.
	Person []*Person `json:"person,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedPerson map[string][]*Person
}

// PersonOrErr returns the Person value or an error if the edge
// was not loaded in eager-loading.
func (e PersonRoleEdges) PersonOrErr() ([]*Person, error) {
	if e.loadedTypes[0] {
		return e.Person, nil
	}
	return nil, &NotLoadedError{edge: "person"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PersonRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case personrole.FieldID:
			values[i] = new(sql.NullInt64)
		case personrole.FieldCreatedBy, personrole.FieldUpdatedBy, personrole.FieldDisplayName, personrole.FieldAbbreviation, personrole.FieldDescription, personrole.FieldExternalLink, personrole.FieldSlug:
			values[i] = new(sql.NullString)
		case personrole.FieldCreatedAt, personrole.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PersonRole fields.
func (pr *PersonRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case personrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case personrole.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case personrole.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pr.CreatedBy = value.String
			}
		case personrole.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case personrole.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pr.UpdatedBy = value.String
			}
		case personrole.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pr.DisplayName = value.String
			}
		case personrole.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				pr.Abbreviation = value.String
			}
		case personrole.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case personrole.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				pr.ExternalLink = value.String
			}
		case personrole.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				pr.Slug = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PersonRole.
// This includes values selected through modifiers, order, etc.
func (pr *PersonRole) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryPerson queries the "person" edge of the PersonRole entity.
func (pr *PersonRole) QueryPerson() *PersonQuery {
	return NewPersonRoleClient(pr.config).QueryPerson(pr)
}

// Update returns a builder for updating this PersonRole.
// Note that you need to call PersonRole.Unwrap() before calling this method if this PersonRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *PersonRole) Update() *PersonRoleUpdateOne {
	return NewPersonRoleClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the PersonRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *PersonRole) Unwrap() *PersonRole {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PersonRole is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *PersonRole) String() string {
	var builder strings.Builder
	builder.WriteString("PersonRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pr.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pr.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(pr.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(pr.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(pr.ExternalLink)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(pr.Slug)
	builder.WriteByte(')')
	return builder.String()
}

// NamedPerson returns the Person named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *PersonRole) NamedPerson(name string) ([]*Person, error) {
	if pr.Edges.namedPerson == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedPerson[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *PersonRole) appendNamedPerson(name string, edges ...*Person) {
	if pr.Edges.namedPerson == nil {
		pr.Edges.namedPerson = make(map[string][]*Person)
	}
	if len(edges) == 0 {
		pr.Edges.namedPerson[name] = []*Person{}
	} else {
		pr.Edges.namedPerson[name] = append(pr.Edges.namedPerson[name], edges...)
	}
}

// PersonRoles is a parsable slice of PersonRole.
type PersonRoles []*PersonRole
