// Code generated by ent, DO NOT EDIT.

package book

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUpdatedBy, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDisplayName, v))
}

// Abbreviation applies equality check predicate on the "abbreviation" field. It's identical to AbbreviationEQ.
func Abbreviation(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldAbbreviation, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDescription, v))
}

// ExternalLink applies equality check predicate on the "external_link" field. It's identical to ExternalLinkEQ.
func ExternalLink(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldExternalLink, v))
}

// PrimaryImageURL applies equality check predicate on the "primary_image_url" field. It's identical to PrimaryImageURLEQ.
func PrimaryImageURL(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPrimaryImageURL, v))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldYear, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameIsNil applies the IsNil predicate on the "display_name" field.
func DisplayNameIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldDisplayName))
}

// DisplayNameNotNil applies the NotNil predicate on the "display_name" field.
func DisplayNameNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldDisplayName))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldDisplayName, v))
}

// AbbreviationEQ applies the EQ predicate on the "abbreviation" field.
func AbbreviationEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldAbbreviation, v))
}

// AbbreviationNEQ applies the NEQ predicate on the "abbreviation" field.
func AbbreviationNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldAbbreviation, v))
}

// AbbreviationIn applies the In predicate on the "abbreviation" field.
func AbbreviationIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldAbbreviation, vs...))
}

// AbbreviationNotIn applies the NotIn predicate on the "abbreviation" field.
func AbbreviationNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldAbbreviation, vs...))
}

// AbbreviationGT applies the GT predicate on the "abbreviation" field.
func AbbreviationGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldAbbreviation, v))
}

// AbbreviationGTE applies the GTE predicate on the "abbreviation" field.
func AbbreviationGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldAbbreviation, v))
}

// AbbreviationLT applies the LT predicate on the "abbreviation" field.
func AbbreviationLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldAbbreviation, v))
}

// AbbreviationLTE applies the LTE predicate on the "abbreviation" field.
func AbbreviationLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldAbbreviation, v))
}

// AbbreviationContains applies the Contains predicate on the "abbreviation" field.
func AbbreviationContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldAbbreviation, v))
}

// AbbreviationHasPrefix applies the HasPrefix predicate on the "abbreviation" field.
func AbbreviationHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldAbbreviation, v))
}

// AbbreviationHasSuffix applies the HasSuffix predicate on the "abbreviation" field.
func AbbreviationHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldAbbreviation, v))
}

// AbbreviationIsNil applies the IsNil predicate on the "abbreviation" field.
func AbbreviationIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldAbbreviation))
}

// AbbreviationNotNil applies the NotNil predicate on the "abbreviation" field.
func AbbreviationNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldAbbreviation))
}

// AbbreviationEqualFold applies the EqualFold predicate on the "abbreviation" field.
func AbbreviationEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldAbbreviation, v))
}

// AbbreviationContainsFold applies the ContainsFold predicate on the "abbreviation" field.
func AbbreviationContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldAbbreviation, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldDescription, v))
}

// ExternalLinkEQ applies the EQ predicate on the "external_link" field.
func ExternalLinkEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldExternalLink, v))
}

// ExternalLinkNEQ applies the NEQ predicate on the "external_link" field.
func ExternalLinkNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldExternalLink, v))
}

// ExternalLinkIn applies the In predicate on the "external_link" field.
func ExternalLinkIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldExternalLink, vs...))
}

// ExternalLinkNotIn applies the NotIn predicate on the "external_link" field.
func ExternalLinkNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldExternalLink, vs...))
}

// ExternalLinkGT applies the GT predicate on the "external_link" field.
func ExternalLinkGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldExternalLink, v))
}

// ExternalLinkGTE applies the GTE predicate on the "external_link" field.
func ExternalLinkGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldExternalLink, v))
}

// ExternalLinkLT applies the LT predicate on the "external_link" field.
func ExternalLinkLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldExternalLink, v))
}

