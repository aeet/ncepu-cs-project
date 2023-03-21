package service

import (
	"context"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/familyinfo"
)

func FamilyAdd(s domain.FamilyInfo) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.FamilyInfo.Create().
			SetStudentID(s.Edges.Student.ID).
			SetOccupation(s.Occupation).
			SetAge(s.Age).
			SetName(s.Name).
			SetHealth(s.Health).
			SetIDCard(s.IDCard).
			SetWorkUnit(s.WorkUnit).
			SetContactNumber(s.ContactNumber).
			SetPost(s.Post).
			SetRelationship(s.Relationship).SaveX(context.Background()), nil
	})
	return err
}

func FamilyDelete(id int) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.FamilyInfo.Delete().Where(familyinfo.ID(id)).Exec(context.Background())
	})
	return err
}

func FamilyUpdate(s domain.FamilyInfo) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.FamilyInfo.Update().
			Where(familyinfo.ID(s.ID)).
			SetAge(s.Age).
			SetName(s.Name).
			SetHealth(s.Health).
			SetIDCard(s.IDCard).
			SetWorkUnit(s.WorkUnit).
			SetContactNumber(s.ContactNumber).
			SetPost(s.Post).
			SetRelationship(s.Relationship).Save(context.Background())
	})
	return err
}

func FamilyQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.FamilyInfo.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
