// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/bookgenre"
)

// BookGenre is the model entity for the BookGenre schema.
type BookGenre struct {
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
	// The values are being populated by the BookGenreQuery when eager-loading is set.
	Edges        BookGenreEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BookGenreEdges holds the relations/edges for other nodes in the graph.
type BookGenreEdges struct {
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedBooks map[string][]*Book
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e BookGenreEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[0] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BookGenre) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookgenre.FieldID:
			values[i] = new(sql.NullInt64)
		case bookgenre.FieldCreatedBy, bookgenre.FieldUpdatedBy, bookgenre.FieldDisplayName, bookgenre.FieldAbbreviation, bookgenre.FieldDescription, bookgenre.FieldExternalLink, bookgenre.FieldSlug:
			values[i] = new(sql.NullString)
		case bookgenre.FieldCreatedAt, bookgenre.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BookGenre fields.
func (bg *BookGenre) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookgenre.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bg.ID = int(value.Int64)
		case bookgenre.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bg.CreatedAt = value.Time
			}
		case bookgenre.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				bg.CreatedBy = value.String
			}
		case bookgenre.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bg.UpdatedAt = value.Time
			}
		case bookgenre.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				bg.UpdatedBy = value.String
			}
		case bookgenre.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				bg.DisplayName = value.String
			}
		case bookgenre.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				bg.Abbreviation = value.String
			}
		case bookgenre.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				bg.Description = value.String
			}
		case bookgenre.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				bg.ExternalLink = value.String
			}
		case bookgenre.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				bg.Slug = value.String
			}
		default:
			bg.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BookGenre.
// This includes values selected through modifiers, order, etc.
func (bg *BookGenre) Value(name string) (ent.Value, error) {
	return bg.selectValues.Get(name)
}

// QueryBooks queries the "books" edge of the BookGenre entity.
func (bg *BookGenre) QueryBooks() *BookQuery {
	return NewBookGenreClient(bg.config).QueryBooks(bg)
}

// Update returns a builder for updating this BookGenre.
// Note that you need to call BookGenre.Unwrap() before calling this method if this BookGenre
// was returned from a transaction, and the transaction was committed or rolled back.
func (bg *BookGenre) Update() *BookGenreUpdateOne {
	return NewBookGenreClient(bg.config).UpdateOne(bg)
}

// Unwrap unwraps the BookGenre entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bg *BookGenre) Unwrap() *BookGenre {
	_tx, ok := bg.config.driver.(*txDriver)
	if !ok {
		panic("ent: BookGenre is not a transactional entity")
	}
	bg.config.driver = _tx.drv
	return bg
}

// String implements the fmt.Stringer.
func (bg *BookGenre) String() string {
	var builder strings.Builder
	builder.WriteString("BookGenre(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bg.ID))
	builder.WriteString("created_at=")
	builder.WriteString(bg.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(bg.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bg.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(bg.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(bg.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(bg.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(bg.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(bg.ExternalLink)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(bg.Slug)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (bg *BookGenre) NamedBooks(name string) ([]*Book, error) {
	if bg.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := bg.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (bg *BookGenre) appendNamedBooks(name string, edges ...*Book) {
	if bg.Edges.namedBooks == nil {
		bg.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		bg.Edges.namedBooks[name] = []*Book{}
	} else {
		bg.Edges.namedBooks[name] = append(bg.Edges.namedBooks[name], edges...)
	}
}

// BookGenres is a parsable slice of BookGenre.
type BookGenres []*BookGenre
