package service

import (
	"context"
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/campus"
)

func CampusAdd(campus domain.Campus) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Campus.Create().SetName(campus.Name).SetAddress(campus.Address).Save(context.Background())
	})
	return err
}

func CampusDelete(id string) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		return client.Campus.Delete().Where(campus.ID(intID)).Exec(context.Background())
	})
	return err
}

func CampusUpdate(r domain.Campus) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Campus.Update().Where(campus.ID(r.ID)).SetName(r.Name).SetAddress(r.Address).Save(context.Background())
	})
	return err
}

func CampusQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Campus.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
