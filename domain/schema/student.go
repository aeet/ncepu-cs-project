package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
		edge.To("department", Department.Type).Unique(),
		edge.To("major", Major.Type).Unique(),
		edge.To("class", Class.Type).Unique(),
		edge.From("class_leader", ClassLeader.Type).Ref("student").Unique(),
		edge.From("tutor", Tutor.Type).Ref("student").Unique(),
		edge.From("certificate", Certificate.Type).Ref("student"),
		edge.From("education_level", EducationLevel.Type).Ref("student"),
		edge.From("enrollment_status", EnrollmentStatus.Type).Ref("student"),
		edge.From("family_info", FamilyInfo.Type).Ref("student"),
		edge.From("practical_experience", PracticalExperience.Type).Ref("student"),
	}
}
