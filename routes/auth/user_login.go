package auth

import (
	"github.com/devcui/ncepu-cs-project/oauth2"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func AuthorizeUserLoginHandler(context echo.Context) error {
	if context.Request().Form == nil {
		context.Request().ParseForm()
	}
	context.Request().Form.Set("client_id", "auth")
	context.Request().Form.Set("client_secret", "auth")
	context.Request().Form.Set("grant_type", "password")
	context.Request().Form.Set("scope", "all")
	err := oauth2.Server.HandleTokenRequest(context.Response(), context.Request())
	if err != nil {
		return context.JSON(status.Http500("", err))
	}
	return nil
}
