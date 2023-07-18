package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type ImagesMixin struct {
	mixin.Schema
}

func (ImagesMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("primary_image_url").
			Optional(),
		field.JSON("additional_image_urls", []string{}).
			Optional(),
	}
}
