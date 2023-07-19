// Code generated by ent, DO NOT EDIT.

package auditlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the auditlog type in the database.
	Label = "audit_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTable holds the string denoting the table field in the database.
	FieldTable = "table"
	// FieldRefID holds the string denoting the ref_id field in the database.
	FieldRefID = "ref_id"
	// FieldOperation holds the string denoting the operation field in the database.
	FieldOperation = "operation"
	// FieldChanges holds the string denoting the changes field in the database.
	FieldChanges = "changes"
	// FieldAddedIds holds the string denoting the added_ids field in the database.
	FieldAddedIds = "added_ids"
	// FieldRemovedIds holds the string denoting the removed_ids field in the database.
	FieldRemovedIds = "removed_ids"
	// FieldClearedEdges holds the string denoting the cleared_edges field in the database.
	FieldClearedEdges = "cleared_edges"
	// FieldBlame holds the string denoting the blame field in the database.
	FieldBlame = "blame"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the auditlog in the database.
	Table = "audit_logs"
)

// Columns holds all SQL columns for auditlog fields.
var Columns = []string{
	FieldID,
	FieldTable,
	FieldRefID,
	FieldOperation,
	FieldChanges,
	FieldAddedIds,
	FieldRemovedIds,
	FieldClearedEdges,
	FieldBlame,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the AuditLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTable orders the results by the table field.
func ByTable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTable, opts...).ToFunc()
}

// ByRefID orders the results by the ref_id field.
func ByRefID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefID, opts...).ToFunc()
}

// ByOperation orders the results by the operation field.
func ByOperation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperation, opts...).ToFunc()
}

// ByBlame orders the results by the blame field.
func ByBlame(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlame, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
