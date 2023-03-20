package service

import (
	"context"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/practicalexperience"
)

func PracticalAdd(s domain.PracticalExperience) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.PracticalExperience.Create().
			SetStudentID(s.Edges.Student.ID).
			SetDescribe(s.Describe).
			SetEndTime(s.EndTime).
			SetName(s.Name).
			SetStartTime(s.StartTime).
			SetUnit(s.Unit).Save(context.Background())
	})
	return err
}

func PracticalDelete(id int) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.PracticalExperience.Delete().Where(practicalexperience.ID(id)).Exec(context.Background())
	})
	return err
}

func PracticalUpdate(s domain.PracticalExperience) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.PracticalExperience.Update().Where(practicalexperience.ID(s.ID)).
			SetDescribe(s.Describe).
			SetEndTime(s.EndTime).
			SetName(s.Name).
			SetStartTime(s.StartTime).
			SetUnit(s.Unit).Save(context.Background())
	})
	return err
}

func PracticalQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.PracticalExperience.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
