package hook

import (
	"context"
	"fmt"

	"github.com/dkrasnovdev/heritage-api/ent"
	"github.com/dkrasnovdev/heritage-api/ent/hook"
	"github.com/dkrasnovdev/heritage-api/internal/ent/privacy"
	"github.com/dkrasnovdev/heritage-api/pkg/transform"
)

// ArtifactLogger is a hook that logs changes to an Artifact entity.
func ArtifactLogger(client *ent.Client) ent.Hook {
	// Register the ArtifactFunc hook.
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.ArtifactFunc(func(ctx context.Context, m *ent.ArtifactMutation) (ent.Value, error) {
			// Track changes to field values.
			changes := make(map[string]ent.Value)
			for _, key := range m.Fields() {
				nextValue, exists := m.Field(key)
				prevValue, err := m.OldField(ctx, key)
				if exists {
					// Handle the case when there was an error fetching the previous value.
					if err != nil {
						changes[key] = fmt.Sprintf("%v", nextValue)
					} else {
						changes[key] = fmt.Sprintf("%v -> %v", prevValue, nextValue)
					}
				}
			}

			// Track IDs of added edges.
			addedIDs := make(map[string]ent.Value)
			for _, key := range m.AddedEdges() {
				addedIDs[key] = m.AddedIDs(key)
			}

			// Track IDs of removed edges.
			removedIDs := make(map[string]ent.Value)
			for _, key := range m.RemovedEdges() {
				removedIDs[key] = m.RemovedIDs(key)
			}

			// Execute the mutation.
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, fmt.Errorf("failed to execute the mutation: %w", err)
			}

			// Retrieve the Artifact ID.
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
					SetAddedIds(transform.KeyValue(addedIDs)).
					SetRemovedIds(transform.KeyValue(removedIDs)).
					SetClearedEdges(m.ClearedEdges()).
					SetBlame(usr).
					Exec(ctx)
			}()

			return value, nil
		})
	}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne)
}
