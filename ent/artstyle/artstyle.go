// Code generated by ent, DO NOT EDIT.

package artstyle

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the artstyle type in the database.
	Label = "art_style"
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
	// EdgeArt holds the string denoting the art edge name in mutations.
	EdgeArt = "art"
	// Table holds the table name of the artstyle in the database.
	Table = "art_styles"
	// ArtTable is the table that holds the art relation/edge. The primary key declared below.
	ArtTable = "art_style_art"
	// ArtInverseTable is the table name for the Art entity.
	// It exists in this package in order to avoid circular dependency with the "art" package.
	ArtInverseTable = "arts"
)

// Columns holds all SQL columns for artstyle fields.
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
}

var (
	// ArtPrimaryKey and ArtColumn2 are the table columns denoting the
	// primary key for the art relation (M2M).
	ArtPrimaryKey = []string{"art_style_id", "art_id"}
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
	Hooks  [2]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the ArtStyle queries.
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
func newArtStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ArtTable, ArtPrimaryKey...),
	)
}
