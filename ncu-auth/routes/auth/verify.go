package auth

import (
	"time"

	"github.com/devcui/ncepu-cs-project/ncu-auth/oauth2"
	"github.com/devcui/ncepu-cs-project/ncu-auth/status"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func AuthorizeVerifyHandler(context echo.Context) error {
	log.Infof("AuthorizeVerifyHandler: %s", context.Request().URL.Path)
	token, err := oauth2.Server.ValidationBearerToken(context.Request())
	if err != nil {
		return context.JSON(status.Http401("invalid token", nil))
	}
	client, err := oauth2.Manager.GetClient(context.Request().Context(), token.GetClientID())
	if err != nil {
		return context.JSON(status.Http401("invalid client", nil))
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"user_id":    token.GetUserID(),
		"client_id":  token.GetClientID(),
		"scope":      token.GetScope(),
		"domain":     client.GetDomain(),
	}
	return context.JSON(status.Http200("verify token succeed", data))
}
