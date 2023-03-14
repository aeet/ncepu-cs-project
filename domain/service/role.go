package service

import (
	"context"
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/role"
)

func RoleAdd(role domain.Role) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Role.Create().SetRoleName(role.RoleName).SetRoleValue(role.RoleValue).Save(context.Background())
	})
	return err
}

func RoleDelete(id string) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		return client.Role.Delete().Where(role.ID(intID)).Exec(context.Background())
	})
	return err
}

func RoleUpdate(r domain.Role) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Role.Update().Where(role.ID(r.ID)).SetRoleName(r.RoleName).SetRoleValue(r.RoleValue).Save(context.Background())
	})
	return err
}

func RoleQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Role.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
