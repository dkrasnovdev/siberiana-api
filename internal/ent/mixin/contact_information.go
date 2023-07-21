package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type ContactInformationMixin struct {
	mixin.Schema
}

func (ContactInformationMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").
			Optional(),
		field.JSON("phone_numbers", []string{}).
			Optional(),
		field.JSON("emails", []string{}).
			Optional(),
	}
}
