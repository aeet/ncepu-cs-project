package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 实践经历
type PracticalExperience struct {
	ent.Schema
}

func (PracticalExperience) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("实践名称"),
		field.String("unit").Comment("实践单位"),
		field.Time("start_time").Comment("开始时间"),
		field.Time("end_time").Comment("结束时间"),
		field.String("describe").Comment("实践描述"),
	}
}

func (PracticalExperience) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
