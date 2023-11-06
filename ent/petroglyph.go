// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/culture"
	"github.com/dkrasnovdev/siberiana-api/ent/location"
	"github.com/dkrasnovdev/siberiana-api/ent/model"
	"github.com/dkrasnovdev/siberiana-api/ent/mound"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/petroglyph"
	"github.com/dkrasnovdev/siberiana-api/ent/region"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/types"
)

// Petroglyph is the model entity for the Petroglyph schema.
type Petroglyph struct {
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
	// Status holds the value of the "status" field.
	Status petroglyph.Status `json:"status,omitempty"`
	// PrimaryImageURL holds the value of the "primary_image_url" field.
	PrimaryImageURL string `json:"primary_image_url,omitempty"`
	// AdditionalImagesUrls holds the value of the "additional_images_urls" field.
	AdditionalImagesUrls []string `json:"additional_images_urls,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Dating holds the value of the "dating" field.
	Dating string `json:"dating,omitempty"`
	// DatingStart holds the value of the "dating_start" field.
	DatingStart int `json:"dating_start,omitempty"`
	// DatingEnd holds the value of the "dating_end" field.
	DatingEnd int `json:"dating_end,omitempty"`
	// Orientation holds the value of the "orientation" field.
	Orientation string `json:"orientation,omitempty"`
	// Position holds the value of the "position" field.
	Position string `json:"position,omitempty"`
	// GeometricShape holds the value of the "geometric_shape" field.
	GeometricShape string `json:"geometric_shape,omitempty"`
	// Height holds the value of the "height" field.
	Height float64 `json:"height,omitempty"`
	// Width holds the value of the "width" field.
	Width float64 `json:"width,omitempty"`
	// Length holds the value of the "length" field.
	Length float64 `json:"length,omitempty"`
	// Depth holds the value of the "depth" field.
	Depth float64 `json:"depth,omitempty"`
	// Diameter holds the value of the "diameter" field.
	Diameter float64 `json:"diameter,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight string `json:"weight,omitempty"`
	// Dimensions holds the value of the "dimensions" field.
	Dimensions string `json:"dimensions,omitempty"`
	// PlanePreservation holds the value of the "plane_preservation" field.
	PlanePreservation string `json:"plane_preservation,omitempty"`
	// PhotoCode holds the value of the "photo_code" field.
	PhotoCode string `json:"photo_code,omitempty"`
	// AccountingDocumentationInformation holds the value of the "accounting_documentation_information" field.
	AccountingDocumentationInformation string `json:"accounting_documentation_information,omitempty"`
	// AccountingDocumentationDate holds the value of the "accounting_documentation_date" field.
	AccountingDocumentationDate time.Time `json:"accounting_documentation_date,omitempty"`
	// Geometry holds the value of the "geometry" field.
	Geometry *types.Geometry `json:"geometry,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetroglyphQuery when eager-loading is set.
	Edges                                         PetroglyphEdges `json:"edges"`
	culture_petroglyphs                           *int
	location_petroglyphs_accounting_documentation *int
	model_petroglyphs                             *int
	mound_petroglyphs                             *int
	person_petroglyphs_accounting_documentation   *int
	region_petroglyphs                            *int
	selectValues                                  sql.SelectValues
}

// PetroglyphEdges holds the relations/edges for other nodes in the graph.
type PetroglyphEdges struct {
	// CulturalAffiliation holds the value of the cultural_affiliation edge.
	CulturalAffiliation *Culture `json:"cultural_affiliation,omitempty"`
	// Model holds the value of the model edge.
	Model *Model `json:"model,omitempty"`
	// Mound holds the value of the mound edge.
	Mound *Mound `json:"mound,omitempty"`
	// Publications holds the value of the publications edge.
	Publications []*Publication `json:"publications,omitempty"`
	// Techniques holds the value of the techniques edge.
	Techniques []*Technique `json:"techniques,omitempty"`
	// Region holds the value of the region edge.
	Region *Region `json:"region,omitempty"`
	// AccountingDocumentationAddress holds the value of the accounting_documentation_address edge.
	AccountingDocumentationAddress *Location `json:"accounting_documentation_address,omitempty"`
	// AccountingDocumentationAuthor holds the value of the accounting_documentation_author edge.
	AccountingDocumentationAuthor *Person `json:"accounting_documentation_author,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
	// totalCount holds the count of the edges above.
	totalCount [8]map[string]int

	namedPublications map[string][]*Publication
	namedTechniques   map[string][]*Technique
}

