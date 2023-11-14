package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type SizeMixin struct {
	mixin.Schema
}

func (SizeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Float("height").Optional(),
		field.Float("width").Optional(),
		field.Float("length").Optional(),
		field.Float("depth").Optional(),
		field.Float("diameter").Optional(),
		field.String("weight").Optional(),
		field.String("dimensions").Optional(),
	}
}
