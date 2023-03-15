package service

import (
	"context"
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/class"
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
		return client.Class.Update().Where(class.ID(d.ID)).SetName(d.Name).SetCode(d.Code).SetDescription(d.Description).SetType(d.Type).Save(context.Background())
	})
	return err
}

func ClassQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Class.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