// CulturalAffiliationOrErr returns the CulturalAffiliation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetroglyphEdges) CulturalAffiliationOrErr() (*Culture, error) {
	if e.loadedTypes[0] {
		if e.CulturalAffiliation == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: culture.Label}
		}
		return e.CulturalAffiliation, nil
	}
	return nil, &NotLoadedError{edge: "cultural_affiliation"}
}

// ModelOrErr returns the Model value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetroglyphEdges) ModelOrErr() (*Model, error) {
	if e.loadedTypes[1] {
		if e.Model == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: model.Label}
		}
		return e.Model, nil
	}
	return nil, &NotLoadedError{edge: "model"}
}

// MoundOrErr returns the Mound value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetroglyphEdges) MoundOrErr() (*Mound, error) {
	if e.loadedTypes[2] {
		if e.Mound == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mound.Label}
		}
		return e.Mound, nil
	}
	return nil, &NotLoadedError{edge: "mound"}
}

// PublicationsOrErr returns the Publications value or an error if the edge
// was not loaded in eager-loading.
func (e PetroglyphEdges) PublicationsOrErr() ([]*Publication, error) {
	if e.loadedTypes[3] {
		return e.Publications, nil
	}
	return nil, &NotLoadedError{edge: "publications"}
}

// TechniquesOrErr returns the Techniques value or an error if the edge
// was not loaded in eager-loading.
func (e PetroglyphEdges) TechniquesOrErr() ([]*Technique, error) {
	if e.loadedTypes[4] {
		return e.Techniques, nil
	}
	return nil, &NotLoadedError{edge: "techniques"}
}

// RegionOrErr returns the Region value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetroglyphEdges) RegionOrErr() (*Region, error) {
	if e.loadedTypes[5] {
		if e.Region == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: region.Label}
		}
		return e.Region, nil
	}
	return nil, &NotLoadedError{edge: "region"}
}

// AccountingDocumentationAddressOrErr returns the AccountingDocumentationAddress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetroglyphEdges) AccountingDocumentationAddressOrErr() (*Location, error) {
	if e.loadedTypes[6] {
		if e.AccountingDocumentationAddress == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: location.Label}
		}
		return e.AccountingDocumentationAddress, nil
	}
	return nil, &NotLoadedError{edge: "accounting_documentation_address"}
}

