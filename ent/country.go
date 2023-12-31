// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/country"
)

// Country is the model entity for the Country schema.
type Country struct {
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
	// The values are being populated by the CountryQuery when eager-loading is set.
	Edges        CountryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CountryEdges holds the relations/edges for other nodes in the graph.
type CountryEdges struct {
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
	// Regions holds the value of the regions edge.
	Regions []*Region `json:"regions,omitempty"`
	// Locations holds the value of the locations edge.
	Locations []*Location `json:"locations,omitempty"`
	// KnownAsAfter holds the value of the known_as_after edge.
	KnownAsAfter []*Country `json:"known_as_after,omitempty"`
	// KnownAsBefore holds the value of the known_as_before edge.
	KnownAsBefore []*Country `json:"known_as_before,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
	// totalCount holds the count of the edges above.
	totalCount [9]map[string]int

	namedArt                   map[string][]*Art
	namedArtifacts             map[string][]*Artifact
	namedBooks                 map[string][]*Book
	namedHerbaria              map[string][]*Herbarium
	namedProtectedAreaPictures map[string][]*ProtectedAreaPicture
	namedRegions               map[string][]*Region
	namedLocations             map[string][]*Location
	namedKnownAsAfter          map[string][]*Country
	namedKnownAsBefore         map[string][]*Country
}

// ArtOrErr returns the Art value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) ArtOrErr() ([]*Art, error) {
	if e.loadedTypes[0] {
		return e.Art, nil
	}
	return nil, &NotLoadedError{edge: "art"}
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[1] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[2] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// HerbariaOrErr returns the Herbaria value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) HerbariaOrErr() ([]*Herbarium, error) {
	if e.loadedTypes[3] {
		return e.Herbaria, nil
	}
	return nil, &NotLoadedError{edge: "herbaria"}
}

// ProtectedAreaPicturesOrErr returns the ProtectedAreaPictures value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) ProtectedAreaPicturesOrErr() ([]*ProtectedAreaPicture, error) {
	if e.loadedTypes[4] {
		return e.ProtectedAreaPictures, nil
	}
	return nil, &NotLoadedError{edge: "protected_area_pictures"}
}

// RegionsOrErr returns the Regions value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) RegionsOrErr() ([]*Region, error) {
	if e.loadedTypes[5] {
		return e.Regions, nil
	}
	return nil, &NotLoadedError{edge: "regions"}
}

// LocationsOrErr returns the Locations value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) LocationsOrErr() ([]*Location, error) {
	if e.loadedTypes[6] {
		return e.Locations, nil
	}
	return nil, &NotLoadedError{edge: "locations"}
}

// KnownAsAfterOrErr returns the KnownAsAfter value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) KnownAsAfterOrErr() ([]*Country, error) {
	if e.loadedTypes[7] {
		return e.KnownAsAfter, nil
	}
	return nil, &NotLoadedError{edge: "known_as_after"}
}

// KnownAsBeforeOrErr returns the KnownAsBefore value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) KnownAsBeforeOrErr() ([]*Country, error) {
	if e.loadedTypes[8] {
		return e.KnownAsBefore, nil
	}
	return nil, &NotLoadedError{edge: "known_as_before"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			values[i] = new(sql.NullInt64)
		case country.FieldCreatedBy, country.FieldUpdatedBy, country.FieldDisplayName, country.FieldAbbreviation, country.FieldDescription, country.FieldExternalLink:
			values[i] = new(sql.NullString)
		case country.FieldCreatedAt, country.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country fields.
func (c *Country) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case country.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case country.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = value.String
			}
		case country.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case country.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = value.String
			}
		case country.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				c.DisplayName = value.String
			}
		case country.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				c.Abbreviation = value.String
			}
		case country.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case country.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				c.ExternalLink = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Country.
// This includes values selected through modifiers, order, etc.
func (c *Country) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryArt queries the "art" edge of the Country entity.
func (c *Country) QueryArt() *ArtQuery {
	return NewCountryClient(c.config).QueryArt(c)
}

// QueryArtifacts queries the "artifacts" edge of the Country entity.
func (c *Country) QueryArtifacts() *ArtifactQuery {
	return NewCountryClient(c.config).QueryArtifacts(c)
}

// QueryBooks queries the "books" edge of the Country entity.
func (c *Country) QueryBooks() *BookQuery {
	return NewCountryClient(c.config).QueryBooks(c)
}

// QueryHerbaria queries the "herbaria" edge of the Country entity.
func (c *Country) QueryHerbaria() *HerbariumQuery {
	return NewCountryClient(c.config).QueryHerbaria(c)
}

// QueryProtectedAreaPictures queries the "protected_area_pictures" edge of the Country entity.
func (c *Country) QueryProtectedAreaPictures() *ProtectedAreaPictureQuery {
	return NewCountryClient(c.config).QueryProtectedAreaPictures(c)
}

// QueryRegions queries the "regions" edge of the Country entity.
func (c *Country) QueryRegions() *RegionQuery {
	return NewCountryClient(c.config).QueryRegions(c)
}

// QueryLocations queries the "locations" edge of the Country entity.
func (c *Country) QueryLocations() *LocationQuery {
	return NewCountryClient(c.config).QueryLocations(c)
}

// QueryKnownAsAfter queries the "known_as_after" edge of the Country entity.
func (c *Country) QueryKnownAsAfter() *CountryQuery {
	return NewCountryClient(c.config).QueryKnownAsAfter(c)
}

// QueryKnownAsBefore queries the "known_as_before" edge of the Country entity.
func (c *Country) QueryKnownAsBefore() *CountryQuery {
	return NewCountryClient(c.config).QueryKnownAsBefore(c)
}

// Update returns a builder for updating this Country.
// Note that you need to call Country.Unwrap() before calling this method if this Country
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Country) Update() *CountryUpdateOne {
	return NewCountryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Country entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Country) Unwrap() *Country {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Country is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Country) String() string {
	var builder strings.Builder
	builder.WriteString("Country(")
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
	builder.WriteString("abbreviation=")
	builder.WriteString(c.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(c.ExternalLink)
	builder.WriteByte(')')
	return builder.String()
}

// NamedArt returns the Art named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedArt(name string) ([]*Art, error) {
	if c.Edges.namedArt == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedArt[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedArt(name string, edges ...*Art) {
	if c.Edges.namedArt == nil {
		c.Edges.namedArt = make(map[string][]*Art)
	}
	if len(edges) == 0 {
		c.Edges.namedArt[name] = []*Art{}
	} else {
		c.Edges.namedArt[name] = append(c.Edges.namedArt[name], edges...)
	}
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedArtifacts(name string) ([]*Artifact, error) {
	if c.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedArtifacts(name string, edges ...*Artifact) {
	if c.Edges.namedArtifacts == nil {
		c.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		c.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		c.Edges.namedArtifacts[name] = append(c.Edges.namedArtifacts[name], edges...)
	}
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedBooks(name string) ([]*Book, error) {
	if c.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedBooks(name string, edges ...*Book) {
	if c.Edges.namedBooks == nil {
		c.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		c.Edges.namedBooks[name] = []*Book{}
	} else {
		c.Edges.namedBooks[name] = append(c.Edges.namedBooks[name], edges...)
	}
}

// NamedHerbaria returns the Herbaria named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedHerbaria(name string) ([]*Herbarium, error) {
	if c.Edges.namedHerbaria == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedHerbaria[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedHerbaria(name string, edges ...*Herbarium) {
	if c.Edges.namedHerbaria == nil {
		c.Edges.namedHerbaria = make(map[string][]*Herbarium)
	}
	if len(edges) == 0 {
		c.Edges.namedHerbaria[name] = []*Herbarium{}
	} else {
		c.Edges.namedHerbaria[name] = append(c.Edges.namedHerbaria[name], edges...)
	}
}

// NamedProtectedAreaPictures returns the ProtectedAreaPictures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedProtectedAreaPictures(name string) ([]*ProtectedAreaPicture, error) {
	if c.Edges.namedProtectedAreaPictures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProtectedAreaPictures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedProtectedAreaPictures(name string, edges ...*ProtectedAreaPicture) {
	if c.Edges.namedProtectedAreaPictures == nil {
		c.Edges.namedProtectedAreaPictures = make(map[string][]*ProtectedAreaPicture)
	}
	if len(edges) == 0 {
		c.Edges.namedProtectedAreaPictures[name] = []*ProtectedAreaPicture{}
	} else {
		c.Edges.namedProtectedAreaPictures[name] = append(c.Edges.namedProtectedAreaPictures[name], edges...)
	}
}

// NamedRegions returns the Regions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedRegions(name string) ([]*Region, error) {
	if c.Edges.namedRegions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedRegions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedRegions(name string, edges ...*Region) {
	if c.Edges.namedRegions == nil {
		c.Edges.namedRegions = make(map[string][]*Region)
	}
	if len(edges) == 0 {
		c.Edges.namedRegions[name] = []*Region{}
	} else {
		c.Edges.namedRegions[name] = append(c.Edges.namedRegions[name], edges...)
	}
}

// NamedLocations returns the Locations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedLocations(name string) ([]*Location, error) {
	if c.Edges.namedLocations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedLocations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedLocations(name string, edges ...*Location) {
	if c.Edges.namedLocations == nil {
		c.Edges.namedLocations = make(map[string][]*Location)
	}
	if len(edges) == 0 {
		c.Edges.namedLocations[name] = []*Location{}
	} else {
		c.Edges.namedLocations[name] = append(c.Edges.namedLocations[name], edges...)
	}
}

// NamedKnownAsAfter returns the KnownAsAfter named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedKnownAsAfter(name string) ([]*Country, error) {
	if c.Edges.namedKnownAsAfter == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedKnownAsAfter[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedKnownAsAfter(name string, edges ...*Country) {
	if c.Edges.namedKnownAsAfter == nil {
		c.Edges.namedKnownAsAfter = make(map[string][]*Country)
	}
	if len(edges) == 0 {
		c.Edges.namedKnownAsAfter[name] = []*Country{}
	} else {
		c.Edges.namedKnownAsAfter[name] = append(c.Edges.namedKnownAsAfter[name], edges...)
	}
}

// NamedKnownAsBefore returns the KnownAsBefore named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedKnownAsBefore(name string) ([]*Country, error) {
	if c.Edges.namedKnownAsBefore == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedKnownAsBefore[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedKnownAsBefore(name string, edges ...*Country) {
	if c.Edges.namedKnownAsBefore == nil {
		c.Edges.namedKnownAsBefore = make(map[string][]*Country)
	}
	if len(edges) == 0 {
		c.Edges.namedKnownAsBefore[name] = []*Country{}
	} else {
		c.Edges.namedKnownAsBefore[name] = append(c.Edges.namedKnownAsBefore[name], edges...)
	}
}

// Countries is a parsable slice of Country.
type Countries []*Country
