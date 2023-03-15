// 专业

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 专业
type Major struct {
	ent.Schema
}

func (Major) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("专业名称").Unique(),
		field.String("code").Comment("专业代码").Unique(),
		field.String("description").Comment("专业描述"),
		field.String("special_type").Comment("特殊类型"),
		field.String("enrollment_type").Comment("招生类型"),
		field.Bool("is_major_category").Comment("是否专业大类"),
		field.String("major_category").Comment("专业大类"),
	}

}

func (Major) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("department", Department.Type).Unique(),
		edge.From("student", Student.Type).Ref("major"),
		edge.From("class", Class.Type).Ref("major"),
	}
}
