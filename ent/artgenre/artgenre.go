// Code generated by ent, DO NOT EDIT.

package artgenre

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the artgenre type in the database.
	Label = "art_genre"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the artgenre in the database.
	Table = "art_genres"
)

// Columns holds all SQL columns for artgenre fields.
var Columns = []string{
	FieldID,
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

// OrderOption defines the ordering options for the ArtGenre queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}