// ExternalLinkLTE applies the LTE predicate on the "external_link" field.
func ExternalLinkLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldExternalLink, v))
}

// ExternalLinkContains applies the Contains predicate on the "external_link" field.
func ExternalLinkContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldExternalLink, v))
}

// ExternalLinkHasPrefix applies the HasPrefix predicate on the "external_link" field.
func ExternalLinkHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldExternalLink, v))
}

// ExternalLinkHasSuffix applies the HasSuffix predicate on the "external_link" field.
func ExternalLinkHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldExternalLink, v))
}

// ExternalLinkIsNil applies the IsNil predicate on the "external_link" field.
func ExternalLinkIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldExternalLink))
}

// ExternalLinkNotNil applies the NotNil predicate on the "external_link" field.
func ExternalLinkNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldExternalLink))
}

// ExternalLinkEqualFold applies the EqualFold predicate on the "external_link" field.
func ExternalLinkEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldExternalLink, v))
}

// ExternalLinkContainsFold applies the ContainsFold predicate on the "external_link" field.
func ExternalLinkContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldExternalLink, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldStatus))
}

// PrimaryImageURLEQ applies the EQ predicate on the "primary_image_url" field.
func PrimaryImageURLEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPrimaryImageURL, v))
}

// PrimaryImageURLNEQ applies the NEQ predicate on the "primary_image_url" field.
func PrimaryImageURLNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldPrimaryImageURL, v))
}

// PrimaryImageURLIn applies the In predicate on the "primary_image_url" field.
func PrimaryImageURLIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldPrimaryImageURL, vs...))
}

// PrimaryImageURLNotIn applies the NotIn predicate on the "primary_image_url" field.
func PrimaryImageURLNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldPrimaryImageURL, vs...))
}

// PrimaryImageURLGT applies the GT predicate on the "primary_image_url" field.
func PrimaryImageURLGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldPrimaryImageURL, v))
}

// PrimaryImageURLGTE applies the GTE predicate on the "primary_image_url" field.
func PrimaryImageURLGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldPrimaryImageURL, v))
}

// PrimaryImageURLLT applies the LT predicate on the "primary_image_url" field.
func PrimaryImageURLLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldPrimaryImageURL, v))
}

// PrimaryImageURLLTE applies the LTE predicate on the "primary_image_url" field.
func PrimaryImageURLLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldPrimaryImageURL, v))
}

// PrimaryImageURLContains applies the Contains predicate on the "primary_image_url" field.
func PrimaryImageURLContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldPrimaryImageURL, v))
}

// PrimaryImageURLHasPrefix applies the HasPrefix predicate on the "primary_image_url" field.
func PrimaryImageURLHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldPrimaryImageURL, v))
}

// PrimaryImageURLHasSuffix applies the HasSuffix predicate on the "primary_image_url" field.
func PrimaryImageURLHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldPrimaryImageURL, v))
}

// PrimaryImageURLIsNil applies the IsNil predicate on the "primary_image_url" field.
func PrimaryImageURLIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldPrimaryImageURL))
}

// PrimaryImageURLNotNil applies the NotNil predicate on the "primary_image_url" field.
func PrimaryImageURLNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldPrimaryImageURL))
}

// PrimaryImageURLEqualFold applies the EqualFold predicate on the "primary_image_url" field.
func PrimaryImageURLEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldPrimaryImageURL, v))
}

// PrimaryImageURLContainsFold applies the ContainsFold predicate on the "primary_image_url" field.
func PrimaryImageURLContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldPrimaryImageURL, v))
}

// AdditionalImagesUrlsIsNil applies the IsNil predicate on the "additional_images_urls" field.
func AdditionalImagesUrlsIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldAdditionalImagesUrls))
}

// AdditionalImagesUrlsNotNil applies the NotNil predicate on the "additional_images_urls" field.
func AdditionalImagesUrlsNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldAdditionalImagesUrls))
}

