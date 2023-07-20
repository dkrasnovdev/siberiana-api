package schema

import "entgo.io/ent"

// BookGenre holds the schema definition for the BookGenre entity.
type BookGenre struct {
	ent.Schema
}

// Fields of the BookGenre.
func (BookGenre) Fields() []ent.Field {
	return nil
}

// Edges of the BookGenre.
func (BookGenre) Edges() []ent.Edge {
	return nil
}
