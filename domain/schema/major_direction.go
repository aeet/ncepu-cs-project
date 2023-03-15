package schema

// 专业方向
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type MajorDirection struct {
	ent.Schema
}

func (MajorDirection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("专业方向名称").Unique(),
	}
}

func (MajorDirection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class", Class.Type).Unique(),
	}
}
