package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// 家庭信息
type FamilyInfo struct {
	ent.Schema
}

func (FamilyInfo) Fields() []ent.Field {
	return []ent.Field{}
}

func (FamilyInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
