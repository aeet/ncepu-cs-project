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
