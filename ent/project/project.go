// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
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
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldExternalLinks holds the string denoting the external_links field in the database.
	FieldExternalLinks = "external_links"
	// EdgeArtifacts holds the string denoting the artifacts edge name in mutations.
	EdgeArtifacts = "artifacts"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// ArtifactsTable is the table that holds the artifacts relation/edge. The primary key declared below.
	ArtifactsTable = "project_artifacts"
	// ArtifactsInverseTable is the table name for the Artifact entity.
	// It exists in this package in order to avoid circular dependency with the "artifact" package.
	ArtifactsInverseTable = "artifacts"
	// TeamTable is the table that holds the team relation/edge. The primary key declared below.
	TeamTable = "person_projects"
	// TeamInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	TeamInverseTable = "persons"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDisplayName,
	FieldDescription,
	FieldExternalLinks,
}

var (
	// ArtifactsPrimaryKey and ArtifactsColumn2 are the table columns denoting the
	// primary key for the artifacts relation (M2M).
	ArtifactsPrimaryKey = []string{"project_id", "artifact_id"}
	// TeamPrimaryKey and TeamColumn2 are the table columns denoting the
	// primary key for the team relation (M2M).
	TeamPrimaryKey = []string{"person_id", "project_id"}
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

// OrderOption defines the ordering options for the Project queries.
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

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
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

// ByTeamCount orders the results by team count.
func ByTeamCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamStep(), opts...)
	}
}

// ByTeam orders the results by team terms.
func ByTeam(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newArtifactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtifactsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ArtifactsTable, ArtifactsPrimaryKey...),
	)
}
func newTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TeamTable, TeamPrimaryKey...),
	)
}
