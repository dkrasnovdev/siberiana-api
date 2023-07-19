// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/category"
	"github.com/dkrasnovdev/heritage-api/ent/collection"
)

// Collection is the model entity for the Collection schema.
type Collection struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CollectionQuery when eager-loading is set.
	Edges                CollectionEdges `json:"edges"`
	category_collections *int
	selectValues         sql.SelectValues
}

// CollectionEdges holds the relations/edges for other nodes in the graph.
type CollectionEdges struct {
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedArtifacts map[string][]*Artifact
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[0] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CollectionEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[1] {
		if e.Category == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Collection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case collection.FieldID:
			values[i] = new(sql.NullInt64)
		case collection.FieldCreatedBy, collection.FieldUpdatedBy, collection.FieldDisplayName, collection.FieldDescription:
			values[i] = new(sql.NullString)
		case collection.FieldCreatedAt, collection.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case collection.ForeignKeys[0]: // category_collections
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Collection fields.
func (c *Collection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case collection.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case collection.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case collection.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = value.String
			}
		case collection.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case collection.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = value.String
			}
		case collection.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				c.DisplayName = value.String
			}
		case collection.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case collection.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_collections", value)
			} else if value.Valid {
				c.category_collections = new(int)
				*c.category_collections = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Collection.
// This includes values selected through modifiers, order, etc.
func (c *Collection) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryArtifacts queries the "artifacts" edge of the Collection entity.
func (c *Collection) QueryArtifacts() *ArtifactQuery {
	return NewCollectionClient(c.config).QueryArtifacts(c)
}

// QueryCategory queries the "category" edge of the Collection entity.
func (c *Collection) QueryCategory() *CategoryQuery {
	return NewCollectionClient(c.config).QueryCategory(c)
}

// Update returns a builder for updating this Collection.
// Note that you need to call Collection.Unwrap() before calling this method if this Collection
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Collection) Update() *CollectionUpdateOne {
	return NewCollectionClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Collection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Collection) Unwrap() *Collection {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Collection is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Collection) String() string {
	var builder strings.Builder
	builder.WriteString("Collection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(c.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(c.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(c.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteByte(')')
	return builder.String()
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Collection) NamedArtifacts(name string) ([]*Artifact, error) {
	if c.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Collection) appendNamedArtifacts(name string, edges ...*Artifact) {
	if c.Edges.namedArtifacts == nil {
		c.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		c.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		c.Edges.namedArtifacts[name] = append(c.Edges.namedArtifacts[name], edges...)
	}
}

// Collections is a parsable slice of Collection.
type Collections []*Collection
