// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"
	"time"

	"github.com/dkrasnovdev/heritage-api/ent/artifact"
	"github.com/dkrasnovdev/heritage-api/ent/auditlog"
	"github.com/dkrasnovdev/heritage-api/ent/category"
	"github.com/dkrasnovdev/heritage-api/ent/collection"
	"github.com/dkrasnovdev/heritage-api/ent/culture"
	"github.com/dkrasnovdev/heritage-api/ent/district"
	"github.com/dkrasnovdev/heritage-api/ent/holder"
	"github.com/dkrasnovdev/heritage-api/ent/holderresponsibility"
	"github.com/dkrasnovdev/heritage-api/ent/license"
	"github.com/dkrasnovdev/heritage-api/ent/location"
	"github.com/dkrasnovdev/heritage-api/ent/medium"
	"github.com/dkrasnovdev/heritage-api/ent/model"
	"github.com/dkrasnovdev/heritage-api/ent/monument"
	"github.com/dkrasnovdev/heritage-api/ent/organization"
	"github.com/dkrasnovdev/heritage-api/ent/organizationtype"
	"github.com/dkrasnovdev/heritage-api/ent/person"
	"github.com/dkrasnovdev/heritage-api/ent/personrole"
	"github.com/dkrasnovdev/heritage-api/ent/project"
	"github.com/dkrasnovdev/heritage-api/ent/projecttype"
	"github.com/dkrasnovdev/heritage-api/ent/publication"
	"github.com/dkrasnovdev/heritage-api/ent/region"
	"github.com/dkrasnovdev/heritage-api/ent/schema"
	"github.com/dkrasnovdev/heritage-api/ent/set"
	"github.com/dkrasnovdev/heritage-api/ent/settlement"
	"github.com/dkrasnovdev/heritage-api/ent/technique"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	artifactMixin := schema.Artifact{}.Mixin()
	artifact.Policy = privacy.NewPolicies(schema.Artifact{})
	artifact.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := artifact.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	artifactMixinHooks0 := artifactMixin[0].Hooks()
	artifactMixinHooks3 := artifactMixin[3].Hooks()

	artifact.Hooks[1] = artifactMixinHooks0[0]

	artifact.Hooks[2] = artifactMixinHooks3[0]
	artifactMixinInters3 := artifactMixin[3].Interceptors()
	artifact.Interceptors[0] = artifactMixinInters3[0]
	artifactMixinFields0 := artifactMixin[0].Fields()
	_ = artifactMixinFields0
	artifactFields := schema.Artifact{}.Fields()
	_ = artifactFields
	// artifactDescCreatedAt is the schema descriptor for created_at field.
	artifactDescCreatedAt := artifactMixinFields0[0].Descriptor()
	// artifact.DefaultCreatedAt holds the default value on creation for the created_at field.
	artifact.DefaultCreatedAt = artifactDescCreatedAt.Default.(func() time.Time)
	// artifactDescUpdatedAt is the schema descriptor for updated_at field.
	artifactDescUpdatedAt := artifactMixinFields0[2].Descriptor()
	// artifact.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	artifact.DefaultUpdatedAt = artifactDescUpdatedAt.Default.(func() time.Time)
	// artifact.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	artifact.UpdateDefaultUpdatedAt = artifactDescUpdatedAt.UpdateDefault.(func() time.Time)
	auditlog.Policy = privacy.NewPolicies(schema.AuditLog{})
	auditlog.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := auditlog.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	auditlogFields := schema.AuditLog{}.Fields()
	_ = auditlogFields
	// auditlogDescCreatedAt is the schema descriptor for created_at field.
	auditlogDescCreatedAt := auditlogFields[8].Descriptor()
	// auditlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	auditlog.DefaultCreatedAt = auditlogDescCreatedAt.Default.(func() time.Time)
	categoryMixin := schema.Category{}.Mixin()
	category.Policy = privacy.NewPolicies(schema.Category{})
	category.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := category.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	categoryMixinHooks0 := categoryMixin[0].Hooks()

	category.Hooks[1] = categoryMixinHooks0[0]
	categoryMixinFields0 := categoryMixin[0].Fields()
	_ = categoryMixinFields0
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryMixinFields0[0].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescUpdatedAt is the schema descriptor for updated_at field.
	categoryDescUpdatedAt := categoryMixinFields0[2].Descriptor()
	// category.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	category.DefaultUpdatedAt = categoryDescUpdatedAt.Default.(func() time.Time)
	// category.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	category.UpdateDefaultUpdatedAt = categoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	collectionMixin := schema.Collection{}.Mixin()
	collection.Policy = privacy.NewPolicies(schema.Collection{})
	collection.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := collection.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	collectionMixinHooks0 := collectionMixin[0].Hooks()

	collection.Hooks[1] = collectionMixinHooks0[0]
	collectionMixinFields0 := collectionMixin[0].Fields()
	_ = collectionMixinFields0
	collectionFields := schema.Collection{}.Fields()
	_ = collectionFields
	// collectionDescCreatedAt is the schema descriptor for created_at field.
	collectionDescCreatedAt := collectionMixinFields0[0].Descriptor()
	// collection.DefaultCreatedAt holds the default value on creation for the created_at field.
	collection.DefaultCreatedAt = collectionDescCreatedAt.Default.(func() time.Time)
	// collectionDescUpdatedAt is the schema descriptor for updated_at field.
	collectionDescUpdatedAt := collectionMixinFields0[2].Descriptor()
	// collection.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	collection.DefaultUpdatedAt = collectionDescUpdatedAt.Default.(func() time.Time)
	// collection.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	collection.UpdateDefaultUpdatedAt = collectionDescUpdatedAt.UpdateDefault.(func() time.Time)
	cultureMixin := schema.Culture{}.Mixin()
	culture.Policy = privacy.NewPolicies(schema.Culture{})
	culture.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := culture.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	cultureMixinHooks0 := cultureMixin[0].Hooks()

	culture.Hooks[1] = cultureMixinHooks0[0]
	cultureMixinFields0 := cultureMixin[0].Fields()
	_ = cultureMixinFields0
	cultureFields := schema.Culture{}.Fields()
	_ = cultureFields
	// cultureDescCreatedAt is the schema descriptor for created_at field.
	cultureDescCreatedAt := cultureMixinFields0[0].Descriptor()
	// culture.DefaultCreatedAt holds the default value on creation for the created_at field.
	culture.DefaultCreatedAt = cultureDescCreatedAt.Default.(func() time.Time)
	// cultureDescUpdatedAt is the schema descriptor for updated_at field.
	cultureDescUpdatedAt := cultureMixinFields0[2].Descriptor()
	// culture.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	culture.DefaultUpdatedAt = cultureDescUpdatedAt.Default.(func() time.Time)
	// culture.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	culture.UpdateDefaultUpdatedAt = cultureDescUpdatedAt.UpdateDefault.(func() time.Time)
	districtMixin := schema.District{}.Mixin()
	district.Policy = privacy.NewPolicies(schema.District{})
	district.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := district.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	districtMixinHooks0 := districtMixin[0].Hooks()

	district.Hooks[1] = districtMixinHooks0[0]
	districtMixinFields0 := districtMixin[0].Fields()
	_ = districtMixinFields0
	districtFields := schema.District{}.Fields()
	_ = districtFields
	// districtDescCreatedAt is the schema descriptor for created_at field.
	districtDescCreatedAt := districtMixinFields0[0].Descriptor()
	// district.DefaultCreatedAt holds the default value on creation for the created_at field.
	district.DefaultCreatedAt = districtDescCreatedAt.Default.(func() time.Time)
	// districtDescUpdatedAt is the schema descriptor for updated_at field.
	districtDescUpdatedAt := districtMixinFields0[2].Descriptor()
	// district.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	district.DefaultUpdatedAt = districtDescUpdatedAt.Default.(func() time.Time)
	// district.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	district.UpdateDefaultUpdatedAt = districtDescUpdatedAt.UpdateDefault.(func() time.Time)
	holderMixin := schema.Holder{}.Mixin()
	holder.Policy = privacy.NewPolicies(schema.Holder{})
	holder.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := holder.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	holderMixinHooks0 := holderMixin[0].Hooks()

	holder.Hooks[1] = holderMixinHooks0[0]
	holderMixinFields0 := holderMixin[0].Fields()
	_ = holderMixinFields0
	holderFields := schema.Holder{}.Fields()
	_ = holderFields
	// holderDescCreatedAt is the schema descriptor for created_at field.
	holderDescCreatedAt := holderMixinFields0[0].Descriptor()
	// holder.DefaultCreatedAt holds the default value on creation for the created_at field.
	holder.DefaultCreatedAt = holderDescCreatedAt.Default.(func() time.Time)
	// holderDescUpdatedAt is the schema descriptor for updated_at field.
	holderDescUpdatedAt := holderMixinFields0[2].Descriptor()
	// holder.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	holder.DefaultUpdatedAt = holderDescUpdatedAt.Default.(func() time.Time)
	// holder.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	holder.UpdateDefaultUpdatedAt = holderDescUpdatedAt.UpdateDefault.(func() time.Time)
	holderresponsibilityMixin := schema.HolderResponsibility{}.Mixin()
	holderresponsibilityMixinHooks0 := holderresponsibilityMixin[0].Hooks()
	holderresponsibility.Hooks[0] = holderresponsibilityMixinHooks0[0]
	holderresponsibilityMixinFields0 := holderresponsibilityMixin[0].Fields()
	_ = holderresponsibilityMixinFields0
	holderresponsibilityFields := schema.HolderResponsibility{}.Fields()
	_ = holderresponsibilityFields
	// holderresponsibilityDescCreatedAt is the schema descriptor for created_at field.
	holderresponsibilityDescCreatedAt := holderresponsibilityMixinFields0[0].Descriptor()
	// holderresponsibility.DefaultCreatedAt holds the default value on creation for the created_at field.
	holderresponsibility.DefaultCreatedAt = holderresponsibilityDescCreatedAt.Default.(func() time.Time)
	// holderresponsibilityDescUpdatedAt is the schema descriptor for updated_at field.
	holderresponsibilityDescUpdatedAt := holderresponsibilityMixinFields0[2].Descriptor()
	// holderresponsibility.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	holderresponsibility.DefaultUpdatedAt = holderresponsibilityDescUpdatedAt.Default.(func() time.Time)
	// holderresponsibility.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	holderresponsibility.UpdateDefaultUpdatedAt = holderresponsibilityDescUpdatedAt.UpdateDefault.(func() time.Time)
	licenseMixin := schema.License{}.Mixin()
	license.Policy = privacy.NewPolicies(schema.License{})
	license.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := license.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	licenseMixinHooks0 := licenseMixin[0].Hooks()

	license.Hooks[1] = licenseMixinHooks0[0]
	licenseMixinFields0 := licenseMixin[0].Fields()
	_ = licenseMixinFields0
	licenseFields := schema.License{}.Fields()
	_ = licenseFields
	// licenseDescCreatedAt is the schema descriptor for created_at field.
	licenseDescCreatedAt := licenseMixinFields0[0].Descriptor()
	// license.DefaultCreatedAt holds the default value on creation for the created_at field.
	license.DefaultCreatedAt = licenseDescCreatedAt.Default.(func() time.Time)
	// licenseDescUpdatedAt is the schema descriptor for updated_at field.
	licenseDescUpdatedAt := licenseMixinFields0[2].Descriptor()
	// license.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	license.DefaultUpdatedAt = licenseDescUpdatedAt.Default.(func() time.Time)
	// license.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	license.UpdateDefaultUpdatedAt = licenseDescUpdatedAt.UpdateDefault.(func() time.Time)
	locationMixin := schema.Location{}.Mixin()
	location.Policy = privacy.NewPolicies(schema.Location{})
	location.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := location.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	locationMixinHooks0 := locationMixin[0].Hooks()

	location.Hooks[1] = locationMixinHooks0[0]
	locationMixinFields0 := locationMixin[0].Fields()
	_ = locationMixinFields0
	locationFields := schema.Location{}.Fields()
	_ = locationFields
	// locationDescCreatedAt is the schema descriptor for created_at field.
	locationDescCreatedAt := locationMixinFields0[0].Descriptor()
	// location.DefaultCreatedAt holds the default value on creation for the created_at field.
	location.DefaultCreatedAt = locationDescCreatedAt.Default.(func() time.Time)
	// locationDescUpdatedAt is the schema descriptor for updated_at field.
	locationDescUpdatedAt := locationMixinFields0[2].Descriptor()
	// location.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	location.DefaultUpdatedAt = locationDescUpdatedAt.Default.(func() time.Time)
	// location.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	location.UpdateDefaultUpdatedAt = locationDescUpdatedAt.UpdateDefault.(func() time.Time)
	mediumMixin := schema.Medium{}.Mixin()
	medium.Policy = privacy.NewPolicies(schema.Medium{})
	medium.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := medium.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	mediumMixinHooks0 := mediumMixin[0].Hooks()

	medium.Hooks[1] = mediumMixinHooks0[0]
	mediumMixinFields0 := mediumMixin[0].Fields()
	_ = mediumMixinFields0
	mediumFields := schema.Medium{}.Fields()
	_ = mediumFields
	// mediumDescCreatedAt is the schema descriptor for created_at field.
	mediumDescCreatedAt := mediumMixinFields0[0].Descriptor()
	// medium.DefaultCreatedAt holds the default value on creation for the created_at field.
	medium.DefaultCreatedAt = mediumDescCreatedAt.Default.(func() time.Time)
	// mediumDescUpdatedAt is the schema descriptor for updated_at field.
	mediumDescUpdatedAt := mediumMixinFields0[2].Descriptor()
	// medium.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	medium.DefaultUpdatedAt = mediumDescUpdatedAt.Default.(func() time.Time)
	// medium.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	medium.UpdateDefaultUpdatedAt = mediumDescUpdatedAt.UpdateDefault.(func() time.Time)
	modelMixin := schema.Model{}.Mixin()
	model.Policy = privacy.NewPolicies(schema.Model{})
	model.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := model.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	modelMixinHooks0 := modelMixin[0].Hooks()

	model.Hooks[1] = modelMixinHooks0[0]
	modelMixinFields0 := modelMixin[0].Fields()
	_ = modelMixinFields0
	modelFields := schema.Model{}.Fields()
	_ = modelFields
	// modelDescCreatedAt is the schema descriptor for created_at field.
	modelDescCreatedAt := modelMixinFields0[0].Descriptor()
	// model.DefaultCreatedAt holds the default value on creation for the created_at field.
	model.DefaultCreatedAt = modelDescCreatedAt.Default.(func() time.Time)
	// modelDescUpdatedAt is the schema descriptor for updated_at field.
	modelDescUpdatedAt := modelMixinFields0[2].Descriptor()
	// model.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	model.DefaultUpdatedAt = modelDescUpdatedAt.Default.(func() time.Time)
	// model.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	model.UpdateDefaultUpdatedAt = modelDescUpdatedAt.UpdateDefault.(func() time.Time)
	monumentMixin := schema.Monument{}.Mixin()
	monument.Policy = privacy.NewPolicies(schema.Monument{})
	monument.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := monument.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	monumentMixinHooks0 := monumentMixin[0].Hooks()

	monument.Hooks[1] = monumentMixinHooks0[0]
	monumentMixinFields0 := monumentMixin[0].Fields()
	_ = monumentMixinFields0
	monumentFields := schema.Monument{}.Fields()
	_ = monumentFields
	// monumentDescCreatedAt is the schema descriptor for created_at field.
	monumentDescCreatedAt := monumentMixinFields0[0].Descriptor()
	// monument.DefaultCreatedAt holds the default value on creation for the created_at field.
	monument.DefaultCreatedAt = monumentDescCreatedAt.Default.(func() time.Time)
	// monumentDescUpdatedAt is the schema descriptor for updated_at field.
	monumentDescUpdatedAt := monumentMixinFields0[2].Descriptor()
	// monument.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	monument.DefaultUpdatedAt = monumentDescUpdatedAt.Default.(func() time.Time)
	// monument.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	monument.UpdateDefaultUpdatedAt = monumentDescUpdatedAt.UpdateDefault.(func() time.Time)
	organizationMixin := schema.Organization{}.Mixin()
	organization.Policy = privacy.NewPolicies(schema.Organization{})
	organization.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := organization.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	organizationMixinHooks0 := organizationMixin[0].Hooks()

	organization.Hooks[1] = organizationMixinHooks0[0]
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescCreatedAt is the schema descriptor for created_at field.
	organizationDescCreatedAt := organizationMixinFields0[0].Descriptor()
	// organization.DefaultCreatedAt holds the default value on creation for the created_at field.
	organization.DefaultCreatedAt = organizationDescCreatedAt.Default.(func() time.Time)
	// organizationDescUpdatedAt is the schema descriptor for updated_at field.
	organizationDescUpdatedAt := organizationMixinFields0[2].Descriptor()
	// organization.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	organization.DefaultUpdatedAt = organizationDescUpdatedAt.Default.(func() time.Time)
	// organization.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	organization.UpdateDefaultUpdatedAt = organizationDescUpdatedAt.UpdateDefault.(func() time.Time)
	organizationtypeMixin := schema.OrganizationType{}.Mixin()
	organizationtypeMixinHooks0 := organizationtypeMixin[0].Hooks()
	organizationtype.Hooks[0] = organizationtypeMixinHooks0[0]
	organizationtypeMixinFields0 := organizationtypeMixin[0].Fields()
	_ = organizationtypeMixinFields0
	organizationtypeFields := schema.OrganizationType{}.Fields()
	_ = organizationtypeFields
	// organizationtypeDescCreatedAt is the schema descriptor for created_at field.
	organizationtypeDescCreatedAt := organizationtypeMixinFields0[0].Descriptor()
	// organizationtype.DefaultCreatedAt holds the default value on creation for the created_at field.
	organizationtype.DefaultCreatedAt = organizationtypeDescCreatedAt.Default.(func() time.Time)
	// organizationtypeDescUpdatedAt is the schema descriptor for updated_at field.
	organizationtypeDescUpdatedAt := organizationtypeMixinFields0[2].Descriptor()
	// organizationtype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	organizationtype.DefaultUpdatedAt = organizationtypeDescUpdatedAt.Default.(func() time.Time)
	// organizationtype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	organizationtype.UpdateDefaultUpdatedAt = organizationtypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	personMixin := schema.Person{}.Mixin()
	person.Policy = privacy.NewPolicies(schema.Person{})
	person.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := person.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	personMixinHooks0 := personMixin[0].Hooks()

	person.Hooks[1] = personMixinHooks0[0]
	personMixinFields0 := personMixin[0].Fields()
	_ = personMixinFields0
	personFields := schema.Person{}.Fields()
	_ = personFields
	// personDescCreatedAt is the schema descriptor for created_at field.
	personDescCreatedAt := personMixinFields0[0].Descriptor()
	// person.DefaultCreatedAt holds the default value on creation for the created_at field.
	person.DefaultCreatedAt = personDescCreatedAt.Default.(func() time.Time)
	// personDescUpdatedAt is the schema descriptor for updated_at field.
	personDescUpdatedAt := personMixinFields0[2].Descriptor()
	// person.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	person.DefaultUpdatedAt = personDescUpdatedAt.Default.(func() time.Time)
	// person.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	person.UpdateDefaultUpdatedAt = personDescUpdatedAt.UpdateDefault.(func() time.Time)
	personroleMixin := schema.PersonRole{}.Mixin()
	personroleMixinHooks0 := personroleMixin[0].Hooks()
	personrole.Hooks[0] = personroleMixinHooks0[0]
	personroleMixinFields0 := personroleMixin[0].Fields()
	_ = personroleMixinFields0
	personroleFields := schema.PersonRole{}.Fields()
	_ = personroleFields
	// personroleDescCreatedAt is the schema descriptor for created_at field.
	personroleDescCreatedAt := personroleMixinFields0[0].Descriptor()
	// personrole.DefaultCreatedAt holds the default value on creation for the created_at field.
	personrole.DefaultCreatedAt = personroleDescCreatedAt.Default.(func() time.Time)
	// personroleDescUpdatedAt is the schema descriptor for updated_at field.
	personroleDescUpdatedAt := personroleMixinFields0[2].Descriptor()
	// personrole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	personrole.DefaultUpdatedAt = personroleDescUpdatedAt.Default.(func() time.Time)
	// personrole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	personrole.UpdateDefaultUpdatedAt = personroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	projectMixin := schema.Project{}.Mixin()
	project.Policy = privacy.NewPolicies(schema.Project{})
	project.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := project.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	projectMixinHooks0 := projectMixin[0].Hooks()

	project.Hooks[1] = projectMixinHooks0[0]
	projectMixinFields0 := projectMixin[0].Fields()
	_ = projectMixinFields0
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescCreatedAt is the schema descriptor for created_at field.
	projectDescCreatedAt := projectMixinFields0[0].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the created_at field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	// projectDescUpdatedAt is the schema descriptor for updated_at field.
	projectDescUpdatedAt := projectMixinFields0[2].Descriptor()
	// project.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	project.DefaultUpdatedAt = projectDescUpdatedAt.Default.(func() time.Time)
	// project.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	project.UpdateDefaultUpdatedAt = projectDescUpdatedAt.UpdateDefault.(func() time.Time)
	projecttypeMixin := schema.ProjectType{}.Mixin()
	projecttypeMixinHooks0 := projecttypeMixin[0].Hooks()
	projecttype.Hooks[0] = projecttypeMixinHooks0[0]
	projecttypeMixinFields0 := projecttypeMixin[0].Fields()
	_ = projecttypeMixinFields0
	projecttypeFields := schema.ProjectType{}.Fields()
	_ = projecttypeFields
	// projecttypeDescCreatedAt is the schema descriptor for created_at field.
	projecttypeDescCreatedAt := projecttypeMixinFields0[0].Descriptor()
	// projecttype.DefaultCreatedAt holds the default value on creation for the created_at field.
	projecttype.DefaultCreatedAt = projecttypeDescCreatedAt.Default.(func() time.Time)
	// projecttypeDescUpdatedAt is the schema descriptor for updated_at field.
	projecttypeDescUpdatedAt := projecttypeMixinFields0[2].Descriptor()
	// projecttype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	projecttype.DefaultUpdatedAt = projecttypeDescUpdatedAt.Default.(func() time.Time)
	// projecttype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	projecttype.UpdateDefaultUpdatedAt = projecttypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	publicationMixin := schema.Publication{}.Mixin()
	publication.Policy = privacy.NewPolicies(schema.Publication{})
	publication.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := publication.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	publicationMixinHooks0 := publicationMixin[0].Hooks()

	publication.Hooks[1] = publicationMixinHooks0[0]
	publicationMixinFields0 := publicationMixin[0].Fields()
	_ = publicationMixinFields0
	publicationFields := schema.Publication{}.Fields()
	_ = publicationFields
	// publicationDescCreatedAt is the schema descriptor for created_at field.
	publicationDescCreatedAt := publicationMixinFields0[0].Descriptor()
	// publication.DefaultCreatedAt holds the default value on creation for the created_at field.
	publication.DefaultCreatedAt = publicationDescCreatedAt.Default.(func() time.Time)
	// publicationDescUpdatedAt is the schema descriptor for updated_at field.
	publicationDescUpdatedAt := publicationMixinFields0[2].Descriptor()
	// publication.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	publication.DefaultUpdatedAt = publicationDescUpdatedAt.Default.(func() time.Time)
	// publication.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	publication.UpdateDefaultUpdatedAt = publicationDescUpdatedAt.UpdateDefault.(func() time.Time)
	regionMixin := schema.Region{}.Mixin()
	region.Policy = privacy.NewPolicies(schema.Region{})
	region.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := region.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	regionMixinHooks0 := regionMixin[0].Hooks()

	region.Hooks[1] = regionMixinHooks0[0]
	regionMixinFields0 := regionMixin[0].Fields()
	_ = regionMixinFields0
	regionFields := schema.Region{}.Fields()
	_ = regionFields
	// regionDescCreatedAt is the schema descriptor for created_at field.
	regionDescCreatedAt := regionMixinFields0[0].Descriptor()
	// region.DefaultCreatedAt holds the default value on creation for the created_at field.
	region.DefaultCreatedAt = regionDescCreatedAt.Default.(func() time.Time)
	// regionDescUpdatedAt is the schema descriptor for updated_at field.
	regionDescUpdatedAt := regionMixinFields0[2].Descriptor()
	// region.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	region.DefaultUpdatedAt = regionDescUpdatedAt.Default.(func() time.Time)
	// region.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	region.UpdateDefaultUpdatedAt = regionDescUpdatedAt.UpdateDefault.(func() time.Time)
	setMixin := schema.Set{}.Mixin()
	set.Policy = privacy.NewPolicies(schema.Set{})
	set.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := set.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	setMixinHooks0 := setMixin[0].Hooks()

	set.Hooks[1] = setMixinHooks0[0]
	setMixinFields0 := setMixin[0].Fields()
	_ = setMixinFields0
	setFields := schema.Set{}.Fields()
	_ = setFields
	// setDescCreatedAt is the schema descriptor for created_at field.
	setDescCreatedAt := setMixinFields0[0].Descriptor()
	// set.DefaultCreatedAt holds the default value on creation for the created_at field.
	set.DefaultCreatedAt = setDescCreatedAt.Default.(func() time.Time)
	// setDescUpdatedAt is the schema descriptor for updated_at field.
	setDescUpdatedAt := setMixinFields0[2].Descriptor()
	// set.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	set.DefaultUpdatedAt = setDescUpdatedAt.Default.(func() time.Time)
	// set.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	set.UpdateDefaultUpdatedAt = setDescUpdatedAt.UpdateDefault.(func() time.Time)
	settlementMixin := schema.Settlement{}.Mixin()
	settlement.Policy = privacy.NewPolicies(schema.Settlement{})
	settlement.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := settlement.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	settlementMixinHooks0 := settlementMixin[0].Hooks()

	settlement.Hooks[1] = settlementMixinHooks0[0]
	settlementMixinFields0 := settlementMixin[0].Fields()
	_ = settlementMixinFields0
	settlementFields := schema.Settlement{}.Fields()
	_ = settlementFields
	// settlementDescCreatedAt is the schema descriptor for created_at field.
	settlementDescCreatedAt := settlementMixinFields0[0].Descriptor()
	// settlement.DefaultCreatedAt holds the default value on creation for the created_at field.
	settlement.DefaultCreatedAt = settlementDescCreatedAt.Default.(func() time.Time)
	// settlementDescUpdatedAt is the schema descriptor for updated_at field.
	settlementDescUpdatedAt := settlementMixinFields0[2].Descriptor()
	// settlement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	settlement.DefaultUpdatedAt = settlementDescUpdatedAt.Default.(func() time.Time)
	// settlement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	settlement.UpdateDefaultUpdatedAt = settlementDescUpdatedAt.UpdateDefault.(func() time.Time)
	techniqueMixin := schema.Technique{}.Mixin()
	technique.Policy = privacy.NewPolicies(schema.Technique{})
	technique.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := technique.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	techniqueMixinHooks0 := techniqueMixin[0].Hooks()

	technique.Hooks[1] = techniqueMixinHooks0[0]
	techniqueMixinFields0 := techniqueMixin[0].Fields()
	_ = techniqueMixinFields0
	techniqueFields := schema.Technique{}.Fields()
	_ = techniqueFields
	// techniqueDescCreatedAt is the schema descriptor for created_at field.
	techniqueDescCreatedAt := techniqueMixinFields0[0].Descriptor()
	// technique.DefaultCreatedAt holds the default value on creation for the created_at field.
	technique.DefaultCreatedAt = techniqueDescCreatedAt.Default.(func() time.Time)
	// techniqueDescUpdatedAt is the schema descriptor for updated_at field.
	techniqueDescUpdatedAt := techniqueMixinFields0[2].Descriptor()
	// technique.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	technique.DefaultUpdatedAt = techniqueDescUpdatedAt.Default.(func() time.Time)
	// technique.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	technique.UpdateDefaultUpdatedAt = techniqueDescUpdatedAt.UpdateDefault.(func() time.Time)
}

const (
	Version = "v0.12.3"                                         // Version of ent codegen.
	Sum     = "h1:N5lO2EOrHpCH5HYfiMOCHYbo+oh5M8GjT0/cx5x6xkk=" // Sum of ent codegen.
)
