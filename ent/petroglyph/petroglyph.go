// Code generated by ent, DO NOT EDIT.

package petroglyph

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
	// Label holds the string label denoting the petroglyph type in the database.
	Label = "petroglyph"
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
	// FieldDating holds the string denoting the dating field in the database.
	FieldDating = "dating"
	// FieldDatingStart holds the string denoting the dating_start field in the database.
	FieldDatingStart = "dating_start"
	// FieldDatingEnd holds the string denoting the dating_end field in the database.
	FieldDatingEnd = "dating_end"
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
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// FieldDepth holds the string denoting the depth field in the database.
	FieldDepth = "depth"
	// FieldDiameter holds the string denoting the diameter field in the database.
	FieldDiameter = "diameter"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldDimensions holds the string denoting the dimensions field in the database.
	FieldDimensions = "dimensions"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldOrientation holds the string denoting the orientation field in the database.
	FieldOrientation = "orientation"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldGeometricShape holds the string denoting the geometric_shape field in the database.
	FieldGeometricShape = "geometric_shape"
	// FieldPlanePreservation holds the string denoting the plane_preservation field in the database.
	FieldPlanePreservation = "plane_preservation"
	// FieldPhotoCode holds the string denoting the photo_code field in the database.
	FieldPhotoCode = "photo_code"
	// FieldAccountingDocumentationInformation holds the string denoting the accounting_documentation_information field in the database.
	FieldAccountingDocumentationInformation = "accounting_documentation_information"
	// FieldAccountingDocumentationDate holds the string denoting the accounting_documentation_date field in the database.
	FieldAccountingDocumentationDate = "accounting_documentation_date"
	// FieldGeometry holds the string denoting the geometry field in the database.
	FieldGeometry = "geometry"
	// EdgeCulturalAffiliation holds the string denoting the cultural_affiliation edge name in mutations.
	EdgeCulturalAffiliation = "cultural_affiliation"
	// EdgeModel holds the string denoting the model edge name in mutations.
	EdgeModel = "model"
	// EdgeMound holds the string denoting the mound edge name in mutations.
	EdgeMound = "mound"
	// EdgePublications holds the string denoting the publications edge name in mutations.
	EdgePublications = "publications"
	// EdgeTechniques holds the string denoting the techniques edge name in mutations.
	EdgeTechniques = "techniques"
	// EdgeRegion holds the string denoting the region edge name in mutations.
	EdgeRegion = "region"
	// EdgeAccountingDocumentationAddress holds the string denoting the accounting_documentation_address edge name in mutations.
	EdgeAccountingDocumentationAddress = "accounting_documentation_address"
	// EdgeAccountingDocumentationAuthor holds the string denoting the accounting_documentation_author edge name in mutations.
	EdgeAccountingDocumentationAuthor = "accounting_documentation_author"
	// EdgeCollection holds the string denoting the collection edge name in mutations.
	EdgeCollection = "collection"
	// Table holds the table name of the petroglyph in the database.
	Table = "petroglyphs"
	// CulturalAffiliationTable is the table that holds the cultural_affiliation relation/edge.
	CulturalAffiliationTable = "petroglyphs"
	// CulturalAffiliationInverseTable is the table name for the Culture entity.
	// It exists in this package in order to avoid circular dependency with the "culture" package.
	CulturalAffiliationInverseTable = "cultures"
	// CulturalAffiliationColumn is the table column denoting the cultural_affiliation relation/edge.
	CulturalAffiliationColumn = "culture_petroglyphs"
	// ModelTable is the table that holds the model relation/edge.
	ModelTable = "petroglyphs"
	// ModelInverseTable is the table name for the Model entity.
	// It exists in this package in order to avoid circular dependency with the "model" package.
	ModelInverseTable = "models"
	// ModelColumn is the table column denoting the model relation/edge.
	ModelColumn = "model_petroglyphs"
	// MoundTable is the table that holds the mound relation/edge.
	MoundTable = "petroglyphs"
	// MoundInverseTable is the table name for the Mound entity.
	// It exists in this package in order to avoid circular dependency with the "mound" package.
	MoundInverseTable = "mounds"
	// MoundColumn is the table column denoting the mound relation/edge.
	MoundColumn = "mound_petroglyphs"
	// PublicationsTable is the table that holds the publications relation/edge. The primary key declared below.
	PublicationsTable = "publication_petroglyphs"
	// PublicationsInverseTable is the table name for the Publication entity.
	// It exists in this package in order to avoid circular dependency with the "publication" package.
	PublicationsInverseTable = "publications"
	// TechniquesTable is the table that holds the techniques relation/edge. The primary key declared below.
	TechniquesTable = "technique_petroglyphs"
	// TechniquesInverseTable is the table name for the Technique entity.
	// It exists in this package in order to avoid circular dependency with the "technique" package.
	TechniquesInverseTable = "techniques"
	// RegionTable is the table that holds the region relation/edge.
	RegionTable = "petroglyphs"
	// RegionInverseTable is the table name for the Region entity.
	// It exists in this package in order to avoid circular dependency with the "region" package.
	RegionInverseTable = "regions"
	// RegionColumn is the table column denoting the region relation/edge.
	RegionColumn = "region_petroglyphs"
	// AccountingDocumentationAddressTable is the table that holds the accounting_documentation_address relation/edge.
	AccountingDocumentationAddressTable = "petroglyphs"
	// AccountingDocumentationAddressInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	AccountingDocumentationAddressInverseTable = "locations"
	// AccountingDocumentationAddressColumn is the table column denoting the accounting_documentation_address relation/edge.
	AccountingDocumentationAddressColumn = "location_petroglyphs_accounting_documentation"
	// AccountingDocumentationAuthorTable is the table that holds the accounting_documentation_author relation/edge.
	AccountingDocumentationAuthorTable = "petroglyphs"
	// AccountingDocumentationAuthorInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	AccountingDocumentationAuthorInverseTable = "persons"
	// AccountingDocumentationAuthorColumn is the table column denoting the accounting_documentation_author relation/edge.
	AccountingDocumentationAuthorColumn = "person_petroglyphs_accounting_documentation"
	// CollectionTable is the table that holds the collection relation/edge.
	CollectionTable = "petroglyphs"
	// CollectionInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionInverseTable = "collections"
	// CollectionColumn is the table column denoting the collection relation/edge.
	CollectionColumn = "collection_petroglyphs"
)

