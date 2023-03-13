package endpoint

import (
	"context"
	"strconv"

	"github.com/devcui/ncepu-cs-project/ncu-domain/service"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/labstack/gommon/log"
)

func PasswordAuthorizationHandler(ctx context.Context, clientID, account, password string) (userID string, err error) {
	log.Infof("PasswordAuthorizationHandler: clientID=%s, account=%s, password=%s", clientID, account, password)
	exist, e := service.AuthorizationExists(clientID)
	if e != nil {
		err = e
		return
	}
	if !exist {
		err = errors.New("client not found")
		return
	}
	user, e := service.User(account, password)
	if e != nil {
		err = e
		return
	}
	return strconv.Itoa(user.ID), nil
}
