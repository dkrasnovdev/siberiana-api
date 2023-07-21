// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedBy, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldAddress, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDisplayName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDescription, v))
}

// PrimaryImageURL applies equality check predicate on the "primary_image_url" field. It's identical to PrimaryImageURLEQ.
func PrimaryImageURL(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldPrimaryImageURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldAddress, v))
}

// PhoneNumbersIsNil applies the IsNil predicate on the "phone_numbers" field.
func PhoneNumbersIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldPhoneNumbers))
}

// PhoneNumbersNotNil applies the NotNil predicate on the "phone_numbers" field.
func PhoneNumbersNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldPhoneNumbers))
}

// EmailsIsNil applies the IsNil predicate on the "emails" field.
func EmailsIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldEmails))
}

// EmailsNotNil applies the NotNil predicate on the "emails" field.
func EmailsNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldEmails))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameIsNil applies the IsNil predicate on the "display_name" field.
func DisplayNameIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldDisplayName))
}

// DisplayNameNotNil applies the NotNil predicate on the "display_name" field.
func DisplayNameNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldDisplayName))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldDisplayName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldDescription, v))
}

// ExternalLinksIsNil applies the IsNil predicate on the "external_links" field.
func ExternalLinksIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldExternalLinks))
}

// ExternalLinksNotNil applies the NotNil predicate on the "external_links" field.
func ExternalLinksNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldExternalLinks))
}

// PrimaryImageURLEQ applies the EQ predicate on the "primary_image_url" field.
func PrimaryImageURLEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldPrimaryImageURL, v))
}

// PrimaryImageURLNEQ applies the NEQ predicate on the "primary_image_url" field.
func PrimaryImageURLNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldPrimaryImageURL, v))
}

// PrimaryImageURLIn applies the In predicate on the "primary_image_url" field.
func PrimaryImageURLIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldPrimaryImageURL, vs...))
}

// PrimaryImageURLNotIn applies the NotIn predicate on the "primary_image_url" field.
func PrimaryImageURLNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldPrimaryImageURL, vs...))
}

// PrimaryImageURLGT applies the GT predicate on the "primary_image_url" field.
func PrimaryImageURLGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldPrimaryImageURL, v))
}

// PrimaryImageURLGTE applies the GTE predicate on the "primary_image_url" field.
func PrimaryImageURLGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldPrimaryImageURL, v))
}

// PrimaryImageURLLT applies the LT predicate on the "primary_image_url" field.
func PrimaryImageURLLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldPrimaryImageURL, v))
}

// PrimaryImageURLLTE applies the LTE predicate on the "primary_image_url" field.
func PrimaryImageURLLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldPrimaryImageURL, v))
}

// PrimaryImageURLContains applies the Contains predicate on the "primary_image_url" field.
func PrimaryImageURLContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldPrimaryImageURL, v))
}

// PrimaryImageURLHasPrefix applies the HasPrefix predicate on the "primary_image_url" field.
func PrimaryImageURLHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldPrimaryImageURL, v))
}

// PrimaryImageURLHasSuffix applies the HasSuffix predicate on the "primary_image_url" field.
func PrimaryImageURLHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldPrimaryImageURL, v))
}

// PrimaryImageURLIsNil applies the IsNil predicate on the "primary_image_url" field.
func PrimaryImageURLIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldPrimaryImageURL))
}

// PrimaryImageURLNotNil applies the NotNil predicate on the "primary_image_url" field.
func PrimaryImageURLNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldPrimaryImageURL))
}

// PrimaryImageURLEqualFold applies the EqualFold predicate on the "primary_image_url" field.
func PrimaryImageURLEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldPrimaryImageURL, v))
}

// PrimaryImageURLContainsFold applies the ContainsFold predicate on the "primary_image_url" field.
func PrimaryImageURLContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldPrimaryImageURL, v))
}

// AdditionalImagesUrlsIsNil applies the IsNil predicate on the "additional_images_urls" field.
func AdditionalImagesUrlsIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldAdditionalImagesUrls))
}

// AdditionalImagesUrlsNotNil applies the NotNil predicate on the "additional_images_urls" field.
func AdditionalImagesUrlsNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldAdditionalImagesUrls))
}

// HasHolder applies the HasEdge predicate on the "holder" edge.
func HasHolder() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, HolderTable, HolderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHolderWith applies the HasEdge predicate on the "holder" edge with a given conditions (other predicates).
func HasHolderWith(preds ...predicate.Holder) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newHolderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
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
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		p(s.Not())
	})
}
