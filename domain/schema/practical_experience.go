package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// 实践经历
type PracticalExperience struct {
	ent.Schema
}

func (PracticalExperience) Fields() []ent.Field {
	return []ent.Field{}
}

func (PracticalExperience) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
