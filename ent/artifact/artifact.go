// Code generated by ent, DO NOT EDIT.

package artifact

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
	// Label holds the string label denoting the artifact type in the database.
	Label = "artifact"
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
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldDating holds the string denoting the dating field in the database.
	FieldDating = "dating"
	// FieldDatingStart holds the string denoting the dating_start field in the database.
	FieldDatingStart = "dating_start"
	// FieldDatingEnd holds the string denoting the dating_end field in the database.
	FieldDatingEnd = "dating_end"
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
	// FieldChemicalComposition holds the string denoting the chemical_composition field in the database.
	FieldChemicalComposition = "chemical_composition"
	// FieldGoskatalogNumber holds the string denoting the goskatalog_number field in the database.
	FieldGoskatalogNumber = "goskatalog_number"
	// FieldInventoryNumber holds the string denoting the inventory_number field in the database.
	FieldInventoryNumber = "inventory_number"
	// FieldTypology holds the string denoting the typology field in the database.
	FieldTypology = "typology"
	// FieldAdmissionDate holds the string denoting the admission_date field in the database.
	FieldAdmissionDate = "admission_date"
	// EdgeAuthors holds the string denoting the authors edge name in mutations.
	EdgeAuthors = "authors"
	// EdgeMediums holds the string denoting the mediums edge name in mutations.
	EdgeMediums = "mediums"
	// EdgeTechniques holds the string denoting the techniques edge name in mutations.
	EdgeTechniques = "techniques"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgePublications holds the string denoting the publications edge name in mutations.
	EdgePublications = "publications"
	// EdgeCulturalAffiliation holds the string denoting the cultural_affiliation edge name in mutations.
	EdgeCulturalAffiliation = "cultural_affiliation"
	// EdgeMonument holds the string denoting the monument edge name in mutations.
	EdgeMonument = "monument"
	// EdgeModel holds the string denoting the model edge name in mutations.
	EdgeModel = "model"
	// EdgeSet holds the string denoting the set edge name in mutations.
	EdgeSet = "set"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeCollection holds the string denoting the collection edge name in mutations.
	EdgeCollection = "collection"
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
	// Table holds the table name of the artifact in the database.
	Table = "artifacts"
	// AuthorsTable is the table that holds the authors relation/edge. The primary key declared below.
	AuthorsTable = "person_artifacts"
	// AuthorsInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	AuthorsInverseTable = "persons"
	// MediumsTable is the table that holds the mediums relation/edge. The primary key declared below.
	MediumsTable = "medium_artifacts"
	// MediumsInverseTable is the table name for the Medium entity.
	// It exists in this package in order to avoid circular dependency with the "medium" package.
	MediumsInverseTable = "media"
	// TechniquesTable is the table that holds the techniques relation/edge. The primary key declared below.
	TechniquesTable = "technique_artifacts"
	// TechniquesInverseTable is the table name for the Technique entity.
	// It exists in this package in order to avoid circular dependency with the "technique" package.
	TechniquesInverseTable = "techniques"
	// ProjectsTable is the table that holds the projects relation/edge. The primary key declared below.
	ProjectsTable = "project_artifacts"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// PublicationsTable is the table that holds the publications relation/edge. The primary key declared below.
	PublicationsTable = "publication_artifacts"
	// PublicationsInverseTable is the table name for the Publication entity.
	// It exists in this package in order to avoid circular dependency with the "publication" package.
	PublicationsInverseTable = "publications"
	// CulturalAffiliationTable is the table that holds the cultural_affiliation relation/edge.
	CulturalAffiliationTable = "artifacts"
	// CulturalAffiliationInverseTable is the table name for the Culture entity.
	// It exists in this package in order to avoid circular dependency with the "culture" package.
	CulturalAffiliationInverseTable = "cultures"
	// CulturalAffiliationColumn is the table column denoting the cultural_affiliation relation/edge.
	CulturalAffiliationColumn = "culture_artifacts"
	// MonumentTable is the table that holds the monument relation/edge.
	MonumentTable = "artifacts"
	// MonumentInverseTable is the table name for the Monument entity.
	// It exists in this package in order to avoid circular dependency with the "monument" package.
	MonumentInverseTable = "monuments"
	// MonumentColumn is the table column denoting the monument relation/edge.
	MonumentColumn = "monument_artifacts"
	// ModelTable is the table that holds the model relation/edge.
	ModelTable = "artifacts"
	// ModelInverseTable is the table name for the Model entity.
	// It exists in this package in order to avoid circular dependency with the "model" package.
	ModelInverseTable = "models"
	// ModelColumn is the table column denoting the model relation/edge.
	ModelColumn = "model_artifacts"
	// SetTable is the table that holds the set relation/edge.
	SetTable = "artifacts"
	// SetInverseTable is the table name for the Set entity.
	// It exists in this package in order to avoid circular dependency with the "set" package.
	SetInverseTable = "sets"
	// SetColumn is the table column denoting the set relation/edge.
	SetColumn = "set_artifacts"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "artifacts"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "locations"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "location_artifacts"
	// CollectionTable is the table that holds the collection relation/edge.
	CollectionTable = "artifacts"
	// CollectionInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionInverseTable = "collections"
	// CollectionColumn is the table column denoting the collection relation/edge.
	CollectionColumn = "collection_artifacts"
	// LicenseTable is the table that holds the license relation/edge.
	LicenseTable = "artifacts"
	// LicenseInverseTable is the table name for the License entity.
	// It exists in this package in order to avoid circular dependency with the "license" package.
	LicenseInverseTable = "licenses"
	// LicenseColumn is the table column denoting the license relation/edge.
	LicenseColumn = "license_artifacts"
	// CountryTable is the table that holds the country relation/edge.
	CountryTable = "artifacts"
	// CountryInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	CountryInverseTable = "countries"
	// CountryColumn is the table column denoting the country relation/edge.
	CountryColumn = "country_artifacts"
	// SettlementTable is the table that holds the settlement relation/edge.
	SettlementTable = "artifacts"
	// SettlementInverseTable is the table name for the Settlement entity.
	// It exists in this package in order to avoid circular dependency with the "settlement" package.
	SettlementInverseTable = "settlements"
	// SettlementColumn is the table column denoting the settlement relation/edge.
	SettlementColumn = "settlement_artifacts"
	// DistrictTable is the table that holds the district relation/edge.
	DistrictTable = "artifacts"
	// DistrictInverseTable is the table name for the District entity.
	// It exists in this package in order to avoid circular dependency with the "district" package.
	DistrictInverseTable = "districts"
	// DistrictColumn is the table column denoting the district relation/edge.
	DistrictColumn = "district_artifacts"
	// RegionTable is the table that holds the region relation/edge.
	RegionTable = "artifacts"
	// RegionInverseTable is the table name for the Region entity.
	// It exists in this package in order to avoid circular dependency with the "region" package.
	RegionInverseTable = "regions"
	// RegionColumn is the table column denoting the region relation/edge.
	RegionColumn = "region_artifacts"
)

