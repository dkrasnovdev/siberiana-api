// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/district"
	"github.com/dkrasnovdev/siberiana-api/ent/region"
	"github.com/dkrasnovdev/siberiana-api/ent/settlement"
)

// Settlement is the model entity for the Settlement schema.
type Settlement struct {
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
	// The values are being populated by the SettlementQuery when eager-loading is set.
	Edges                SettlementEdges `json:"edges"`
	district_settlements *int
	region_settlements   *int
	selectValues         sql.SelectValues
}

// SettlementEdges holds the relations/edges for other nodes in the graph.
type SettlementEdges struct {
	// Art holds the value of the art edge.
	Art []*Art `json:"art,omitempty"`
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// Herbaria holds the value of the herbaria edge.
	Herbaria []*Herbarium `json:"herbaria,omitempty"`
	// ProtectedAreaPictures holds the value of the protected_area_pictures edge.
	ProtectedAreaPictures []*ProtectedAreaPicture `json:"protected_area_pictures,omitempty"`
	// Locations holds the value of the locations edge.
	Locations []*Location `json:"locations,omitempty"`
	// Region holds the value of the region edge.
	Region *Region `json:"region,omitempty"`
	// District holds the value of the district edge.
	District *District `json:"district,omitempty"`
	// KnownAsAfter holds the value of the known_as_after edge.
	KnownAsAfter []*Settlement `json:"known_as_after,omitempty"`
	// KnownAsBefore holds the value of the known_as_before edge.
	KnownAsBefore []*Settlement `json:"known_as_before,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
	// totalCount holds the count of the edges above.
	totalCount [10]map[string]int

	namedArt                   map[string][]*Art
	namedArtifacts             map[string][]*Artifact
	namedBooks                 map[string][]*Book
	namedHerbaria              map[string][]*Herbarium
	namedProtectedAreaPictures map[string][]*ProtectedAreaPicture
	namedLocations             map[string][]*Location
	namedKnownAsAfter          map[string][]*Settlement
	namedKnownAsBefore         map[string][]*Settlement
}

// ArtOrErr returns the Art value or an error if the edge
// was not loaded in eager-loading.
func (e SettlementEdges) ArtOrErr() ([]*Art, error) {
	if e.loadedTypes[0] {
		return e.Art, nil
	}
	return nil, &NotLoadedError{edge: "art"}
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e SettlementEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[1] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e SettlementEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[2] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// HerbariaOrErr returns the Herbaria value or an error if the edge
// was not loaded in eager-loading.
func (e SettlementEdges) HerbariaOrErr() ([]*Herbarium, error) {
	if e.loadedTypes[3] {
		return e.Herbaria, nil
	}
	return nil, &NotLoadedError{edge: "herbaria"}
}

// ProtectedAreaPicturesOrErr returns the ProtectedAreaPictures value or an error if the edge
// was not loaded in eager-loading.
func (e SettlementEdges) ProtectedAreaPicturesOrErr() ([]*ProtectedAreaPicture, error) {
	if e.loadedTypes[4] {
		return e.ProtectedAreaPictures, nil
	}
	return nil, &NotLoadedError{edge: "protected_area_pictures"}
}

// LocationsOrErr returns the Locations value or an error if the edge
// was not loaded in eager-loading.
func (e SettlementEdges) LocationsOrErr() ([]*Location, error) {
	if e.loadedTypes[5] {
		return e.Locations, nil
	}
	return nil, &NotLoadedError{edge: "locations"}
}

// RegionOrErr returns the Region value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SettlementEdges) RegionOrErr() (*Region, error) {
	if e.loadedTypes[6] {
		if e.Region == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: region.Label}
		}
		return e.Region, nil
	}
	return nil, &NotLoadedError{edge: "region"}
}

// DistrictOrErr returns the District value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SettlementEdges) DistrictOrErr() (*District, error) {
	if e.loadedTypes[7] {
		if e.District == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: district.Label}
		}
		return e.District, nil
	}
	return nil, &NotLoadedError{edge: "district"}
}

// KnownAsAfterOrErr returns the KnownAsAfter value or an error if the edge
// was not loaded in eager-loading.
func (e SettlementEdges) KnownAsAfterOrErr() ([]*Settlement, error) {
	if e.loadedTypes[8] {
		return e.KnownAsAfter, nil
	}
	return nil, &NotLoadedError{edge: "known_as_after"}
}

// KnownAsBeforeOrErr returns the KnownAsBefore value or an error if the edge
// was not loaded in eager-loading.
func (e SettlementEdges) KnownAsBeforeOrErr() ([]*Settlement, error) {
	if e.loadedTypes[9] {
		return e.KnownAsBefore, nil
	}
	return nil, &NotLoadedError{edge: "known_as_before"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Settlement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case settlement.FieldID:
			values[i] = new(sql.NullInt64)
		case settlement.FieldCreatedBy, settlement.FieldUpdatedBy, settlement.FieldDisplayName, settlement.FieldAbbreviation, settlement.FieldDescription, settlement.FieldExternalLink:
			values[i] = new(sql.NullString)
		case settlement.FieldCreatedAt, settlement.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case settlement.ForeignKeys[0]: // district_settlements
			values[i] = new(sql.NullInt64)
		case settlement.ForeignKeys[1]: // region_settlements
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Settlement fields.
func (s *Settlement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case settlement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case settlement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case settlement.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				s.CreatedBy = value.String
			}
		case settlement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case settlement.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				s.UpdatedBy = value.String
			}
		case settlement.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				s.DisplayName = value.String
			}
		case settlement.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				s.Abbreviation = value.String
			}
		case settlement.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case settlement.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				s.ExternalLink = value.String
			}
		case settlement.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field district_settlements", value)
			} else if value.Valid {
				s.district_settlements = new(int)
				*s.district_settlements = int(value.Int64)
			}
		case settlement.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field region_settlements", value)
			} else if value.Valid {
				s.region_settlements = new(int)
				*s.region_settlements = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Settlement.
// This includes values selected through modifiers, order, etc.
func (s *Settlement) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryArt queries the "art" edge of the Settlement entity.
func (s *Settlement) QueryArt() *ArtQuery {
	return NewSettlementClient(s.config).QueryArt(s)
}

// QueryArtifacts queries the "artifacts" edge of the Settlement entity.
func (s *Settlement) QueryArtifacts() *ArtifactQuery {
	return NewSettlementClient(s.config).QueryArtifacts(s)
}

// QueryBooks queries the "books" edge of the Settlement entity.
func (s *Settlement) QueryBooks() *BookQuery {
	return NewSettlementClient(s.config).QueryBooks(s)
}

// QueryHerbaria queries the "herbaria" edge of the Settlement entity.
func (s *Settlement) QueryHerbaria() *HerbariumQuery {
	return NewSettlementClient(s.config).QueryHerbaria(s)
}

// QueryProtectedAreaPictures queries the "protected_area_pictures" edge of the Settlement entity.
func (s *Settlement) QueryProtectedAreaPictures() *ProtectedAreaPictureQuery {
	return NewSettlementClient(s.config).QueryProtectedAreaPictures(s)
}

// QueryLocations queries the "locations" edge of the Settlement entity.
func (s *Settlement) QueryLocations() *LocationQuery {
	return NewSettlementClient(s.config).QueryLocations(s)
}

// QueryRegion queries the "region" edge of the Settlement entity.
func (s *Settlement) QueryRegion() *RegionQuery {
	return NewSettlementClient(s.config).QueryRegion(s)
}

// QueryDistrict queries the "district" edge of the Settlement entity.
func (s *Settlement) QueryDistrict() *DistrictQuery {
	return NewSettlementClient(s.config).QueryDistrict(s)
}

// QueryKnownAsAfter queries the "known_as_after" edge of the Settlement entity.
func (s *Settlement) QueryKnownAsAfter() *SettlementQuery {
	return NewSettlementClient(s.config).QueryKnownAsAfter(s)
}

// QueryKnownAsBefore queries the "known_as_before" edge of the Settlement entity.
func (s *Settlement) QueryKnownAsBefore() *SettlementQuery {
	return NewSettlementClient(s.config).QueryKnownAsBefore(s)
}

// Update returns a builder for updating this Settlement.
// Note that you need to call Settlement.Unwrap() before calling this method if this Settlement
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Settlement) Update() *SettlementUpdateOne {
	return NewSettlementClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Settlement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Settlement) Unwrap() *Settlement {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Settlement is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Settlement) String() string {
	var builder strings.Builder
	builder.WriteString("Settlement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(s.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(s.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(s.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(s.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(s.ExternalLink)
	builder.WriteByte(')')
	return builder.String()
}

// NamedArt returns the Art named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Settlement) NamedArt(name string) ([]*Art, error) {
	if s.Edges.namedArt == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedArt[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Settlement) appendNamedArt(name string, edges ...*Art) {
	if s.Edges.namedArt == nil {
		s.Edges.namedArt = make(map[string][]*Art)
	}
	if len(edges) == 0 {
		s.Edges.namedArt[name] = []*Art{}
	} else {
		s.Edges.namedArt[name] = append(s.Edges.namedArt[name], edges...)
	}
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Settlement) NamedArtifacts(name string) ([]*Artifact, error) {
	if s.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Settlement) appendNamedArtifacts(name string, edges ...*Artifact) {
	if s.Edges.namedArtifacts == nil {
		s.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		s.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		s.Edges.namedArtifacts[name] = append(s.Edges.namedArtifacts[name], edges...)
	}
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Settlement) NamedBooks(name string) ([]*Book, error) {
	if s.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Settlement) appendNamedBooks(name string, edges ...*Book) {
	if s.Edges.namedBooks == nil {
		s.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		s.Edges.namedBooks[name] = []*Book{}
	} else {
		s.Edges.namedBooks[name] = append(s.Edges.namedBooks[name], edges...)
	}
}

// NamedHerbaria returns the Herbaria named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Settlement) NamedHerbaria(name string) ([]*Herbarium, error) {
	if s.Edges.namedHerbaria == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedHerbaria[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Settlement) appendNamedHerbaria(name string, edges ...*Herbarium) {
	if s.Edges.namedHerbaria == nil {
		s.Edges.namedHerbaria = make(map[string][]*Herbarium)
	}
	if len(edges) == 0 {
		s.Edges.namedHerbaria[name] = []*Herbarium{}
	} else {
		s.Edges.namedHerbaria[name] = append(s.Edges.namedHerbaria[name], edges...)
	}
}

// NamedProtectedAreaPictures returns the ProtectedAreaPictures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Settlement) NamedProtectedAreaPictures(name string) ([]*ProtectedAreaPicture, error) {
	if s.Edges.namedProtectedAreaPictures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedProtectedAreaPictures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Settlement) appendNamedProtectedAreaPictures(name string, edges ...*ProtectedAreaPicture) {
	if s.Edges.namedProtectedAreaPictures == nil {
		s.Edges.namedProtectedAreaPictures = make(map[string][]*ProtectedAreaPicture)
	}
	if len(edges) == 0 {
		s.Edges.namedProtectedAreaPictures[name] = []*ProtectedAreaPicture{}
	} else {
		s.Edges.namedProtectedAreaPictures[name] = append(s.Edges.namedProtectedAreaPictures[name], edges...)
	}
}

// NamedLocations returns the Locations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Settlement) NamedLocations(name string) ([]*Location, error) {
	if s.Edges.namedLocations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedLocations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Settlement) appendNamedLocations(name string, edges ...*Location) {
	if s.Edges.namedLocations == nil {
		s.Edges.namedLocations = make(map[string][]*Location)
	}
	if len(edges) == 0 {
		s.Edges.namedLocations[name] = []*Location{}
	} else {
		s.Edges.namedLocations[name] = append(s.Edges.namedLocations[name], edges...)
	}
}

// NamedKnownAsAfter returns the KnownAsAfter named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Settlement) NamedKnownAsAfter(name string) ([]*Settlement, error) {
	if s.Edges.namedKnownAsAfter == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedKnownAsAfter[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Settlement) appendNamedKnownAsAfter(name string, edges ...*Settlement) {
	if s.Edges.namedKnownAsAfter == nil {
		s.Edges.namedKnownAsAfter = make(map[string][]*Settlement)
	}
	if len(edges) == 0 {
		s.Edges.namedKnownAsAfter[name] = []*Settlement{}
	} else {
		s.Edges.namedKnownAsAfter[name] = append(s.Edges.namedKnownAsAfter[name], edges...)
	}
}

// NamedKnownAsBefore returns the KnownAsBefore named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Settlement) NamedKnownAsBefore(name string) ([]*Settlement, error) {
	if s.Edges.namedKnownAsBefore == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedKnownAsBefore[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Settlement) appendNamedKnownAsBefore(name string, edges ...*Settlement) {
	if s.Edges.namedKnownAsBefore == nil {
		s.Edges.namedKnownAsBefore = make(map[string][]*Settlement)
	}
	if len(edges) == 0 {
		s.Edges.namedKnownAsBefore[name] = []*Settlement{}
	} else {
		s.Edges.namedKnownAsBefore[name] = append(s.Edges.namedKnownAsBefore[name], edges...)
	}
}

// Settlements is a parsable slice of Settlement.
type Settlements []*Settlement
