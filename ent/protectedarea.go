// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/heritage-api/ent/protectedarea"
	"github.com/dkrasnovdev/heritage-api/ent/protectedareacategory"
)

// ProtectedArea is the model entity for the ProtectedArea schema.
type ProtectedArea struct {
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
	// Area holds the value of the "area" field.
	Area string `json:"area,omitempty"`
	// EstablishmentDate holds the value of the "establishment_date" field.
	EstablishmentDate time.Time `json:"establishment_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProtectedAreaQuery when eager-loading is set.
	Edges                                   ProtectedAreaEdges `json:"edges"`
	protected_area_category_protected_areas *int
	selectValues                            sql.SelectValues
}

// ProtectedAreaEdges holds the relations/edges for other nodes in the graph.
type ProtectedAreaEdges struct {
	// ProtectedAreaPictures holds the value of the protected_area_pictures edge.
	ProtectedAreaPictures []*ProtectedAreaPicture `json:"protected_area_pictures,omitempty"`
	// ProtectedAreaCategory holds the value of the protected_area_category edge.
	ProtectedAreaCategory *ProtectedAreaCategory `json:"protected_area_category,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedProtectedAreaPictures map[string][]*ProtectedAreaPicture
}

// ProtectedAreaPicturesOrErr returns the ProtectedAreaPictures value or an error if the edge
// was not loaded in eager-loading.
func (e ProtectedAreaEdges) ProtectedAreaPicturesOrErr() ([]*ProtectedAreaPicture, error) {
	if e.loadedTypes[0] {
		return e.ProtectedAreaPictures, nil
	}
	return nil, &NotLoadedError{edge: "protected_area_pictures"}
}

// ProtectedAreaCategoryOrErr returns the ProtectedAreaCategory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProtectedAreaEdges) ProtectedAreaCategoryOrErr() (*ProtectedAreaCategory, error) {
	if e.loadedTypes[1] {
		if e.ProtectedAreaCategory == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: protectedareacategory.Label}
		}
		return e.ProtectedAreaCategory, nil
	}
	return nil, &NotLoadedError{edge: "protected_area_category"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProtectedArea) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case protectedarea.FieldID:
			values[i] = new(sql.NullInt64)
		case protectedarea.FieldCreatedBy, protectedarea.FieldUpdatedBy, protectedarea.FieldDisplayName, protectedarea.FieldAbbreviation, protectedarea.FieldDescription, protectedarea.FieldExternalLink, protectedarea.FieldArea:
			values[i] = new(sql.NullString)
		case protectedarea.FieldCreatedAt, protectedarea.FieldUpdatedAt, protectedarea.FieldEstablishmentDate:
			values[i] = new(sql.NullTime)
		case protectedarea.ForeignKeys[0]: // protected_area_category_protected_areas
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProtectedArea fields.
func (pa *ProtectedArea) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case protectedarea.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case protectedarea.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		case protectedarea.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pa.CreatedBy = value.String
			}
		case protectedarea.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Time
			}
		case protectedarea.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pa.UpdatedBy = value.String
			}
		case protectedarea.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pa.DisplayName = value.String
			}
		case protectedarea.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				pa.Abbreviation = value.String
			}
		case protectedarea.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pa.Description = value.String
			}
		case protectedarea.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				pa.ExternalLink = value.String
			}
		case protectedarea.FieldArea:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field area", values[i])
			} else if value.Valid {
				pa.Area = value.String
			}
		case protectedarea.FieldEstablishmentDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field establishment_date", values[i])
			} else if value.Valid {
				pa.EstablishmentDate = value.Time
			}
		case protectedarea.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field protected_area_category_protected_areas", value)
			} else if value.Valid {
				pa.protected_area_category_protected_areas = new(int)
				*pa.protected_area_category_protected_areas = int(value.Int64)
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProtectedArea.
// This includes values selected through modifiers, order, etc.
func (pa *ProtectedArea) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryProtectedAreaPictures queries the "protected_area_pictures" edge of the ProtectedArea entity.
func (pa *ProtectedArea) QueryProtectedAreaPictures() *ProtectedAreaPictureQuery {
	return NewProtectedAreaClient(pa.config).QueryProtectedAreaPictures(pa)
}

// QueryProtectedAreaCategory queries the "protected_area_category" edge of the ProtectedArea entity.
func (pa *ProtectedArea) QueryProtectedAreaCategory() *ProtectedAreaCategoryQuery {
	return NewProtectedAreaClient(pa.config).QueryProtectedAreaCategory(pa)
}

// Update returns a builder for updating this ProtectedArea.
// Note that you need to call ProtectedArea.Unwrap() before calling this method if this ProtectedArea
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *ProtectedArea) Update() *ProtectedAreaUpdateOne {
	return NewProtectedAreaClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the ProtectedArea entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *ProtectedArea) Unwrap() *ProtectedArea {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProtectedArea is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *ProtectedArea) String() string {
	var builder strings.Builder
	builder.WriteString("ProtectedArea(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pa.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pa.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(pa.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(pa.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pa.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(pa.ExternalLink)
	builder.WriteString(", ")
	builder.WriteString("area=")
	builder.WriteString(pa.Area)
	builder.WriteString(", ")
	builder.WriteString("establishment_date=")
	builder.WriteString(pa.EstablishmentDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedProtectedAreaPictures returns the ProtectedAreaPictures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pa *ProtectedArea) NamedProtectedAreaPictures(name string) ([]*ProtectedAreaPicture, error) {
	if pa.Edges.namedProtectedAreaPictures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pa.Edges.namedProtectedAreaPictures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pa *ProtectedArea) appendNamedProtectedAreaPictures(name string, edges ...*ProtectedAreaPicture) {
	if pa.Edges.namedProtectedAreaPictures == nil {
		pa.Edges.namedProtectedAreaPictures = make(map[string][]*ProtectedAreaPicture)
	}
	if len(edges) == 0 {
		pa.Edges.namedProtectedAreaPictures[name] = []*ProtectedAreaPicture{}
	} else {
		pa.Edges.namedProtectedAreaPictures[name] = append(pa.Edges.namedProtectedAreaPictures[name], edges...)
	}
}

// ProtectedAreas is a parsable slice of ProtectedArea.
type ProtectedAreas []*ProtectedArea
