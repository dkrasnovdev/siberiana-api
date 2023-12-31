// Code generated by ent, DO NOT EDIT.

package protectedareapicture

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the protectedareapicture type in the database.
	Label = "protected_area_picture"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldAbbreviation holds the string denoting the abbreviation field in the database.
	FieldAbbreviation = "abbreviation"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldExternalLink holds the string denoting the external_link field in the database.
	FieldExternalLink = "external_link"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPrimaryImageURL holds the string denoting the primary_image_url field in the database.
	FieldPrimaryImageURL = "primary_image_url"
	// FieldAdditionalImagesUrls holds the string denoting the additional_images_urls field in the database.
	FieldAdditionalImagesUrls = "additional_images_urls"
	// FieldShootingDate holds the string denoting the shooting_date field in the database.
	FieldShootingDate = "shooting_date"
	// FieldGeometry holds the string denoting the geometry field in the database.
	FieldGeometry = "geometry"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeCollection holds the string denoting the collection edge name in mutations.
	EdgeCollection = "collection"
	// EdgeProtectedArea holds the string denoting the protected_area edge name in mutations.
	EdgeProtectedArea = "protected_area"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeLicense holds the string denoting the license edge name in mutations.
	EdgeLicense = "license"
	// EdgeCountry holds the string denoting the country edge name in mutations.
	EdgeCountry = "country"
	// EdgeSettlement holds the string denoting the settlement edge name in mutations.
	EdgeSettlement = "settlement"
	// EdgeDistrict holds the string denoting the district edge name in mutations.
	EdgeDistrict = "district"
	// EdgeRegion holds the string denoting the region edge name in mutations.
	EdgeRegion = "region"
	// EdgePersonalCollection holds the string denoting the personal_collection edge name in mutations.
	EdgePersonalCollection = "personal_collection"
	// Table holds the table name of the protectedareapicture in the database.
	Table = "protected_area_pictures"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "protected_area_pictures"
	// AuthorInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	AuthorInverseTable = "persons"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "person_protected_area_pictures"
	// CollectionTable is the table that holds the collection relation/edge.
	CollectionTable = "protected_area_pictures"
	// CollectionInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionInverseTable = "collections"
	// CollectionColumn is the table column denoting the collection relation/edge.
	CollectionColumn = "collection_protected_area_pictures"
	// ProtectedAreaTable is the table that holds the protected_area relation/edge.
	ProtectedAreaTable = "protected_area_pictures"
	// ProtectedAreaInverseTable is the table name for the ProtectedArea entity.
	// It exists in this package in order to avoid circular dependency with the "protectedarea" package.
	ProtectedAreaInverseTable = "protected_areas"
	// ProtectedAreaColumn is the table column denoting the protected_area relation/edge.
	ProtectedAreaColumn = "protected_area_protected_area_pictures"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "protected_area_pictures"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "locations"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "location_protected_area_pictures"
	// LicenseTable is the table that holds the license relation/edge.
	LicenseTable = "protected_area_pictures"
	// LicenseInverseTable is the table name for the License entity.
	// It exists in this package in order to avoid circular dependency with the "license" package.
	LicenseInverseTable = "licenses"
	// LicenseColumn is the table column denoting the license relation/edge.
	LicenseColumn = "license_protected_area_pictures"
	// CountryTable is the table that holds the country relation/edge.
	CountryTable = "protected_area_pictures"
	// CountryInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	CountryInverseTable = "countries"
	// CountryColumn is the table column denoting the country relation/edge.
	CountryColumn = "country_protected_area_pictures"
	// SettlementTable is the table that holds the settlement relation/edge.
	SettlementTable = "protected_area_pictures"
	// SettlementInverseTable is the table name for the Settlement entity.
	// It exists in this package in order to avoid circular dependency with the "settlement" package.
	SettlementInverseTable = "settlements"
	// SettlementColumn is the table column denoting the settlement relation/edge.
	SettlementColumn = "settlement_protected_area_pictures"
	// DistrictTable is the table that holds the district relation/edge.
	DistrictTable = "protected_area_pictures"
	// DistrictInverseTable is the table name for the District entity.
	// It exists in this package in order to avoid circular dependency with the "district" package.
	DistrictInverseTable = "districts"
	// DistrictColumn is the table column denoting the district relation/edge.
	DistrictColumn = "district_protected_area_pictures"
	// RegionTable is the table that holds the region relation/edge.
	RegionTable = "protected_area_pictures"
	// RegionInverseTable is the table name for the Region entity.
	// It exists in this package in order to avoid circular dependency with the "region" package.
	RegionInverseTable = "regions"
	// RegionColumn is the table column denoting the region relation/edge.
	RegionColumn = "region_protected_area_pictures"
	// PersonalCollectionTable is the table that holds the personal_collection relation/edge. The primary key declared below.
	PersonalCollectionTable = "personal_collection_protected_area_pictures"
	// PersonalCollectionInverseTable is the table name for the PersonalCollection entity.
	// It exists in this package in order to avoid circular dependency with the "personalcollection" package.
	PersonalCollectionInverseTable = "personal_collections"
)

