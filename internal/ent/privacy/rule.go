package privacy

import (
	"context"

	"github.com/dkrasnovdev/heritage-api/ent/privacy"
)

// DenyIfNoViewer is a rule that returns deny decision if the viewer is missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer not found in the context or user not authenticated")
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfAdministrator is a rule that returns allow decision if the viewer is moderator.
func AllowIfAdministrator() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := FromContext(ctx)
		if view.IsAdministrator() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfModerator is a rule that returns allow decision if the viewer is moderator.
func AllowIfModerator() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := FromContext(ctx)
		if view.IsModerator() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfResearcher is a rule that returns allow decision if the viewer is researcher.
func AllowIfResearcher() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := FromContext(ctx)
		if view.IsResearcher() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}