// FilesIsNil applies the IsNil predicate on the "files" field.
func FilesIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldFiles))
}

// FilesNotNil applies the NotNil predicate on the "files" field.
func FilesNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldFiles))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldYear, v))
}

// YearIsNil applies the IsNil predicate on the "year" field.
func YearIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldYear))
}

// YearNotNil applies the NotNil predicate on the "year" field.
func YearNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldYear))
}

// HasAuthors applies the HasEdge predicate on the "authors" edge.
func HasAuthors() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AuthorsTable, AuthorsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorsWith applies the HasEdge predicate on the "authors" edge with a given conditions (other predicates).
func HasAuthorsWith(preds ...predicate.Person) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newAuthorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBookGenres applies the HasEdge predicate on the "book_genres" edge.
func HasBookGenres() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, BookGenresTable, BookGenresPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookGenresWith applies the HasEdge predicate on the "book_genres" edge with a given conditions (other predicates).
func HasBookGenresWith(preds ...predicate.BookGenre) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newBookGenresStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCollection applies the HasEdge predicate on the "collection" edge.
func HasCollection() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CollectionTable, CollectionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCollectionWith applies the HasEdge predicate on the "collection" edge with a given conditions (other predicates).
func HasCollectionWith(preds ...predicate.Collection) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newCollectionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPeriodical applies the HasEdge predicate on the "periodical" edge.
func HasPeriodical() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PeriodicalTable, PeriodicalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPeriodicalWith applies the HasEdge predicate on the "periodical" edge with a given conditions (other predicates).
func HasPeriodicalWith(preds ...predicate.Periodical) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newPeriodicalStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPublisher applies the HasEdge predicate on the "publisher" edge.
func HasPublisher() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PublisherTable, PublisherColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPublisherWith applies the HasEdge predicate on the "publisher" edge with a given conditions (other predicates).
func HasPublisherWith(preds ...predicate.Publisher) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newPublisherStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLicense applies the HasEdge predicate on the "license" edge.
func HasLicense() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LicenseTable, LicenseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLicenseWith applies the HasEdge predicate on the "license" edge with a given conditions (other predicates).
func HasLicenseWith(preds ...predicate.License) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newLicenseStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.Location) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newLocationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLibrary applies the HasEdge predicate on the "library" edge.
func HasLibrary() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LibraryTable, LibraryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLibraryWith applies the HasEdge predicate on the "library" edge with a given conditions (other predicates).
func HasLibraryWith(preds ...predicate.Organization) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newLibraryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCountry applies the HasEdge predicate on the "country" edge.
func HasCountry() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CountryTable, CountryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCountryWith applies the HasEdge predicate on the "country" edge with a given conditions (other predicates).
func HasCountryWith(preds ...predicate.Country) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newCountryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSettlement applies the HasEdge predicate on the "settlement" edge.
func HasSettlement() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SettlementTable, SettlementColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSettlementWith applies the HasEdge predicate on the "settlement" edge with a given conditions (other predicates).
func HasSettlementWith(preds ...predicate.Settlement) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newSettlementStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDistrict applies the HasEdge predicate on the "district" edge.
func HasDistrict() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DistrictTable, DistrictColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDistrictWith applies the HasEdge predicate on the "district" edge with a given conditions (other predicates).
func HasDistrictWith(preds ...predicate.District) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newDistrictStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRegion applies the HasEdge predicate on the "region" edge.
func HasRegion() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RegionTable, RegionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRegionWith applies the HasEdge predicate on the "region" edge with a given conditions (other predicates).
func HasRegionWith(preds ...predicate.Region) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newRegionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonalCollection applies the HasEdge predicate on the "personal_collection" edge.
func HasPersonalCollection() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PersonalCollectionTable, PersonalCollectionPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonalCollectionWith applies the HasEdge predicate on the "personal_collection" edge with a given conditions (other predicates).
func HasPersonalCollectionWith(preds ...predicate.PersonalCollection) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newPersonalCollectionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(sql.NotPredicates(p))
}
