package mixin

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
)

// AuditMixin holds the schema definition for the mixin.
type AuditMixin struct {
	mixin.Schema
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		// created_at field represents the creation time of the record.
		// It is set to be immutable and initialized with the current time.
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),

		// created_by field represents the user who created the record.
		// It is optional and can be set to track the creator of the record.
		field.String("created_by").
			Optional(),

		// updated_at field represents the last update time of the record.
		// It is initialized with the current time on creation and update.
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entgql.OrderField("UPDATED_AT")),

		// updated_by field represents the user who last updated the record.
		// It is optional and can be set to track the updater of the record.
		field.String("updated_by").
			Optional(),
	}
}

// Hooks returns the audit log hook for the mixin.
func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}

// AuditHook is a hook that performs the audit logging.
// It receives the next ent.Mutator and wraps it to perform audit logging.
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger interface defines methods to access audit log fields.
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedBy(string)
		CreatedBy() (usr string, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedBy(string)
		UpdatedBy() (usr string, exists bool)
	}

	// Return the wrapped ent.MutateFunc.
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		// Try to assert the incoming mutation to the AuditLogger interface.
		mx, ok := m.(AuditLogger)
		if !ok {
			// If the assertion fails, return an error.
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
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

		// Check the operation type and perform audit logging accordingly.
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			// For create operations, set the created_at timestamp to the current time.
			mx.SetCreatedAt(time.Now())

			// If the created_by field is not set, set it to the preferred username.
			if _, exists := mx.CreatedBy(); !exists {
				mx.SetCreatedBy(usr)
			}

		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			// For update operations, set the updated_at timestamp to the current time.
			mx.SetUpdatedAt(time.Now())

			// If the updated_by field is not set, set it to the preferred username.
			if _, exists := mx.UpdatedBy(); !exists {
				mx.SetUpdatedBy(usr)
			}
		}

		// Call the next mutator in the chain.
		return next.Mutate(ctx, m)
	})
}
