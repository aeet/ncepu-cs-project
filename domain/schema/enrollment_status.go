package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 学籍状态
type EnrollmentStatus struct {
	ent.Schema
}

func (EnrollmentStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("学籍状态名称").Unique(),
	}
}

func (EnrollmentStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
