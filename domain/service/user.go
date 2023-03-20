package service

import (
	"context"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/user"
	"github.com/devcui/ncepu-cs-project/mysql"
)

func User(account string, passwd string) (data *domain.User, err error) {
	client, e := mysql.Open()
	if e != nil {
		err = e
		return
	}
	data, err = client.User.Query().Where(user.Account(account), user.Passwd(passwd)).First(context.Background())
	client.Close()
	return
}

func UserByID(id int) (data *domain.User, err error) {
	client, e := mysql.Open()
	if e != nil {
		err = e
		return
	}
	data, err = client.User.Query().Where(user.ID(id)).First(context.Background())
	client.Close()
	return
}

func UserAdd(s domain.User) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.User.Create().
			SetAccount(s.Account).
			SetPasswd(s.Passwd).
			SetAvatar(*s.Avatar).
			SetEmail(s.Email).
			SetUsername(s.Username).Save(context.Background())
	})
	return err
}

func UserDelete(id int) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.User.Delete().Where(user.ID(id)).Exec(context.Background())
	})
	return err
}

func UserUpdate(s domain.User) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.User.Update().
			Where(user.ID(s.ID)).
			SetAccount(s.Account).
			SetPasswd(s.Passwd).
			SetAvatar(*s.Avatar).
			SetEmail(s.Email).
			SetUsername(s.Username).Save(context.Background())
	})
	return err
}

func UserQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.User.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
