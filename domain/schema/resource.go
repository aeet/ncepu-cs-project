package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Resource struct {
	ent.Schema
}

func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("resource_name").NotEmpty(),
		field.String("resource_value").NotEmpty(),
	}
}

func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("resource"),
		edge.From("user", User.Type).Ref("resource"),
		edge.From("authorization", Authorization.Type).Ref("resource"),
	}
}
