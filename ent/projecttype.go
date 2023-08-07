// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/projecttype"
)

// ProjectType is the model entity for the ProjectType schema.
type ProjectType struct {
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
	// The values are being populated by the ProjectTypeQuery when eager-loading is set.
	Edges        ProjectTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProjectTypeEdges holds the relations/edges for other nodes in the graph.
type ProjectTypeEdges struct {
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedProjects map[string][]*Project
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectTypeEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProjectType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case projecttype.FieldID:
			values[i] = new(sql.NullInt64)
		case projecttype.FieldCreatedBy, projecttype.FieldUpdatedBy, projecttype.FieldDisplayName, projecttype.FieldAbbreviation, projecttype.FieldDescription, projecttype.FieldExternalLink:
			values[i] = new(sql.NullString)
		case projecttype.FieldCreatedAt, projecttype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProjectType fields.
func (pt *ProjectType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case projecttype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pt.ID = int(value.Int64)
		case projecttype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pt.CreatedAt = value.Time
			}
		case projecttype.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pt.CreatedBy = value.String
			}
		case projecttype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pt.UpdatedAt = value.Time
			}
		case projecttype.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pt.UpdatedBy = value.String
			}
		case projecttype.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pt.DisplayName = value.String
			}
		case projecttype.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				pt.Abbreviation = value.String
			}
		case projecttype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pt.Description = value.String
			}
		case projecttype.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				pt.ExternalLink = value.String
			}
		default:
			pt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProjectType.
// This includes values selected through modifiers, order, etc.
func (pt *ProjectType) Value(name string) (ent.Value, error) {
	return pt.selectValues.Get(name)
}

// QueryProjects queries the "projects" edge of the ProjectType entity.
func (pt *ProjectType) QueryProjects() *ProjectQuery {
	return NewProjectTypeClient(pt.config).QueryProjects(pt)
}

// Update returns a builder for updating this ProjectType.
// Note that you need to call ProjectType.Unwrap() before calling this method if this ProjectType
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *ProjectType) Update() *ProjectTypeUpdateOne {
	return NewProjectTypeClient(pt.config).UpdateOne(pt)
}

// Unwrap unwraps the ProjectType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *ProjectType) Unwrap() *ProjectType {
	_tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProjectType is not a transactional entity")
	}
	pt.config.driver = _tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *ProjectType) String() string {
	var builder strings.Builder
	builder.WriteString("ProjectType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pt.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pt.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(pt.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(pt.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pt.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(pt.ExternalLink)
	builder.WriteByte(')')
	return builder.String()
}

// NamedProjects returns the Projects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pt *ProjectType) NamedProjects(name string) ([]*Project, error) {
	if pt.Edges.namedProjects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pt.Edges.namedProjects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pt *ProjectType) appendNamedProjects(name string, edges ...*Project) {
	if pt.Edges.namedProjects == nil {
		pt.Edges.namedProjects = make(map[string][]*Project)
	}
	if len(edges) == 0 {
		pt.Edges.namedProjects[name] = []*Project{}
	} else {
		pt.Edges.namedProjects[name] = append(pt.Edges.namedProjects[name], edges...)
	}
}

// ProjectTypes is a parsable slice of ProjectType.
type ProjectTypes []*ProjectType
