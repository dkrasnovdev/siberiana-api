package schema

import "entgo.io/ent"

// AuditLog holds the schema definition for the AuditLog entity.
type AuditLog struct {
	ent.Schema
}

// Fields of the AuditLog.
func (AuditLog) Fields() []ent.Field {
	return nil
}

// Edges of the AuditLog.
func (AuditLog) Edges() []ent.Edge {
	return nil
}
