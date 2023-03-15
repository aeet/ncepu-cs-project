package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 家庭信息
type FamilyInfo struct {
	ent.Schema
}

func (FamilyInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("家庭成员姓名"),
		field.String("relationship").Comment("与学生关系"),
		field.String("id_card").Comment("身份证号"),
		field.String("age").Comment("年龄"),
		field.String("occupation").Comment("职业"),
		field.String("post").Comment("职务"),
		field.String("work_unit").Comment("工作单位"),
		field.String("contact_number").Comment("联系电话"),
		field.String("health").Comment("健康状况"),
	}
}

func (FamilyInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
