package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// 学籍状态
type EnrollmentStatus struct {
	ent.Schema
}

func (EnrollmentStatus) Fields() []ent.Field {
	return []ent.Field{}
}

func (EnrollmentStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
