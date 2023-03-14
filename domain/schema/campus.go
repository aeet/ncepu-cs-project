package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// 校区
type Campus struct {
	ent.Schema
}

func (Campus) Fields() []ent.Field {
	return []ent.Field{}
}

func (Campus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class", Class.Type).Unique(),
	}
}
