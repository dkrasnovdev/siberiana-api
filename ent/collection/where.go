// Code generated by ent, DO NOT EDIT.

package collection

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldUpdatedBy, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldDisplayName, v))
}

// Abbreviation applies equality check predicate on the "abbreviation" field. It's identical to AbbreviationEQ.
func Abbreviation(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldAbbreviation, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldDescription, v))
}

// ExternalLink applies equality check predicate on the "external_link" field. It's identical to ExternalLinkEQ.
func ExternalLink(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldExternalLink, v))
}

// PrimaryImageURL applies equality check predicate on the "primary_image_url" field. It's identical to PrimaryImageURLEQ.
func PrimaryImageURL(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldPrimaryImageURL, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldSlug, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameIsNil applies the IsNil predicate on the "display_name" field.
func DisplayNameIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldDisplayName))
}

// DisplayNameNotNil applies the NotNil predicate on the "display_name" field.
func DisplayNameNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldDisplayName))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldDisplayName, v))
}

// AbbreviationEQ applies the EQ predicate on the "abbreviation" field.
func AbbreviationEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldAbbreviation, v))
}

// AbbreviationNEQ applies the NEQ predicate on the "abbreviation" field.
func AbbreviationNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldAbbreviation, v))
}

// AbbreviationIn applies the In predicate on the "abbreviation" field.
func AbbreviationIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldAbbreviation, vs...))
}

// AbbreviationNotIn applies the NotIn predicate on the "abbreviation" field.
func AbbreviationNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldAbbreviation, vs...))
}

// AbbreviationGT applies the GT predicate on the "abbreviation" field.
func AbbreviationGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldAbbreviation, v))
}

// AbbreviationGTE applies the GTE predicate on the "abbreviation" field.
func AbbreviationGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldAbbreviation, v))
}

// AbbreviationLT applies the LT predicate on the "abbreviation" field.
func AbbreviationLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldAbbreviation, v))
}

// AbbreviationLTE applies the LTE predicate on the "abbreviation" field.
func AbbreviationLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldAbbreviation, v))
}

// AbbreviationContains applies the Contains predicate on the "abbreviation" field.
func AbbreviationContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldAbbreviation, v))
}

// AbbreviationHasPrefix applies the HasPrefix predicate on the "abbreviation" field.
func AbbreviationHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldAbbreviation, v))
}

// AbbreviationHasSuffix applies the HasSuffix predicate on the "abbreviation" field.
func AbbreviationHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldAbbreviation, v))
}

// AbbreviationIsNil applies the IsNil predicate on the "abbreviation" field.
func AbbreviationIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldAbbreviation))
}

// AbbreviationNotNil applies the NotNil predicate on the "abbreviation" field.
func AbbreviationNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldAbbreviation))
}

// AbbreviationEqualFold applies the EqualFold predicate on the "abbreviation" field.
func AbbreviationEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldAbbreviation, v))
}

// AbbreviationContainsFold applies the ContainsFold predicate on the "abbreviation" field.
func AbbreviationContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldAbbreviation, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldDescription, v))
}

// ExternalLinkEQ applies the EQ predicate on the "external_link" field.
func ExternalLinkEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldExternalLink, v))
}

// ExternalLinkNEQ applies the NEQ predicate on the "external_link" field.
func ExternalLinkNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldExternalLink, v))
}

// ExternalLinkIn applies the In predicate on the "external_link" field.
func ExternalLinkIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldExternalLink, vs...))
}

// ExternalLinkNotIn applies the NotIn predicate on the "external_link" field.
func ExternalLinkNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldExternalLink, vs...))
}

// ExternalLinkGT applies the GT predicate on the "external_link" field.
func ExternalLinkGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldExternalLink, v))
}

// ExternalLinkGTE applies the GTE predicate on the "external_link" field.
func ExternalLinkGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldExternalLink, v))
}

// ExternalLinkLT applies the LT predicate on the "external_link" field.
func ExternalLinkLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldExternalLink, v))
}

// ExternalLinkLTE applies the LTE predicate on the "external_link" field.
func ExternalLinkLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldExternalLink, v))
}

// ExternalLinkContains applies the Contains predicate on the "external_link" field.
func ExternalLinkContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldExternalLink, v))
}

// ExternalLinkHasPrefix applies the HasPrefix predicate on the "external_link" field.
func ExternalLinkHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldExternalLink, v))
}

// ExternalLinkHasSuffix applies the HasSuffix predicate on the "external_link" field.
func ExternalLinkHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldExternalLink, v))
}

// ExternalLinkIsNil applies the IsNil predicate on the "external_link" field.
func ExternalLinkIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldExternalLink))
}

// ExternalLinkNotNil applies the NotNil predicate on the "external_link" field.
func ExternalLinkNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldExternalLink))
}

// ExternalLinkEqualFold applies the EqualFold predicate on the "external_link" field.
func ExternalLinkEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldExternalLink, v))
}

