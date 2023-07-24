// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
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
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPhoneNumbers holds the string denoting the phone_numbers field in the database.
	FieldPhoneNumbers = "phone_numbers"
	// FieldEmails holds the string denoting the emails field in the database.
	FieldEmails = "emails"
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
	// FieldPreviousNames holds the string denoting the previous_names field in the database.
	FieldPreviousNames = "previous_names"
	// FieldIsInAConsortium holds the string denoting the is_in_a_consortium field in the database.
	FieldIsInAConsortium = "is_in_a_consortium"
	// FieldConsortiumDocumentURL holds the string denoting the consortium_document_url field in the database.
	FieldConsortiumDocumentURL = "consortium_document_url"
	// EdgePeople holds the string denoting the people edge name in mutations.
	EdgePeople = "people"
	// EdgeHolder holds the string denoting the holder edge name in mutations.
	EdgeHolder = "holder"
	// EdgeOrganizationType holds the string denoting the organization_type edge name in mutations.
	EdgeOrganizationType = "organization_type"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// PeopleTable is the table that holds the people relation/edge.
	PeopleTable = "persons"
	// PeopleInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	PeopleInverseTable = "persons"
	// PeopleColumn is the table column denoting the people relation/edge.
	PeopleColumn = "organization_people"
	// HolderTable is the table that holds the holder relation/edge.
	HolderTable = "organizations"
	// HolderInverseTable is the table name for the Holder entity.
	// It exists in this package in order to avoid circular dependency with the "holder" package.
	HolderInverseTable = "holders"
	// HolderColumn is the table column denoting the holder relation/edge.
	HolderColumn = "holder_organization"
	// OrganizationTypeTable is the table that holds the organization_type relation/edge.
	OrganizationTypeTable = "organizations"
	// OrganizationTypeInverseTable is the table name for the OrganizationType entity.
	// It exists in this package in order to avoid circular dependency with the "organizationtype" package.
	OrganizationTypeInverseTable = "organization_types"
	// OrganizationTypeColumn is the table column denoting the organization_type relation/edge.
	OrganizationTypeColumn = "organization_type_organizations"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldAddress,
	FieldPhoneNumbers,
	FieldEmails,
	FieldDisplayName,
	FieldAbbreviation,
	FieldDescription,
	FieldExternalLinks,
	FieldPrimaryImageURL,
	FieldAdditionalImagesUrls,
	FieldPreviousNames,
	FieldIsInAConsortium,
	FieldConsortiumDocumentURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "organizations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"holder_organization",
	"organization_type_organizations",
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

// OrderOption defines the ordering options for the Organization queries.
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

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
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

// ByIsInAConsortium orders the results by the is_in_a_consortium field.
func ByIsInAConsortium(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsInAConsortium, opts...).ToFunc()
}

// ByConsortiumDocumentURL orders the results by the consortium_document_url field.
func ByConsortiumDocumentURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConsortiumDocumentURL, opts...).ToFunc()
}

// ByPeopleCount orders the results by people count.
func ByPeopleCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPeopleStep(), opts...)
	}
}

// ByPeople orders the results by people terms.
func ByPeople(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPeopleStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByHolderField orders the results by holder field.
func ByHolderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHolderStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrganizationTypeField orders the results by organization_type field.
func ByOrganizationTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationTypeStep(), sql.OrderByField(field, opts...))
	}
}
func newPeopleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PeopleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PeopleTable, PeopleColumn),
	)
}
func newHolderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HolderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, HolderTable, HolderColumn),
	)
}
func newOrganizationTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTypeTable, OrganizationTypeColumn),
	)
}
