package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DatingMixin struct {
	mixin.Schema
}

func (DatingMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("dating").Optional(),
		field.Int("dating_start").Optional(),
		field.Int("dating_end").Optional(),
	}
}
