// Code generated by ent, DO NOT EDIT.

package personalcollection

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the personalcollection type in the database.
	Label = "personal_collection"
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
	// FieldIsPublic holds the string denoting the is_public field in the database.
	FieldIsPublic = "is_public"
	// EdgeArt holds the string denoting the art edge name in mutations.
	EdgeArt = "art"
	// EdgeArtifacts holds the string denoting the artifacts edge name in mutations.
	EdgeArtifacts = "artifacts"
	// EdgePetroglyphs holds the string denoting the petroglyphs edge name in mutations.
	EdgePetroglyphs = "petroglyphs"
	// EdgeBooks holds the string denoting the books edge name in mutations.
	EdgeBooks = "books"
	// EdgeProtectedAreaPictures holds the string denoting the protected_area_pictures edge name in mutations.
	EdgeProtectedAreaPictures = "protected_area_pictures"
	// Table holds the table name of the personalcollection in the database.
	Table = "personal_collections"
	// ArtTable is the table that holds the art relation/edge. The primary key declared below.
	ArtTable = "personal_collection_art"
	// ArtInverseTable is the table name for the Art entity.
	// It exists in this package in order to avoid circular dependency with the "art" package.
	ArtInverseTable = "arts"
	// ArtifactsTable is the table that holds the artifacts relation/edge. The primary key declared below.
	ArtifactsTable = "personal_collection_artifacts"
	// ArtifactsInverseTable is the table name for the Artifact entity.
	// It exists in this package in order to avoid circular dependency with the "artifact" package.
	ArtifactsInverseTable = "artifacts"
	// PetroglyphsTable is the table that holds the petroglyphs relation/edge. The primary key declared below.
	PetroglyphsTable = "personal_collection_petroglyphs"
	// PetroglyphsInverseTable is the table name for the Petroglyph entity.
	// It exists in this package in order to avoid circular dependency with the "petroglyph" package.
	PetroglyphsInverseTable = "petroglyphs"
	// BooksTable is the table that holds the books relation/edge. The primary key declared below.
	BooksTable = "personal_collection_books"
	// BooksInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BooksInverseTable = "books"
	// ProtectedAreaPicturesTable is the table that holds the protected_area_pictures relation/edge. The primary key declared below.
	ProtectedAreaPicturesTable = "personal_collection_protected_area_pictures"
	// ProtectedAreaPicturesInverseTable is the table name for the ProtectedAreaPicture entity.
	// It exists in this package in order to avoid circular dependency with the "protectedareapicture" package.
	ProtectedAreaPicturesInverseTable = "protected_area_pictures"
)

// Columns holds all SQL columns for personalcollection fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDisplayName,
	FieldIsPublic,
}

var (
	// ArtPrimaryKey and ArtColumn2 are the table columns denoting the
	// primary key for the art relation (M2M).
	ArtPrimaryKey = []string{"personal_collection_id", "art_id"}
	// ArtifactsPrimaryKey and ArtifactsColumn2 are the table columns denoting the
	// primary key for the artifacts relation (M2M).
	ArtifactsPrimaryKey = []string{"personal_collection_id", "artifact_id"}
	// PetroglyphsPrimaryKey and PetroglyphsColumn2 are the table columns denoting the
	// primary key for the petroglyphs relation (M2M).
	PetroglyphsPrimaryKey = []string{"personal_collection_id", "petroglyph_id"}
	// BooksPrimaryKey and BooksColumn2 are the table columns denoting the
	// primary key for the books relation (M2M).
	BooksPrimaryKey = []string{"personal_collection_id", "book_id"}
	// ProtectedAreaPicturesPrimaryKey and ProtectedAreaPicturesColumn2 are the table columns denoting the
	// primary key for the protected_area_pictures relation (M2M).
	ProtectedAreaPicturesPrimaryKey = []string{"personal_collection_id", "protected_area_picture_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// DefaultIsPublic holds the default value on creation for the "is_public" field.
	DefaultIsPublic bool
)

// OrderOption defines the ordering options for the PersonalCollection queries.
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

// ByIsPublic orders the results by the is_public field.
func ByIsPublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPublic, opts...).ToFunc()
}

// ByArtCount orders the results by art count.
func ByArtCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArtStep(), opts...)
	}
}

// ByArt orders the results by art terms.
func ByArt(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArtStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByArtifactsCount orders the results by artifacts count.
func ByArtifactsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArtifactsStep(), opts...)
	}
}

// ByArtifacts orders the results by artifacts terms.
func ByArtifacts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArtifactsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPetroglyphsCount orders the results by petroglyphs count.
func ByPetroglyphsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPetroglyphsStep(), opts...)
	}
}

// ByPetroglyphs orders the results by petroglyphs terms.
func ByPetroglyphs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPetroglyphsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBooksCount orders the results by books count.
func ByBooksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBooksStep(), opts...)
	}
}

// ByBooks orders the results by books terms.
func ByBooks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBooksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProtectedAreaPicturesCount orders the results by protected_area_pictures count.
func ByProtectedAreaPicturesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProtectedAreaPicturesStep(), opts...)
	}
}

// ByProtectedAreaPictures orders the results by protected_area_pictures terms.
func ByProtectedAreaPictures(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProtectedAreaPicturesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newArtStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ArtTable, ArtPrimaryKey...),
	)
}
func newArtifactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtifactsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ArtifactsTable, ArtifactsPrimaryKey...),
	)
}
func newPetroglyphsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PetroglyphsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PetroglyphsTable, PetroglyphsPrimaryKey...),
	)
}
func newBooksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BooksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, BooksTable, BooksPrimaryKey...),
	)
}
func newProtectedAreaPicturesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProtectedAreaPicturesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProtectedAreaPicturesTable, ProtectedAreaPicturesPrimaryKey...),
	)
}