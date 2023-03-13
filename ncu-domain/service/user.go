package service

import (
	"context"

	"ncu-domain/domain"
	"ncu-domain/domain/user"
	"ncu-domain/mysql"
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
