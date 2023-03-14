package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 系部

type Department struct {
	ent.Schema
}

func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("系部名称").Unique(),
		field.String("code").Comment("系部代码").Unique(),
		field.String("description").Comment("系部描述"),
	}
}

func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("major", Major.Type).Ref("department"),
		edge.From("class", Class.Type).Ref("department"),
		edge.From("student", Student.Type).Ref("department"),
	}
}
