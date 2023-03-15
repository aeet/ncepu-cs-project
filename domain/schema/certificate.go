package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 获得证书
type Certificate struct {
	ent.Schema
}

func (Certificate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("证书名称").Unique(),
		field.String("code").Comment("证书代码").Unique(),
		field.String("description").Comment("证书描述"),
		field.String("department").Comment("颁发部门"),
		field.Time("issue_date").Comment("颁发日期"),
		field.String("certificate_type").Comment("证书类型"),
		field.String("certificate_level").Comment("证书级别"),
		field.String("certificate_type2").Comment("证书类别"),
		field.String("award_category").Comment("获奖类别"),
		field.Bytes("certificate_image").Comment("证书图片"),
	}
}

func (Certificate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
