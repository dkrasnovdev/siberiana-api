package mixin

import (
	"entgo.io/contrib/entgql"
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
			Optional().
			Annotations(
				entgql.OrderField("DISPLAY_NAME"),
			),
		field.String("abbreviation").
			Optional(),
		field.String("description").
			Optional(),
		field.String("external_link").
			Optional(),
	}
}
