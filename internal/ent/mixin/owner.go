package mixin

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
)

// OwnerMixin holds the schema definition for the mixin.
type OwnerMixin struct {
	mixin.Schema
}

// Fields of the OwnerMixin.
func (OwnerMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner_id").
			NotEmpty().
			Immutable(),
	}
}

// Hooks returns the owner hook for the mixin.
func (OwnerMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		OwnerHook,
	}
}

// Receives the next ent.Mutator and wraps it to perform owner assertation.
func OwnerHook(next ent.Mutator) ent.Mutator {
	// OwnerAssertation interface defines methods to access owner field.
	type OwnerAssertation interface {
		OwnerId() (usr string, exists bool)
		SetOwnerId(string)
	}

	// Return the wrapped ent.MutateFunc.
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		// Try to assert the incoming mutation to the OwnerAssertation interface.
		mx, ok := m.(OwnerAssertation)
		if !ok {
			// If the assertion fails, return an error.
			return nil, fmt.Errorf("unexpected call from mutation type %T", m)
		}

		// Retrieve the Viewer from the context.
		v := privacy.FromContext(ctx)

		// Check if the Viewer is present in the context.
		// If the Viewer is not found, it means that the user information is missing or not properly authenticated.
		if v == nil {
			// Return an error indicating that the Viewer is not available or the user is not authenticated.
			return nil, fmt.Errorf("not authenticated")
		}

		// Get the preferred username from the privacy policy.
		usr := v.GetPreferredUsername()

		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			// Set the owner_id field to the username from the context if it doesn't already exist.
			if _, exists := mx.OwnerId(); !exists {
				mx.SetOwnerId(usr)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			// Verify if the current owner matches the username and returns an error if not.
			owner, _ := mx.OwnerId()
			if owner != usr {
				return nil, fmt.Errorf("forbidden")
			}
		}

		// Call the next mutator in the chain.
		return next.Mutate(ctx, m)
	})
}
