package auth

import (
	"github.com/devcui/ncepu-cs-project/ncu-auth/oauth2"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func AuthorizeTokenHandler(context echo.Context) error {
	log.Infof("AuthorizeTokenHandler: %s", context.Request().URL.Path)
	log.Infof("%v", context.QueryParams())
	return oauth2.Server.HandleTokenRequest(context.Response(), context.Request())
}
