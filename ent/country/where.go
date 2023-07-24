// Code generated by ent, DO NOT EDIT.

package country

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldUpdatedBy, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldDisplayName, v))
}

// Abbreviation applies equality check predicate on the "abbreviation" field. It's identical to AbbreviationEQ.
func Abbreviation(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldAbbreviation, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameIsNil applies the IsNil predicate on the "display_name" field.
func DisplayNameIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldDisplayName))
}

// DisplayNameNotNil applies the NotNil predicate on the "display_name" field.
func DisplayNameNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldDisplayName))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldDisplayName, v))
}

// AbbreviationEQ applies the EQ predicate on the "abbreviation" field.
func AbbreviationEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldAbbreviation, v))
}

// AbbreviationNEQ applies the NEQ predicate on the "abbreviation" field.
func AbbreviationNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldAbbreviation, v))
}

// AbbreviationIn applies the In predicate on the "abbreviation" field.
func AbbreviationIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldAbbreviation, vs...))
}

// AbbreviationNotIn applies the NotIn predicate on the "abbreviation" field.
func AbbreviationNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldAbbreviation, vs...))
}

// AbbreviationGT applies the GT predicate on the "abbreviation" field.
func AbbreviationGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldAbbreviation, v))
}

// AbbreviationGTE applies the GTE predicate on the "abbreviation" field.
func AbbreviationGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldAbbreviation, v))
}

// AbbreviationLT applies the LT predicate on the "abbreviation" field.
func AbbreviationLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldAbbreviation, v))
}

// AbbreviationLTE applies the LTE predicate on the "abbreviation" field.
func AbbreviationLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldAbbreviation, v))
}

// AbbreviationContains applies the Contains predicate on the "abbreviation" field.
func AbbreviationContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldAbbreviation, v))
}

// AbbreviationHasPrefix applies the HasPrefix predicate on the "abbreviation" field.
func AbbreviationHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldAbbreviation, v))
}

// AbbreviationHasSuffix applies the HasSuffix predicate on the "abbreviation" field.
func AbbreviationHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldAbbreviation, v))
}

// AbbreviationIsNil applies the IsNil predicate on the "abbreviation" field.
func AbbreviationIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldAbbreviation))
}

// AbbreviationNotNil applies the NotNil predicate on the "abbreviation" field.
func AbbreviationNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldAbbreviation))
}

// AbbreviationEqualFold applies the EqualFold predicate on the "abbreviation" field.
func AbbreviationEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldAbbreviation, v))
}

// AbbreviationContainsFold applies the ContainsFold predicate on the "abbreviation" field.
func AbbreviationContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldAbbreviation, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldDescription, v))
}

// ExternalLinksIsNil applies the IsNil predicate on the "external_links" field.
func ExternalLinksIsNil() predicate.Country {
	return predicate.Country(sql.FieldIsNull(FieldExternalLinks))
}

// ExternalLinksNotNil applies the NotNil predicate on the "external_links" field.
func ExternalLinksNotNil() predicate.Country {
	return predicate.Country(sql.FieldNotNull(FieldExternalLinks))
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.Location) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := newLocationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		p(s.Not())
	})
}