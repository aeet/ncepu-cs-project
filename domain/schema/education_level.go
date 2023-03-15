package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 培养层次
type EducationLevel struct {
	ent.Schema
}

func (EducationLevel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("培养层次名称").Unique(),
	}
}

func (EducationLevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
