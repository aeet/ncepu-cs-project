package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// 班级负责人

type ClassLeader struct {
	ent.Schema
}

func (ClassLeader) Fields() []ent.Field {
	return []ent.Field{}
}

func (ClassLeader) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class", Class.Type).Unique(),
		edge.To("student", Student.Type).Unique(),
	}
}
