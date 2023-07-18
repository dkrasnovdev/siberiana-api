// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/dkrasnovdev/heritage-api/ent"

	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The ArtifactQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ArtifactQueryRuleFunc func(context.Context, *ent.ArtifactQuery) error

// EvalQuery return f(ctx, q).
func (f ArtifactQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ArtifactQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ArtifactQuery", q)
}

// The ArtifactMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ArtifactMutationRuleFunc func(context.Context, *ent.ArtifactMutation) error

// EvalMutation calls f(ctx, m).
func (f ArtifactMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ArtifactMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ArtifactMutation", m)
}

// The CategoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CategoryQueryRuleFunc func(context.Context, *ent.CategoryQuery) error

// EvalQuery return f(ctx, q).
func (f CategoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CategoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CategoryQuery", q)
}

// The CategoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CategoryMutationRuleFunc func(context.Context, *ent.CategoryMutation) error

// EvalMutation calls f(ctx, m).
func (f CategoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CategoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CategoryMutation", m)
}

// The CollectionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CollectionQueryRuleFunc func(context.Context, *ent.CollectionQuery) error

// EvalQuery return f(ctx, q).
func (f CollectionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CollectionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CollectionQuery", q)
}

// The CollectionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CollectionMutationRuleFunc func(context.Context, *ent.CollectionMutation) error

// EvalMutation calls f(ctx, m).
func (f CollectionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CollectionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CollectionMutation", m)
}

// The CultureQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CultureQueryRuleFunc func(context.Context, *ent.CultureQuery) error

// EvalQuery return f(ctx, q).
func (f CultureQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CultureQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CultureQuery", q)
}

// The CultureMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CultureMutationRuleFunc func(context.Context, *ent.CultureMutation) error

// EvalMutation calls f(ctx, m).
func (f CultureMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CultureMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CultureMutation", m)
}

// The DistrictQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DistrictQueryRuleFunc func(context.Context, *ent.DistrictQuery) error

// EvalQuery return f(ctx, q).
func (f DistrictQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DistrictQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DistrictQuery", q)
}

// The DistrictMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DistrictMutationRuleFunc func(context.Context, *ent.DistrictMutation) error

// EvalMutation calls f(ctx, m).
func (f DistrictMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DistrictMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DistrictMutation", m)
}

// The HolderQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type HolderQueryRuleFunc func(context.Context, *ent.HolderQuery) error

// EvalQuery return f(ctx, q).
func (f HolderQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.HolderQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.HolderQuery", q)
}

// The HolderMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type HolderMutationRuleFunc func(context.Context, *ent.HolderMutation) error

// EvalMutation calls f(ctx, m).
func (f HolderMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.HolderMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.HolderMutation", m)
}

// The LicenseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LicenseQueryRuleFunc func(context.Context, *ent.LicenseQuery) error

// EvalQuery return f(ctx, q).
func (f LicenseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LicenseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LicenseQuery", q)
}

// The LicenseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LicenseMutationRuleFunc func(context.Context, *ent.LicenseMutation) error

// EvalMutation calls f(ctx, m).
func (f LicenseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LicenseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LicenseMutation", m)
}

// The LocationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LocationQueryRuleFunc func(context.Context, *ent.LocationQuery) error

// EvalQuery return f(ctx, q).
func (f LocationQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LocationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LocationQuery", q)
}

// The LocationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LocationMutationRuleFunc func(context.Context, *ent.LocationMutation) error

// EvalMutation calls f(ctx, m).
func (f LocationMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LocationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LocationMutation", m)
}

// The MediumQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MediumQueryRuleFunc func(context.Context, *ent.MediumQuery) error

// EvalQuery return f(ctx, q).
func (f MediumQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MediumQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MediumQuery", q)
}

// The MediumMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MediumMutationRuleFunc func(context.Context, *ent.MediumMutation) error

// EvalMutation calls f(ctx, m).
func (f MediumMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MediumMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MediumMutation", m)
}

// The ModelQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ModelQueryRuleFunc func(context.Context, *ent.ModelQuery) error

