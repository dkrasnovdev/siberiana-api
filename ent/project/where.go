// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedBy, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDisplayName, v))
}

// Abbreviation applies equality check predicate on the "abbreviation" field. It's identical to AbbreviationEQ.
func Abbreviation(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldAbbreviation, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDescription, v))
}

// ExternalLink applies equality check predicate on the "external_link" field. It's identical to ExternalLinkEQ.
func ExternalLink(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldExternalLink, v))
}

// BeginDate applies equality check predicate on the "begin_date" field. It's identical to BeginDateEQ.
func BeginDate(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldBeginDate, v))
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldEndDate, v))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldYear, v))
}

// BeginYear applies equality check predicate on the "begin_year" field. It's identical to BeginYearEQ.
func BeginYear(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldBeginYear, v))
}

// EndYear applies equality check predicate on the "end_year" field. It's identical to EndYearEQ.
func EndYear(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldEndYear, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameIsNil applies the IsNil predicate on the "display_name" field.
func DisplayNameIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDisplayName))
}

// DisplayNameNotNil applies the NotNil predicate on the "display_name" field.
func DisplayNameNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDisplayName))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldDisplayName, v))
}

// AbbreviationEQ applies the EQ predicate on the "abbreviation" field.
func AbbreviationEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldAbbreviation, v))
}

// AbbreviationNEQ applies the NEQ predicate on the "abbreviation" field.
func AbbreviationNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldAbbreviation, v))
}

// AbbreviationIn applies the In predicate on the "abbreviation" field.
func AbbreviationIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldAbbreviation, vs...))
}

// AbbreviationNotIn applies the NotIn predicate on the "abbreviation" field.
func AbbreviationNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldAbbreviation, vs...))
}

// AbbreviationGT applies the GT predicate on the "abbreviation" field.
func AbbreviationGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldAbbreviation, v))
}

// AbbreviationGTE applies the GTE predicate on the "abbreviation" field.
func AbbreviationGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldAbbreviation, v))
}

// AbbreviationLT applies the LT predicate on the "abbreviation" field.
func AbbreviationLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldAbbreviation, v))
}

// AbbreviationLTE applies the LTE predicate on the "abbreviation" field.
func AbbreviationLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldAbbreviation, v))
}

// AbbreviationContains applies the Contains predicate on the "abbreviation" field.
func AbbreviationContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldAbbreviation, v))
}

// AbbreviationHasPrefix applies the HasPrefix predicate on the "abbreviation" field.
func AbbreviationHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldAbbreviation, v))
}

// AbbreviationHasSuffix applies the HasSuffix predicate on the "abbreviation" field.
func AbbreviationHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldAbbreviation, v))
}

// AbbreviationIsNil applies the IsNil predicate on the "abbreviation" field.
func AbbreviationIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldAbbreviation))
}

// AbbreviationNotNil applies the NotNil predicate on the "abbreviation" field.
func AbbreviationNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldAbbreviation))
}

// AbbreviationEqualFold applies the EqualFold predicate on the "abbreviation" field.
func AbbreviationEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldAbbreviation, v))
}

// AbbreviationContainsFold applies the ContainsFold predicate on the "abbreviation" field.
func AbbreviationContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldAbbreviation, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldDescription, v))
}

// ExternalLinkEQ applies the EQ predicate on the "external_link" field.
func ExternalLinkEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldExternalLink, v))
}

// ExternalLinkNEQ applies the NEQ predicate on the "external_link" field.
func ExternalLinkNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldExternalLink, v))
}

// ExternalLinkIn applies the In predicate on the "external_link" field.
func ExternalLinkIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldExternalLink, vs...))
}

// ExternalLinkNotIn applies the NotIn predicate on the "external_link" field.
func ExternalLinkNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldExternalLink, vs...))
}

// ExternalLinkGT applies the GT predicate on the "external_link" field.
func ExternalLinkGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldExternalLink, v))
}

// ExternalLinkGTE applies the GTE predicate on the "external_link" field.
func ExternalLinkGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldExternalLink, v))
}

// ExternalLinkLT applies the LT predicate on the "external_link" field.
func ExternalLinkLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldExternalLink, v))
}

// ExternalLinkLTE applies the LTE predicate on the "external_link" field.
func ExternalLinkLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldExternalLink, v))
}

// ExternalLinkContains applies the Contains predicate on the "external_link" field.
func ExternalLinkContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldExternalLink, v))
}

// ExternalLinkHasPrefix applies the HasPrefix predicate on the "external_link" field.
func ExternalLinkHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldExternalLink, v))
}

// ExternalLinkHasSuffix applies the HasSuffix predicate on the "external_link" field.
func ExternalLinkHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldExternalLink, v))
}

// ExternalLinkIsNil applies the IsNil predicate on the "external_link" field.
func ExternalLinkIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldExternalLink))
}

