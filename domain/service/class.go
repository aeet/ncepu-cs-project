package service

import (
	"context"
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/class"
	"github.com/devcui/ncepu-cs-project/domain/major"
)

// 班级

func ClassAdd(d domain.Class) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Class.Create().SetName(d.Name).SetCode(d.Code).SetDescription(d.Description).SetType(d.Type).Save(context.Background())
	})
	return err
}

func ClassDelete(id string) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		return client.Class.Delete().Where(class.ID(intID)).Exec(context.Background())
	})
	return err
}

func ClassUpdate(d domain.Class) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		a := client.Class.Update().Where(class.ID(d.ID)).SetName(d.Name).SetCode(d.Code).SetDescription(d.Description).SetType(d.Type)
		if d.Edges.Major != nil {
			a.SetMajorID(d.Edges.Major.ID)
			major, _ := client.Major.Query().Where(major.ID(d.Edges.Major.ID)).WithDepartment().First(context.Background())
			if major != nil && major.Edges.Department != nil {
				a.SetDepartmentID(major.Edges.Department.ID)
			}
			return a.Save(context.Background())
		}
		return a.Save(context.Background())
	})
	return err
}

func ClassQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Class.Query().WithMajor().WithDepartment().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
