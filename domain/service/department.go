package service

import (
	"context"
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/department"
)

// 系部

func DepartmentAdd(d domain.Department) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Department.Create().SetName(d.Name).SetCode(d.Code).SetDescription(d.Description).Save(context.Background())
	})
	return err
}

func DepartmentDelete(id string) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		return client.Department.Delete().Where(department.ID(intID)).Exec(context.Background())
	})
	return err
}

func DepartmentUpdate(d domain.Department) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Department.Update().Where(department.ID(d.ID)).SetName(d.Name).SetCode(d.Code).SetDescription(d.Description).Save(context.Background())
	})
	return err
}

func DepartmentQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Department.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
