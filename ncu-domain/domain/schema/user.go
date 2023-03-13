package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("account").NotEmpty(),
		field.String("passwd").NotEmpty(),
		field.String("username").NotEmpty(),
		field.Bytes("avatar").Nillable(),
		field.String("email").NotEmpty(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type),
		edge.To("resource", Resource.Type),
	}
}
