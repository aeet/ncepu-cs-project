package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// 获得证书
type Certificate struct {
	ent.Schema
}

func (Certificate) Fields() []ent.Field {
	return []ent.Field{}
}

func (Certificate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
	}

}
