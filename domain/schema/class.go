package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 班级

type Class struct {
	ent.Schema
}

func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("班级名称"),
		field.String("code").Comment("班级编号"),
		field.String("description").Comment("班级描述"),
		field.String("type").Comment("班级类型"),
	}
}

func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("major", Major.Type).Unique(),
		edge.To("department", Department.Type).Unique(),
		edge.From("campus", Campus.Type).Ref("class").Unique(),
		edge.From("student", Student.Type).Ref("class"),
		edge.From("class_leader", ClassLeader.Type).Ref("class").Unique(),
		edge.From("tutor", Tutor.Type).Ref("class").Unique(),
		edge.From("major_direction", MajorDirection.Type).Ref("class").Unique(),
	}
}
