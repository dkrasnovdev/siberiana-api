package schema

import "entgo.io/ent"

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return nil
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return nil
}
