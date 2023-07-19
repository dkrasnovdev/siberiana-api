// Code generated by ent, DO NOT EDIT.

package person

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the person type in the database.
	Label = "person"
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
	// EdgeArtifacts holds the string denoting the artifacts edge name in mutations.
	EdgeArtifacts = "artifacts"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgePublications holds the string denoting the publications edge name in mutations.
	EdgePublications = "publications"
	// EdgeHolder holds the string denoting the holder edge name in mutations.
	EdgeHolder = "holder"
	// Table holds the table name of the person in the database.
	Table = "persons"
	// ArtifactsTable is the table that holds the artifacts relation/edge. The primary key declared below.
	ArtifactsTable = "person_artifacts"
	// ArtifactsInverseTable is the table name for the Artifact entity.
	// It exists in this package in order to avoid circular dependency with the "artifact" package.
	ArtifactsInverseTable = "artifacts"
	// ProjectsTable is the table that holds the projects relation/edge. The primary key declared below.
	ProjectsTable = "person_projects"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// PublicationsTable is the table that holds the publications relation/edge. The primary key declared below.
	PublicationsTable = "person_publications"
	// PublicationsInverseTable is the table name for the Publication entity.
	// It exists in this package in order to avoid circular dependency with the "publication" package.
	PublicationsInverseTable = "publications"
	// HolderTable is the table that holds the holder relation/edge.
	HolderTable = "persons"
	// HolderInverseTable is the table name for the Holder entity.
	// It exists in this package in order to avoid circular dependency with the "holder" package.
	HolderInverseTable = "holders"
	// HolderColumn is the table column denoting the holder relation/edge.
	HolderColumn = "holder_person"
)

// Columns holds all SQL columns for person fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDisplayName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "persons"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"holder_person",
}

var (
	// ArtifactsPrimaryKey and ArtifactsColumn2 are the table columns denoting the
	// primary key for the artifacts relation (M2M).
	ArtifactsPrimaryKey = []string{"person_id", "artifact_id"}
	// ProjectsPrimaryKey and ProjectsColumn2 are the table columns denoting the
	// primary key for the projects relation (M2M).
	ProjectsPrimaryKey = []string{"person_id", "project_id"}
	// PublicationsPrimaryKey and PublicationsColumn2 are the table columns denoting the
	// primary key for the publications relation (M2M).
	PublicationsPrimaryKey = []string{"person_id", "publication_id"}
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

// OrderOption defines the ordering options for the Person queries.
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

// ByHolderField orders the results by holder field.
func ByHolderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHolderStep(), sql.OrderByField(field, opts...))
	}
}
func newArtifactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtifactsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ArtifactsTable, ArtifactsPrimaryKey...),
	)
}
func newProjectsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProjectsTable, ProjectsPrimaryKey...),
	)
}
func newPublicationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PublicationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PublicationsTable, PublicationsPrimaryKey...),
	)
}
func newHolderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HolderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, HolderTable, HolderColumn),
	)
}