// EvalQuery return f(ctx, q).
func (f ModelQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ModelQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ModelQuery", q)
}

// The ModelMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ModelMutationRuleFunc func(context.Context, *ent.ModelMutation) error

// EvalMutation calls f(ctx, m).
func (f ModelMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ModelMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ModelMutation", m)
}

// The MonumentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MonumentQueryRuleFunc func(context.Context, *ent.MonumentQuery) error

// EvalQuery return f(ctx, q).
func (f MonumentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MonumentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MonumentQuery", q)
}

// The MonumentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MonumentMutationRuleFunc func(context.Context, *ent.MonumentMutation) error

// EvalMutation calls f(ctx, m).
func (f MonumentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MonumentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MonumentMutation", m)
}

// The OrganizationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrganizationQueryRuleFunc func(context.Context, *ent.OrganizationQuery) error

// EvalQuery return f(ctx, q).
func (f OrganizationQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OrganizationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OrganizationQuery", q)
}

// The OrganizationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrganizationMutationRuleFunc func(context.Context, *ent.OrganizationMutation) error

// EvalMutation calls f(ctx, m).
func (f OrganizationMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OrganizationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OrganizationMutation", m)
}

// The PersonQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PersonQueryRuleFunc func(context.Context, *ent.PersonQuery) error

// EvalQuery return f(ctx, q).
func (f PersonQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PersonQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PersonQuery", q)
}

// The PersonMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PersonMutationRuleFunc func(context.Context, *ent.PersonMutation) error

// EvalMutation calls f(ctx, m).
func (f PersonMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PersonMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PersonMutation", m)
}

// The ProjectQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProjectQueryRuleFunc func(context.Context, *ent.ProjectQuery) error

// EvalQuery return f(ctx, q).
func (f ProjectQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProjectQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProjectQuery", q)
}

// The ProjectMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProjectMutationRuleFunc func(context.Context, *ent.ProjectMutation) error

// EvalMutation calls f(ctx, m).
func (f ProjectMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProjectMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProjectMutation", m)
}

// The PublicationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PublicationQueryRuleFunc func(context.Context, *ent.PublicationQuery) error

// EvalQuery return f(ctx, q).
func (f PublicationQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PublicationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PublicationQuery", q)
}

// The PublicationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PublicationMutationRuleFunc func(context.Context, *ent.PublicationMutation) error

// EvalMutation calls f(ctx, m).
func (f PublicationMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PublicationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PublicationMutation", m)
}

// The RegionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RegionQueryRuleFunc func(context.Context, *ent.RegionQuery) error

// EvalQuery return f(ctx, q).
func (f RegionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RegionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RegionQuery", q)
}

// The RegionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RegionMutationRuleFunc func(context.Context, *ent.RegionMutation) error

// EvalMutation calls f(ctx, m).
func (f RegionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RegionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RegionMutation", m)
}

// The SetQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SetQueryRuleFunc func(context.Context, *ent.SetQuery) error

// EvalQuery return f(ctx, q).
func (f SetQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SetQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SetQuery", q)
}

// The SetMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SetMutationRuleFunc func(context.Context, *ent.SetMutation) error

// EvalMutation calls f(ctx, m).
func (f SetMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SetMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SetMutation", m)
}

// The SettlementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SettlementQueryRuleFunc func(context.Context, *ent.SettlementQuery) error

// EvalQuery return f(ctx, q).
func (f SettlementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SettlementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SettlementQuery", q)
}

// The SettlementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SettlementMutationRuleFunc func(context.Context, *ent.SettlementMutation) error

// EvalMutation calls f(ctx, m).
func (f SettlementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SettlementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SettlementMutation", m)
}

// The TechniqueQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TechniqueQueryRuleFunc func(context.Context, *ent.TechniqueQuery) error

// EvalQuery return f(ctx, q).
func (f TechniqueQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TechniqueQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TechniqueQuery", q)
}

// The TechniqueMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TechniqueMutationRuleFunc func(context.Context, *ent.TechniqueMutation) error

// EvalMutation calls f(ctx, m).
func (f TechniqueMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TechniqueMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TechniqueMutation", m)
}
