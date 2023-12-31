// Code generated by ent, DO NOT EDIT.

package dendrochronologicalanalysis

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the dendrochronologicalanalysis type in the database.
	Label = "dendrochronological_analysis"
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
	// FieldStartYear holds the string denoting the start_year field in the database.
	FieldStartYear = "start_year"
	// FieldEndYear holds the string denoting the end_year field in the database.
	FieldEndYear = "end_year"
	// FieldNumberOfRings holds the string denoting the number_of_rings field in the database.
	FieldNumberOfRings = "number_of_rings"
	// FieldCoefficientCorrelation holds the string denoting the coefficient_correlation field in the database.
	FieldCoefficientCorrelation = "coefficient_correlation"
	// FieldStandardDeviation holds the string denoting the standard_deviation field in the database.
	FieldStandardDeviation = "standard_deviation"
	// FieldSensitivity holds the string denoting the sensitivity field in the database.
	FieldSensitivity = "sensitivity"
	// FieldSamplingLocation holds the string denoting the sampling_location field in the database.
	FieldSamplingLocation = "sampling_location"
	// EdgeDendrochronology holds the string denoting the dendrochronology edge name in mutations.
	EdgeDendrochronology = "dendrochronology"
	// Table holds the table name of the dendrochronologicalanalysis in the database.
	Table = "dendrochronological_analyses"
	// DendrochronologyTable is the table that holds the dendrochronology relation/edge.
	DendrochronologyTable = "dendrochronological_analyses"
	// DendrochronologyInverseTable is the table name for the Dendrochronology entity.
	// It exists in this package in order to avoid circular dependency with the "dendrochronology" package.
	DendrochronologyInverseTable = "dendrochronologies"
	// DendrochronologyColumn is the table column denoting the dendrochronology relation/edge.
	DendrochronologyColumn = "dendrochronological_analysis_dendrochronology"
)

// Columns holds all SQL columns for dendrochronologicalanalysis fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDisplayName,
	FieldStartYear,
	FieldEndYear,
	FieldNumberOfRings,
	FieldCoefficientCorrelation,
	FieldStandardDeviation,
	FieldSensitivity,
	FieldSamplingLocation,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "dendrochronological_analyses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"dendrochronological_analysis_dendrochronology",
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

// OrderOption defines the ordering options for the DendrochronologicalAnalysis queries.
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

// ByStartYear orders the results by the start_year field.
func ByStartYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartYear, opts...).ToFunc()
}

// ByEndYear orders the results by the end_year field.
func ByEndYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndYear, opts...).ToFunc()
}

// ByNumberOfRings orders the results by the number_of_rings field.
func ByNumberOfRings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumberOfRings, opts...).ToFunc()
}

// ByCoefficientCorrelation orders the results by the coefficient_correlation field.
func ByCoefficientCorrelation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoefficientCorrelation, opts...).ToFunc()
}

// ByStandardDeviation orders the results by the standard_deviation field.
func ByStandardDeviation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStandardDeviation, opts...).ToFunc()
}

// BySensitivity orders the results by the sensitivity field.
func BySensitivity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSensitivity, opts...).ToFunc()
}

// BySamplingLocation orders the results by the sampling_location field.
func BySamplingLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSamplingLocation, opts...).ToFunc()
}

// ByDendrochronologyField orders the results by dendrochronology field.
func ByDendrochronologyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDendrochronologyStep(), sql.OrderByField(field, opts...))
	}
}
func newDendrochronologyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DendrochronologyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DendrochronologyTable, DendrochronologyColumn),
	)
}