// Columns holds all SQL columns for petroglyph fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDating,
	FieldDatingStart,
	FieldDatingEnd,
	FieldDisplayName,
	FieldAbbreviation,
	FieldDescription,
	FieldExternalLink,
	FieldStatus,
	FieldPrimaryImageURL,
	FieldAdditionalImagesUrls,
	FieldHeight,
	FieldWidth,
	FieldLength,
	FieldDepth,
	FieldDiameter,
	FieldWeight,
	FieldDimensions,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldNumber,
	FieldOrientation,
	FieldPosition,
	FieldGeometricShape,
	FieldPlanePreservation,
	FieldPhotoCode,
	FieldAccountingDocumentationInformation,
	FieldAccountingDocumentationDate,
	FieldGeometry,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "petroglyphs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"collection_petroglyphs",
	"culture_petroglyphs",
	"location_petroglyphs_accounting_documentation",
	"model_petroglyphs",
	"mound_petroglyphs",
	"person_petroglyphs_accounting_documentation",
	"region_petroglyphs",
}

var (
	// PublicationsPrimaryKey and PublicationsColumn2 are the table columns denoting the
	// primary key for the publications relation (M2M).
	PublicationsPrimaryKey = []string{"publication_id", "petroglyph_id"}
	// TechniquesPrimaryKey and TechniquesColumn2 are the table columns denoting the
	// primary key for the techniques relation (M2M).
	TechniquesPrimaryKey = []string{"technique_id", "petroglyph_id"}
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
	Hooks        [3]ent.Hook
	Interceptors [1]ent.Interceptor
	Policy       ent.Policy
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
		return fmt.Errorf("petroglyph: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Petroglyph queries.
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

// ByDating orders the results by the dating field.
func ByDating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDating, opts...).ToFunc()
}

