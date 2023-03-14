package schema

// 辅导员

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type Tutor struct {
	ent.Schema
}

func (Tutor) Fields() []ent.Field {
	return []ent.Field{}
}

func (Tutor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class", Class.Type).Unique(),
		edge.To("student", Student.Type).Unique(),
	}
}
