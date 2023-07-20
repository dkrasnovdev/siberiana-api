package schema

import "entgo.io/ent"

// ArtStyle holds the schema definition for the ArtStyle entity.
type ArtStyle struct {
	ent.Schema
}

// Fields of the ArtStyle.
func (ArtStyle) Fields() []ent.Field {
	return nil
}

// Edges of the ArtStyle.
func (ArtStyle) Edges() []ent.Edge {
	return nil
}