// Columns holds all SQL columns for protectedareapicture fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDisplayName,
	FieldAbbreviation,
	FieldDescription,
	FieldExternalLink,
	FieldStatus,
	FieldPrimaryImageURL,
	FieldAdditionalImagesUrls,
	FieldShootingDate,
	FieldGeometry,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "protected_area_pictures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"collection_protected_area_pictures",
	"country_protected_area_pictures",
	"district_protected_area_pictures",
	"license_protected_area_pictures",
	"location_protected_area_pictures",
	"person_protected_area_pictures",
	"protected_area_protected_area_pictures",
	"region_protected_area_pictures",
	"settlement_protected_area_pictures",
}

var (
	// PersonalCollectionPrimaryKey and PersonalCollectionColumn2 are the table columns denoting the
	// primary key for the personal_collection relation (M2M).
	PersonalCollectionPrimaryKey = []string{"personal_collection_id", "protected_area_picture_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/dkrasnovdev/siberiana-api/ent/runtime"
var (
	Hooks  [2]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusDraft is the default value of the Status enum.
const DefaultStatus = StatusDraft

// Status values.
const (
	StatusListed   Status = "listed"
	StatusUnlisted Status = "unlisted"
	StatusDraft    Status = "draft"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusListed, StatusUnlisted, StatusDraft:
		return nil
	default:
		return fmt.Errorf("protectedareapicture: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the ProtectedAreaPicture queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByAbbreviation orders the results by the abbreviation field.
func ByAbbreviation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAbbreviation, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByExternalLink orders the results by the external_link field.
func ByExternalLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExternalLink, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPrimaryImageURL orders the results by the primary_image_url field.
func ByPrimaryImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrimaryImageURL, opts...).ToFunc()
}

// ByShootingDate orders the results by the shooting_date field.
func ByShootingDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShootingDate, opts...).ToFunc()
}

// ByGeometry orders the results by the geometry field.
func ByGeometry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeometry, opts...).ToFunc()
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByCollectionField orders the results by collection field.
func ByCollectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCollectionStep(), sql.OrderByField(field, opts...))
	}
}

// ByProtectedAreaField orders the results by protected_area field.
func ByProtectedAreaField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProtectedAreaStep(), sql.OrderByField(field, opts...))
	}
}

// ByLocationField orders the results by location field.
func ByLocationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLocationStep(), sql.OrderByField(field, opts...))
	}
}

// ByLicenseField orders the results by license field.
func ByLicenseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLicenseStep(), sql.OrderByField(field, opts...))
	}
}

// ByCountryField orders the results by country field.
func ByCountryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCountryStep(), sql.OrderByField(field, opts...))
	}
}

// BySettlementField orders the results by settlement field.
func BySettlementField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSettlementStep(), sql.OrderByField(field, opts...))
	}
}

// ByDistrictField orders the results by district field.
func ByDistrictField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDistrictStep(), sql.OrderByField(field, opts...))
	}
}

// ByRegionField orders the results by region field.
func ByRegionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRegionStep(), sql.OrderByField(field, opts...))
	}
}

// ByPersonalCollectionCount orders the results by personal_collection count.
func ByPersonalCollectionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPersonalCollectionStep(), opts...)
	}
}

// ByPersonalCollection orders the results by personal_collection terms.
func ByPersonalCollection(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPersonalCollectionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
	)
}
func newCollectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CollectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CollectionTable, CollectionColumn),
	)
}
func newProtectedAreaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProtectedAreaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProtectedAreaTable, ProtectedAreaColumn),
	)
}
func newLocationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LocationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
	)
}
func newLicenseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LicenseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LicenseTable, LicenseColumn),
	)
}
func newCountryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CountryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CountryTable, CountryColumn),
	)
}
func newSettlementStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SettlementInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SettlementTable, SettlementColumn),
	)
}
func newDistrictStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DistrictInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DistrictTable, DistrictColumn),
	)
}
func newRegionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RegionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RegionTable, RegionColumn),
	)
}
func newPersonalCollectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PersonalCollectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PersonalCollectionTable, PersonalCollectionPrimaryKey...),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
