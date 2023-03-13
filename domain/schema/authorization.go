package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Authorization struct {
	ent.Schema
}

func (Authorization) Fields() []ent.Field {
	return []ent.Field{
		field.String("client_id").NotEmpty().Unique(),
		field.String("client_secret").NotEmpty(),
		field.String("client_name").NotEmpty().Unique(),
		field.JSON("grant_type", []string{}),
		field.JSON("scope", []string{}),
		field.String("redirect_url"),
		field.String("domain"),
	}
}

func (Authorization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("resource", Resource.Type),
	}
}
