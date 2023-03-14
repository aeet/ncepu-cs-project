package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// 培养层次
type EducationLevel struct {
	ent.Schema
}

func (EducationLevel) Fields() []ent.Field {
	return []ent.Field{}
}

func (EducationLevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
