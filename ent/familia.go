// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/familia"
)

// Familia is the model entity for the Familia schema.
type Familia struct {
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
	// PrimaryImageURL holds the value of the "primary_image_url" field.
	PrimaryImageURL string `json:"primary_image_url,omitempty"`
	// AdditionalImagesUrls holds the value of the "additional_images_urls" field.
	AdditionalImagesUrls []string `json:"additional_images_urls,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FamiliaQuery when eager-loading is set.
	Edges        FamiliaEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FamiliaEdges holds the relations/edges for other nodes in the graph.
type FamiliaEdges struct {
	// Herbaria holds the value of the herbaria edge.
	Herbaria []*Herbarium `json:"herbaria,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedHerbaria map[string][]*Herbarium
}

// HerbariaOrErr returns the Herbaria value or an error if the edge
// was not loaded in eager-loading.
func (e FamiliaEdges) HerbariaOrErr() ([]*Herbarium, error) {
	if e.loadedTypes[0] {
		return e.Herbaria, nil
	}
	return nil, &NotLoadedError{edge: "herbaria"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Familia) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case familia.FieldAdditionalImagesUrls:
			values[i] = new([]byte)
		case familia.FieldID:
			values[i] = new(sql.NullInt64)
		case familia.FieldCreatedBy, familia.FieldUpdatedBy, familia.FieldDisplayName, familia.FieldAbbreviation, familia.FieldDescription, familia.FieldExternalLink, familia.FieldPrimaryImageURL:
			values[i] = new(sql.NullString)
		case familia.FieldCreatedAt, familia.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Familia fields.
func (f *Familia) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case familia.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case familia.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case familia.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				f.CreatedBy = value.String
			}
		case familia.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case familia.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				f.UpdatedBy = value.String
			}
		case familia.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				f.DisplayName = value.String
			}
		case familia.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				f.Abbreviation = value.String
			}
		case familia.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				f.Description = value.String
			}
		case familia.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				f.ExternalLink = value.String
			}
		case familia.FieldPrimaryImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_image_url", values[i])
			} else if value.Valid {
				f.PrimaryImageURL = value.String
			}
		case familia.FieldAdditionalImagesUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_images_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &f.AdditionalImagesUrls); err != nil {
					return fmt.Errorf("unmarshal field additional_images_urls: %w", err)
				}
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Familia.
// This includes values selected through modifiers, order, etc.
func (f *Familia) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryHerbaria queries the "herbaria" edge of the Familia entity.
func (f *Familia) QueryHerbaria() *HerbariumQuery {
	return NewFamiliaClient(f.config).QueryHerbaria(f)
}

// Update returns a builder for updating this Familia.
// Note that you need to call Familia.Unwrap() before calling this method if this Familia
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Familia) Update() *FamiliaUpdateOne {
	return NewFamiliaClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Familia entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Familia) Unwrap() *Familia {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Familia is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Familia) String() string {
	var builder strings.Builder
	builder.WriteString("Familia(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(f.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(f.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(f.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(f.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(f.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(f.ExternalLink)
	builder.WriteString(", ")
	builder.WriteString("primary_image_url=")
	builder.WriteString(f.PrimaryImageURL)
	builder.WriteString(", ")
	builder.WriteString("additional_images_urls=")
	builder.WriteString(fmt.Sprintf("%v", f.AdditionalImagesUrls))
	builder.WriteByte(')')
	return builder.String()
}

// NamedHerbaria returns the Herbaria named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *Familia) NamedHerbaria(name string) ([]*Herbarium, error) {
	if f.Edges.namedHerbaria == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedHerbaria[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *Familia) appendNamedHerbaria(name string, edges ...*Herbarium) {
	if f.Edges.namedHerbaria == nil {
		f.Edges.namedHerbaria = make(map[string][]*Herbarium)
	}
	if len(edges) == 0 {
		f.Edges.namedHerbaria[name] = []*Herbarium{}
	} else {
		f.Edges.namedHerbaria[name] = append(f.Edges.namedHerbaria[name], edges...)
	}
}

// FamiliaSlice is a parsable slice of Familia.
type FamiliaSlice []*Familia