// AccountingDocumentationAuthorOrErr returns the AccountingDocumentationAuthor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetroglyphEdges) AccountingDocumentationAuthorOrErr() (*Person, error) {
	if e.loadedTypes[7] {
		if e.AccountingDocumentationAuthor == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: person.Label}
		}
		return e.AccountingDocumentationAuthor, nil
	}
	return nil, &NotLoadedError{edge: "accounting_documentation_author"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Petroglyph) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case petroglyph.FieldGeometry:
			values[i] = &sql.NullScanner{S: new(types.Geometry)}
		case petroglyph.FieldAdditionalImagesUrls:
			values[i] = new([]byte)
		case petroglyph.FieldHeight, petroglyph.FieldWidth, petroglyph.FieldLength, petroglyph.FieldDepth, petroglyph.FieldDiameter:
			values[i] = new(sql.NullFloat64)
		case petroglyph.FieldID, petroglyph.FieldDatingStart, petroglyph.FieldDatingEnd:
			values[i] = new(sql.NullInt64)
		case petroglyph.FieldCreatedBy, petroglyph.FieldUpdatedBy, petroglyph.FieldDisplayName, petroglyph.FieldAbbreviation, petroglyph.FieldDescription, petroglyph.FieldExternalLink, petroglyph.FieldStatus, petroglyph.FieldPrimaryImageURL, petroglyph.FieldDeletedBy, petroglyph.FieldNumber, petroglyph.FieldDating, petroglyph.FieldOrientation, petroglyph.FieldPosition, petroglyph.FieldGeometricShape, petroglyph.FieldWeight, petroglyph.FieldDimensions, petroglyph.FieldPlanePreservation, petroglyph.FieldPhotoCode, petroglyph.FieldAccountingDocumentationInformation:
			values[i] = new(sql.NullString)
		case petroglyph.FieldCreatedAt, petroglyph.FieldUpdatedAt, petroglyph.FieldDeletedAt, petroglyph.FieldAccountingDocumentationDate:
			values[i] = new(sql.NullTime)
		case petroglyph.ForeignKeys[0]: // culture_petroglyphs
			values[i] = new(sql.NullInt64)
		case petroglyph.ForeignKeys[1]: // location_petroglyphs_accounting_documentation
			values[i] = new(sql.NullInt64)
		case petroglyph.ForeignKeys[2]: // model_petroglyphs
			values[i] = new(sql.NullInt64)
		case petroglyph.ForeignKeys[3]: // mound_petroglyphs
			values[i] = new(sql.NullInt64)
		case petroglyph.ForeignKeys[4]: // person_petroglyphs_accounting_documentation
			values[i] = new(sql.NullInt64)
		case petroglyph.ForeignKeys[5]: // region_petroglyphs
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Petroglyph fields.
func (pe *Petroglyph) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case petroglyph.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case petroglyph.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Time
			}
		case petroglyph.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pe.CreatedBy = value.String
			}
		case petroglyph.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Time
			}
		case petroglyph.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pe.UpdatedBy = value.String
			}
		case petroglyph.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pe.DisplayName = value.String
			}
		case petroglyph.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				pe.Abbreviation = value.String
			}
		case petroglyph.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = value.String
			}
		case petroglyph.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				pe.ExternalLink = value.String
			}
		case petroglyph.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pe.Status = petroglyph.Status(value.String)
			}
		case petroglyph.FieldPrimaryImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_image_url", values[i])
			} else if value.Valid {
				pe.PrimaryImageURL = value.String
			}
		case petroglyph.FieldAdditionalImagesUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_images_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.AdditionalImagesUrls); err != nil {
					return fmt.Errorf("unmarshal field additional_images_urls: %w", err)
				}
			}
		case petroglyph.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pe.DeletedAt = value.Time
			}
		case petroglyph.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				pe.DeletedBy = value.String
			}
		case petroglyph.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				pe.Number = value.String
			}
		case petroglyph.FieldDating:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dating", values[i])
			} else if value.Valid {
				pe.Dating = value.String
			}
		case petroglyph.FieldDatingStart:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dating_start", values[i])
			} else if value.Valid {
				pe.DatingStart = int(value.Int64)
			}
		case petroglyph.FieldDatingEnd:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dating_end", values[i])
			} else if value.Valid {
				pe.DatingEnd = int(value.Int64)
			}
		case petroglyph.FieldOrientation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orientation", values[i])
			} else if value.Valid {
				pe.Orientation = value.String
			}
		case petroglyph.FieldPosition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field position", values[i])
			} else if value.Valid {
				pe.Position = value.String
			}
		case petroglyph.FieldGeometricShape:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field geometric_shape", values[i])
			} else if value.Valid {
				pe.GeometricShape = value.String
			}
		case petroglyph.FieldHeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				pe.Height = value.Float64
			}
		case petroglyph.FieldWidth:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				pe.Width = value.Float64
			}
		case petroglyph.FieldLength:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field length", values[i])
			} else if value.Valid {
				pe.Length = value.Float64
			}
		case petroglyph.FieldDepth:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field depth", values[i])
			} else if value.Valid {
				pe.Depth = value.Float64
			}
		case petroglyph.FieldDiameter:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field diameter", values[i])
			} else if value.Valid {
				pe.Diameter = value.Float64
			}
		case petroglyph.FieldWeight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				pe.Weight = value.String
			}
		case petroglyph.FieldDimensions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dimensions", values[i])
			} else if value.Valid {
				pe.Dimensions = value.String
			}
		case petroglyph.FieldPlanePreservation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plane_preservation", values[i])
			} else if value.Valid {
				pe.PlanePreservation = value.String
			}
		case petroglyph.FieldPhotoCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_code", values[i])
			} else if value.Valid {
				pe.PhotoCode = value.String
			}
		case petroglyph.FieldAccountingDocumentationInformation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field accounting_documentation_information", values[i])
			} else if value.Valid {
				pe.AccountingDocumentationInformation = value.String
			}
		case petroglyph.FieldAccountingDocumentationDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field accounting_documentation_date", values[i])
			} else if value.Valid {
				pe.AccountingDocumentationDate = value.Time
			}
		case petroglyph.FieldGeometry:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field geometry", values[i])
			} else if value.Valid {
				pe.Geometry = new(types.Geometry)
				*pe.Geometry = *value.S.(*types.Geometry)
			}
		case petroglyph.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field culture_petroglyphs", value)
			} else if value.Valid {
				pe.culture_petroglyphs = new(int)
				*pe.culture_petroglyphs = int(value.Int64)
			}
		case petroglyph.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field location_petroglyphs_accounting_documentation", value)
			} else if value.Valid {
				pe.location_petroglyphs_accounting_documentation = new(int)
				*pe.location_petroglyphs_accounting_documentation = int(value.Int64)
			}
		case petroglyph.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field model_petroglyphs", value)
			} else if value.Valid {
				pe.model_petroglyphs = new(int)
				*pe.model_petroglyphs = int(value.Int64)
			}
		case petroglyph.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field mound_petroglyphs", value)
			} else if value.Valid {
				pe.mound_petroglyphs = new(int)
				*pe.mound_petroglyphs = int(value.Int64)
			}
		case petroglyph.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field person_petroglyphs_accounting_documentation", value)
			} else if value.Valid {
				pe.person_petroglyphs_accounting_documentation = new(int)
				*pe.person_petroglyphs_accounting_documentation = int(value.Int64)
			}
		case petroglyph.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field region_petroglyphs", value)
			} else if value.Valid {
				pe.region_petroglyphs = new(int)
				*pe.region_petroglyphs = int(value.Int64)
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Petroglyph.
// This includes values selected through modifiers, order, etc.
func (pe *Petroglyph) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryCulturalAffiliation queries the "cultural_affiliation" edge of the Petroglyph entity.
func (pe *Petroglyph) QueryCulturalAffiliation() *CultureQuery {
	return NewPetroglyphClient(pe.config).QueryCulturalAffiliation(pe)
}

// QueryModel queries the "model" edge of the Petroglyph entity.
func (pe *Petroglyph) QueryModel() *ModelQuery {
	return NewPetroglyphClient(pe.config).QueryModel(pe)
}

// QueryMound queries the "mound" edge of the Petroglyph entity.
func (pe *Petroglyph) QueryMound() *MoundQuery {
	return NewPetroglyphClient(pe.config).QueryMound(pe)
}

// QueryPublications queries the "publications" edge of the Petroglyph entity.
func (pe *Petroglyph) QueryPublications() *PublicationQuery {
	return NewPetroglyphClient(pe.config).QueryPublications(pe)
}

// QueryTechniques queries the "techniques" edge of the Petroglyph entity.
func (pe *Petroglyph) QueryTechniques() *TechniqueQuery {
	return NewPetroglyphClient(pe.config).QueryTechniques(pe)
}

// QueryRegion queries the "region" edge of the Petroglyph entity.
func (pe *Petroglyph) QueryRegion() *RegionQuery {
	return NewPetroglyphClient(pe.config).QueryRegion(pe)
}

// QueryAccountingDocumentationAddress queries the "accounting_documentation_address" edge of the Petroglyph entity.
func (pe *Petroglyph) QueryAccountingDocumentationAddress() *LocationQuery {
	return NewPetroglyphClient(pe.config).QueryAccountingDocumentationAddress(pe)
}

// QueryAccountingDocumentationAuthor queries the "accounting_documentation_author" edge of the Petroglyph entity.
func (pe *Petroglyph) QueryAccountingDocumentationAuthor() *PersonQuery {
	return NewPetroglyphClient(pe.config).QueryAccountingDocumentationAuthor(pe)
}

// Update returns a builder for updating this Petroglyph.
// Note that you need to call Petroglyph.Unwrap() before calling this method if this Petroglyph
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Petroglyph) Update() *PetroglyphUpdateOne {
	return NewPetroglyphClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Petroglyph entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Petroglyph) Unwrap() *Petroglyph {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Petroglyph is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Petroglyph) String() string {
	var builder strings.Builder
	builder.WriteString("Petroglyph(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pe.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pe.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(pe.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(pe.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pe.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(pe.ExternalLink)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pe.Status))
	builder.WriteString(", ")
	builder.WriteString("primary_image_url=")
	builder.WriteString(pe.PrimaryImageURL)
	builder.WriteString(", ")
	builder.WriteString("additional_images_urls=")
	builder.WriteString(fmt.Sprintf("%v", pe.AdditionalImagesUrls))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pe.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(pe.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(pe.Number)
	builder.WriteString(", ")
	builder.WriteString("dating=")
	builder.WriteString(pe.Dating)
	builder.WriteString(", ")
	builder.WriteString("dating_start=")
	builder.WriteString(fmt.Sprintf("%v", pe.DatingStart))
	builder.WriteString(", ")
	builder.WriteString("dating_end=")
	builder.WriteString(fmt.Sprintf("%v", pe.DatingEnd))
	builder.WriteString(", ")
	builder.WriteString("orientation=")
	builder.WriteString(pe.Orientation)
	builder.WriteString(", ")
	builder.WriteString("position=")
	builder.WriteString(pe.Position)
	builder.WriteString(", ")
	builder.WriteString("geometric_shape=")
	builder.WriteString(pe.GeometricShape)
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(fmt.Sprintf("%v", pe.Height))
	builder.WriteString(", ")
	builder.WriteString("width=")
	builder.WriteString(fmt.Sprintf("%v", pe.Width))
	builder.WriteString(", ")
	builder.WriteString("length=")
	builder.WriteString(fmt.Sprintf("%v", pe.Length))
	builder.WriteString(", ")
	builder.WriteString("depth=")
	builder.WriteString(fmt.Sprintf("%v", pe.Depth))
	builder.WriteString(", ")
	builder.WriteString("diameter=")
	builder.WriteString(fmt.Sprintf("%v", pe.Diameter))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(pe.Weight)
	builder.WriteString(", ")
	builder.WriteString("dimensions=")
	builder.WriteString(pe.Dimensions)
	builder.WriteString(", ")
	builder.WriteString("plane_preservation=")
	builder.WriteString(pe.PlanePreservation)
	builder.WriteString(", ")
	builder.WriteString("photo_code=")
	builder.WriteString(pe.PhotoCode)
	builder.WriteString(", ")
	builder.WriteString("accounting_documentation_information=")
	builder.WriteString(pe.AccountingDocumentationInformation)
	builder.WriteString(", ")
	builder.WriteString("accounting_documentation_date=")
	builder.WriteString(pe.AccountingDocumentationDate.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pe.Geometry; v != nil {
		builder.WriteString("geometry=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedPublications returns the Publications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Petroglyph) NamedPublications(name string) ([]*Publication, error) {
	if pe.Edges.namedPublications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedPublications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Petroglyph) appendNamedPublications(name string, edges ...*Publication) {
	if pe.Edges.namedPublications == nil {
		pe.Edges.namedPublications = make(map[string][]*Publication)
	}
	if len(edges) == 0 {
		pe.Edges.namedPublications[name] = []*Publication{}
	} else {
		pe.Edges.namedPublications[name] = append(pe.Edges.namedPublications[name], edges...)
	}
}

// NamedTechniques returns the Techniques named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Petroglyph) NamedTechniques(name string) ([]*Technique, error) {
	if pe.Edges.namedTechniques == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedTechniques[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Petroglyph) appendNamedTechniques(name string, edges ...*Technique) {
	if pe.Edges.namedTechniques == nil {
		pe.Edges.namedTechniques = make(map[string][]*Technique)
	}
	if len(edges) == 0 {
		pe.Edges.namedTechniques[name] = []*Technique{}
	} else {
		pe.Edges.namedTechniques[name] = append(pe.Edges.namedTechniques[name], edges...)
	}
}

// Petroglyphs is a parsable slice of Petroglyph.
type Petroglyphs []*Petroglyph