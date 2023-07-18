package schema

import "entgo.io/ent"

// License holds the schema definition for the License entity.
type License struct {
	ent.Schema
}

// Fields of the License.
func (License) Fields() []ent.Field {
	return nil
}

// Edges of the License.
func (License) Edges() []ent.Edge {
	return nil
}
