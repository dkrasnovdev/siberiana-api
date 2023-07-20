package hook

import (
	"context"
	"fmt"

	"github.com/dkrasnovdev/heritage-api/ent"
	"github.com/dkrasnovdev/heritage-api/ent/hook"
	"github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/pkg/transform"
)

// LocationLogger is a hook that logs changes to an Location entity.
func LocationLogger(client *ent.Client) ent.Hook {
	// Register the LocationFunc hook.
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.LocationFunc(func(ctx context.Context, m *ent.LocationMutation) (ent.Value, error) {
			var blame string

			// Track changes to field values.
			changes := make(map[string]ent.Value)
			for _, key := range m.Fields() {
				// Skip tracking 'updated_at' and 'updated_by' fields.
				if key == "updated_at" || key == "updated_by" {
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

			// Retrieve the Location ID.
			refID, _ := m.ID()

			// Retrieve the Viewer from the context.
			v := privacy.FromContext(ctx)

			// Determine the blame (responsible user).
			if v == nil {
				blame = "appctl" // Default to "appctl" if the Viewer is not present in the context.
			} else {
				blame = v.GetPreferredUsername() // Get the preferred username from the Viewer.
			}

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
					SetBlame(blame).
					Exec(ctx)
			}()

			return value, nil
		})
	}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne)
}