// ExternalLinkContainsFold applies the ContainsFold predicate on the "external_link" field.
func ExternalLinkContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldExternalLink, v))
}

// PrimaryImageURLEQ applies the EQ predicate on the "primary_image_url" field.
func PrimaryImageURLEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldPrimaryImageURL, v))
}

// PrimaryImageURLNEQ applies the NEQ predicate on the "primary_image_url" field.
func PrimaryImageURLNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldPrimaryImageURL, v))
}

// PrimaryImageURLIn applies the In predicate on the "primary_image_url" field.
func PrimaryImageURLIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldPrimaryImageURL, vs...))
}

// PrimaryImageURLNotIn applies the NotIn predicate on the "primary_image_url" field.
func PrimaryImageURLNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldPrimaryImageURL, vs...))
}

// PrimaryImageURLGT applies the GT predicate on the "primary_image_url" field.
func PrimaryImageURLGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldPrimaryImageURL, v))
}

// PrimaryImageURLGTE applies the GTE predicate on the "primary_image_url" field.
func PrimaryImageURLGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldPrimaryImageURL, v))
}

// PrimaryImageURLLT applies the LT predicate on the "primary_image_url" field.
func PrimaryImageURLLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldPrimaryImageURL, v))
}

// PrimaryImageURLLTE applies the LTE predicate on the "primary_image_url" field.
func PrimaryImageURLLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldPrimaryImageURL, v))
}

// PrimaryImageURLContains applies the Contains predicate on the "primary_image_url" field.
func PrimaryImageURLContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldPrimaryImageURL, v))
}

// PrimaryImageURLHasPrefix applies the HasPrefix predicate on the "primary_image_url" field.
func PrimaryImageURLHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldPrimaryImageURL, v))
}

// PrimaryImageURLHasSuffix applies the HasSuffix predicate on the "primary_image_url" field.
func PrimaryImageURLHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldPrimaryImageURL, v))
}

// PrimaryImageURLIsNil applies the IsNil predicate on the "primary_image_url" field.
func PrimaryImageURLIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldPrimaryImageURL))
}

// PrimaryImageURLNotNil applies the NotNil predicate on the "primary_image_url" field.
func PrimaryImageURLNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldPrimaryImageURL))
}

// PrimaryImageURLEqualFold applies the EqualFold predicate on the "primary_image_url" field.
func PrimaryImageURLEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldPrimaryImageURL, v))
}

// PrimaryImageURLContainsFold applies the ContainsFold predicate on the "primary_image_url" field.
func PrimaryImageURLContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldPrimaryImageURL, v))
}

// AdditionalImagesUrlsIsNil applies the IsNil predicate on the "additional_images_urls" field.
func AdditionalImagesUrlsIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldAdditionalImagesUrls))
}

// AdditionalImagesUrlsNotNil applies the NotNil predicate on the "additional_images_urls" field.
func AdditionalImagesUrlsNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldAdditionalImagesUrls))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldSlug, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldType, vs...))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Collection {
	return predicate.Collection(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Collection {
	return predicate.Collection(sql.FieldNotNull(FieldType))
}

// HasArts applies the HasEdge predicate on the "arts" edge.
func HasArts() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArtsTable, ArtsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtsWith applies the HasEdge predicate on the "arts" edge with a given conditions (other predicates).
func HasArtsWith(preds ...predicate.Art) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newArtsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArtifacts applies the HasEdge predicate on the "artifacts" edge.
func HasArtifacts() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArtifactsTable, ArtifactsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtifactsWith applies the HasEdge predicate on the "artifacts" edge with a given conditions (other predicates).
func HasArtifactsWith(preds ...predicate.Artifact) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newArtifactsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBooks applies the HasEdge predicate on the "books" edge.
func HasBooks() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BooksTable, BooksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBooksWith applies the HasEdge predicate on the "books" edge with a given conditions (other predicates).
func HasBooksWith(preds ...predicate.Book) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newBooksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProtectedAreaPictures applies the HasEdge predicate on the "protected_area_pictures" edge.
func HasProtectedAreaPictures() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProtectedAreaPicturesTable, ProtectedAreaPicturesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProtectedAreaPicturesWith applies the HasEdge predicate on the "protected_area_pictures" edge with a given conditions (other predicates).
func HasProtectedAreaPicturesWith(preds ...predicate.ProtectedAreaPicture) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newProtectedAreaPicturesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthors applies the HasEdge predicate on the "authors" edge.
func HasAuthors() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AuthorsTable, AuthorsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorsWith applies the HasEdge predicate on the "authors" edge with a given conditions (other predicates).
func HasAuthorsWith(preds ...predicate.Person) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newAuthorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Collection) predicate.Collection {
	return predicate.Collection(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Collection) predicate.Collection {
	return predicate.Collection(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Collection) predicate.Collection {
	return predicate.Collection(sql.NotPredicates(p))
}
