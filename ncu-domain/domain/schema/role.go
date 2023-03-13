package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_name").NotEmpty(),
		field.String("role_value").NotEmpty(),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("resource", Resource.Type),
		edge.From("user", User.Type).Ref("role"),
	}
}
