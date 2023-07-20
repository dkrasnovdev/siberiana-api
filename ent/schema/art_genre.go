package schema

import "entgo.io/ent"

// ArtGenre holds the schema definition for the ArtGenre entity.
type ArtGenre struct {
	ent.Schema
}

// Fields of the ArtGenre.
func (ArtGenre) Fields() []ent.Field {
	return nil
}

// Edges of the ArtGenre.
func (ArtGenre) Edges() []ent.Edge {
	return nil
}
