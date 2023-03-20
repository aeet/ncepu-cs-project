package service

import (
	"context"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/student"
)

func StudentAdd(s domain.Student) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Student.Create().
			SetUserID(s.Edges.User.ID).
			SetAge(s.Age).
			SetAvatar(s.Avatar).
			SetCode(s.Code).
			SetName(s.Name).
			Save(context.Background())
	})
	return err
}

func StudentDelete(id int) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Student.Delete().Where(student.ID(id)).Exec(context.Background())
	})
	return err
}

func StudentUpdate(s domain.Student) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Student.Update().
			Where(student.ID(s.ID)).
			SetAge(s.Age).
			SetAvatar(s.Avatar).
			SetCode(s.Code).
			SetName(s.Name).
			Save(context.Background())
	})
	return err
}

func StudentQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Student.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
