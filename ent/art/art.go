// Code generated by ent, DO NOT EDIT.

package art

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the art type in the database.
	Label = "art"
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
	// FieldPrimaryImageURL holds the string denoting the primary_image_url field in the database.
	FieldPrimaryImageURL = "primary_image_url"
	// FieldAdditionalImagesUrls holds the string denoting the additional_images_urls field in the database.
	FieldAdditionalImagesUrls = "additional_images_urls"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldDating holds the string denoting the dating field in the database.
	FieldDating = "dating"
	// FieldDimensions holds the string denoting the dimensions field in the database.
	FieldDimensions = "dimensions"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeArtGenre holds the string denoting the art_genre edge name in mutations.
	EdgeArtGenre = "art_genre"
	// EdgeArtStyle holds the string denoting the art_style edge name in mutations.
	EdgeArtStyle = "art_style"
	// EdgeMediums holds the string denoting the mediums edge name in mutations.
	EdgeMediums = "mediums"
	// EdgeCollection holds the string denoting the collection edge name in mutations.
	EdgeCollection = "collection"
	// EdgeCountry holds the string denoting the country edge name in mutations.
	EdgeCountry = "country"
	// EdgeSettlement holds the string denoting the settlement edge name in mutations.
	EdgeSettlement = "settlement"
	// EdgeDistrict holds the string denoting the district edge name in mutations.
	EdgeDistrict = "district"
	// EdgeRegion holds the string denoting the region edge name in mutations.
	EdgeRegion = "region"
	// Table holds the table name of the art in the database.
	Table = "arts"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "arts"
	// AuthorInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	AuthorInverseTable = "persons"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "person_art"
	// ArtGenreTable is the table that holds the art_genre relation/edge. The primary key declared below.
	ArtGenreTable = "art_genre_art"
	// ArtGenreInverseTable is the table name for the ArtGenre entity.
	// It exists in this package in order to avoid circular dependency with the "artgenre" package.
	ArtGenreInverseTable = "art_genres"
	// ArtStyleTable is the table that holds the art_style relation/edge. The primary key declared below.
	ArtStyleTable = "art_style_art"
	// ArtStyleInverseTable is the table name for the ArtStyle entity.
	// It exists in this package in order to avoid circular dependency with the "artstyle" package.
	ArtStyleInverseTable = "art_styles"
	// MediumsTable is the table that holds the mediums relation/edge. The primary key declared below.
	MediumsTable = "medium_art"
	// MediumsInverseTable is the table name for the Medium entity.
	// It exists in this package in order to avoid circular dependency with the "medium" package.
	MediumsInverseTable = "media"
	// CollectionTable is the table that holds the collection relation/edge.
	CollectionTable = "arts"
	// CollectionInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionInverseTable = "collections"
	// CollectionColumn is the table column denoting the collection relation/edge.
	CollectionColumn = "collection_art"
	// CountryTable is the table that holds the country relation/edge.
	CountryTable = "arts"
	// CountryInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	CountryInverseTable = "countries"
	// CountryColumn is the table column denoting the country relation/edge.
	CountryColumn = "country_art"
	// SettlementTable is the table that holds the settlement relation/edge.
	SettlementTable = "arts"
	// SettlementInverseTable is the table name for the Settlement entity.
	// It exists in this package in order to avoid circular dependency with the "settlement" package.
	SettlementInverseTable = "settlements"
	// SettlementColumn is the table column denoting the settlement relation/edge.
	SettlementColumn = "settlement_art"
	// DistrictTable is the table that holds the district relation/edge.
	DistrictTable = "arts"
	// DistrictInverseTable is the table name for the District entity.
	// It exists in this package in order to avoid circular dependency with the "district" package.
	DistrictInverseTable = "districts"
	// DistrictColumn is the table column denoting the district relation/edge.
	DistrictColumn = "district_art"
	// RegionTable is the table that holds the region relation/edge.
	RegionTable = "arts"
	// RegionInverseTable is the table name for the Region entity.
	// It exists in this package in order to avoid circular dependency with the "region" package.
	RegionInverseTable = "regions"
	// RegionColumn is the table column denoting the region relation/edge.
	RegionColumn = "region_art"
)

// Columns holds all SQL columns for art fields.
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
	FieldPrimaryImageURL,
	FieldAdditionalImagesUrls,
	FieldNumber,
	FieldDating,
	FieldDimensions,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "arts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"collection_art",
	"country_art",
	"district_art",
	"person_art",
	"region_art",
	"settlement_art",
}

var (
	// ArtGenrePrimaryKey and ArtGenreColumn2 are the table columns denoting the
	// primary key for the art_genre relation (M2M).
	ArtGenrePrimaryKey = []string{"art_genre_id", "art_id"}
	// ArtStylePrimaryKey and ArtStyleColumn2 are the table columns denoting the
	// primary key for the art_style relation (M2M).
	ArtStylePrimaryKey = []string{"art_style_id", "art_id"}
	// MediumsPrimaryKey and MediumsColumn2 are the table columns denoting the
	// primary key for the mediums relation (M2M).
	MediumsPrimaryKey = []string{"medium_id", "art_id"}
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

// OrderOption defines the ordering options for the Art queries.
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

// ByPrimaryImageURL orders the results by the primary_image_url field.
func ByPrimaryImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrimaryImageURL, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByDating orders the results by the dating field.
func ByDating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDating, opts...).ToFunc()
}

// ByDimensions orders the results by the dimensions field.
func ByDimensions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDimensions, opts...).ToFunc()
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByArtGenreCount orders the results by art_genre count.
func ByArtGenreCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArtGenreStep(), opts...)
	}
}

// ByArtGenre orders the results by art_genre terms.
func ByArtGenre(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArtGenreStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByArtStyleCount orders the results by art_style count.
func ByArtStyleCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArtStyleStep(), opts...)
	}
}

// ByArtStyle orders the results by art_style terms.
func ByArtStyle(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArtStyleStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMediumsCount orders the results by mediums count.
func ByMediumsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMediumsStep(), opts...)
	}
}

// ByMediums orders the results by mediums terms.
func ByMediums(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMediumsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCollectionField orders the results by collection field.
func ByCollectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCollectionStep(), sql.OrderByField(field, opts...))
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
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
	)
}
func newArtGenreStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtGenreInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ArtGenreTable, ArtGenrePrimaryKey...),
	)
}
func newArtStyleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtStyleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ArtStyleTable, ArtStylePrimaryKey...),
	)
}
func newMediumsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediumsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, MediumsTable, MediumsPrimaryKey...),
	)
}
func newCollectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CollectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CollectionTable, CollectionColumn),
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
