// Code generated by ent, DO NOT EDIT.

package protectedareacategory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldUpdatedBy, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldDisplayName, v))
}

// Abbreviation applies equality check predicate on the "abbreviation" field. It's identical to AbbreviationEQ.
func Abbreviation(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldAbbreviation, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldDescription, v))
}

// ExternalLink applies equality check predicate on the "external_link" field. It's identical to ExternalLinkEQ.
func ExternalLink(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldExternalLink, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameIsNil applies the IsNil predicate on the "display_name" field.
func DisplayNameIsNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIsNull(FieldDisplayName))
}

// DisplayNameNotNil applies the NotNil predicate on the "display_name" field.
func DisplayNameNotNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotNull(FieldDisplayName))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContainsFold(FieldDisplayName, v))
}

// AbbreviationEQ applies the EQ predicate on the "abbreviation" field.
func AbbreviationEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldAbbreviation, v))
}

// AbbreviationNEQ applies the NEQ predicate on the "abbreviation" field.
func AbbreviationNEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldAbbreviation, v))
}

// AbbreviationIn applies the In predicate on the "abbreviation" field.
func AbbreviationIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldAbbreviation, vs...))
}

// AbbreviationNotIn applies the NotIn predicate on the "abbreviation" field.
func AbbreviationNotIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldAbbreviation, vs...))
}

// AbbreviationGT applies the GT predicate on the "abbreviation" field.
func AbbreviationGT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldAbbreviation, v))
}

// AbbreviationGTE applies the GTE predicate on the "abbreviation" field.
func AbbreviationGTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldAbbreviation, v))
}

// AbbreviationLT applies the LT predicate on the "abbreviation" field.
func AbbreviationLT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldAbbreviation, v))
}

// AbbreviationLTE applies the LTE predicate on the "abbreviation" field.
func AbbreviationLTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldAbbreviation, v))
}

// AbbreviationContains applies the Contains predicate on the "abbreviation" field.
func AbbreviationContains(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContains(FieldAbbreviation, v))
}

// AbbreviationHasPrefix applies the HasPrefix predicate on the "abbreviation" field.
func AbbreviationHasPrefix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasPrefix(FieldAbbreviation, v))
}

// AbbreviationHasSuffix applies the HasSuffix predicate on the "abbreviation" field.
func AbbreviationHasSuffix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasSuffix(FieldAbbreviation, v))
}

// AbbreviationIsNil applies the IsNil predicate on the "abbreviation" field.
func AbbreviationIsNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIsNull(FieldAbbreviation))
}

// AbbreviationNotNil applies the NotNil predicate on the "abbreviation" field.
func AbbreviationNotNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotNull(FieldAbbreviation))
}

// AbbreviationEqualFold applies the EqualFold predicate on the "abbreviation" field.
func AbbreviationEqualFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEqualFold(FieldAbbreviation, v))
}

// AbbreviationContainsFold applies the ContainsFold predicate on the "abbreviation" field.
func AbbreviationContainsFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContainsFold(FieldAbbreviation, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContainsFold(FieldDescription, v))
}

// ExternalLinkEQ applies the EQ predicate on the "external_link" field.
func ExternalLinkEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEQ(FieldExternalLink, v))
}

// ExternalLinkNEQ applies the NEQ predicate on the "external_link" field.
func ExternalLinkNEQ(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNEQ(FieldExternalLink, v))
}

// ExternalLinkIn applies the In predicate on the "external_link" field.
func ExternalLinkIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIn(FieldExternalLink, vs...))
}

// ExternalLinkNotIn applies the NotIn predicate on the "external_link" field.
func ExternalLinkNotIn(vs ...string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotIn(FieldExternalLink, vs...))
}

// ExternalLinkGT applies the GT predicate on the "external_link" field.
func ExternalLinkGT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGT(FieldExternalLink, v))
}

// ExternalLinkGTE applies the GTE predicate on the "external_link" field.
func ExternalLinkGTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldGTE(FieldExternalLink, v))
}

// ExternalLinkLT applies the LT predicate on the "external_link" field.
func ExternalLinkLT(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLT(FieldExternalLink, v))
}

// ExternalLinkLTE applies the LTE predicate on the "external_link" field.
func ExternalLinkLTE(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldLTE(FieldExternalLink, v))
}

// ExternalLinkContains applies the Contains predicate on the "external_link" field.
func ExternalLinkContains(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContains(FieldExternalLink, v))
}

// ExternalLinkHasPrefix applies the HasPrefix predicate on the "external_link" field.
func ExternalLinkHasPrefix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasPrefix(FieldExternalLink, v))
}

// ExternalLinkHasSuffix applies the HasSuffix predicate on the "external_link" field.
func ExternalLinkHasSuffix(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldHasSuffix(FieldExternalLink, v))
}

// ExternalLinkIsNil applies the IsNil predicate on the "external_link" field.
func ExternalLinkIsNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldIsNull(FieldExternalLink))
}

// ExternalLinkNotNil applies the NotNil predicate on the "external_link" field.
func ExternalLinkNotNil() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldNotNull(FieldExternalLink))
}

// ExternalLinkEqualFold applies the EqualFold predicate on the "external_link" field.
func ExternalLinkEqualFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldEqualFold(FieldExternalLink, v))
}

// ExternalLinkContainsFold applies the ContainsFold predicate on the "external_link" field.
func ExternalLinkContainsFold(v string) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.FieldContainsFold(FieldExternalLink, v))
}

// HasProtectedAreas applies the HasEdge predicate on the "protected_areas" edge.
func HasProtectedAreas() predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProtectedAreasTable, ProtectedAreasColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProtectedAreasWith applies the HasEdge predicate on the "protected_areas" edge with a given conditions (other predicates).
func HasProtectedAreasWith(preds ...predicate.ProtectedArea) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(func(s *sql.Selector) {
		step := newProtectedAreasStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProtectedAreaCategory) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProtectedAreaCategory) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProtectedAreaCategory) predicate.ProtectedAreaCategory {
	return predicate.ProtectedAreaCategory(sql.NotPredicates(p))
}
