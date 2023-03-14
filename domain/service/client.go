package service

import (
	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/mysql"
)

func HandleByClient(handle func(client *domain.Client) (interface{}, error)) (interface{}, error) {
	client, err := mysql.Open()
	if err != nil {
		return nil, err
	}
	res, e := handle(client)
	if e != nil {
		return nil, err
	}
	return res, client.Close()
}
