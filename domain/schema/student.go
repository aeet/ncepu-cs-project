package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("student").Unique().Required(),
	}
}
