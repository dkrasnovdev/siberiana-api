package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DetailsMixin struct {
	mixin.Schema
}

func (DetailsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name").
			Optional(),
		field.String("abbreviation").
			Optional(),
		field.String("description").
			Optional(),
		field.JSON("external_links", []string{}).
			Optional(),
	}
}