// Columns holds all SQL columns for artifact fields.
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
	FieldDeletedAt,
	FieldDeletedBy,
	FieldDating,
	FieldDatingStart,
	FieldDatingEnd,
	FieldHeight,
	FieldWidth,
	FieldLength,
	FieldDepth,
	FieldDiameter,
	FieldWeight,
	FieldDimensions,
	FieldChemicalComposition,
	FieldGoskatalogNumber,
	FieldInventoryNumber,
	FieldTypology,
	FieldAdmissionDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "artifacts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"collection_artifacts",
	"country_artifacts",
	"culture_artifacts",
	"district_artifacts",
	"license_artifacts",
	"location_artifacts",
	"model_artifacts",
	"monument_artifacts",
	"region_artifacts",
	"set_artifacts",
	"settlement_artifacts",
}

var (
	// AuthorsPrimaryKey and AuthorsColumn2 are the table columns denoting the
	// primary key for the authors relation (M2M).
	AuthorsPrimaryKey = []string{"person_id", "artifact_id"}
	// MediumsPrimaryKey and MediumsColumn2 are the table columns denoting the
	// primary key for the mediums relation (M2M).
	MediumsPrimaryKey = []string{"medium_id", "artifact_id"}
	// TechniquesPrimaryKey and TechniquesColumn2 are the table columns denoting the
	// primary key for the techniques relation (M2M).
	TechniquesPrimaryKey = []string{"technique_id", "artifact_id"}
	// ProjectsPrimaryKey and ProjectsColumn2 are the table columns denoting the
	// primary key for the projects relation (M2M).
	ProjectsPrimaryKey = []string{"project_id", "artifact_id"}
	// PublicationsPrimaryKey and PublicationsColumn2 are the table columns denoting the
	// primary key for the publications relation (M2M).
	PublicationsPrimaryKey = []string{"publication_id", "artifact_id"}
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
		return fmt.Errorf("artifact: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Artifact queries.
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

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
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

// ByChemicalComposition orders the results by the chemical_composition field.
func ByChemicalComposition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChemicalComposition, opts...).ToFunc()
}

