package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DraftMixin struct {
	mixin.Schema
}

func (DraftMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("listed", "unlisted", "draft").
			Optional().
			Default("draft"),
	}
}
