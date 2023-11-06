// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/country"
	"github.com/dkrasnovdev/siberiana-api/ent/district"
	"github.com/dkrasnovdev/siberiana-api/ent/location"
	"github.com/dkrasnovdev/siberiana-api/ent/region"
	"github.com/dkrasnovdev/siberiana-api/ent/settlement"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/types"
)

// Location is the model entity for the Location schema.
type Location struct {
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
	// Geometry holds the value of the "geometry" field.
	Geometry *types.Geometry `json:"geometry,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LocationQuery when eager-loading is set.
	Edges               LocationEdges `json:"edges"`
	location_country    *int
	location_district   *int
	location_settlement *int
	location_region     *int
	selectValues        sql.SelectValues
}

// LocationEdges holds the relations/edges for other nodes in the graph.
type LocationEdges struct {
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// ProtectedAreaPictures holds the value of the protected_area_pictures edge.
	ProtectedAreaPictures []*ProtectedAreaPicture `json:"protected_area_pictures,omitempty"`
	// PetroglyphsAccountingDocumentation holds the value of the petroglyphs_accounting_documentation edge.
	PetroglyphsAccountingDocumentation []*Petroglyph `json:"petroglyphs_accounting_documentation,omitempty"`
	// Country holds the value of the country edge.
	Country *Country `json:"country,omitempty"`
	// District holds the value of the district edge.
	District *District `json:"district,omitempty"`
	// Settlement holds the value of the settlement edge.
	Settlement *Settlement `json:"settlement,omitempty"`
	// Region holds the value of the region edge.
	Region *Region `json:"region,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
	// totalCount holds the count of the edges above.
	totalCount [8]map[string]int

	namedArtifacts                          map[string][]*Artifact
	namedBooks                              map[string][]*Book
	namedProtectedAreaPictures              map[string][]*ProtectedAreaPicture
	namedPetroglyphsAccountingDocumentation map[string][]*Petroglyph
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[0] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[1] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// ProtectedAreaPicturesOrErr returns the ProtectedAreaPictures value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) ProtectedAreaPicturesOrErr() ([]*ProtectedAreaPicture, error) {
	if e.loadedTypes[2] {
		return e.ProtectedAreaPictures, nil
	}
	return nil, &NotLoadedError{edge: "protected_area_pictures"}
}

// PetroglyphsAccountingDocumentationOrErr returns the PetroglyphsAccountingDocumentation value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) PetroglyphsAccountingDocumentationOrErr() ([]*Petroglyph, error) {
	if e.loadedTypes[3] {
		return e.PetroglyphsAccountingDocumentation, nil
	}
	return nil, &NotLoadedError{edge: "petroglyphs_accounting_documentation"}
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) CountryOrErr() (*Country, error) {
	if e.loadedTypes[4] {
		if e.Country == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: country.Label}
		}
		return e.Country, nil
	}
	return nil, &NotLoadedError{edge: "country"}
}

// DistrictOrErr returns the District value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) DistrictOrErr() (*District, error) {
	if e.loadedTypes[5] {
		if e.District == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: district.Label}
		}
		return e.District, nil
	}
	return nil, &NotLoadedError{edge: "district"}
}

// SettlementOrErr returns the Settlement value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) SettlementOrErr() (*Settlement, error) {
	if e.loadedTypes[6] {
		if e.Settlement == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: settlement.Label}
		}
		return e.Settlement, nil
	}
	return nil, &NotLoadedError{edge: "settlement"}
}

