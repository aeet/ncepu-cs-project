package schema

// 专业方向
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type MajorDirection struct {
	ent.Schema
}

func (MajorDirection) Fields() []ent.Field {
	return []ent.Field{}
}

func (MajorDirection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class", Class.Type).Unique(),
	}
}
