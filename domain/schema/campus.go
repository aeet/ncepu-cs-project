package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 校区
type Campus struct {
	ent.Schema
}

func (Campus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("校区名称"),
		field.String("address").Comment("校区地址"),
	}
}

func (Campus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class", Class.Type).Unique(),
	}
}