// RegionOrErr returns the Region value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) RegionOrErr() (*Region, error) {
	if e.loadedTypes[7] {
		if e.Region == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: region.Label}
		}
		return e.Region, nil
	}
	return nil, &NotLoadedError{edge: "region"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Location) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case location.FieldGeometry:
			values[i] = &sql.NullScanner{S: new(types.Geometry)}
		case location.FieldID:
			values[i] = new(sql.NullInt64)
		case location.FieldCreatedBy, location.FieldUpdatedBy, location.FieldDisplayName, location.FieldAbbreviation, location.FieldDescription, location.FieldExternalLink:
			values[i] = new(sql.NullString)
		case location.FieldCreatedAt, location.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case location.ForeignKeys[0]: // location_country
			values[i] = new(sql.NullInt64)
		case location.ForeignKeys[1]: // location_district
			values[i] = new(sql.NullInt64)
		case location.ForeignKeys[2]: // location_settlement
			values[i] = new(sql.NullInt64)
		case location.ForeignKeys[3]: // location_region
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Location fields.
func (l *Location) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case location.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case location.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case location.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				l.CreatedBy = value.String
			}
		case location.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case location.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				l.UpdatedBy = value.String
			}
		case location.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				l.DisplayName = value.String
			}
		case location.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				l.Abbreviation = value.String
			}
		case location.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				l.Description = value.String
			}
		case location.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				l.ExternalLink = value.String
			}
		case location.FieldGeometry:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field geometry", values[i])
			} else if value.Valid {
				l.Geometry = new(types.Geometry)
				*l.Geometry = *value.S.(*types.Geometry)
			}
		case location.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field location_country", value)
			} else if value.Valid {
				l.location_country = new(int)
				*l.location_country = int(value.Int64)
			}
		case location.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field location_district", value)
			} else if value.Valid {
				l.location_district = new(int)
				*l.location_district = int(value.Int64)
			}
		case location.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field location_settlement", value)
			} else if value.Valid {
				l.location_settlement = new(int)
				*l.location_settlement = int(value.Int64)
			}
		case location.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field location_region", value)
			} else if value.Valid {
				l.location_region = new(int)
				*l.location_region = int(value.Int64)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Location.
// This includes values selected through modifiers, order, etc.
func (l *Location) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryArtifacts queries the "artifacts" edge of the Location entity.
func (l *Location) QueryArtifacts() *ArtifactQuery {
	return NewLocationClient(l.config).QueryArtifacts(l)
}

// QueryBooks queries the "books" edge of the Location entity.
func (l *Location) QueryBooks() *BookQuery {
	return NewLocationClient(l.config).QueryBooks(l)
}

// QueryProtectedAreaPictures queries the "protected_area_pictures" edge of the Location entity.
func (l *Location) QueryProtectedAreaPictures() *ProtectedAreaPictureQuery {
	return NewLocationClient(l.config).QueryProtectedAreaPictures(l)
}

// QueryPetroglyphsAccountingDocumentation queries the "petroglyphs_accounting_documentation" edge of the Location entity.
func (l *Location) QueryPetroglyphsAccountingDocumentation() *PetroglyphQuery {
	return NewLocationClient(l.config).QueryPetroglyphsAccountingDocumentation(l)
}

// QueryCountry queries the "country" edge of the Location entity.
func (l *Location) QueryCountry() *CountryQuery {
	return NewLocationClient(l.config).QueryCountry(l)
}

// QueryDistrict queries the "district" edge of the Location entity.
func (l *Location) QueryDistrict() *DistrictQuery {
	return NewLocationClient(l.config).QueryDistrict(l)
}

// QuerySettlement queries the "settlement" edge of the Location entity.
func (l *Location) QuerySettlement() *SettlementQuery {
	return NewLocationClient(l.config).QuerySettlement(l)
}

// QueryRegion queries the "region" edge of the Location entity.
func (l *Location) QueryRegion() *RegionQuery {
	return NewLocationClient(l.config).QueryRegion(l)
}

// Update returns a builder for updating this Location.
// Note that you need to call Location.Unwrap() before calling this method if this Location
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Location) Update() *LocationUpdateOne {
	return NewLocationClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Location entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Location) Unwrap() *Location {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Location is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Location) String() string {
	var builder strings.Builder
	builder.WriteString("Location(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(l.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(l.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(l.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(l.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(l.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(l.ExternalLink)
	builder.WriteString(", ")
	if v := l.Geometry; v != nil {
		builder.WriteString("geometry=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Location) NamedArtifacts(name string) ([]*Artifact, error) {
	if l.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Location) appendNamedArtifacts(name string, edges ...*Artifact) {
	if l.Edges.namedArtifacts == nil {
		l.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		l.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		l.Edges.namedArtifacts[name] = append(l.Edges.namedArtifacts[name], edges...)
	}
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Location) NamedBooks(name string) ([]*Book, error) {
	if l.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Location) appendNamedBooks(name string, edges ...*Book) {
	if l.Edges.namedBooks == nil {
		l.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		l.Edges.namedBooks[name] = []*Book{}
	} else {
		l.Edges.namedBooks[name] = append(l.Edges.namedBooks[name], edges...)
	}
}

// NamedProtectedAreaPictures returns the ProtectedAreaPictures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Location) NamedProtectedAreaPictures(name string) ([]*ProtectedAreaPicture, error) {
	if l.Edges.namedProtectedAreaPictures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedProtectedAreaPictures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Location) appendNamedProtectedAreaPictures(name string, edges ...*ProtectedAreaPicture) {
	if l.Edges.namedProtectedAreaPictures == nil {
		l.Edges.namedProtectedAreaPictures = make(map[string][]*ProtectedAreaPicture)
	}
	if len(edges) == 0 {
		l.Edges.namedProtectedAreaPictures[name] = []*ProtectedAreaPicture{}
	} else {
		l.Edges.namedProtectedAreaPictures[name] = append(l.Edges.namedProtectedAreaPictures[name], edges...)
	}
}

// NamedPetroglyphsAccountingDocumentation returns the PetroglyphsAccountingDocumentation named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Location) NamedPetroglyphsAccountingDocumentation(name string) ([]*Petroglyph, error) {
	if l.Edges.namedPetroglyphsAccountingDocumentation == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedPetroglyphsAccountingDocumentation[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Location) appendNamedPetroglyphsAccountingDocumentation(name string, edges ...*Petroglyph) {
	if l.Edges.namedPetroglyphsAccountingDocumentation == nil {
		l.Edges.namedPetroglyphsAccountingDocumentation = make(map[string][]*Petroglyph)
	}
	if len(edges) == 0 {
		l.Edges.namedPetroglyphsAccountingDocumentation[name] = []*Petroglyph{}
	} else {
		l.Edges.namedPetroglyphsAccountingDocumentation[name] = append(l.Edges.namedPetroglyphsAccountingDocumentation[name], edges...)
	}
}

// Locations is a parsable slice of Location.
type Locations []*Location