// ByGoskatalogNumber orders the results by the goskatalog_number field.
func ByGoskatalogNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGoskatalogNumber, opts...).ToFunc()
}

// ByInventoryNumber orders the results by the inventory_number field.
func ByInventoryNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInventoryNumber, opts...).ToFunc()
}

// ByTypology orders the results by the typology field.
func ByTypology(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTypology, opts...).ToFunc()
}

// ByAdmissionDate orders the results by the admission_date field.
func ByAdmissionDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdmissionDate, opts...).ToFunc()
}

// ByAuthorsCount orders the results by authors count.
func ByAuthorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorsStep(), opts...)
	}
}

// ByAuthors orders the results by authors terms.
func ByAuthors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByProjectsCount orders the results by projects count.
func ByProjectsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProjectsStep(), opts...)
	}
}

// ByProjects orders the results by projects terms.
func ByProjects(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByCulturalAffiliationField orders the results by cultural_affiliation field.
func ByCulturalAffiliationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCulturalAffiliationStep(), sql.OrderByField(field, opts...))
	}
}

// ByMonumentField orders the results by monument field.
func ByMonumentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMonumentStep(), sql.OrderByField(field, opts...))
	}
}

// ByModelField orders the results by model field.
func ByModelField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModelStep(), sql.OrderByField(field, opts...))
	}
}

// BySetField orders the results by set field.
func BySetField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSetStep(), sql.OrderByField(field, opts...))
	}
}

// ByLocationField orders the results by location field.
func ByLocationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLocationStep(), sql.OrderByField(field, opts...))
	}
}

// ByCollectionField orders the results by collection field.
func ByCollectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCollectionStep(), sql.OrderByField(field, opts...))
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
func newAuthorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AuthorsTable, AuthorsPrimaryKey...),
	)
}
func newMediumsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediumsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, MediumsTable, MediumsPrimaryKey...),
	)
}
func newTechniquesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TechniquesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TechniquesTable, TechniquesPrimaryKey...),
	)
}
func newProjectsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProjectsTable, ProjectsPrimaryKey...),
	)
}
func newPublicationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PublicationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PublicationsTable, PublicationsPrimaryKey...),
	)
}
func newCulturalAffiliationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CulturalAffiliationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CulturalAffiliationTable, CulturalAffiliationColumn),
	)
}
func newMonumentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MonumentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MonumentTable, MonumentColumn),
	)
}
func newModelStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ModelInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ModelTable, ModelColumn),
	)
}
func newSetStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SetInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SetTable, SetColumn),
	)
}
func newLocationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LocationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
	)
}
func newCollectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CollectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CollectionTable, CollectionColumn),
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