// ByDatingStart orders the results by the dating_start field.
func ByDatingStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDatingStart, opts...).ToFunc()
}

// ByDatingEnd orders the results by the dating_end field.
func ByDatingEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDatingEnd, opts...).ToFunc()
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

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByWidth orders the results by the width field.
func ByWidth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWidth, opts...).ToFunc()
}

// ByLength orders the results by the length field.
func ByLength(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLength, opts...).ToFunc()
}

// ByDepth orders the results by the depth field.
func ByDepth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepth, opts...).ToFunc()
}

// ByDiameter orders the results by the diameter field.
func ByDiameter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDiameter, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByDimensions orders the results by the dimensions field.
func ByDimensions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDimensions, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByOrientation orders the results by the orientation field.
func ByOrientation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrientation, opts...).ToFunc()
}

// ByPosition orders the results by the position field.
func ByPosition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPosition, opts...).ToFunc()
}

// ByGeometricShape orders the results by the geometric_shape field.
func ByGeometricShape(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeometricShape, opts...).ToFunc()
}

// ByPlanePreservation orders the results by the plane_preservation field.
func ByPlanePreservation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlanePreservation, opts...).ToFunc()
}

// ByPhotoCode orders the results by the photo_code field.
func ByPhotoCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoCode, opts...).ToFunc()
}

// ByAccountingDocumentationInformation orders the results by the accounting_documentation_information field.
func ByAccountingDocumentationInformation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountingDocumentationInformation, opts...).ToFunc()
}

// ByAccountingDocumentationDate orders the results by the accounting_documentation_date field.
func ByAccountingDocumentationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountingDocumentationDate, opts...).ToFunc()
}

// ByGeometry orders the results by the geometry field.
func ByGeometry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeometry, opts...).ToFunc()
}

// ByCulturalAffiliationField orders the results by cultural_affiliation field.
func ByCulturalAffiliationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCulturalAffiliationStep(), sql.OrderByField(field, opts...))
	}
}

// ByModelField orders the results by model field.
func ByModelField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModelStep(), sql.OrderByField(field, opts...))
	}
}

// ByMoundField orders the results by mound field.
func ByMoundField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMoundStep(), sql.OrderByField(field, opts...))
	}
}

// ByPublicationsCount orders the results by publications count.
func ByPublicationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPublicationsStep(), opts...)
	}
}

// ByPublications orders the results by publications terms.
func ByPublications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPublicationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTechniquesCount orders the results by techniques count.
func ByTechniquesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTechniquesStep(), opts...)
	}
}

// ByTechniques orders the results by techniques terms.
func ByTechniques(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTechniquesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRegionField orders the results by region field.
func ByRegionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRegionStep(), sql.OrderByField(field, opts...))
	}
}

// ByAccountingDocumentationAddressField orders the results by accounting_documentation_address field.
func ByAccountingDocumentationAddressField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountingDocumentationAddressStep(), sql.OrderByField(field, opts...))
	}
}

// ByAccountingDocumentationAuthorField orders the results by accounting_documentation_author field.
func ByAccountingDocumentationAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountingDocumentationAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByCollectionField orders the results by collection field.
func ByCollectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCollectionStep(), sql.OrderByField(field, opts...))
	}
}
func newCulturalAffiliationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CulturalAffiliationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CulturalAffiliationTable, CulturalAffiliationColumn),
	)
}
func newModelStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ModelInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ModelTable, ModelColumn),
	)
}
func newMoundStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MoundInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MoundTable, MoundColumn),
	)
}
func newPublicationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PublicationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PublicationsTable, PublicationsPrimaryKey...),
	)
}
func newTechniquesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TechniquesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TechniquesTable, TechniquesPrimaryKey...),
	)
}
func newRegionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RegionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RegionTable, RegionColumn),
	)
}
func newAccountingDocumentationAddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountingDocumentationAddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountingDocumentationAddressTable, AccountingDocumentationAddressColumn),
	)
}
func newAccountingDocumentationAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountingDocumentationAuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountingDocumentationAuthorTable, AccountingDocumentationAuthorColumn),
	)
}
func newCollectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CollectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CollectionTable, CollectionColumn),
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
