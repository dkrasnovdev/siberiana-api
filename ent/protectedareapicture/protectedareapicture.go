// Code generated by ent, DO NOT EDIT.

package protectedareapicture

import (
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
	// FieldExternalLinks holds the string denoting the external_links field in the database.
	FieldExternalLinks = "external_links"
	// FieldPrimaryImageURL holds the string denoting the primary_image_url field in the database.
	FieldPrimaryImageURL = "primary_image_url"
	// FieldAdditionalImagesUrls holds the string denoting the additional_images_urls field in the database.
	FieldAdditionalImagesUrls = "additional_images_urls"
	// FieldShootingDate holds the string denoting the shooting_date field in the database.
	FieldShootingDate = "shooting_date"
	// FieldGeometry holds the string denoting the geometry field in the database.
	FieldGeometry = "geometry"
	// EdgeCollection holds the string denoting the collection edge name in mutations.
	EdgeCollection = "collection"
	// EdgeProtectedArea holds the string denoting the protected_area edge name in mutations.
	EdgeProtectedArea = "protected_area"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeLicense holds the string denoting the license edge name in mutations.
	EdgeLicense = "license"
	// Table holds the table name of the protectedareapicture in the database.
	Table = "protected_area_pictures"
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
	FieldExternalLinks,
	FieldPrimaryImageURL,
	FieldAdditionalImagesUrls,
	FieldShootingDate,
	FieldGeometry,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "protected_area_pictures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"collection_protected_area_pictures",
	"license_protected_area_pictures",
	"location_protected_area_pictures",
	"protected_area_protected_area_pictures",
}

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
//	import _ "github.com/dkrasnovdev/heritage-api/ent/runtime"
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
