package hook

import (
	"context"
	"fmt"

	"github.com/dkrasnovdev/heritage-api/ent"
	"github.com/dkrasnovdev/heritage-api/ent/hook"
	"github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/pkg/transform"
)

// SetLogger is a hook that logs changes to an Set entity.
func SetLogger(client *ent.Client) ent.Hook {
	// Register the SetFunc hook.
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.SetFunc(func(ctx context.Context, m *ent.SetMutation) (ent.Value, error) {
			// Track changes to field values.
			changes := make(map[string]ent.Value)
			for _, key := range m.Fields() {
				// Skip tracking 'updated_at' field.
				if key == "updated_at" {
					continue
				}
				nextValue, exists := m.Field(key)
				prevValue, err := m.OldField(ctx, key)
				if exists && nextValue != prevValue {
					// Handle the case when there was an error fetching the previous value.
					if err != nil || prevValue == nil || prevValue == "" {
						changes[key] = fmt.Sprintf("(new) %v", nextValue)
					} else {
						if nextValue == "" || nextValue == nil {
							changes[key] = fmt.Sprintf("%v -> (removed)", prevValue)
						} else {
							changes[key] = fmt.Sprintf("%v -> %v", prevValue, nextValue)
						}
					}
				}
			}

			// Track IDs of added edges.
			addedEdges := make(map[string]ent.Value)
			for _, key := range m.AddedEdges() {
				addedEdges[key] = m.AddedIDs(key)
			}

			// Track IDs of removed edges.
			removedEdges := make(map[string]ent.Value)
			for _, key := range m.RemovedEdges() {
				removedEdges[key] = m.RemovedIDs(key)
			}

			// Execute the mutation.
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, fmt.Errorf("failed to execute the mutation: %w", err)
			}

			// If there are no changes or edge modifications, skip logging.
			if len(changes) == 0 && len(addedEdges) == 0 && len(removedEdges) == 0 && len(m.ClearedEdges()) == 0 {
				return value, nil
			}

			// Retrieve the Set ID.
			refID, _ := m.ID()

			// Retrieve the Viewer from the context.
			v := privacy.FromContext(ctx)

			// Check if the Viewer is present in the context.
			// If the Viewer is not found, it means that the user information is missing or not properly authenticated.
			if v == nil {
				// Return an error indicating that the Viewer is not available or the user is not authenticated.
				return nil, fmt.Errorf("viewer not found in the context or user not authenticated")
			}

			// Get the preferred username from the privacy policy.
			usr := v.GetPreferredUsername()

			// Create an audit log entry for the mutation.
			defer func() {
				err = client.AuditLog.
					Create().
					SetRefID(refID).
					SetTable(m.Type()).
					SetOperation(m.Op().String()).
					SetChanges(transform.KeyValue(changes)).
					SetAddedEdges(transform.KeyValue(addedEdges)).
					SetRemovedEdges(transform.KeyValue(removedEdges)).
					SetClearedEdges(m.ClearedEdges()).
					SetBlame(usr).
					Exec(ctx)
			}()

			return value, nil
		})
	}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne)
}