// ExternalLinkNotNil applies the NotNil predicate on the "external_link" field.
func ExternalLinkNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldExternalLink))
}

// ExternalLinkEqualFold applies the EqualFold predicate on the "external_link" field.
func ExternalLinkEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldExternalLink, v))
}

// ExternalLinkContainsFold applies the ContainsFold predicate on the "external_link" field.
func ExternalLinkContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldExternalLink, v))
}

// BeginDateEQ applies the EQ predicate on the "begin_date" field.
func BeginDateEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldBeginDate, v))
}

// BeginDateNEQ applies the NEQ predicate on the "begin_date" field.
func BeginDateNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldBeginDate, v))
}

// BeginDateIn applies the In predicate on the "begin_date" field.
func BeginDateIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldBeginDate, vs...))
}

// BeginDateNotIn applies the NotIn predicate on the "begin_date" field.
func BeginDateNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldBeginDate, vs...))
}

// BeginDateGT applies the GT predicate on the "begin_date" field.
func BeginDateGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldBeginDate, v))
}

// BeginDateGTE applies the GTE predicate on the "begin_date" field.
func BeginDateGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldBeginDate, v))
}

// BeginDateLT applies the LT predicate on the "begin_date" field.
func BeginDateLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldBeginDate, v))
}

// BeginDateLTE applies the LTE predicate on the "begin_date" field.
func BeginDateLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldBeginDate, v))
}

// BeginDateIsNil applies the IsNil predicate on the "begin_date" field.
func BeginDateIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldBeginDate))
}

// BeginDateNotNil applies the NotNil predicate on the "begin_date" field.
func BeginDateNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldBeginDate))
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldEndDate, v))
}

// EndDateIsNil applies the IsNil predicate on the "end_date" field.
func EndDateIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldEndDate))
}

// EndDateNotNil applies the NotNil predicate on the "end_date" field.
func EndDateNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldEndDate))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldYear, v))
}

// YearIsNil applies the IsNil predicate on the "year" field.
func YearIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldYear))
}

// YearNotNil applies the NotNil predicate on the "year" field.
func YearNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldYear))
}

// BeginYearEQ applies the EQ predicate on the "begin_year" field.
func BeginYearEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldBeginYear, v))
}

// BeginYearNEQ applies the NEQ predicate on the "begin_year" field.
func BeginYearNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldBeginYear, v))
}

// BeginYearIn applies the In predicate on the "begin_year" field.
func BeginYearIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldBeginYear, vs...))
}

// BeginYearNotIn applies the NotIn predicate on the "begin_year" field.
func BeginYearNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldBeginYear, vs...))
}

// BeginYearGT applies the GT predicate on the "begin_year" field.
func BeginYearGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldBeginYear, v))
}

// BeginYearGTE applies the GTE predicate on the "begin_year" field.
func BeginYearGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldBeginYear, v))
}

// BeginYearLT applies the LT predicate on the "begin_year" field.
func BeginYearLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldBeginYear, v))
}

// BeginYearLTE applies the LTE predicate on the "begin_year" field.
func BeginYearLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldBeginYear, v))
}

// BeginYearIsNil applies the IsNil predicate on the "begin_year" field.
func BeginYearIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldBeginYear))
}

// BeginYearNotNil applies the NotNil predicate on the "begin_year" field.
func BeginYearNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldBeginYear))
}

// EndYearEQ applies the EQ predicate on the "end_year" field.
func EndYearEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldEndYear, v))
}

// EndYearNEQ applies the NEQ predicate on the "end_year" field.
func EndYearNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldEndYear, v))
}

// EndYearIn applies the In predicate on the "end_year" field.
func EndYearIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldEndYear, vs...))
}

// EndYearNotIn applies the NotIn predicate on the "end_year" field.
func EndYearNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldEndYear, vs...))
}

// EndYearGT applies the GT predicate on the "end_year" field.
func EndYearGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldEndYear, v))
}

// EndYearGTE applies the GTE predicate on the "end_year" field.
func EndYearGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldEndYear, v))
}

// EndYearLT applies the LT predicate on the "end_year" field.
func EndYearLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldEndYear, v))
}

// EndYearLTE applies the LTE predicate on the "end_year" field.
func EndYearLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldEndYear, v))
}

// EndYearIsNil applies the IsNil predicate on the "end_year" field.
func EndYearIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldEndYear))
}

// EndYearNotNil applies the NotNil predicate on the "end_year" field.
func EndYearNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldEndYear))
}

// HasArtifacts applies the HasEdge predicate on the "artifacts" edge.
func HasArtifacts() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ArtifactsTable, ArtifactsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtifactsWith applies the HasEdge predicate on the "artifacts" edge with a given conditions (other predicates).
func HasArtifactsWith(preds ...predicate.Artifact) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newArtifactsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TeamTable, TeamPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Person) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newTeamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(sql.NotPredicates(p))
}
