package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/privacy"
	rule "github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
)

// AuditLog holds the schema definition for the AuditLog entity.
type AuditLog struct {
	ent.Schema
}

// Privacy policy of the AuditLog.
func (AuditLog) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdministrator(),
			privacy.AlwaysDenyRule(),
		},
		Mutation: privacy.MutationPolicy{},
	}
}

// Annotations of the AuditLog.
func (AuditLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
	}
}

// Fields of the AuditLog.
func (AuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("table").Optional(),
		field.Int("ref_id").Optional(),
		field.String("operation").Optional(),
		field.JSON("changes", []string{}).Optional(),
		field.JSON("added_edges", []string{}).Optional(),
		field.JSON("removed_edges", []string{}).Optional(),
		field.JSON("cleared_edges", []string{}).Optional(),
		field.String("blame").Optional(),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}
